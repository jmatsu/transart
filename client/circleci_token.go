package client

import (
	"encoding/base64"
	"fmt"
	"gopkg.in/guregu/null.v3"
	"net/http"
	"net/url"
)

type circleCIToken struct {
	token string
}

func newCircleCIToken(t null.String) *circleCIToken {
	if !t.Valid {
		return nil
	}

	return &circleCIToken{
		token: t.String,
	}
}

func (t *circleCIToken) SetToHeader(request *http.Request) {
	if t != nil {
		token := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:", t.token)))
		request.Header.Set("Authorization", fmt.Sprintf("Basic %s", token))
	}
}

func (t *circleCIToken) ToParam() url.Values {
	if t == nil {
		return nil
	}

	return url.Values(
		map[string][]string{
			"circle-token": {t.token},
		})
}
