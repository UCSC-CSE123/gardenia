package main

import (
	"log"
	"net"
	"os"

	"github.com/UCSC-CSE123/gardenia/internal/beavertail"
	"github.com/UCSC-CSE123/gardenia/internal/config"
	"github.com/UCSC-CSE123/gardenia/internal/sunflower"
	"google.golang.org/grpc"
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

	grpcConn, err := grpc.Dial(net.JoinHostPort(args.GRPCHost, args.GRPCPort), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}
	grpcClient := beavertail.NewPushDatagramClient(grpcConn)

	client := sunflower.NewClient(args, grpcClient)
	err = client.StressCSV(os.Stdout)
	if err != nil {
		log.Fatal(err)
	}

}
