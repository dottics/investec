package investec

import (
	"encoding/base64"
	"fmt"
)

// encodeBasic to encode the client id and client secret into a base64 encoded
// string.
func encodeBasic(id, secret string) string {
	data := []byte(fmt.Sprintf("%s:%s", id, secret))
	dst := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(dst, data)
	return string(dst)
}

// Auth authenticates a user based on their authentication credentials. The
// credentials are read from the environmental variables.
func Auth() {
	//id := os.Getenv("INVID")
	//sec := os.Getenv("INVSECRET")
	//key := os.Getenv("INVKEY")
}
