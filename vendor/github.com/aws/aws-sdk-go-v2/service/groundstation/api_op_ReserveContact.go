// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package groundstation

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/groundstation-2019-05-23/ReserveContactRequest
type ReserveContactInput struct {
	_ struct{} `type:"structure"`

	// EndTime is a required field
	EndTime *time.Time `locationName:"endTime" type:"timestamp" required:"true"`

	// GroundStation is a required field
	GroundStation *string `locationName:"groundStation" type:"string" required:"true"`

	// MissionProfileArn is a required field
	MissionProfileArn *string `locationName:"missionProfileArn" type:"string" required:"true"`

	// SatelliteArn is a required field
	SatelliteArn *string `locationName:"satelliteArn" type:"string" required:"true"`

	// StartTime is a required field
	StartTime *time.Time `locationName:"startTime" type:"timestamp" required:"true"`

	Tags map[string]string `locationName:"tags" type:"map"`
}

// String returns the string representation
func (s ReserveContactInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ReserveContactInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ReserveContactInput"}

	if s.EndTime == nil {
		invalidParams.Add(aws.NewErrParamRequired("EndTime"))
	}

	if s.GroundStation == nil {
		invalidParams.Add(aws.NewErrParamRequired("GroundStation"))
	}

	if s.MissionProfileArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("MissionProfileArn"))
	}

	if s.SatelliteArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("SatelliteArn"))
	}

	if s.StartTime == nil {
		invalidParams.Add(aws.NewErrParamRequired("StartTime"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ReserveContactInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.EndTime != nil {
		v := *s.EndTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "endTime",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.GroundStation != nil {
		v := *s.GroundStation

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "groundStation", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MissionProfileArn != nil {
		v := *s.MissionProfileArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "missionProfileArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SatelliteArn != nil {
		v := *s.SatelliteArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "satelliteArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.StartTime != nil {
		v := *s.StartTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "startTime",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
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

// Please also see https://docs.aws.amazon.com/goto/WebAPI/groundstation-2019-05-23/ContactIdResponse
type ReserveContactOutput struct {
	_ struct{} `type:"structure"`

	ContactId *string `locationName:"contactId" type:"string"`
}

// String returns the string representation
func (s ReserveContactOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ReserveContactOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ContactId != nil {
		v := *s.ContactId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "contactId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opReserveContact = "ReserveContact"

// ReserveContactRequest returns a request value for making API operation for
// AWS Ground Station.
//
// Reserves a contact using specified parameters.
//
//    // Example sending a request using ReserveContactRequest.
//    req := client.ReserveContactRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/groundstation-2019-05-23/ReserveContact
func (c *Client) ReserveContactRequest(input *ReserveContactInput) ReserveContactRequest {
	op := &aws.Operation{
		Name:       opReserveContact,
		HTTPMethod: "POST",
		HTTPPath:   "/contact",
	}

	if input == nil {
		input = &ReserveContactInput{}
	}

	req := c.newRequest(op, input, &ReserveContactOutput{})
	return ReserveContactRequest{Request: req, Input: input, Copy: c.ReserveContactRequest}
}

// ReserveContactRequest is the request type for the
// ReserveContact API operation.
type ReserveContactRequest struct {
	*aws.Request
	Input *ReserveContactInput
	Copy  func(*ReserveContactInput) ReserveContactRequest
}

// Send marshals and sends the ReserveContact API request.
func (r ReserveContactRequest) Send(ctx context.Context) (*ReserveContactResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ReserveContactResponse{
		ReserveContactOutput: r.Request.Data.(*ReserveContactOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ReserveContactResponse is the response type for the
// ReserveContact API operation.
type ReserveContactResponse struct {
	*ReserveContactOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ReserveContact request.
func (r *ReserveContactResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
