package keybase

// UserLookup wraps the user/lookup API endpoint.
func UserLookup(params UserLookupParams) (*UserLookupResponse, error) {
	r := new(UserLookupResponse)
	err := get("user/lookup", params, r)

	return r, err
}

// UserAutocomplete calls the user/autocomplete API endpoint.
func UserAutocomplete(query string) (*UserAutocompleteResponse, error) {
	r := new(UserAutocompleteResponse)
	err := get("user/autocomplete", struct {
		Query string `url:"q"`
	}{query}, r)

	return r, err
}

// UserDiscover wraps the user/discover API endpoint.
func UserDiscover(params UserDiscoverParams) (*UserDiscoverResponse, error) {
	r := new(UserDiscoverResponse)
	err := get("user/discover", params, r)

	return r, err
}

// User is what the Keybase API defines as a "User Object".
// Documentation is available at https://keybase.io/docs/api/1.0/user_objects.
type User struct {
	ID     string `json:"id"`
	Basics struct {
		Ctime    int    `json:"ctime"`
		Mtime    int    `json:"mtime"`
		Salt     string `json:"salt"`
		UID      string `json:"uid"`
		Username string `json:"username"`
	} `json:"basics"`
	InvitationStats struct {
		Available int `json:"available"`
		Open      int `json:"open"`
		Power     int `json:"power"`
		Used      int `json:"used"`
	} `json:"invitation_stats"`
	Profile struct {
		Bio      string `json:"bio"`
		FullName string `json:"full_name"`
		Location string `json:"location"`
		Mtime    int    `json:"mtime"`
	} `json:"profile"`
	Emails struct {
		Primary struct {
			Email      string `json:"email"`
			IsVerified int    `json:"is_verified"`
		} `json:"primary"`
	} `json:"emails"`
	PublicKeys struct {
		Primary struct {
			KeyFingerprint string `json:"key_fingerprint"`
			KID            string `json:"kid"`
			KeyType        int    `json:"key_type"`
			Bundle         string `json:"bundle"`
			Ctime          int    `json:"ctime"`
			Mtime          int    `json:"mtime"`
		} `json:"primary"`
	} `json:"public_keys"`
	PrivateKeys struct {
		Bundle  string `json:"bundle"`
		KeyType int    `json:"key_type"`
		KID     string `json:"kid"`
		Ctime   int    `json:"ctime"`
		Mtime   int    `json:"mtime"`
	} `json:"private_keys"`
	CryptoCurrencyAddress struct {
		Bitcoin struct {
			Address string `json:"address"`
			SigID   string `json:"sig_id"`
		} `json:"bitcoin"`
	} `json:"cryptocurrency_addresses"`
}
