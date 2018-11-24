package lib

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

//type AuthType int
//
//const (
//	HeaderAuth AuthType = iota
//	ParameterAuth
//)
//
//type Endpoint struct {
//	Url      string
//	AuthType AuthType
//	Accept   string
//}
//
//type Token interface {
//	SetToHeader(request *http.Request)
//	ToParam() url.Values
//}

func launchServer(onRequest func(w http.ResponseWriter, r *http.Request)) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(onRequest))
}

var testMergeParamsTests = []struct {
	lhs url.Values
	rhs url.Values
	out url.Values
}{
	{
		url.Values{
			"x": []string{
				"x",
			},
		},
		url.Values{
			"y": []string{
				"y",
			},
			"z": []string{},
		},
		url.Values{
			"x": []string{
				"x",
			},
			"y": []string{
				"y",
			},
			"z": []string{},
		},
	},
	{
		url.Values{
			"x": []string{
				"x",
			},
		},
		nil,
		url.Values{
			"x": []string{
				"x",
			},
		},
	},
	{
		nil,
		url.Values{
			"x": []string{
				"x",
			},
		},
		url.Values{
			"x": []string{
				"x",
			},
		},
	},
}

func TestMergeParams(t *testing.T) {
	for i, c := range testMergeParamsTests {
		t.Run(fmt.Sprintf("TestMergeParams %d", i), func(t *testing.T) {
			params := MergeParams(c.lhs, c.rhs)

			if params == nil {
				t.Errorf("params is nil")
			}

			for k, v := range c.lhs {
				assert.EqualValues(t, v, params[k], "%s => %v | %v", k, v, params[k])
			}

			for k, v := range c.rhs {
				assert.EqualValues(t, v, params[k], "%s => %v | %v", k, v, params[k])
			}
		})
	}
}

type token struct {
	value string
}

func (t token) SetToHeader(request *http.Request) {
	if t.value != "" {
		request.Header.Set("Authorization", t.value)
	}
}

func (t token) ToParam() url.Values {
	return url.Values{
		"auth": []string{
			t.value,
		},
	}
}

var testGetRequestTests = []struct {
	endpoint Endpoint
	token    token
	values   url.Values
}{
	{
		Endpoint{
			Accept:   "application/json",
			AuthType: ParameterAuth,
		},
		token{
			"this is a token",
		},
		nil,
	},
	{
		Endpoint{
			Accept:   "plain/text",
			AuthType: HeaderAuth,
		},
		token{
			"this is also a token",
		},
		url.Values{
			"key": []string{
				"value",
			},
		},
	},
}

func TestGetRequest(t *testing.T) {
	for i, c := range testGetRequestTests {
		t.Run(fmt.Sprintf("TestGetRequest %d", i), func(t *testing.T) {
			server := launchServer(func(w http.ResponseWriter, r *http.Request) {
				assert.EqualValues(t, http.MethodGet, r.Method)

				assert.EqualValues(t, c.endpoint.Accept, r.Header.Get("Accept"))

				query := r.URL.Query()

				if c.values != nil {
					for k := range query {
						if k == "auth" {
							continue
						}

						assert.EqualValues(t, c.values.Get(k), query.Get(k))
					}
				}

				switch c.endpoint.AuthType {
				case HeaderAuth:
					assert.EqualValues(t, c.token.value, r.Header.Get("Authorization"))
				case ParameterAuth:
					assert.EqualValues(t, c.token.value, query.Get("auth"))
				default:
					t.Fatal(fmt.Errorf("unknown auth type is found : %v", c.endpoint.AuthType))
				}

				w.WriteHeader(http.StatusOK)
				w.Write([]byte("{}"))
			})

			defer server.Close()

			c.endpoint.Url = server.URL

			if _, err := GetRequest(c.endpoint, c.token, c.values); err != nil {
				t.Error(err)
			}
		})
	}
}

var testPostRequestTests = []struct {
	endpoint Endpoint
	token    token
	values   url.Values
}{
	{
		Endpoint{
			Accept:   "application/json",
			AuthType: ParameterAuth,
		},
		token{
			"this is a token",
		},
		nil,
	},
	{
		Endpoint{
			Accept:   "plain/text",
			AuthType: HeaderAuth,
		},
		token{
			"this is also a token",
		},
		url.Values{
			"key": []string{
				"value",
			},
		},
	},
}

func TestPostRequest(t *testing.T) {
	for i, c := range testPostRequestTests {
		t.Run(fmt.Sprintf("TestPostRequest %d", i), func(t *testing.T) {
			server := launchServer(func(w http.ResponseWriter, r *http.Request) {
				assert.EqualValues(t, http.MethodPost, r.Method)

				assert.EqualValues(t, c.endpoint.Accept, r.Header.Get("Accept"))

				query := r.URL.Query()

				if c.values != nil {
					for k := range query {
						if k == "auth" {
							continue
						}

						assert.EqualValues(t, c.values.Get(k), query.Get(k))
					}
				}

				switch c.endpoint.AuthType {
				case HeaderAuth:
					assert.EqualValues(t, c.token.value, r.Header.Get("Authorization"))
				case ParameterAuth:
					assert.EqualValues(t, c.token.value, query.Get("auth"))
				default:
					t.Fatal(fmt.Errorf("unknown auth type is found : %v", c.endpoint.AuthType))
				}

				w.WriteHeader(http.StatusOK)
				w.Write([]byte("{}"))
			})

			defer server.Close()

			c.endpoint.Url = server.URL

			if _, err := PostRequest(c.endpoint, c.token, c.values, []byte{}); err != nil {
				t.Error(err)
			}
		})
	}
}
