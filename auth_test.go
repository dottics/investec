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
			Body:   `{"accessToken":"red"}`,
		},
	}
	ms.Append(exchange)

	to, _ := s.Auth("", "", "")
	if to != accessToken {
		t.Errorf("expected access token %s got %s", accessToken, to)
	}
}
