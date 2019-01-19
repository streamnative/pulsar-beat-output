package pulsar

import (
	"testing"
)

func Test_pulsarConfig_Validate(t *testing.T) {
	type fields struct {
		URL                                string
		UseTls                             bool
		TLSTrustCertsFilePath              string
		Topic                              string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
	// TODO: Add test cases.
		{
			"test url",
			fields{
				"",
				false,
				"",
				"test",
			},
			true,
		},
		{
			"test topic",
			fields{
				"pulsar://localhost:6650",
				false,
				"",
				"",
			},
			true,
		},
		{
			"test use tls",
			fields{
				"pulsar://localhost:6650",
				true,
				"/go/src/github.com/AmateurEvents/filebeat-ouput-pulsar/certs/ca.cert.pem",
				"persistent://public/default/my-topic1",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &pulsarConfig{
				URL:                      tt.fields.URL,
				UseTls:                     tt.fields.UseTls,
				TLSTrustCertsFilePath:      tt.fields.TLSTrustCertsFilePath,
				Topic:                              tt.fields.Topic,
			}
			if err := c.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("pulsarConfig.Validate() error = %v, wantErr=%v", err, tt.wantErr)
			}
		})
	}
}
