package token

import (
	"math/rand"
	"strings"
	"time"
)

// MakeBearer https://tools.ietf.org/html/rfc6750#section-2.1
// b64token    = 1*( ALPHA / DIGIT /"-" / "." / "_" / "~" / "+" / "/" ) *"="
func MakeBearer(length int) string {
	charlist := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	t := make([]string, length)
	rand.Seed(time.Now().UnixNano() + int64(length) + rand.Int63n(10000))
	for i := 0; i < length; i++ {
		rn := rand.Intn(len(charlist))
		w := charlist[rn : rn+1]
		t = append(t, w)
	}

	token := strings.Join(t, "")
	return token
}
