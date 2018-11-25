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
	if !t.Valid {
		return nil
	}

	return &gitHubToken{
		token: t.String,
	}
}

func (t *gitHubToken) SetToHeader(request *http.Request) {
	if t != nil {
		request.Header.Set("Authorization", fmt.Sprintf("token %s", t.token))
	}
}

func (t *gitHubToken) ToParam() url.Values {
	panic("not implemented")
}
