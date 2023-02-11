package investec

import "testing"

func TestEncodeBasic(t *testing.T) {
	id := "client id"
	secret := "client secret"
	out := "Y2xpZW50IGlkOmNsaWVudCBzZWNyZXQ="

	o := encodeBasic(id, secret)
	if o != out {
		t.Errorf("expected %s got %s", out, o)
	}
}
