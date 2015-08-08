package keybase

// UserLookupResponse provides the user lookup response.
// Per Keybase documentation, if the user making the request is the user being
// fetched, the 'Me' field will be populated. Otherwise, the 'them' field will be.
type UserLookupResponse struct {
	Status status `json:"status"`
	Them   []User `json:"them,omitempty"`
	Me     User   `json:"me,omitempty"`
}

// UserAutocompleteResponse contains a response to user/autocomplete requests.
type UserAutocompleteResponse struct {
	Status      status `json:"status"`
	Completions []struct {
		TotalScore float64 `json:"total_score"`
		Components struct {
			Username       acComponent `json:"username"`
			KeyFingerprint struct {
				acComponent
				Algo  int `json:"algo"`
				NBits int `json:"nbits"`
			} `json:"key_fingerprint"`
			FullName   acComponent `json:"full_name"`
			Github     acComponent `json:"github"`
			Reddit     acComponent `json:"reddit"`
			Twitter    acComponent `json:"twitter"`
			Coinbase   acComponent `json:"coinbase"`
			Hackernews acComponent `json:"hackernews"`
			Websites   []struct {
				acComponent
				Protocol string `json:"protocol"`
			} `json:"websites"`
		} `json:"components"`
		UID        string `json:"uid"`
		Thumbnail  string `json:"thumbnail"`
		IsFollowee bool   `json:"is_followee"`
	} `json:"completions"`
}

type acComponent struct {
	Val   string  `json:"val"`
	Score float64 `json:"score"`
}

// UserDiscoverResponse contains a user/discover response.
type UserDiscoverResponse struct {
	Status  status `json:"status"`
	Matches struct {
		Twitter    [][]discoverAccount `json:"twitter"`
		Github     [][]discoverAccount `json:"github"`
		Hackernews [][]discoverAccount `json:"hackernews"`
		Web        [][]discoverAccount `json:"web"`
		Coinbase   [][]discoverAccount `json:"coinbase"`
	} `json:"matches"`
}

type discoverAccount struct {
	Thumbnail string `json:"thumbnail"`
	Username  string `json:"username"`
	PublicKey struct {
		KeyFingerprint string `json:"key_fingerprint"`
		Bits           int    `json:"bits"`
		Algo           int    `json:"algo"`
	} `json:"public_key"`
	FullName     string `json:"full_name"`
	CTime        int    `json:"ctime"`
	RemoteProofs struct {
		DNS            []string `json:"dns"`
		GenericWebSite []struct {
			Hostname   string `json:"hostname"`
			Protocol   string `json:"protocol"`
			Searchable string `json:"searchable"`
		} `json:"generic_web_site"`
		Twitter    string `json:"twitter"`
		Github     string `json:"github"`
		Reddit     string `json:"reddit"`
		Hackernews string `json:"hackernews"`
		Coinbase   string `json:"coinbase"`
	} `json:"remote_proofs"`
}
