package rsa

type Options struct {
	privateKey []byte
	publicKey []byte
}

type Option func(*Options)

func WithPrivateKey(key []byte) Option {
	return func(o *Options) {
		o.privateKey = key
	}
}

func WithPublicKey(key []byte) Option {
	return func(o *Options) {
		o.publicKey = key
	}
}