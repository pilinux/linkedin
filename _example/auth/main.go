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

	// create a new session
	session := GlobalApp.Session(token.AccessToken)

	// token introspection
	accessTokenData, err := session.Introspect(session.AccessToken())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("access token introspection:")
	fmt.Println("Active:", accessTokenData.Active)
	fmt.Println("ClientID:", accessTokenData.ClientID)
	fmt.Println("AuthorizedAt:", accessTokenData.AuthorizedAt)
	fmt.Println("CreatedAt:", accessTokenData.CreatedAt)
	fmt.Println("Status:", accessTokenData.Status)
	fmt.Println("ExpiresAt:", accessTokenData.ExpiresAt)
	fmt.Println("Scope:", accessTokenData.Scope)
	fmt.Println("AuthType:", accessTokenData.AuthType)

	refreshTokenData, err := session.Introspect(token.RefreshToken)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("refresh token introspection:")
	fmt.Println("Active:", refreshTokenData.Active)
	fmt.Println("ClientID:", refreshTokenData.ClientID)
	fmt.Println("AuthorizedAt:", refreshTokenData.AuthorizedAt)
	fmt.Println("CreatedAt:", refreshTokenData.CreatedAt)
	fmt.Println("Status:", refreshTokenData.Status)
	fmt.Println("ExpiresAt:", refreshTokenData.ExpiresAt)
	fmt.Println("Scope:", refreshTokenData.Scope)
	fmt.Println("AuthType:", refreshTokenData.AuthType)

	// refresh tokens
	newTokens, err := session.RefreshToken(token.RefreshToken)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("New Access Token:", newTokens.AccessToken)
	fmt.Println("New Refresh Token:", newTokens.RefreshToken)
}
