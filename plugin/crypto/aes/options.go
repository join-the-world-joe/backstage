package aes

type Options struct {
	padding string // "PKCS5", "PKCS7", "ZEROS"
	key     string
	iv      string
}

type Option func(*Options)

func WithPadding(padding string) Option { // "PKCS5", "PKCS7", "ZEROS"
	return func(o *Options) {
		o.padding = padding
	}
}

func WithKey(key string) Option {
	return func(o *Options) {
		o.key = key
	}
}

func WithIV(iv string) Option {
	return func(o *Options) {
		o.iv = iv
	}
}
