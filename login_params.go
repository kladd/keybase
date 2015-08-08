package keybase

// LoginParams defines a request for a call to the login API endpoint.
type LoginParams struct {
	Username string
	Salt     GetSaltResponse
}

// GetSaltParams define parameters for calls to the getsalt API endpoint.
type GetSaltParams struct {
	Username string `url:"email_or_username"`
}
