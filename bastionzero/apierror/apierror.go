package apierror

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// An ErrorResponse reports the error caused by an API request
type ErrorResponse struct {
	// HTTP response that caused this error
	Response *http.Response

	// Error message
	ErrorMessage string `json:"errorMsg"`

	// Error type
	ErrorType string `json:"errorType"`

	// Validation errors
	ValidationErrors map[string][]string `json:"errors"`
}

func (r *ErrorResponse) Error() string {
	if r.ErrorType != "" {
		// If errorType is provided, then we always have an errorMsg
		return fmt.Sprintf("%v %v: %v: %v (%v)", r.Response.Request.Method, r.Response.Request.URL, r.Response.Status, r.ErrorMessage, r.ErrorType)
	} else if r.ErrorMessage != "" {
		// If we fail to decode response into valid JSON, then ErrorMessage will
		// be set to the body of the response
		return fmt.Sprintf("%v %v: %v: %v", r.Response.Request.Method, r.Response.Request.URL, r.Response.Status, r.ErrorMessage)
	} else if len(r.ValidationErrors) != 0 {
		// Model binding error on the request populates validation errors

		var prettyMsg string = "Bad Request:"
		for prop, errors := range r.ValidationErrors {
			prettyMsg += fmt.Sprintf(" %v: %v", prop, strings.Join(errors[:], ", "))
		}

		return fmt.Sprintf("%v %v: %d %v", r.Response.Request.Method, r.Response.Request.URL, r.Response.StatusCode, prettyMsg)
	} else {
		// Otherwise, display HTTP status code and associated HTTP status code
		// message
		return fmt.Sprintf("%v %v: %v", r.Response.Request.Method, r.Response.Request.URL, r.Response.Status)
	}
}

// IsAPIErrorStatusCode returns true when the error is of
// *apierror.ErrorResponse type and the underlying HTTP status code error
// matches the one passed as the argument.
func IsAPIErrorStatusCode(err error, code int) bool {
	bzeroError := &ErrorResponse{}
	if errors.As(err, &bzeroError) {
		return bzeroError.Response.StatusCode == code
	} else {
		return false
	}
}
