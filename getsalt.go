package keybase

// GetSaltResponse contains a response from a getsalt api call
type GetSaltResponse struct {
	Status       status `json:"status"`
	Salt         string `json:"salt"`
	CsrfToken    string `json:"csrf_token"`
	LoginSession string `json:"login_token"`
}

type getSaltParams struct {
	Username string `url:"email_or_username"`
}

// GetSalt corresponds to the getsalt rpc
func GetSalt(username string) (*GetSaltResponse, error) {
	r := new(GetSaltResponse)

	err := call("getsalt", getSaltParams{username}, r)

	return r, err
}
