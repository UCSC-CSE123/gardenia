package config

import (
	"fmt"
	"io"
	"time"

	"gopkg.in/yaml.v2"
)

// Args represent
type Args struct {
	Host          string `yaml:"Host"`
	Port          string `yaml:"Port"`
	TotalCalls    int    `yaml:"Total-Calls"`
	Duration      int    `yaml:"Duration"`
	DurationUnits string `yaml:"DurationUnits"`
	GoDuration    time.Duration
}

// FromYAML reads YAML from rd, and returns the represented args.
// Returns an error if the time is not parsable.
func FromYAML(rd io.Reader) (*Args, error) {
	// Get the args.
	var parsedArgs Args
	if err := yaml.NewDecoder(rd).Decode(&parsedArgs); err != nil {
		return nil, err
	}

	// Verify the units
	var err error
	durationString := fmt.Sprintf("%d%s", parsedArgs.Duration, parsedArgs.DurationUnits)
	parsedArgs.GoDuration, err = time.ParseDuration(durationString)
	if err != nil {
		return nil, err
	}

	return &parsedArgs, nil
}
