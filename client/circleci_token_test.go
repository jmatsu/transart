package client

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"gopkg.in/guregu/null.v3"
	"net/http"
	"testing"
)

var testNewCircleCITokenTests = []struct {
	in  null.String
	out string
}{
	{
		null.StringFrom("token"),
		"token",
	},
	{
		null.StringFromPtr(nil),
		"",
	},
}

func TestNewCircleCIToken(t *testing.T) {
	for i, c := range testNewCircleCITokenTests {
		t.Run(fmt.Sprintf("TestNewCircleCIToken %d", i), func(t *testing.T) {
			token := newCircleCIToken(c.in)

			if c.out == "" {
				assert.Nil(t, token)
			} else {
				assert.EqualValues(t, c.out, token.token)
			}
		})
	}
}

var testCircleCIToken_SetToHeaderTests = []struct {
	in  *circleCIToken
	out string
}{
	{
		&circleCIToken{
			token: "XYZ",
		},
		"Basic WFlaOg==",
	},
	{
		nil,
		"",
	},
}

func TestCircleCIToken_SetToHeader(t *testing.T) {
	for i, c := range testCircleCIToken_SetToHeaderTests {
		t.Run(fmt.Sprintf("TestCircleCIToken_SetToHeader %d", i), func(t *testing.T) {
			request := http.Request{
				Header: make(http.Header),
			}

			c.in.SetToHeader(&request)

			assert.EqualValues(t, c.out, request.Header.Get("Authorization"))
		})
	}
}

var testToken_ToParamTests = []struct {
	in  *circleCIToken
	out string
}{
	{
		&circleCIToken{
			token: "XYZ",
		},
		"XYZ",
	},
	{
		nil,
		"",
	},
}

func TestCircleCIToken_ToParam(t *testing.T) {
	for i, c := range testToken_ToParamTests {
		t.Run(fmt.Sprintf("TestCircleCIToken_ToParam %d", i), func(t *testing.T) {
			param := c.in.ToParam()

			if c.out == "" {
				assert.Nil(t, param)
			} else {
				assert.EqualValues(t, c.out, param.Get("circle-token"))
			}
		})
	}
}
