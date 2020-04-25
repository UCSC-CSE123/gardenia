package config

import (
	"io"

	"gopkg.in/yaml.v2"
)

// Args represent
type Args struct {
	Host       string `yaml:"Sunflower-Host"`
	Port       string `yaml:"Sunflower-Port"`
	TotalCalls int    `yaml:"Sunflower-Calls"`
	GRPCHost   string `yaml:"GRPC-Host"`
	GRPCPort   string `yaml:"GRPC-Port"`
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
