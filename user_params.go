package keybase

// UserLookupParams provides user lookup request parameters.
type UserLookupParams struct {
	Usernames      string `url:"usernames,omitempty"`
	Domain         string `url:"domain,omitempty"`
	Twitter        string `url:"twitter,omitempty"`
	Github         string `url:"github,omitempty"`
	Reddit         string `url:"reddit,omitempty"`
	HackerNews     string `url:"hackernews,omitempty"`
	Coinbase       string `url:"coinbase,omitempty"`
	KeyFingerprint string `url:"key_fingerprint,omitempty"`

	Fields string `url:"fields,omitempty"`
}

// UserDiscoverParams provides user/discover params.
type UserDiscoverParams UserLookupParams
