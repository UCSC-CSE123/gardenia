package main

import (
	"fmt"
	"log"
	"os"

	"github.com/UCSC-CSE123/gardenia/internal/config"
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

	fmt.Println(args)
}
