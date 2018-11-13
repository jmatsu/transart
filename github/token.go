package github

import (
	"fmt"
	"gopkg.in/guregu/null.v3"
	"net/http"
	"net/url"
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
		request.Header.Set("Authorization", fmt.Sprintf("token %s", t.token))
	}
}

func (t *Token) ToParam() url.Values {
	panic("not implemented")
}
