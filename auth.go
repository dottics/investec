package investec

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// Credentials is the required information to be able to authenticate the user
// trying to access that data.
//
// NB: it is recommended that the user only allow "read" permissions to the
// account and only when really necessary allow "write" permissions (or to
// allow the processing payments).
type Credentials struct {
	ClientID string `json:"client_id"`
	Secret   string `json:"secret"`
	APIKey   string `json:"api_key"`
}

// encodeBasic to encode the client id and client secret into a base64 encoded
// string.
func encodeBasic(id, secret string) string {
	data := []byte(fmt.Sprintf("%s:%s", id, secret))
	dst := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(dst, data)
	return string(dst)
}

// Auth authenticates a user based on their authentication credentials.
func (s *Service) Auth(credentials *Credentials) (string, error) {
	// set the auth path
	s.URL.Path = "identity/v2/oauth2/token"
	basic := encodeBasic(credentials.ClientID, credentials.Secret)
	// create the request form payload
	payload := strings.NewReader("grant_type=client_credentials")

	req, err := http.NewRequest("POST", s.URL.String(), payload)
	if err != nil {
		return "", err
	}
	// Headers are case-insensitive
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("authorization", fmt.Sprintf("Basic %s", basic))
	req.Header.Add("x-api-key", credentials.APIKey)

	res, err := s.DoRequest(req)
	if err != nil {
		return "", err
	}
	if res.StatusCode != 200 {
		log.Printf("Investec Error: %s", res.Status)
		return "", fmt.Errorf("investec: authentication failed status: %s", res.Status)
	}

	body := struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		ExpiresIn   int    `json:"expires_in"`
		Scope       string `json:"scope"`
	}{}
	err = s.MarshalResponseJSON(res, &body)
	if err != nil {
		fmt.Println("error marshalling response: to print json body enable env DEBUG=1")
		return "", err
	}

	s.Token = body.AccessToken

	return body.AccessToken, nil
}
