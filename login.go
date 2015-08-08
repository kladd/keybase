package keybase

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"net/url"

	"golang.org/x/crypto/scrypt"
)

// GetSaltParams define parameters for calls to the getsalt API endpoint.
type GetSaltParams struct {
	Username string `url:"email_or_username"`
}

// GetSaltResponse defines a response to a getsalt request.
type GetSaltResponse struct {
	Status       status `json:"status"`
	Salt         string `json:"salt"`
	CsrfToken    string `json:"csrf_token"`
	LoginSession string `json:"login_session"`
}

// GetSalt calls the getsalt API endpoint.
func GetSalt(params GetSaltParams) (GetSaltResponse, error) {
	r := new(GetSaltResponse)
	err := get("getsalt", params, r)

	return *r, err
}

// LoginParams defines a request for a call to the login API endpoint.
type LoginParams struct {
	Username string
	Salt     GetSaltResponse
}

// LoginResponse defines a response to a login request.
type LoginResponse struct {
	Status  status `json:"status"`
	Session string `json:"session"`
	Me      User   `json:"me"`
}

// Login encrypts password using a salt created with GetSalt() and transmits it
// to keybase in exchange for a session.
func Login(l LoginParams, password []byte) (*LoginResponse, error) {
	salt, err := hex.DecodeString(l.Salt.Salt)
	ls, err := base64.StdEncoding.DecodeString(l.Salt.LoginSession)

	pwh, err := scrypt.Key(password, salt, 32768, 8, 1, 224)
	hm := hmac.New(sha512.New, pwh[192:224])
	hm.Write(ls)

	v := url.Values{}
	v.Set("hmac_pwh", fmt.Sprintf("%x", hm.Sum(nil)))
	v.Set("email_or_username", l.Username)
	v.Set("csrf_token", l.Salt.CsrfToken)
	v.Set("login_session", l.Salt.LoginSession)

	r := new(LoginResponse)
	post("login", v, r)

	session = r.Session

	return r, err
}
