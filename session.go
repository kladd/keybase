package keybase

// Session is a type alias intended to store a user's session after login.
type Session string

// Destroy clears a user's session
func (s *Session) Destroy() {
	session = ""
}

func (s *Session) String() string {
	return string(session)
}
