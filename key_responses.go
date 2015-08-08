package keybase

// KeyFetchResponse defines a response to a request to the key/fetch api
type KeyFetchResponse struct {
	Status status `json:"status"`
	Keys   []Key  `json:"keys"`
}
