package linkedin

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
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

// sendAuthRequest sends an auth request to LinkedIn and returns new tokens.
func (session *Session) sendAuthRequest(uri string, params Params) (Token, error) {
	grantType := params["grant_type"].(string)
	if grantType == "" {
		return Token{}, fmt.Errorf("linkedIn: grant_type is required to receive new tokens")
	}
	clientID := params["client_id"].(string)
	if clientID == "" {
		return Token{}, fmt.Errorf("linkedIn: client_id is required to receive new tokens")
	}
	clientSecret := params["client_secret"].(string)
	if clientSecret == "" {
		return Token{}, fmt.Errorf("linkedIn: client_secret is required to receive new tokens")
	}
	redirectURI := params["redirect_uri"].(string)
	if redirectURI == "" && grantType == "authorization_code" {
		return Token{}, fmt.Errorf("linkedIn: redirect_uri is required to redeem auth code")
	}
	code := params["code"].(string)
	if code == "" && grantType == "authorization_code" {
		return Token{}, fmt.Errorf("linkedIn: auth code is required to receive new tokens")
	}
	refToken := params["refresh_token"].(string)
	if refToken == "" && grantType == "refresh_token" {
		return Token{}, fmt.Errorf("linkedIn: refresh_token is required to refresh tokens")
	}

	oauthURL := session.BaseURL + uri

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
	response.Body.Close()
	if err != nil {
		err = fmt.Errorf("linkedIn: cannot read linkedIn response; %w", err)
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
