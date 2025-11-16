package linkedin

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// Session holds a LinkedIn session with an access token.
// Session should be created by App.Session.
type Session struct {
	HTTPClient             HTTPClient      // HTTP client to send requests
	BaseURL                string          // set to override API base URL
	accessToken            string          // linkedIn access token, can be empty
	app                    *App            // linkedIn app
	LinkedInVersion        string          // e.g. 202404
	useAuthorizationHeader bool            // pass accessToken in headers
	context                context.Context // session context
}

// HTTPClient is an interface to send http request.
// It is compatible with type `*http.Client`.
type HTTPClient interface {
	Do(req *http.Request) (resp *http.Response, err error)
	Get(url string) (resp *http.Response, err error)
	Post(url string, bodyType string, body io.Reader) (resp *http.Response, err error)
}

// Params to construct the request payload.
type Params map[string]interface{}

// default LinkedIn session
var (
	defaultSession = &Session{}
)

// App returns associated App.
func (session *Session) App() *App {
	return session.app
}

// AccessToken returns current access token.
func (session *Session) AccessToken() string {
	return session.accessToken
}

// SetAccessToken sets a new access token.
func (session *Session) SetAccessToken(token string) {
	if token != session.accessToken {
		session.accessToken = token
	}
}

// UseAuthorizationHeader passes `access_token` in HTTP Authorization header.
func (session *Session) UseAuthorizationHeader() {
	session.useAuthorizationHeader = true
}

// Get sends a GET request to LinkedIn API and returns the response.
func (session *Session) Get(uri string) (response *http.Response, data []byte, err error) {
	// uri must start with `/`
	if !strings.HasPrefix(uri, "/") {
		uri = "/" + uri
	}
	url := session.BaseURL + uri

	// create a new HTTP request
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		err = fmt.Errorf("linkedIn: cannot create new request; %w", err)
		return nil, nil, err
	}

	// set headers
	request.Header.Set(string(ContentType), string(JSON))
	request.Header.Set(string(RestLiProtocolVersion), "2.0.0")
	request.Header.Set(string(LinkedInVersion), session.LinkedInVersion)

	// send the request
	response, data, err = session.sendRequest(request)
	return
}

// sendAuthRequest sends an auth request to LinkedIn and returns new tokens.
func (session *Session) sendAuthRequest(uri string, params Params) (Token, error) {
	if params == nil {
		return Token{}, fmt.Errorf("linkedIn: required params are missing")
	}

	if params["grant_type"] == nil {
		return Token{}, fmt.Errorf("linkedIn: grant_type is missing")
	}
	grantType := params["grant_type"].(string)
	if grantType == "" {
		return Token{}, fmt.Errorf("linkedIn: grant_type is required to receive new tokens")
	}

	if params["client_id"] == nil {
		return Token{}, fmt.Errorf("linkedIn: client_id is missing")
	}
	clientID := params["client_id"].(string)
	if clientID == "" {
		return Token{}, fmt.Errorf("linkedIn: client_id is required to receive new tokens")
	}

	if params["client_secret"] == nil {
		return Token{}, fmt.Errorf("linkedIn: client_secret is missing")
	}
	clientSecret := params["client_secret"].(string)
	if clientSecret == "" {
		return Token{}, fmt.Errorf("linkedIn: client_secret is required to receive new tokens")
	}

	redirectURI := ""
	code := ""
	if grantType == "authorization_code" {
		if params["redirect_uri"] == nil {
			return Token{}, fmt.Errorf("linkedIn: redirect_uri is missing")
		}
		redirectURI = params["redirect_uri"].(string)
		if redirectURI == "" {
			return Token{}, fmt.Errorf("linkedIn: redirect_uri is required to redeem auth code")
		}

		if params["code"] == nil {
			return Token{}, fmt.Errorf("linkedIn: auth code is missing")
		}
		code = params["code"].(string)
		if code == "" {
			return Token{}, fmt.Errorf("linkedIn: auth code is required to receive new tokens")
		}
	}

	refToken := ""
	if grantType == "refresh_token" {
		if params["refresh_token"] == nil {
			return Token{}, fmt.Errorf("linkedIn: refresh_token is missing")
		}
		refToken = params["refresh_token"].(string)
		if refToken == "" {
			return Token{}, fmt.Errorf("linkedIn: refresh_token is required to refresh tokens")
		}
	}

	oauthURL := OauthBaseURL + uri

	// data to be sent in the body (x-www-form-urlencoded)
	data := url.Values{}
	data.Set("grant_type", grantType)
	data.Add("client_id", clientID)
	data.Add("client_secret", clientSecret)
	if grantType == "authorization_code" {
		data.Add("redirect_uri", redirectURI)
		data.Add("code", code)
	}
	if grantType == "refresh_token" {
		data.Add("refresh_token", refToken)
	}

	// encode data into appropriate format
	requestBody := bytes.NewBufferString(data.Encode())

	// create a new HTTP request
	request, err := http.NewRequest("POST", oauthURL, requestBody)
	if err != nil {
		return Token{}, err
	}

	// set headers
	request.Header.Set(string(ContentType), string(URLEncoded))

	// send the request
	response, responseData, err := session.sendRequest(request)
	if err != nil {
		return Token{}, err
	}
	if response.StatusCode != http.StatusOK {
		return Token{}, fmt.Errorf("linkedIn: failed to receive tokens with status %d", response.StatusCode)
	}

	// parse the response body
	var token Token
	err = json.Unmarshal(responseData, &token)
	if err != nil {
		return Token{}, err
	}

	return token, nil
}

