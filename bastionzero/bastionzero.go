package bastionzero

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/apierror"
	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/service/apikeys"
	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/service/autodiscoveryscripts"
	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/service/environments"
	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/service/organization"
	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/service/policies"
	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/service/serviceaccounts"
	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/service/targets"
	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/service/users"
	"github.com/bastionzero/bastionzero-sdk-go/internal"
	"github.com/bastionzero/bastionzero-sdk-go/internal/client"
)

const (
	libraryVersion = "v0.2.0"
	// DefaultBaseURL is the default BastionZero API URL the Client communicates
	// with
	DefaultBaseURL   = "https://cloud.bastionzero.com/"
	defaultUserAgent = "bastionzero-sdk-go/" + libraryVersion
	mediaType        = "application/json"
)

// Client manages communication with the BastionZero API
type Client struct {
	// HTTP client used to communicate with the BastionZero API
	client *http.Client

	// Base URL for API requests
	baseURL *url.URL

	// User agent for client
	userAgent string

	// Optional extra HTTP headers to set on every request to the API
	headers map[string]string

	common client.Service // Reuse a single struct instead of allocating one for each service on the heap.

	// Services used for talking to different parts of the BastionZero API.

	ApiKeys              *apikeys.ApiKeysService
	AutodiscoveryScripts *autodiscoveryscripts.AutodiscoveryScriptsService
	Environments         *environments.EnvironmentsService
	Organization         *organization.OrganizationService
	Policies             *policies.PoliciesService
	ServiceAccounts      *serviceaccounts.ServiceAccountsService
	Targets              *targets.TargetsService
	Users                *users.UsersService
}

// NewClient returns a new BastionZero API client, using the given http.Client
// to perform all requests. If httpClient is nil, the default HTTP client
// provided by the http package is used.
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	// Set Client defaults here
	baseURL, _ := url.Parse(DefaultBaseURL)
	c := &Client{
		client:    httpClient,
		baseURL:   baseURL,
		userAgent: defaultUserAgent,
		headers:   make(map[string]string),
	}
	c.common.Client = c
	c.ApiKeys = (*apikeys.ApiKeysService)(&c.common)
	c.AutodiscoveryScripts = (*autodiscoveryscripts.AutodiscoveryScriptsService)(&c.common)
	c.Environments = (*environments.EnvironmentsService)(&c.common)
	c.Organization = (*organization.OrganizationService)(&c.common)
	c.Policies = (*policies.PoliciesService)(&c.common)
	c.ServiceAccounts = (*serviceaccounts.ServiceAccountsService)(&c.common)
	c.Targets = (*targets.TargetsService)(&c.common)
	c.Users = (*users.UsersService)(&c.common)

	return c
}

// Do sends an API request and returns the API response. The API response is
// JSON decoded and stored in the value pointed to by v, or returned as an error
// if an API error has occurred. If v implements the io.Writer interface, the
// raw response will be written to v, without attempting to decode it.
//
// The error type will be *ErrorResponse if the API response is considered an
// error
func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = checkResponse(resp)
	if err != nil {
		return resp, err
	}

	switch v := v.(type) {
	case nil:
	case io.Writer:
		_, err = io.Copy(v, resp.Body)
	default:
		decErr := json.NewDecoder(resp.Body).Decode(v)
		if decErr == io.EOF {
			// ignore EOF errors caused by empty response body
			decErr = nil
		}
		if decErr != nil {
			err = decErr
		}
	}
	return resp, err
}

