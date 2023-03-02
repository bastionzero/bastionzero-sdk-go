package client

import (
	"context"
	"net/http"
	"net/url"
	"reflect"

	"github.com/google/go-querystring/query"
)

// ClientInterface is an interface for making and executing BastionZero API
// requests
type ClientInterface interface {
	NewRequest(ctx context.Context, method, urlStr string, body interface{}) (*http.Request, error)
	Do(req *http.Request, v interface{}) (*http.Response, error)
}

// Service manages communication with exactly one BastionZero API service. It
// contains a client that can be used to make API requests to BastionZero.
type Service struct {
	Client ClientInterface
}

// AddOptions adds the parameters in opts as URL query parameters to s. opts
// must be a struct whose fields may contain "url" tags.
func AddOptions(s string, opts interface{}) (string, error) {
	v := reflect.ValueOf(opts)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	qs, err := query.Values(opts)
	if err != nil {
		return s, err
	}

	u.RawQuery = qs.Encode()
	return u.String(), nil
}
