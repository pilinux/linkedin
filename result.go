package linkedin

// Result - response body from LinkedIn API call request.
type Result map[string]interface{}

// Token - access and refresh tokens.
//
// See: https://learn.microsoft.com/en-us/linkedin/shared/authentication/authorization-code-flow?toc=%2Flinkedin%2Fmarketing%2Ftoc.json&bc=%2Flinkedin%2Fbreadcrumb%2Ftoc.json&view=li-lms-2024-04&tabs=HTTPS1#step-3-exchange-authorization-code-for-an-access-token
type Token struct {
	AccessToken           string `json:"access_token"`
	ExpiresIn             int64  `json:"expires_in"`
	RefreshToken          string `json:"refresh_token"`
	RefreshTokenExpiresIn int64  `json:"refresh_token_expires_in"`
	Scope                 string `json:"scope"`
}

// TokenData - access and refresh token data after token inspection.
//
// See: https://learn.microsoft.com/en-us/linkedin/shared/authentication/token-introspection?context=linkedin/context
type TokenData struct {
	Active       bool   `json:"active"`
	ClientID     string `json:"client_id"`
	AuthorizedAt int64  `json:"authorized_at"`
	CreatedAt    int64  `json:"created_at"`
	Status       string `json:"status"`
	ExpiresAt    int64  `json:"expires_at"`
	Scope        string `json:"scope"`
	AuthType     string `json:"auth_type"`
}
