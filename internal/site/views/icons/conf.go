package icons

import "strconv"

type Config struct {
	Height      uint
	Width       uint
	StrokeWidth uint
}

func (c Config) Into() StrConfig {
	return StrConfig{
		Height:      strconv.FormatUint(uint64(c.Height), 10),
		Width:       strconv.FormatUint(uint64(c.Width), 10),
		StrokeWidth: strconv.FormatUint(uint64(c.StrokeWidth), 10),
	}
}

type StrConfig struct {
	Height      string
	Width       string
	StrokeWidth string
}

type Option func(*Config)

func WithHeight(h uint) Option {
	return func(c *Config) {
		c.Height = h
	}
}

func WithWidth(w uint) Option {
	return func(c *Config) {
		c.Width = w
	}
}

func WithStrokeWidth(sw uint) Option {
	return func(c *Config) {
		c.StrokeWidth = sw
	}
}

var DefaultConfig = Config{
	Height:      24,
	Width:       24,
	StrokeWidth: 2,
}

func Apply(opts ...Option) StrConfig {
	c := DefaultConfig

	for _, opt := range opts {
		opt(&c)
	}

	return c.Into()
}
