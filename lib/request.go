package lib

import (
	"bytes"
	"fmt"
	"github.com/jmatsu/transart/version"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
)

type AuthType int

const (
	HeaderAuth AuthType = iota
	ParameterAuth
)

type Endpoint struct {
	Url      string
	AuthType AuthType
	Accept   string
}

type Token interface {
	SetToHeader(request *http.Request)
	ToParam() url.Values
}

func MergeParams(lhs url.Values, rhs url.Values) url.Values {
	if lhs == nil && rhs == nil {
		return url.Values{}
	} else if lhs == nil {
		return rhs
	} else if rhs == nil {
		return lhs
	}

	result := url.Values{}

	for key, values := range lhs {
		if len(values) != 0 {
			for _, value := range values {
				if result[key] == nil {
					result.Set(key, value)
				} else {
					result.Add(key, value)
				}
			}
		} else {
			result[key] = []string{}
		}
	}

	for key, values := range rhs {
		if len(values) != 0 {
			for _, value := range values {
				if result[key] == nil {
					result.Set(key, value)
				} else {
					result.Add(key, value)
				}
			}
		} else {
			result[key] = []string{}
		}
	}

	return result
}

func GetRequest(endpoint Endpoint, token Token, values url.Values) ([]byte, error) {
	if values == nil {
		values = url.Values{}
	}

	if endpoint.AuthType == ParameterAuth {
		values = MergeParams(values, token.ToParam())
	}

	query := values.Encode()

	uri := fmt.Sprintf("%s?%s", endpoint.Url, query)

	logrus.Debugf("Request to %s\n", uri)

	req, err := http.NewRequest(http.MethodGet, uri, nil)

	if err != nil {
		return nil, err
	}

	if endpoint.Accept != "" {
		req.Header.Set("Accept", endpoint.Accept)
	} else {
		panic("accept header is not found")
	}

	req.Header.Set("User-Agent", version.UserAgent())

	if endpoint.AuthType == HeaderAuth {
		token.SetToHeader(req)
	}

	resp, err := new(http.Client).Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if bytes, err := ioutil.ReadAll(resp.Body); err != nil {
		return nil, err
	} else {
		return bytes, nil
	}
}

func PostRequest(endpoint Endpoint, token Token, values url.Values, body []byte) ([]byte, error) {
	if values == nil {
		values = url.Values{}
	}

	if endpoint.AuthType == ParameterAuth {
		values = MergeParams(values, token.ToParam())
	}

	query := values.Encode()

	uri := fmt.Sprintf("%s?%s", endpoint.Url, query)

	logrus.Debugf("Request to %s\n", uri)

	req, err := http.NewRequest(http.MethodPost, uri, bytes.NewReader(body))

	if err != nil {
		return nil, err
	}

	if endpoint.Accept != "" {
		req.Header.Set("Accept", endpoint.Accept)
	} else {
		req.Header.Set("Accept", "application/json")
	}

	req.Header.Set("Content-Type", http.DetectContentType(body))
	req.Header.Set("User-Agent", version.UserAgent())

	if endpoint.AuthType == HeaderAuth {
		token.SetToHeader(req)
	}

	resp, err := new(http.Client).Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if bodyBytes, err := ioutil.ReadAll(resp.Body); err != nil {
		return nil, err
	} else {
		return bodyBytes, nil
	}
}
