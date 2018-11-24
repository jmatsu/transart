package github

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"gopkg.in/guregu/null.v3"
	"net/http"
	"testing"
)

var testNewTokenTests = []struct {
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

func TestNewToken(t *testing.T) {
	for i, c := range testNewTokenTests {
		t.Run(fmt.Sprintf("TestNewToken %d", i), func(t *testing.T) {
			token := NewToken(c.in)

			if c.out == "" {
				assert.Nil(t, token)
			} else {
				assert.EqualValues(t, c.out, token.token)
			}
		})
	}
}

var testToken_SetToHeaderTests = []struct {
	in  *Token
	out string
}{
	{
		&Token{
			token: "XYZ",
		},
		"token XYZ",
	},
	{
		nil,
		"",
	},
}

func TestToken_SetToHeader(t *testing.T) {
	for i, c := range testToken_SetToHeaderTests {
		t.Run(fmt.Sprintf("TestToken_SetToHeader %d", i), func(t *testing.T) {
			request := http.Request{
				Header: make(http.Header),
			}

			c.in.SetToHeader(&request)

			assert.EqualValues(t, c.out, request.Header.Get("Authorization"))
		})
	}
}

func TestToken_ToParam(t *testing.T) {
	// no-op
}
