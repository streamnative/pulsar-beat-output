package main

import (
	"os"

	_ "github.com/elastic/beats/v7/heartbeat/include"
	"github.com/elastic/beats/v7/x-pack/heartbeat/cmd"
	_ "github.com/streamnative/pulsar-beat-output/pulsar"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
