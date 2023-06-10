package investec

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Service struct {
	// Token is the Bearer token used to authenticate requests to the Investec
	// API.
	Token string
	URL   url.URL
}

func NewService(envName string) *Service {
	scheme := os.Getenv(fmt.Sprintf("%s_SCHEME", strings.ToUpper(envName)))
	host := os.Getenv(fmt.Sprintf("%s_HOST", strings.ToUpper(envName)))
	return &Service{
		URL: url.URL{
			Scheme: scheme,
			Host:   host,
		},
	}
}

// SetURL is what makes this service a mock-able service using microtest
func (s *Service) SetURL(sc string, h string) {
	s.URL.Scheme = sc
	s.URL.Host = h
}

func (s *Service) DoRequest(r *http.Request) (*http.Response, error) {
	client := http.Client{}
	res, err := client.Do(r)
	return res, err
}

// prettyJSONOut is a helper function to pretty print JSON data as a slice of bytes.
func prettyJSONOut(xb []byte) {
	var out bytes.Buffer
	err := json.Indent(&out, append(xb, "\n"...), "", "  ")
	if err != nil {
		fmt.Println("JSON:", err)
	}
	_, err = out.WriteTo(os.Stdout)
	if err != nil {
		fmt.Println("JSON:", err)
	}
}

// equalError returns true if two errors are equal.
func equalError(a, b error) bool {
	if a == nil && b == nil {
		return true
	}
	if a != nil && b != nil {
		if a.Error() == b.Error() {
			return true
		}
	}
	return false
}

// MarshalResponseJSON is a wrapper function to remove some boilerplate code
// when parsing the JSON body response from investec.
func (s *Service) MarshalResponseJSON(res *http.Response, v interface{}) error {
	debug := os.Getenv("DEBUG")
	defer func(rc io.ReadCloser) {
		err := rc.Close()
		if err != nil {
			log.Println(err.Error())
		}
	}(res.Body)
	xb, err := io.ReadAll(res.Body)
	if debug != "" {
		prettyJSONOut(xb)
	}
	if err != nil {
		return err
	}
	err = json.Unmarshal(xb, v)
	if err != nil {
		return err
	}
	return nil
}
