package examples

import (
	"encoding/json"
	"fmt"
	"github.com/siimtalts/v3-go-sdk"
	"io/ioutil"
	"log"

	"golang.org/x/oauth2"
)

func GetExampleScoroApiClient() scoro.APIClient {
	return scoro.GetClient(scoro.GetAPIClientConfigFromEnvFile(), exampleClient{})
}

type exampleClient struct{}

func (exampleClient) HandleAuthorization(config oauth2.Config) string {
	url := config.AuthCodeURL("state")
	fmt.Printf("Visit the URL for the auth dialog: %v\nEnter the code: ", url)

	// Use the authorization code that is pushed to the redirect
	// URL. Exchange will do the handshake to retrieve the
	// initial access token. The HTTP APIClient returned by
	// conf.APIClient will refresh the token as necessary.
	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("CODE: %v", code)

	return code
}

func (exampleClient) SaveTokens(token *oauth2.Token) {
	file, _ := json.MarshalIndent(token, "", " ")

	err := ioutil.WriteFile("test.json", file, 0644)

	if err != nil {
		log.Fatal(err)
	}
}

func (exampleClient) FetchTokens() *oauth2.Token {
	file, _ := ioutil.ReadFile("test.json")

	tok := &oauth2.Token{}
	_ = json.Unmarshal([]byte(file), &tok)

	return tok
}
