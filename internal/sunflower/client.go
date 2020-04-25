package sunflower

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/UCSC-CSE123/gardenia/internal/beavertail"
	"github.com/UCSC-CSE123/gardenia/internal/config"
)

// Response represents the response from a call
// to sunflower
type Response struct {
	State struct {
		NumAutos int `json:"NumAutos"`
		Autos    []struct {
			ID     string `json:"ID"`
			Count  int    `json:"Count"`
			Status string `json:"Status"`
		} `json:"Autos"`
	} `json:"State"`
	DebugInfo struct {
		StopPeriodicity string `json:"StopPeriodicity"`
		InitialCount    int    `json:"InitialCount"`
		ElapsedTime     string `json:"ElapsedTime"`
	} `json:"DebugInfo"`
}

// Client represents a connection
// to the sunflower and grpc servers.
type Client struct {
	Host       string
	Port       string
	Endpoint   string
	Frequency  time.Duration
	TotalCalls int
	GRPCClient beavertail.PushDatagramClient
}

// NewClient returns a new client given the arguments.
func NewClient(args *config.Args, grpcClient beavertail.PushDatagramClient) Client {
	return Client{
		Host:       args.Host,
		Port:       args.Port,
		TotalCalls: args.TotalCalls,
		GRPCClient: grpcClient,
		Endpoint:   fmt.Sprintf("http://%s:%s/api/state", args.Host, args.Port),
	}
}

// Call calls the sunflower api once.
func (cli Client) Call() (*Response, error) {
	sfResponse := new(Response)
	httpResponse, err := http.Get(cli.Endpoint)
	if err != nil {
		return nil, err
	}
	defer httpResponse.Body.Close()

	err = json.NewDecoder(httpResponse.Body).Decode(sfResponse)
	return sfResponse, err
}

// Sample calls Call() n times.
func (cli Client) Sample(n int) ([]*Response, error) {
	responses := make([]*Response, cli.TotalCalls)
	var err error
	for i := 0; i < n; i++ {
		responses[i], err = cli.Call()
		if err != nil {
			return nil, err
		}
	}

	return responses, nil
}

// StressCSV stresses the GRPC server by calling
// it cli.TotalCalls times and writes timing retuts to wt.
//
// It can optionally takes in timeout that serves as the max time that the server can respond in.
func (cli Client) StressCSV(wt io.Writer, timeout ...time.Duration) error {
	// Get all the calls.
	tResponses, err := cli.Sample(cli.TotalCalls)
	if err != nil {
		return err
	}

	// Make a csv writer.
	csvWT := csv.NewWriter(wt)

	// Make the header
	if err := csvWT.Write([]string{"Call #", "Duration", "Acknowledgment"}); err != nil {
		return err
	}

	// Loop over each individual auto.
	for i, resp := range tResponses {
		for j, bus := range resp.State.Autos {
			// Get the current time.
			start := time.Now()

			// Make a context.
			var (
				ctx                       = context.Background()
				cancel context.CancelFunc = nil
			)
			if len(timeout) > 0 {
				ctx, cancel = context.WithTimeout(ctx, timeout[0])
			}

			// Make a push datagram, and send it over.
			push := beavertail.DatagramPush{
				BusID:                    bus.ID,
				PassengerCount:           uint32(bus.Count),
				Timestamp:                time.Now().UnixNano(),
				PassengerCountConfidence: rand.Float64() + float64(rand.Intn(100-90)+90),
			}
			ack, err := cli.GRPCClient.Push(ctx, &push)

			// Once the call is over, cancel the context.
			// If it the timeout was called before
			// this does nothing.
			if cancel != nil {
				cancel()
			}

			// Check for errors.
			if err != nil {
				return err
			}

			// Make the results.
			results := []string{
				strconv.Itoa(((resp.State.NumAutos * i) + j) + 1),
				time.Since(start).String(),
				ack.Acknowledgment.String(),
			}

			// Write the result.
			if err := csvWT.Write(results); err != nil {
				return err
			}
			csvWT.Flush()
		}
	}

	return nil
}
