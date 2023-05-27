package investec

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
)

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
	req, err := http.NewRequest("POST", s.URL.String(), nil)
	if err != nil {
		return "", err
	}
	// Headers are case-insensitive
	req.Header.Add("accept", "application/json")
	req.Header.Add("authentication", fmt.Sprintf("Basic %s", basic))
	req.Header.Add("x-api-key", credentials.APIKey)

	res, err := s.DoRequest(req)
	if err != nil {
		return "", err
	}
	body := struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		ExpiresIn   int    `json:"expires_in"`
		Scope       string `json:"scope"`
	}{}
	err = s.MarshalResponseJSON(res, &body)
	if err != nil {
		return "", err
	}
	if res.StatusCode != 200 {
		log.Printf("Investec Error: %d %s", res.StatusCode, res.Status)
		return "", nil
	}
	return body.AccessToken, nil
}
