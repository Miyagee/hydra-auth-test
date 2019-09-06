## Test implementation of ORY Hydra in GO

Tested with Ory Hydra v1.0.0 with a simple Go implementation. This should not be done in a production environment. This repository is purely for testing purposes. 

### Setting up a docker network

    docker network create hydrago

### Setting up the database

    docker run \
    --network hydrago \
    --name ory-hydra-example--postgres \
    -e POSTGRES_USER=hydra \
    -e POSTGRES_PASSWORD=secret \
    -e POSTGRES_DB=hydra \
    -d postgres:9.6

### Migrate database 

    docker run -it --rm \
    --network hydrago \
    oryd/hydra \
    migrate sql --yes postgres://hydra:secret@ory-hydra-example--postgres:5432/hydra?sslmode=disable

### Start hydra instance

    docker run -d \
    --name ory-hydra-example--hydra \
    --network hydrago \
    -p 9000:4444 \
    -p 9001:4445 \
    -e SECRETS_SYSTEM=this_needs_to_be_the_same_always_and_also_very_$3cuR3-._ \
    -e DSN=postgres://hydra:secret@ory-hydra-example--postgres:5432/hydra?sslmode=disable \
    -e URLS_SELF_ISSUER=https://localhost:9000/ \
    -e URLS_CONSENT=http://localhost:9020/consent \
    -e URLS_LOGIN=http://localhost:9020/login \
    oryd/hydra serve all

### Create client

    docker run --rm -it \
    -e HYDRA_ADMIN_URL=https://ory-hydra-example--hydra:4445 \
    --network hydrago \
    oryd/hydra \
    clients create --skip-tls-verify \
    --id my-client \
    --secret my-secret \
    --grant-types authorization_code,refresh_token,client_credentials,implicit \
    --response-types token,code,id_token \
    --scope openid,offline,photos.read \
    --callbacks http://127.0.0.1:9010/callback

### Delete client if something got fudged

    docker run --rm -it \
    -e HYDRA_ADMIN_URL=https://ory-hydra-example--hydra:4445 \
    --network hydrago \
    oryd/hydra \
    clients delete --skip-tls-verify \
    my-client
