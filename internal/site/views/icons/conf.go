package icons

import "strconv"

type Config struct {
	Height      string
	Width       string
	StrokeWidth string
}

type Option func(*Config)

func WithHeight(h uint64) Option {
	return func(c *Config) {
		c.Height = strconv.FormatUint(h, 10)
	}
}

func WithWidth(w uint64) Option {
	return func(c *Config) {
		c.Width = strconv.FormatUint(w, 10)
	}
}

func WithStrokeWidth(sw uint64) Option {
	return func(c *Config) {
		c.StrokeWidth = strconv.FormatUint(sw, 10)
	}
}

var DefaultConfig = Config{
	Height:      "24",
	Width:       "24",
	StrokeWidth: "2",
}

func Apply(opts ...Option) Config {
	c := DefaultConfig

	for _, opt := range opts {
		opt(&c)
	}

	return c
}
