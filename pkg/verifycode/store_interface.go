package verifycode

type Store interface {
	// Set verification code
	Set(id string, value string) bool

	// Get verification code
	Get(id string, clear bool) string

	// Verify verification code
	Verify(id, answer string, clear bool) bool
}
