package config

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestFromYAML(t *testing.T) {
	type args struct {
		rd io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    *Args
		wantErr bool
	}{
		{
			name: "Valid YAML",
			args: args{
				rd: strings.NewReader(`Host: localhost
GRPC-Host: localhost
Port: 8080
GRPC-Port: 8081
Total-Calls: 100`),
			},
			want: &Args{
				Host:       "localhost",
				Port:       "8080",
				GRPCHost:   "localhost",
				GRPCPort:   "8081",
				TotalCalls: 100,
			},
			wantErr: false,
		},
		{
			name: "Inalid YAML",
			args: args{
				rd: strings.NewReader(`
				Host: localhost
Port: 8080
TotalCalls: 100`),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FromYAML(tt.args.rd)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromYAML() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromYAML() = %v, want %v", got, tt.want)
			}
		})
	}
}
