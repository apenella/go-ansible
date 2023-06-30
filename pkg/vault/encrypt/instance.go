package encrypt

type PasswordReader interface {
	Read() (string, error)
}
