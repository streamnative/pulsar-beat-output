package main

import (
        "os"

        _ "github.com/AmateurEvents/filebeat-ouput-pulsar/pulsar"

        "github.com/elastic/beats/filebeat/cmd"
)

func main() {
    if err := cmd.RootCmd.Execute(); err != nil {
        os.Exit(1)
    }
}
