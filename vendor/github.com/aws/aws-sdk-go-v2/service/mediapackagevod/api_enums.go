// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediapackagevod

type AdMarkers string

// Enum values for AdMarkers
const (
	AdMarkersNone           AdMarkers = "NONE"
	AdMarkersScte35Enhanced AdMarkers = "SCTE35_ENHANCED"
	AdMarkersPassthrough    AdMarkers = "PASSTHROUGH"
)

func (enum AdMarkers) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum AdMarkers) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type EncryptionMethod string

// Enum values for EncryptionMethod
const (
	EncryptionMethodAes128    EncryptionMethod = "AES_128"
	EncryptionMethodSampleAes EncryptionMethod = "SAMPLE_AES"
)

func (enum EncryptionMethod) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum EncryptionMethod) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type Profile string

// Enum values for Profile
const (
	ProfileNone    Profile = "NONE"
	ProfileHbbtv15 Profile = "HBBTV_1_5"
)

func (enum Profile) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum Profile) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type StreamOrder string

// Enum values for StreamOrder
const (
	StreamOrderOriginal               StreamOrder = "ORIGINAL"
	StreamOrderVideoBitrateAscending  StreamOrder = "VIDEO_BITRATE_ASCENDING"
	StreamOrderVideoBitrateDescending StreamOrder = "VIDEO_BITRATE_DESCENDING"
)

func (enum StreamOrder) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum StreamOrder) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}
