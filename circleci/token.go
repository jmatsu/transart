package circleci

import (
	"encoding/base64"
	"fmt"
	"gopkg.in/guregu/null.v3"
	"net/http"
	"net/url"
	"unsafe"
)

type Token struct {
	token string
}

func NewToken(t null.String) *Token {
	if !t.Valid {
		return nil
	}

	return &Token{
		token: t.String,
	}
}

func (t *Token) SetToHeader(request *http.Request) {
	if t != nil {
		bytes := *(*[]byte)(unsafe.Pointer(&t.token))
		token := base64.StdEncoding.EncodeToString(bytes)
		request.Header.Set("Authorization", fmt.Sprintf("Basic %s", token))
	}
}

func (t *Token) ToParam() url.Values {
	if t == nil {
		return nil
	}

	return url.Values(
		map[string][]string{
			"circle-token": {t.token},
		})
}
