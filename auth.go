package investec

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
)

// encodeBasic to encode the client id and client secret into a base64 encoded
// string.
func encodeBasic(id, secret string) string {
	data := []byte(fmt.Sprintf("%s:%s", id, secret))
	dst := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(dst, data)
	return string(dst)
}

// Auth authenticates a user based on their authentication credentials.
func (s *Service) Auth(id, secret, key string) (string, error) {
	basic := encodeBasic(id, secret)
	req, err := http.NewRequest("POST", s.URL.String(), nil)
	if err != nil {
		return "", err
	}
	// Headers are case-insensitive
	req.Header.Add("accept", "application/json")
	req.Header.Add("authentication", fmt.Sprintf("Basic %s", basic))
	req.Header.Add("x-api-key", key)

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
