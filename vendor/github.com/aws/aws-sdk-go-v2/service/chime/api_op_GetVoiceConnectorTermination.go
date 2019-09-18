// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/GetVoiceConnectorTerminationRequest
type GetVoiceConnectorTerminationInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Chime Voice Connector ID.
	//
	// VoiceConnectorId is a required field
	VoiceConnectorId *string `location:"uri" locationName:"voiceConnectorId" type:"string" required:"true"`
}

// String returns the string representation
func (s GetVoiceConnectorTerminationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetVoiceConnectorTerminationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetVoiceConnectorTerminationInput"}

	if s.VoiceConnectorId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VoiceConnectorId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetVoiceConnectorTerminationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.VoiceConnectorId != nil {
		v := *s.VoiceConnectorId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "voiceConnectorId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/GetVoiceConnectorTerminationResponse
type GetVoiceConnectorTerminationOutput struct {
	_ struct{} `type:"structure"`

	// The termination setting details.
	Termination *Termination `type:"structure"`
}

// String returns the string representation
func (s GetVoiceConnectorTerminationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetVoiceConnectorTerminationOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Termination != nil {
		v := s.Termination

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Termination", v, metadata)
	}
	return nil
}

const opGetVoiceConnectorTermination = "GetVoiceConnectorTermination"

// GetVoiceConnectorTerminationRequest returns a request value for making API operation for
// Amazon Chime.
//
// Retrieves termination setting details for the specified Amazon Chime Voice
// Connector.
//
//    // Example sending a request using GetVoiceConnectorTerminationRequest.
//    req := client.GetVoiceConnectorTerminationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/GetVoiceConnectorTermination
func (c *Client) GetVoiceConnectorTerminationRequest(input *GetVoiceConnectorTerminationInput) GetVoiceConnectorTerminationRequest {
	op := &aws.Operation{
		Name:       opGetVoiceConnectorTermination,
		HTTPMethod: "GET",
		HTTPPath:   "/voice-connectors/{voiceConnectorId}/termination",
	}

	if input == nil {
		input = &GetVoiceConnectorTerminationInput{}
	}

	req := c.newRequest(op, input, &GetVoiceConnectorTerminationOutput{})
	return GetVoiceConnectorTerminationRequest{Request: req, Input: input, Copy: c.GetVoiceConnectorTerminationRequest}
}

// GetVoiceConnectorTerminationRequest is the request type for the
// GetVoiceConnectorTermination API operation.
type GetVoiceConnectorTerminationRequest struct {
	*aws.Request
	Input *GetVoiceConnectorTerminationInput
	Copy  func(*GetVoiceConnectorTerminationInput) GetVoiceConnectorTerminationRequest
}

// Send marshals and sends the GetVoiceConnectorTermination API request.
func (r GetVoiceConnectorTerminationRequest) Send(ctx context.Context) (*GetVoiceConnectorTerminationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetVoiceConnectorTerminationResponse{
		GetVoiceConnectorTerminationOutput: r.Request.Data.(*GetVoiceConnectorTerminationOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetVoiceConnectorTerminationResponse is the response type for the
// GetVoiceConnectorTermination API operation.
type GetVoiceConnectorTerminationResponse struct {
	*GetVoiceConnectorTerminationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetVoiceConnectorTermination request.
func (r *GetVoiceConnectorTerminationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
