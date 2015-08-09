package keybase

// KeyFetchParams defines params for the key/fetch API endpoint.
type KeyFetchParams struct {
	PGPKeyIDs string `url:"pgp_key_ids"`
	Ops       int    `url:"ops"`
}

// KeyFetchResponse defines a response to a request to the key/fetch api
type KeyFetchResponse struct {
	Status status `json:"status"`
	Keys   []Key  `json:"keys"`
}

// KeyFetch fetches keys using the key/fetch API endpoint.
func KeyFetch(params KeyFetchParams) (*KeyFetchResponse, error) {
	r := new(KeyFetchResponse)
	err := get("key/fetch", params, r)

	return r, err
}

// Key outlines the structure of a Key object in responses containing keys.
type Key struct {
	Bundle                 string `json:"bundle"`
	UID                    string `json:"uid"`
	Username               string `json:"username"`
	KeyType                int    `json:"key_type"`
	KID                    string `json:"kid"`
	PrimaryBundleInKeyring int    `json:"primary_bundle_in_keyring"`
	Secret                 int    `json:"secret"`
	SelfSignType           int    `json:"self_sign_type"`
	SelfSigned             int    `json:"self_signed"`
	SubKeys                map[string]struct {
		Flags     int `json:"flags"`
		IsPrimary int `json:"is_primary"`
	} `json:"subkeys"`
}
