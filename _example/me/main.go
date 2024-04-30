// Package main - LinkedIn HTTP GET example.
// Get user profile.
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/pilinux/linkedin"
)

// GlobalApp - LinkedIn application
var GlobalApp *linkedin.App

// Init - Initialize the application
func Init() error {
	clientID := strings.TrimSpace(os.Getenv("LINKEDIN_CLIENT_ID"))
	clientSecret := strings.TrimSpace(os.Getenv("LINKEDIN_CLIENT_SECRET"))
	redirectURI := strings.TrimSpace(os.Getenv("LINKEDIN_REDIRECT_URI"))

	if clientID == "" || clientSecret == "" || redirectURI == "" {
		return fmt.Errorf("LinkedIn: clientID, clientSecret, or redirectURI is empty")
	}

	GlobalApp = linkedin.New(clientID, clientSecret)
	GlobalApp.RedirectURI = redirectURI

	return nil
}

func main() {
	// initialize the application
	err := Init()
	if err != nil {
		fmt.Println(err)
		return
	}

	token := linkedin.Token{}
	token.AccessToken = strings.TrimSpace(os.Getenv("LINKEDIN_ACCESS_TOKEN"))

	// create a new LinkedIn session
	session := GlobalApp.Session(token.AccessToken)

	// set Authorization header
	session.UseAuthorizationHeader()

	// get user profile
	response, data, err := session.Get("/me")
	if err != nil {
		fmt.Println(err)
		return
	}

	if response.StatusCode != http.StatusOK {
		fmt.Println("LinkedIn: unexpected status code", response.StatusCode)
		return
	}

	// parse the response body
	var result linkedin.Result
	err = json.Unmarshal(data, &result)
	if err != nil {
		fmt.Println(err)
		return
	}

	// accessing keys and values
	for key, value := range result {
		fmt.Printf("Key: %s, Value: %v\n", key, value)
	}
}
