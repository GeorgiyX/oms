package config

import (
	"os"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

func Read[Config any](file string) (Config, error) {
	var config Config
	rawYAML, err := os.ReadFile(file)
	if err != nil {
		return config, errors.WithMessage(err, "reading config file")
	}

	err = yaml.Unmarshal(rawYAML, &config)
	if err != nil {
		return config, errors.WithMessage(err, "parsing yaml")
	}

	return config, nil
}
