package keybase

// KeyFetchParams defines params for the key/fetch API endpoint.
type KeyFetchParams struct {
	PGPKeyIDs string `url:"pgp_key_ids"`
	Ops       int    `url:"ops"`
}
