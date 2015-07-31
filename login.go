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

// GetSaltResponse contains a response from a getsalt api call
type GetSaltResponse struct {
	Status       status `json:"status"`
	Salt         string `json:"salt"`
	CsrfToken    string `json:"csrf_token"`
	LoginSession string `json:"login_session"`
}

// GetSaltParams contain parameters for get salt call
type GetSaltParams struct {
	Username string `url:"email_or_username"`
}

// GetSalt corresponds to the getsalt rpc
func GetSalt(params GetSaltParams) (GetSaltResponse, error) {
	r := new(GetSaltResponse)

	err := get("getsalt", params, r)

	return *r, err
}

// LoginParams contain login api parameters
type LoginParams struct {
	Username     string          `url:"email_or_username"`
	HmacPwh      string          `url:"hmac_pwh"`
	LoginSession string          `url:"login_session"`
	CsrfToken    string          `url:"csrf_token"`
	Salt         GetSaltResponse `url:"-"`
}

// LoginResponse wraps a login response
type LoginResponse struct {
	Status  status `json:"status"`
	Session string `json:"session"`
	Me      string `json:"me"`
}

// Login encrypts password and transmits it to keybase
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

	return r, err
}
