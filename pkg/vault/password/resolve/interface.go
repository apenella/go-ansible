package resolve

type PasswordReader interface {
	Read() (string, error)
}
