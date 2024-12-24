package svg

import "strconv"

type Config struct {
	Height         string
	Width          string
	Fill           string
	Stroke         string
	StrokeWidth    string
	StrokeLineCap  string
	StrokeLineJoin string
}

type Option func(*Config)

func Height(h uint64) Option {
	return func(c *Config) {
		c.Height = strconv.FormatUint(h, 10)
	}
}

func Width(w uint64) Option {
	return func(c *Config) {
		c.Width = strconv.FormatUint(w, 10)
	}
}

func Square(s uint64) Option {
	return func(c *Config) {
		c.Height = strconv.FormatUint(s, 10)
		c.Width = strconv.FormatUint(s, 10)
	}
}

func Fill(f string) Option {
	return func(c *Config) {
		c.Fill = f
	}
}

func Stroke(s string) Option {
	return func(c *Config) {
		c.Stroke = s
	}
}

func StrokeWidth(sw uint64) Option {
	return func(c *Config) {
		c.StrokeWidth = strconv.FormatUint(sw, 10)
	}
}

func StrokeLineCap(sl string) Option {
	return func(c *Config) {
		c.StrokeLineCap = sl
	}
}

func StrokeLineJoin(sl string) Option {
	return func(c *Config) {
		c.StrokeLineJoin = sl
	}
}

var DefaultConfig = Config{
	Height:         "24",
	Width:          "24",
	Fill:           "none",
	Stroke:         "currentColor",
	StrokeWidth:    "1",
	StrokeLineCap:  "round",
	StrokeLineJoin: "round",
}

func Apply(opts ...Option) (c Config) {
	if len(opts) == 0 {
		return DefaultConfig
	}

	for _, opt := range opts {
		opt(&c)
	}

	if c.Height == "" {
		c.Height = DefaultConfig.Height
	}
	if c.Width == "" {
		c.Width = DefaultConfig.Width
	}
	if c.Fill == "" {
		c.Fill = DefaultConfig.Fill
	}
	if c.Stroke == "" {
		c.Stroke = DefaultConfig.Stroke
	}
	if c.StrokeWidth == "" {
		c.StrokeWidth = DefaultConfig.StrokeWidth
	}
	if c.StrokeLineCap == "" {
		c.StrokeLineCap = DefaultConfig.StrokeLineCap
	}
	if c.StrokeLineJoin == "" {
		c.StrokeLineJoin = DefaultConfig.StrokeLineJoin
	}

	return c
}
