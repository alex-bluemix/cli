package uaa

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// UAAErrorResponse represents a generic UAA error response.
type UAAErrorResponse struct {
	Type        string `json:"error"`
	Description string `json:"error_description"`
}

func (e UAAErrorResponse) Error() string {
	return fmt.Sprintf("Error Type: %s\nDescription: %s", e.Type, e.Description)
}

type errorWrapper struct {
	connection Connection
}

func NewErrorWrapper() *errorWrapper {
	return new(errorWrapper)
}

func (e *errorWrapper) Wrap(innerconnection Connection) Connection {
	e.connection = innerconnection
	return e
}

func (e *errorWrapper) Make(request *http.Request, passedResponse *Response) error {
	err := e.connection.Make(request, passedResponse)

	if rawHTTPStatusErr, ok := err.(RawHTTPStatusError); ok {
		return convert(rawHTTPStatusErr)
	}

	return err
}

func convert(rawHTTPStatusErr RawHTTPStatusError) error {
	// Try to unmarshal the raw http status error into a UAA error. If
	// unmarshaling fails, return the raw error.
	var uaaErrorResponse UAAErrorResponse
	err := json.Unmarshal(rawHTTPStatusErr.RawResponse, &uaaErrorResponse)
	if err != nil {
		return rawHTTPStatusErr
	}

	switch rawHTTPStatusErr.StatusCode {
	case http.StatusUnauthorized: // 401
		if uaaErrorResponse.Type == "invalid_token" {
			return InvalidAuthTokenError{Message: uaaErrorResponse.Description}
		}
		return uaaErrorResponse
	default:
		return uaaErrorResponse
	}
}
