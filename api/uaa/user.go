package uaa

import (
	"bytes"
	"encoding/json"
	"net/http"

	"code.cloudfoundry.org/cli/api/uaa/internal"
)

// User represents an UAA user account.
type User struct {
	ID string
}

type newUserRequest struct {
	Username string   `json:"userName"`
	Password string   `json:"password"`
	Name     userName `json:"name"`
	Emails   []email  `json:"emails"`
}

type userName struct {
	FamilyName string `json:"familyName"`
	GivenName  string `json:"givenName"`
}

type email struct {
	Value   string `json:"value"`
	Primary bool   `json:"primary"`
}

type newUserResponse struct {
	ID       string `json:"id"`
	Username string `json:"userName"`
}

// NewUser creates a new UAA user account with the provided password.
func (client *Client) NewUser(user string, password string) (User, error) {
	userRequest := newUserRequest{
		Username: user,
		Password: password,
		Name: userName{
			FamilyName: user,
			GivenName:  user,
		},
		Emails: []email{
			{
				Value: user,
				// TODO: should we set primary to true? it's set to false in the old code
				Primary: false,
			},
		},
	}

	body, err := json.Marshal(userRequest)
	if err != nil {
		return User{}, err
	}

	request, err := client.newRequest(requestOptions{
		RequestName: internal.NewUserRequest,
		Header: http.Header{
			"Content-Type": {"application/json"},
		},
		Body: bytes.NewBuffer(body),
	})
	if err != nil {
		return User{}, err
	}

	var userResponse newUserResponse
	response := Response{
		Result: &userResponse,
	}
	err = client.connection.Make(request, &response)
	if err != nil {
		return User{}, err
	}

	return User{ID: userResponse.ID}, nil
}
