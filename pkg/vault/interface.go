package vault

type Encrypter interface {
	Encrypt(plainText string) (string, error)
}
