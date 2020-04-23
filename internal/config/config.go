package config

import (
	"io"

	"gopkg.in/yaml.v2"
)

// Args represent
type Args struct {
	Host       string `yaml:"Host"`
	Port       string `yaml:"Port"`
	TotalCalls int    `yaml:"Total-Calls"`
}

// FromYAML reads YAML from rd, and returns the represented args.
func FromYAML(rd io.Reader) (*Args, error) {
	// Get the args.
	var parsedArgs Args
	if err := yaml.NewDecoder(rd).Decode(&parsedArgs); err != nil {
		return nil, err
	}

	return &parsedArgs, nil
}
