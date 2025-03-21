package config

import (
	"github.com/alhaos/enroll/internal/logging"
	"github.com/alhaos/enroll/internal/webServer"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Logging   logging.Config   `yaml:"logging"`
	WebServer webServer.Config `yaml:"webServer"`
}

// MustLoadFromFile constructor from config
func MustLoadFromFile(filename string) (*Config, error) {

	c := Config{}

	err := cleanenv.ReadConfig(filename, &c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
