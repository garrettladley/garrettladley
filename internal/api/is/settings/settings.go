package settings

import "github.com/caarlos0/env/v11"

type Settings struct {
	AI `envPrefix:"AI_"`
}

func Load() (Settings, error) {
	return env.ParseAs[Settings]()
}
