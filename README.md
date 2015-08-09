# keybase - A Go library for accessing the Keybase API

[![GoDoc](https://godoc.org/github.com/kladd/keybase?status.svg)](https://godoc.org/github.com/kladd/keybase)

## Keybase API

The Keybase API documentation is available [on their website](https://keybase.io/docs/api/1.0).

## Example

> Note: See [Sessions](Sessions) section.

```go

var username string = "some_username"
var passphrase string = "some passphrase"

func main() {
	// Log a user in.
	// This function will execute both stages of the 2 stage authentication
	// process for keybase.
	// API calls following this will be authenticated given the global state
	// described in the note above.
	// Login() returns a LoginResponse.
	login, _ := keybase.Login(username, passphrase)

	// Make some API call.
	response, _ := keybase.KeyFetch(

		// KeyFetchParams mirror API call parameters in name and format as
		// exactly as possible.
		keybase.KeyFetchParams{
			PGPKeyIDs: "comma,separated,list,of,key,ids",
			Ops: keybase.OpEncrypt,
		}
	)

	// Do something
	fmt.Println(response.Keys[0].Username)
}

```

## Sessions

Sessions are very hacky right now. First, on a successful call to the login API,
the session is stored in a package-wide session variable. API calls made after
Login() will be authenticated automatically.

Keybase has a rate limit for calls to the login API. In order to avoid this rate
limit I'm saving sessions to disk, and reusing them in later calls to Login().
Sessions on disk can be destroyed with Session.Destroy().

This is all very hacky, and mostly for testing while I develop this API client
library. In the end, Login() will do nothing other than return its response with
the session, and every auth protected API call will take a session as a parameter.
Storing sessions to allow them to persist will be up to the user of the library.

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
