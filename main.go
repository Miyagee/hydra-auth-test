package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func jwks() {
	headers := map[string][]string{
		"Accept": []string{"application/json"},
	}

	var body []byte
	var content interface{}

	req, err := http.NewRequest("GET", "https://localhost:9000/.well-known/jwks.json", bytes.NewBuffer(body))
	if err != nil {
		fmt.Println("Error NewRequest")
		log.Fatal(err)
	}
	req.Header = headers

	client := &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
	}}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Error executing request")
		log.Fatal(err)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		fmt.Println("Error ReadAll")
		log.Fatal(err)
	}

	json.Unmarshal(respBody, &content)

	fmt.Println(content.(map[string]interface{}))
}

func openIDConfig() {
	headers := map[string][]string{
		"Accept": []string{"application/json"},
	}

	var body []byte
	var content interface{}
	// body = ...

	req, err := http.NewRequest("GET", "https://localhost:9000/.well-known/openid-configuration", bytes.NewBuffer(body))
	if err != nil {
		fmt.Println("Error NewRequest")
		log.Fatal(err)
	}
	req.Header = headers

	client := &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore SSL certificates
	}}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Error executing request")
		log.Fatal(err)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		fmt.Println("Error ReadAll")
		log.Fatal(err)
	}

	json.Unmarshal(respBody, &content)

	fmt.Println(content.(map[string]interface{}))
	// ...
}

func authEndpoint() {
	headers := map[string][]string{
		"Accept": []string{"application/json"},
	}

	var body []byte
	// body = ...

	req, err := http.NewRequest("GET", "https://localhost:9000/oauth2/auth", bytes.NewBuffer(body))
	if err != nil {
		fmt.Println("Error NewRequest")
		log.Fatal(err)
	}
	req.Header = headers

	client := &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore SSL certificates
	}}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Error executing request")
		log.Fatal(err)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		fmt.Println("Error parsing html")
		log.Fatal(err)
	}

	fmt.Println(string(respBody))
	// ...
}

func getAccessToken() string {
	headers := map[string][]string{
		"Accept":       []string{"application/json"},
		"Content-Type": []string{"application/x-www-form-urlencoded"},
	}

	var content interface{}
	body := []byte(`grant_type=client_credentials&scope=offline+openid`)

	req, err := http.NewRequest("POST", "https://localhost:9000/oauth2/token", bytes.NewBuffer(body))
	if err != nil {
		fmt.Println("Error NewRequest")
		log.Fatal(err)
	}
	req.Header = headers
	req.SetBasicAuth("my-client", "my-secret")

	client := &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore SSL certificates
	}}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Error executing request")
		log.Fatal(err)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		fmt.Println("Error parsing response")
		log.Fatal(err)
	}

	json.Unmarshal(respBody, &content)

	if str, ok := (content.(map[string]interface{})["access_token"]).(string); ok {
		fmt.Println(str)
		return str
	}

	return "Failed"
	// ...
}

func main() {
	fmt.Println("--------------------------------------------------------------------")
	fmt.Println("Getting JWKs")
	fmt.Println("--------------------------------------------------------------------")
	jwks()
	fmt.Println()
	fmt.Println("--------------------------------------------------------------------")
	fmt.Println("Getting OpenID Configs")
	fmt.Println("--------------------------------------------------------------------")
	openIDConfig()
	fmt.Println()
	fmt.Println("--------------------------------------------------------------------")
	fmt.Println("Checking Auth Endpoint")
	fmt.Println("--------------------------------------------------------------------")
	authEndpoint()
	fmt.Println()
	fmt.Println("--------------------------------------------------------------------")
	fmt.Println("Getting Auth Token")
	fmt.Println("--------------------------------------------------------------------")
	getAccessToken()
	fmt.Println()
	fmt.Println("--------------------------------------------------------------------")
	fmt.Println("Getting Auth Token")
	fmt.Println("--------------------------------------------------------------------")
}
