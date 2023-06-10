package investec

import (
	"github.com/johannesscr/micro/microtest"
	"testing"
)

func TestEncodeBasic(t *testing.T) {
	id := "client id"
	secret := "client secret"
	out := "Y2xpZW50IGlkOmNsaWVudCBzZWNyZXQ="

	o := encodeBasic(id, secret)
	if o != out {
		t.Errorf("expected %s got %s", out, o)
	}
}

func TestService_Auth(t *testing.T) {
	s := NewService("investec")
	ms := microtest.MockServer(s)
	defer ms.Server.Close()

	accessToken := "access-token"

	exchange := &microtest.Exchange{
		Response: microtest.Response{
			Status: 200,
			Body:   `{"access_token":"access-token","token_type":"Bearer","expires_in":1799,"scope":"accounts"}`,
		},
	}
	ms.Append(exchange)
	credentials := &Credentials{}

	to, _ := s.Auth(credentials)
	if to != accessToken {
		t.Errorf("expected access token %s got %s", accessToken, to)
	}
	// also check that the auth token has been stored on the service
	if s.Token != accessToken {
		t.Errorf("expected access token %s got %s", accessToken, s.Token)
	}
}
