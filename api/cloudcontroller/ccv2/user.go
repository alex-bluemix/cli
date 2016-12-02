package ccv2

// User represents a cloud controller user.
type User struct {
	GUID string
}

// NewUser creates a new cloud controller user from the provided UAA user ID.
func (client *Client) NewUser(uaaUserID string) (User, Warnings, error) {
	return User{}, nil, nil
}
