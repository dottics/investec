package investec

import (
	"os"
	"testing"
)

func TestNewService(t *testing.T) {
	_ = os.Unsetenv("INVESTEC_SCHEME")
	_ = os.Unsetenv("INVESTEC_HOST")
	/*
		Given we want to be able to mock and simplify the setup of each instance
		of the integration we want to be able to create a quick mock.
	*/
	s := NewService("investec")
	if s.URL.Scheme != "" {
		t.Errorf("expected '' got %s", s.URL.Scheme)
	}
	if s.URL.Host != "" {
		t.Errorf("expected '' got %s", s.URL.Host)
	}

	/*
		Given a new service is to be created
		When the instance is created it should read then scheme and host
			variables from the environmental variables
		Then should set them as the service URL parameters
	*/
	_ = os.Setenv("INVESTEC_SCHEME", "https")
	_ = os.Setenv("INVESTEC_HOST", "openapi.investec.com")
	s = NewService("investec")
	if s.URL.Scheme != "https" {
		t.Errorf("expected 'https' got %s", s.URL.Scheme)
	}
	if s.URL.Host != "openapi.investec.com" {
		t.Errorf("expected 'openapi.investec.com' got %s", s.URL.Host)
	}
}
