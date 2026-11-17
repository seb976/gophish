package imap

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// oauth2TokenResponse models the JSON returned by the Microsoft identity
// platform token endpoint.
type oauth2TokenResponse struct {
	AccessToken      string `json:"access_token"`
	TokenType        string `json:"token_type"`
	ExpiresIn        int    `json:"expires_in"`
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

// getOAuth2Token requests an access token from Azure AD / Microsoft 365 using
// the OAuth2 client credentials flow (app-only). The returned token is suitable
// for IMAP authentication via the XOAUTH2 SASL mechanism.
func getOAuth2Token(tenantID, clientID, clientSecret string) (string, error) {
	tokenURL := fmt.Sprintf("https://login.microsoftonline.com/%s/oauth2/v2.0/token", tenantID)

	form := url.Values{}
	form.Set("client_id", clientID)
	form.Set("client_secret", clientSecret)
	form.Set("grant_type", "client_credentials")
	form.Set("scope", "https://outlook.office365.com/.default")

	req, err := http.NewRequest("POST", tokenURL, strings.NewReader(form.Encode()))
	if err != nil {
		return "", fmt.Errorf("failed to build OAuth2 token request: %s", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	httpClient := &http.Client{Timeout: 30 * time.Second}
	resp, err := httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("OAuth2 token request failed: %s", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read OAuth2 token response: %s", err)
	}

	var tokenResp oauth2TokenResponse
	if err := json.Unmarshal(body, &tokenResp); err != nil {
		return "", fmt.Errorf("failed to parse OAuth2 token response: %s (body: %s)", err, string(body))
	}

	if tokenResp.Error != "" {
		return "", fmt.Errorf("OAuth2 error: %s - %s", tokenResp.Error, tokenResp.ErrorDescription)
	}
	if tokenResp.AccessToken == "" {
		return "", fmt.Errorf("OAuth2 token endpoint returned no access_token (HTTP %d): %s", resp.StatusCode, string(body))
	}

	return tokenResp.AccessToken, nil
}

// xoauth2Client implements the emersion/go-sasl Client interface for the
// XOAUTH2 mechanism, used by Microsoft 365 / Outlook for OAuth2 IMAP auth.
type xoauth2Client struct {
	Username string
	Token    string
}

// Start returns the XOAUTH2 mechanism name and the initial client response.
// The format is: user=<email>^Aauth=Bearer <token>^A^A  (^A = 0x01)
func (a *xoauth2Client) Start() (mech string, ir []byte, err error) {
	resp := fmt.Sprintf("user=%s\x01auth=Bearer %s\x01\x01", a.Username, a.Token)
	return "XOAUTH2", []byte(resp), nil
}

// Next is only called if the server issues a challenge, which for XOAUTH2
// means an authentication error. We surface it rather than continuing.
func (a *xoauth2Client) Next(challenge []byte) (response []byte, err error) {
	return nil, fmt.Errorf("unexpected XOAUTH2 server challenge: %s", string(challenge))
}
