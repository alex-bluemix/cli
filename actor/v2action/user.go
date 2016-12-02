package v2action

import (
	"fmt"

	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv2"
)

// User represents a CLI user.
type User ccv2.User

// NewUser creates a new user in UAA and registers it with cloud controller.
func (actor Actor) NewUser(username string, password string) (User, Warnings, error) {
	// register user with UAA client
	uaaUser, err := actor.UAAClient.NewUser(username, password)
	if err != nil {
		// TODO: display warning and return nil for error when we get a 409 status code from UAA client (user already exists)
		return User{}, nil, err
	}

	fmt.Println("UAA User", uaaUser)

	// register UID with CAPI API
	ccUser, ccWarnings, err := actor.CloudControllerClient.NewUser(uaaUser.ID)

	return User(ccUser), Warnings(ccWarnings), err
}
