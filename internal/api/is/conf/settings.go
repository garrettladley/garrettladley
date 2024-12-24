package conf

import "github.com/caarlos0/env/v11"

type Config struct {
	AI `envPrefix:"AI_"`
}

func Load() (Config, error) {
	return env.ParseAs[Config]()
}
