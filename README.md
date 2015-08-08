# keybase - A Go library for accessing the Keybase API

[![GoDoc](https://godoc.org/github.com/kladd/keybase?status.svg)](https://godoc.org/github.com/kladd/keybase)

## Keybase API

The Keybase API documentation is available [on their website](https://keybase.io/docs/api/1.0).

## Example

> Note: the Login flow/usage will be changing soon. Currently the Keybase session is stored package-wide as a side effect of calling Login() successfully.

```go

var username string = "some_username"
var passphrase []byte = []byte("some passphrase")

func main() {
	// Obtain a salt for the user
	salt, _ := keybase.GetSalt(GetSaltParams{Username: username})

	// Using the salt log the user in.
	// API calls following this will be authenticated given the global state
	// described in the note above.
	// Login() returns a LoginResponse.
	login, _ := keybase.Login(LoginParams{username,salt}, passphrase))

	// Make some API call.
	response, _ := keybase.KeyFetch(

		// KeyFetchParams mirror API call parameters in name and format as
		// exactly as possible.
		keybase.KeyFetchParams{
			PGPKeyIDs: "comma,separated,list,of,key,ids",
			Ops: keybase.OpEncrypt
		}
	)

	// Do something
	fmt.Println(response.Keys[0].Username)
}

```

## Methods implemented so far

* getsalt
* login
* user/lookup
* user/autocomplete
* user/discover
* key/fetch

## TODO

* Implement more api methods
* Make session not a global state
* ...
