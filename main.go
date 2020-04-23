package main

import (
	"log"
	"os"

	"github.com/UCSC-CSE123/gardenia/internal/config"
	"github.com/UCSC-CSE123/gardenia/internal/sunflower"
)

func main() {

	// Get the config file.
	var yamlFile = "config.yaml"
	const minArgs = 1
	if len(os.Args) > minArgs {
		yamlFile = os.Args[1]
	}
	yamlFD, err := os.Open(yamlFile)
	if err != nil {
		log.Fatalf("could not open %s: %v\n", yamlFile, err)
	}

	args, err := config.FromYAML(yamlFD)
	if err != nil {
		log.Fatalf("could not parse YAML: %v\n", err)
	}

	client := sunflower.NewClient(args)
	resp, err := client.Sample()
	if err != nil {
		log.Fatal(err)
	}
	_ = resp
}
