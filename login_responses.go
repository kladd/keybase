package keybase

// GetSaltResponse defines a response to a getsalt request.
type GetSaltResponse struct {
	Status       status `json:"status"`
	Salt         string `json:"salt"`
	CsrfToken    string `json:"csrf_token"`
	LoginSession string `json:"login_session"`
}

// LoginResponse defines a response to a login request.
type LoginResponse struct {
	Status  status `json:"status"`
	Session string `json:"session"`
	Me      User   `json:"me"`
}
