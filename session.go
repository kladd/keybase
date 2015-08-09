package keybase

import (
	"io/ioutil"
	"os"
)

// Session is a type alias intended to store a user's session after login.
type Session string

// LoadSession loads an existing session if it exists
func LoadSession() Session {
	buf, err := ioutil.ReadFile(os.Getenv("HOME") + "/.keybase_session")

	s := Session(buf)
	if err != nil {
		s = Session("")
	}

	if s != "" {
		session = s
	}

	return s
}

// Save writes a login session to disk for use later
func (s *Session) Save() {
	err := ioutil.WriteFile(os.Getenv("HOME")+"/.keybase_session", []byte(s.String()), 0644)
	if err != nil {
		panic("Unable to save session to disk")
	}
	session = *s
}

// Destroy clears a user's session
func (s *Session) Destroy() {
	session = ""
	*s = ""
	os.Remove(os.Getenv("HOME") + "/.keybase_session")
}

func (s *Session) String() string {
	return string(*s)
}
