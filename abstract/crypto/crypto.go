package crypto

type Crypto interface {
	Name() string
	Encrypt([]byte)([]byte, error)
	Decrypt([]byte)([]byte, error)
}
