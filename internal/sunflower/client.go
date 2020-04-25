package sunflower

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/UCSC-CSE123/gardenia/internal/beavertail"
	"github.com/UCSC-CSE123/gardenia/internal/config"
)

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

type Client struct {
	Host       string
	Port       string
	Endpoint   string
	Frequency  time.Duration
	TotalCalls int
	GRPCClient beavertail.PushDatagramClient
}

func NewClient(args *config.Args, grpcClient beavertail.PushDatagramClient) Client {
	return Client{
		Host:       args.Host,
		Port:       args.Port,
		TotalCalls: args.TotalCalls,
		GRPCClient: grpcClient,
		Endpoint:   fmt.Sprintf("http://%s:%s/api/state", args.Host, args.Port),
	}
}

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

func (cli Client) Sample() ([]*Response, error) {
	responses := make([]*Response, cli.TotalCalls)
	var err error
	for i := 0; i < cli.TotalCalls; i++ {
		responses[i], err = cli.Call()
		if err != nil {
			return nil, err
		}
	}

	return responses, nil
}

func (cli Client) Stress() ([]*beavertail.DatagramAck, error) {
	tResponses, err := cli.Sample()
	if err != nil {
		return nil, err
	}
	var tAcks []*beavertail.DatagramAck

	for _, resp := range tResponses {
		for _, bus := range resp.State.Autos {
			push := beavertail.DatagramPush{
				BusID:                    bus.ID,
				PassengerCount:           uint32(bus.Count),
				Timestamp:                time.Now().UnixNano(),
				PassengerCountConfidence: rand.Float64() + float64(rand.Intn(100-90)+90),
			}
			ack, err := cli.GRPCClient.Push(context.Background(), &push)
			if err != nil {
				return nil, err
			}
			tAcks = append(tAcks, ack)
		}
	}

	return tAcks, nil
}
