package keybase

// Key represents an API response key object
type Key struct {
	Status                 status        `json:"status"`
	Bundle                 string        `json:"bundle"`
	KeyType                int           `json:"key_type"`
	Kid                    string        `json:"kid"`
	PrimaryBundleInKeyring int           `json:"primary_bundle_in_keyring"`
	Secret                 int           `json:"secret"`
	SelfSignType           int           `json:"self_sign_type"`
	SelfSigned             int           `json:"self_signed"`
	SubKeys                []interface{} `json:"subkeys"`
}

// KeyFetchResponse contains a response to a request to the key/fetch api
type KeyFetchResponse struct {
	Status status `json:"status"`
	Keys   []Key  `json:"keys"`
}

// KeyFetchParams contains params for the key/fetch API endpoint
type KeyFetchParams struct {
	PGPKeyIDs string `url:"pgp_key_ids"`
	Ops       int    `url:"ops"`
}

// KeyFetch fetches keys using the key/fetch.json API
func KeyFetch(params KeyFetchParams) (*KeyFetchResponse, error) {
	r := new(KeyFetchResponse)
	err := call("key/fetch", params, r)

	return r, err
}
