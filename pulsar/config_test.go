package pulsar

import (
	"testing"
	"time"

	"github.com/elastic/beats/libbeat/outputs/codec"
)

func Test_pulsarConfig_Validate(t *testing.T) {
	type fields struct {
		URL                                string
		IOThreads                          int
		OperationTimeoutSeconds            time.Duration
		MessageListenerThreads             int
		ConcurrentLookupRequests           int
		TLSTrustCertsFilePath              string
		TLSAllowInsecureConnection         bool
		StatsIntervalInSeconds             int
		Codec                              codec.Config
		BulkMaxSize                        int
		MaxRetries                         int
		Topic                              string
		Name                               string
		Properties                         map[string]string
		SendTimeout                        time.Duration
		MaxPendingMessages                 int
		MaxPendingMessagesAcrossPartitions int
		BlockIfQueueFull                   bool
		MessageRoutingMode                 int
		HashingScheme                      int
		CompressionType                    int
		Batching                           bool
		BatchingMaxPublishDelay            time.Duration
		BatchingMaxMessages                uint
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &pulsarConfig{
				URL:                        tt.fields.URL,
				IOThreads:                  tt.fields.IOThreads,
				OperationTimeoutSeconds:    tt.fields.OperationTimeoutSeconds,
				MessageListenerThreads:     tt.fields.MessageListenerThreads,
				ConcurrentLookupRequests:   tt.fields.ConcurrentLookupRequests,
				TLSTrustCertsFilePath:      tt.fields.TLSTrustCertsFilePath,
				TLSAllowInsecureConnection: tt.fields.TLSAllowInsecureConnection,
				StatsIntervalInSeconds:     tt.fields.StatsIntervalInSeconds,
				Codec:                              tt.fields.Codec,
				BulkMaxSize:                        tt.fields.BulkMaxSize,
				MaxRetries:                         tt.fields.MaxRetries,
				Topic:                              tt.fields.Topic,
				Name:                               tt.fields.Name,
				Properties:                         tt.fields.Properties,
				SendTimeout:                        tt.fields.SendTimeout,
				MaxPendingMessages:                 tt.fields.MaxPendingMessages,
				MaxPendingMessagesAcrossPartitions: tt.fields.MaxPendingMessagesAcrossPartitions,
				BlockIfQueueFull:                   tt.fields.BlockIfQueueFull,
				MessageRoutingMode:                 tt.fields.MessageRoutingMode,
				HashingScheme:                      tt.fields.HashingScheme,
				CompressionType:                    tt.fields.CompressionType,
				Batching:                           tt.fields.Batching,
				BatchingMaxPublishDelay:            tt.fields.BatchingMaxPublishDelay,
				BatchingMaxMessages:                tt.fields.BatchingMaxMessages,
			}
			if err := c.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("pulsarConfig.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