// Introspect checks the Time to Live (TTL) and status (active/expired) for the given token.
//
// See: https://learn.microsoft.com/en-us/linkedin/shared/authentication/token-introspection?toc=%2Flinkedin%2Fmarketing%2Ftoc.json&bc=%2Flinkedin%2Fbreadcrumb%2Ftoc.json&view=li-lms-2024-04&tabs=http
func (session *Session) Introspect(token string) (TokenData, error) {
	token = strings.TrimSpace(token)
	if token == "" {
		err := fmt.Errorf("linkedIn: token is empty")
		return TokenData{}, err
	}

	tokenData, err := session.introspect("/introspectToken", Params{
		"token": token,
	})

	return tokenData, err
}

// token introspection
func (session *Session) introspect(uri string, params Params) (TokenData, error) {
	if params == nil {
		return TokenData{}, fmt.Errorf("linkedIn: required params are missing")
	}

	if params["token"] == nil {
		return TokenData{}, fmt.Errorf("linkedIn: token is missing")
	}
	token := params["token"].(string)
	if token == "" {
		return TokenData{}, fmt.Errorf("linkedIn: token is required for introspection")
	}

	oauthURL := OauthBaseURL + uri

	// data to be sent in the body (x-www-form-urlencoded)
	data := url.Values{}
	data.Set("client_id", session.App().ClientID)
	data.Add("client_secret", session.App().ClientSecret)
	data.Add("token", token)

	// encode data into appropriate format
	requestBody := bytes.NewBufferString(data.Encode())

	// create a new HTTP request
	request, err := http.NewRequest("POST", oauthURL, requestBody)
	if err != nil {
		return TokenData{}, err
	}

	// set headers
	request.Header.Set(string(ContentType), string(URLEncoded))

	// send the request
	response, responseData, err := session.sendRequest(request)
	if err != nil {
		return TokenData{}, err
	}
	if response.StatusCode != http.StatusOK {
		return TokenData{}, fmt.Errorf("linkedIn: failed to introspect token with status %d", response.StatusCode)
	}

	// parse the response body
	var tokenData TokenData
	err = json.Unmarshal(responseData, &tokenData)
	if err != nil {
		return TokenData{}, err
	}

	return tokenData, nil
}

// sendRequest sends an API request and returns the response.
func (session *Session) sendRequest(request *http.Request) (response *http.Response, data []byte, err error) {
	if session.context != nil {
		request = request.WithContext(session.context)
	}

	if session.useAuthorizationHeader {
		request.Header.Set(string(Authorization), "Bearer "+session.accessToken)
	}

	if session.HTTPClient == nil {
		response, err = http.DefaultClient.Do(request)
	} else {
		response, err = session.HTTPClient.Do(request)
	}
	if err != nil {
		err = fmt.Errorf("linkedIn: cannot reach linkedIn server; %w", err)
		return
	}

	buf := &bytes.Buffer{}
	_, err = io.Copy(buf, response.Body)
	closeErr := response.Body.Close()
	if err != nil {
		err = fmt.Errorf("linkedIn: cannot read linkedIn response; %w", err)
	} else if closeErr != nil {
		err = fmt.Errorf("linkedIn: error closing response body; %w", closeErr)
	}

	data = buf.Bytes()
	return
}

// Context returns the session's context.
// To change the context, use `Session#WithContext`.
//
// The returned context is always non-nil; it defaults to the background context.
// For outgoing LinkedIn API requests, the context controls timeout/deadline and cancellation.
func (session *Session) Context() context.Context {
	if session.context != nil {
		return session.context
	}

	return context.Background()
}

// WithContext returns a shallow copy of session with its context changed to ctx.
// The provided ctx must be non-nil.
func (session *Session) WithContext(ctx context.Context) *Session {
	s := *session
	s.context = ctx
	return &s
}
