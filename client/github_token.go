package client

import (
	"fmt"
	"github.com/jmatsu/transart/lib"
	"gopkg.in/guregu/null.v3"
	"net/http"
	"net/url"
)

type gitHubToken struct {
	token string
}

func newGitHubToken(t null.String) lib.Token {
	var token *gitHubToken

	if t.Valid {
		token = &gitHubToken{
			token: t.String,
		}
	}

	return token
}

func (t *gitHubToken) SetToHeader(request *http.Request) {
	if t != nil {
		request.Header.Set("Authorization", fmt.Sprintf("token %s", t.token))
	}
}

func (t *gitHubToken) ToParam() url.Values {
	if t == nil {
		return nil
	}

	return url.Values(
		map[string][]string{
			"access_token": {t.token},
		})
}
