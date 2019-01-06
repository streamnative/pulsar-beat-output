package pulsar

import (
    "github.com/elastic/beats/libbeat/beat"
    "github.com/elastic/beats/libbeat/common"
    "github.com/elastic/beats/libbeat/logp"
    "github.com/elastic/beats/libbeat/outputs"
)


func init() {
    outputs.RegisterType("pulsar", makePulsar)
}

func makePulsar(beat beat.Info,
    observer outputs.Observer,
    cfg *common.Config,
) (outputs.Group, error) {
    config := defaultConfig()
    logp.Info("initialize pulsar output")
    if err := cfg.Unpack(&config); err != nil {
        return outputs.Fail(err)
    }

    logp.Info("init config %v", config)
    clientOptions, producerOptions, err := initOptions(&config)
    client, err := newPulsarClient(beat, observer, clientOptions, producerOptions, &config)
    if err != nil {
        return outputs.Fail(err)
    }
    retry := 0
    if config.MaxRetries < 0 {
        retry = -1
    }
    return outputs.Success(config.BulkMaxSize, retry, client)
}