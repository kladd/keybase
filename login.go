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
type getSaltParams struct {
	Username string `url:"email_or_username"`
}

// GetSaltResponse defines a response to a getsalt request.
type getSaltResponse struct {
	Status       status `json:"status" url:"-"`
	Salt         string `json:"salt" url:"-"`
	CsrfToken    string `json:"csrf_token" url:"-"`
	LoginSession string `json:"login_session" url:"login_session"`
}

func getSalt(params getSaltParams) (*getSaltResponse, error) {
	r := new(getSaltResponse)
	err := get("getsalt", params, r)

	return r, err
}

// LoginResponse defines a response to a login request.
type LoginResponse struct {
	Status  status  `json:"status"`
	Session Session `json:"session"`
	Me      User    `json:"me"`
}

// LoginParams defines a request for a call to the login API endpoint.
type loginParams struct {
	getSaltParams
	Salt    getSaltResponse
	hmacPwh string `url:"hmac_pwh"`
}

// Login encrypts password using a salt created with GetSalt() and transmits it
// to keybase in exchange for a session.
func Login(username string, password string) (*LoginResponse, error) {
	gsp := getSaltParams{username}
	gspr, err := getSalt(gsp)

	l := loginParams{gsp, *gspr, ""}

	salt, err := hex.DecodeString(l.Salt.Salt)
	ls, err := base64.StdEncoding.DecodeString(l.Salt.LoginSession)

	pwh, err := scrypt.Key([]byte(password), salt, 32768, 8, 1, 224)
	hm := hmac.New(sha512.New, pwh[192:224])
	hm.Write(ls)

	l.hmacPwh = fmt.Sprintf("%x", hm.Sum(nil))

	r := new(LoginResponse)
	post("login", l, r)

	session = r.Session

	return r, err
}
