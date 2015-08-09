package keybase

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/scrypt"
)

// GetSaltParams define parameters for calls to the getsalt API endpoint.
type GetSaltParams struct {
	Username string `url:"email_or_username"`
}

// GetSaltResponse defines a response to a getsalt request.
type GetSaltResponse struct {
	Status       status `json:"status" url:"-"`
	Salt         string `json:"salt" url:"-"`
	CsrfToken    string `json:"csrf_token" url:"-"`
	LoginSession string `json:"login_session" url:"login_session"`
}

// GetSalt calls the getsalt API endpoint.
func GetSalt(params GetSaltParams) (*GetSaltResponse, error) {
	r := new(GetSaltResponse)
	err := get("getsalt", params, r)

	return r, err
}

// LoginResponse defines a response to a login request.
type LoginResponse struct {
	Status  status `json:"status"`
	Session string `json:"session"`
	Me      User   `json:"me"`
}

// LoginParams defines a request for a call to the login API endpoint.
type LoginParams struct {
	GetSaltParams
	Salt    GetSaltResponse
	hmacPwh string `url:"hmac_pwh"`
}

// Login encrypts password using a salt created with GetSalt() and transmits it
// to keybase in exchange for a session.
func Login(l LoginParams, password []byte) (*LoginResponse, error) {
	salt, err := hex.DecodeString(l.Salt.Salt)
	ls, err := base64.StdEncoding.DecodeString(l.Salt.LoginSession)

	pwh, err := scrypt.Key(password, salt, 32768, 8, 1, 224)
	hm := hmac.New(sha512.New, pwh[192:224])
	hm.Write(ls)

	l.hmacPwh = fmt.Sprintf("%x", hm.Sum(nil))

	r := new(LoginResponse)
	post("login", l, r)

	session = r.Session

	return r, err
}
