// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/DeleteVoiceChannelRequest
type DeleteVoiceChannelInput struct {
	_ struct{} `type:"structure"`

	// ApplicationId is a required field
	ApplicationId *string `location:"uri" locationName:"application-id" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteVoiceChannelInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteVoiceChannelInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteVoiceChannelInput"}

	if s.ApplicationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteVoiceChannelInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ApplicationId != nil {
		v := *s.ApplicationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "application-id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/DeleteVoiceChannelResponse
type DeleteVoiceChannelOutput struct {
	_ struct{} `type:"structure" payload:"VoiceChannelResponse"`

	// Provides information about the status and settings of the voice channel for
	// an application.
	//
	// VoiceChannelResponse is a required field
	VoiceChannelResponse *VoiceChannelResponse `type:"structure" required:"true"`
}

// String returns the string representation
func (s DeleteVoiceChannelOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteVoiceChannelOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.VoiceChannelResponse != nil {
		v := s.VoiceChannelResponse

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "VoiceChannelResponse", v, metadata)
	}
	return nil
}

const opDeleteVoiceChannel = "DeleteVoiceChannel"

// DeleteVoiceChannelRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Disables the voice channel for an application and deletes any existing settings
// for the channel.
//
//    // Example sending a request using DeleteVoiceChannelRequest.
//    req := client.DeleteVoiceChannelRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/DeleteVoiceChannel
func (c *Client) DeleteVoiceChannelRequest(input *DeleteVoiceChannelInput) DeleteVoiceChannelRequest {
	op := &aws.Operation{
		Name:       opDeleteVoiceChannel,
		HTTPMethod: "DELETE",
		HTTPPath:   "/v1/apps/{application-id}/channels/voice",
	}

	if input == nil {
		input = &DeleteVoiceChannelInput{}
	}

	req := c.newRequest(op, input, &DeleteVoiceChannelOutput{})
	return DeleteVoiceChannelRequest{Request: req, Input: input, Copy: c.DeleteVoiceChannelRequest}
}

// DeleteVoiceChannelRequest is the request type for the
// DeleteVoiceChannel API operation.
type DeleteVoiceChannelRequest struct {
	*aws.Request
	Input *DeleteVoiceChannelInput
	Copy  func(*DeleteVoiceChannelInput) DeleteVoiceChannelRequest
}

// Send marshals and sends the DeleteVoiceChannel API request.
func (r DeleteVoiceChannelRequest) Send(ctx context.Context) (*DeleteVoiceChannelResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteVoiceChannelResponse{
		DeleteVoiceChannelOutput: r.Request.Data.(*DeleteVoiceChannelOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteVoiceChannelResponse is the response type for the
// DeleteVoiceChannel API operation.
type DeleteVoiceChannelResponse struct {
	*DeleteVoiceChannelOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteVoiceChannel request.
func (r *DeleteVoiceChannelResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
