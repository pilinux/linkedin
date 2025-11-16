package linkedin

import (
	"fmt"
	"strings"
)

// App holds the configuration for the LinkedIn application.
//
// ClientID and ClientSecret:
// https://www.linkedin.com/developers/apps/{appID}/auth
type App struct {
	ClientID     string
	ClientSecret string
	RedirectURI  string
	session      *Session
}

// New creates a new LinkedIn application and sets clientID and clientSecret.
func New(clientID, clientSecret string) *App {
	return &App{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		session:      defaultSession,
	}
}

// ParseCode redeems authorization code for access and refresh tokens.
//
// See https://learn.microsoft.com/en-us/linkedin/shared/authentication/authorization-code-flow?context=linkedin%2Fcontext&tabs=HTTPS1
func (app *App) ParseCode(code string) (Token, error) {
	code = strings.TrimSpace(code)
	if code == "" {
		err := fmt.Errorf("linkedIn: authorization code is empty")
		return Token{}, err
	}

	token, err := app.session.sendAuthRequest("/accessToken", Params{
		"grant_type":    "authorization_code",
		"client_id":     app.ClientID,
		"client_secret": app.ClientSecret,
		"redirect_uri":  app.RedirectURI,
		"code":          code,
	})

	return token, err
}

// RefreshToken redeems refresh token for new access and refresh tokens.
//
// See: https://learn.microsoft.com/en-us/linkedin/shared/authentication/programmatic-refresh-tokens?toc=%2Flinkedin%2Fmarketing%2Ftoc.json&bc=%2Flinkedin%2Fbreadcrumb%2Ftoc.json&view=li-lms-2024-04
func (app *App) RefreshToken(refreshToken string) (Token, error) {
	refreshToken = strings.TrimSpace(refreshToken)
	if refreshToken == "" {
		err := fmt.Errorf("linkedIn: refresh token is empty")
		return Token{}, err
	}

	token, err := app.session.sendAuthRequest("/accessToken", Params{
		"grant_type":    "refresh_token",
		"client_id":     app.ClientID,
		"client_secret": app.ClientSecret,
		"refresh_token": refreshToken,
	})

	return token, err
}

// Session creates a new LinkedIn session based on the app configuration.
func (app *App) Session(accessToken string) *Session {
	return &Session{
		BaseURL:         VersionedBaseURL,
		accessToken:     accessToken,
		app:             app,
		LinkedInVersion: "202510",
	}
}

// SetLinkedInVersion overrides the default LinkedIn API version.
func (session *Session) SetLinkedInVersion(version string) {
	session.LinkedInVersion = version
}
