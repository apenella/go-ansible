package resolve

// PasswordReader defines the implementation of a password reader
type PasswordReader interface {
	Read() (string, error)
}
