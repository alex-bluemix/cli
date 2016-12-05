package uaa

import "fmt"

// UnverifiedServerError replaces x509.UnknownAuthorityError when the server
// has SSL but the client is unable to verify it's certificate
type UnverifiedServerError struct {
	URL string
}

func (e UnverifiedServerError) Error() string {
	return "x509: certificate signed by unknown authority"
}

// RequestError represents a connection level request error
type RequestError struct {
	Err error
}

func (e RequestError) Error() string {
	return e.Err.Error()
}

// InvalidAuthTokenError is returned when the client has an invalid
// authorization header.
type InvalidAuthTokenError struct {
	Message string
}

func (e InvalidAuthTokenError) Error() string {
	return e.Message
}

// RawHTTPStatusError wraps HTTP responses with 4xx and 5xx status codes.
type RawHTTPStatusError struct {
	StatusCode  int
	RawResponse []byte
}

func (r RawHTTPStatusError) Error() string {
	return fmt.Sprintf("Error Code: %d\nRaw Response: %s", r.StatusCode, r.RawResponse)
}