// NewRequest creates an API request. A relative URL can be provided in urlStr,
// which will be resolved to the BaseURL of the Client. Relative URLS should
// always be specified without a preceding slash. If specified, the value
// pointed to by body is JSON encoded and included in as the request body if
// applicable for the provided HTTP method.
func (c *Client) NewRequest(ctx context.Context, method, urlStr string, body interface{}) (*http.Request, error) {
	u, err := c.baseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var req *http.Request
	switch method {
	case http.MethodGet, http.MethodHead, http.MethodOptions, http.MethodDelete:
		req, err = http.NewRequestWithContext(ctx, method, u.String(), nil)
		if err != nil {
			return nil, err
		}

	default:
		buf := new(bytes.Buffer)
		if body != nil {
			err = json.NewEncoder(buf).Encode(body)
			if err != nil {
				return nil, err
			}
		}

		req, err = http.NewRequestWithContext(ctx, method, u.String(), buf)
		if err != nil {
			return nil, err
		}
		req.Header.Set("Content-Type", mediaType)
	}

	for k, v := range c.headers {
		req.Header.Add(k, v)
	}

	req.Header.Set("Accept", mediaType)
	req.Header.Set("User-Agent", c.userAgent)

	return req, nil
}

// checkResponse checks the API response for errors, and returns them if
// present. A response is considered an error if it has a status code outside
// the 200 range. API error responses are expected to have either no response
// body, or a JSON response body that maps to ErrorResponse. Any other response
// body will be silently ignored. If the body contains invalid JSON, then
// ErrorMessage is set to the text of the response body.
//
// The error type will be *apierror.ErrorResponse if the response is considered
// an error
func checkResponse(r *http.Response) error {
	if c := r.StatusCode; c >= 200 && c <= 299 {
		return nil
	}

	errorResponse := &apierror.ErrorResponse{Response: r}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && len(data) > 0 {
		err := json.Unmarshal(data, errorResponse)
		if err != nil {
			errorResponse.ErrorMessage = string(data)
		}
	}

	return errorResponse
}

// Host returns the BastionZero server URL this client is configured to issue
// API requests against.
func (c *Client) Host() string {
	return c.baseURL.String()
}

type clientOpt func(*Client) error

// ClientOpt are options for New.
type ClientOpt clientOpt

// New returns a new BastionZero API client instance. If httpClient is nil, the
// default HTTP client provided by the http package is used.
func New(httpClient *http.Client, opts ...ClientOpt) (*Client, error) {
	c := NewClient(httpClient)

	// Functional options: https://www.sohamkamani.com/golang/options-pattern/
	for _, opt := range opts {
		if err := opt(c); err != nil {
			return nil, err
		}
	}

	return c, nil
}

// NewFromAPISecret returns a new BastionZero API client that authenticates its
// requests with the given API secret. If httpClient is nil, the default HTTP
// client provided by the http package is used.
func NewFromAPISecret(httpClient *http.Client, apiSecret string, opts ...ClientOpt) (*Client, error) {
	// Validate API secret
	_, err := base64.StdEncoding.DecodeString(apiSecret)
	if err != nil {
		return nil, fmt.Errorf("apiSecret (%v) is not valid base64", apiSecret)
	}

	// Add our own WithRequestHeaders option at the end of opts to ensure that
	// the API secret header is added
	return New(httpClient, append(opts, WithRequestHeaders(map[string]string{"X-API-KEY": apiSecret}))...)
}

// WithBaseURL is a client option for setting the base URL.
func WithBaseURL(baseUrl string) ClientOpt {
	return func(c *Client) error {
		u, err := url.Parse(baseUrl)
		if err != nil {
			return err
		}

		c.baseURL = u
		return nil
	}
}

// WithUserAgent is a client option for seting the user agent.
func WithUserAgent(userAgent string) ClientOpt {
	return func(c *Client) error {
		// Preserve default user agent by prepending provided user agent
		c.userAgent = fmt.Sprintf("%s %s", userAgent, c.userAgent)
		return nil
	}
}

// WithRequestHeaders sets optional HTTP headers on the client that are sent on
// each HTTP request.
func WithRequestHeaders(headers map[string]string) ClientOpt {
	return func(c *Client) error {
		for k, v := range headers {
			c.headers[k] = v
		}
		return nil
	}
}

// PtrTo returns a pointer to the provided input.
func PtrTo[T any](v T) *T {
	return &v
}

// ToStringSlice returns a list of strings given any list of a type that is
// aliased to a string.
func ToStringSlice[T ~string](elems []T) []string {
	return internal.MapSlice(elems, func(t T) string { return string(t) })
}
