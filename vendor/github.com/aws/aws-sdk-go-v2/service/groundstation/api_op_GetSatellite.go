// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package groundstation

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/groundstation-2019-05-23/GetSatelliteRequest
type GetSatelliteInput struct {
	_ struct{} `type:"structure"`

	// SatelliteId is a required field
	SatelliteId *string `location:"uri" locationName:"satelliteId" type:"string" required:"true"`
}

// String returns the string representation
func (s GetSatelliteInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetSatelliteInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetSatelliteInput"}

	if s.SatelliteId == nil {
		invalidParams.Add(aws.NewErrParamRequired("SatelliteId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetSatelliteInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.SatelliteId != nil {
		v := *s.SatelliteId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "satelliteId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/groundstation-2019-05-23/GetSatelliteResponse
type GetSatelliteOutput struct {
	_ struct{} `type:"structure"`

	DateCreated *time.Time `locationName:"dateCreated" type:"timestamp"`

	LastUpdated *time.Time `locationName:"lastUpdated" type:"timestamp"`

	NoradSatelliteID *int64 `locationName:"noradSatelliteID" min:"1" type:"integer"`

	SatelliteArn *string `locationName:"satelliteArn" type:"string"`

	SatelliteId *string `locationName:"satelliteId" min:"1" type:"string"`

	Tags map[string]string `locationName:"tags" type:"map"`
}

// String returns the string representation
func (s GetSatelliteOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetSatelliteOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.DateCreated != nil {
		v := *s.DateCreated

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "dateCreated",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.LastUpdated != nil {
		v := *s.LastUpdated

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "lastUpdated",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.NoradSatelliteID != nil {
		v := *s.NoradSatelliteID

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "noradSatelliteID", protocol.Int64Value(v), metadata)
	}
	if s.SatelliteArn != nil {
		v := *s.SatelliteArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "satelliteArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SatelliteId != nil {
		v := *s.SatelliteId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "satelliteId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	return nil
}

const opGetSatellite = "GetSatellite"

// GetSatelliteRequest returns a request value for making API operation for
// AWS Ground Station.
//
// Returns a satellite.
//
//    // Example sending a request using GetSatelliteRequest.
//    req := client.GetSatelliteRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/groundstation-2019-05-23/GetSatellite
func (c *Client) GetSatelliteRequest(input *GetSatelliteInput) GetSatelliteRequest {
	op := &aws.Operation{
		Name:       opGetSatellite,
		HTTPMethod: "GET",
		HTTPPath:   "/satellite/{satelliteId}",
	}

	if input == nil {
		input = &GetSatelliteInput{}
	}

	req := c.newRequest(op, input, &GetSatelliteOutput{})
	return GetSatelliteRequest{Request: req, Input: input, Copy: c.GetSatelliteRequest}
}

// GetSatelliteRequest is the request type for the
// GetSatellite API operation.
type GetSatelliteRequest struct {
	*aws.Request
	Input *GetSatelliteInput
	Copy  func(*GetSatelliteInput) GetSatelliteRequest
}

// Send marshals and sends the GetSatellite API request.
func (r GetSatelliteRequest) Send(ctx context.Context) (*GetSatelliteResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetSatelliteResponse{
		GetSatelliteOutput: r.Request.Data.(*GetSatelliteOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetSatelliteResponse is the response type for the
// GetSatellite API operation.
type GetSatelliteResponse struct {
	*GetSatelliteOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetSatellite request.
func (r *GetSatelliteResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
