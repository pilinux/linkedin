// Package main - LinkedIn OAuth authorization code flow example.
// Redeem authorization code for access and refresh tokens.
package main

import (
	"fmt"
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

// ExchangeCode - Redeem authorization code for access and refresh tokens
func ExchangeCode(app *linkedin.App, code string) (linkedin.Token, error) {
	token, err := app.ParseCode(code)
	if err != nil {
		return linkedin.Token{}, err
	}

	return token, nil
}

func main() {
	// initialize the application
	err := Init()
	if err != nil {
		fmt.Println(err)
		return
	}

	// get authorization code
	// GET https://www.linkedin.com/oauth/v2/authorization?response_type=code&client_id={client_id}&redirect_uri={redirect_uri_with_https}&state={foobar}&scope={scopes}

	// redeem authorization code for access and refresh tokens
	code := strings.TrimSpace(os.Getenv("LINKEDIN_AUTH_CODE"))
	if code == "" {
		fmt.Println("LinkedIn: authorization code is empty")
		return
	}

	token, err := ExchangeCode(GlobalApp, code)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Access Token:", token.AccessToken)
	fmt.Println("Refresh Token:", token.RefreshToken)
}
