package geo_lite

type Options struct {
	file string
}

type Option func(*Options)

func WithFile(file string) Option {
	return func(o *Options) {
		o.file = file
	}
}