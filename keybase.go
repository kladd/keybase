// Package keybase wraps the Keybase API.
package keybase

const (
	kbURL = "https://keybase.io/_/api/1.0"
)

// Bitmasks for key operations
const (
	OpEncrypt = 0x1
	OpDecrypt = 0x2
	OpVerify  = 0x4
	OpSign    = 0x8
)

var (
	session Session
)

type status struct {
	Code int    `json:"code"`
	Name string `json:"name"`
}
