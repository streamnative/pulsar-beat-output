// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ram

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ram-2018-01-04/AcceptResourceShareInvitationRequest
type AcceptResourceShareInvitationInput struct {
	_ struct{} `type:"structure"`

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request.
	ClientToken *string `locationName:"clientToken" type:"string"`

	// The Amazon Resource Name (ARN) of the invitation.
	//
	// ResourceShareInvitationArn is a required field
	ResourceShareInvitationArn *string `locationName:"resourceShareInvitationArn" type:"string" required:"true"`
}

// String returns the string representation
func (s AcceptResourceShareInvitationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AcceptResourceShareInvitationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AcceptResourceShareInvitationInput"}

	if s.ResourceShareInvitationArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceShareInvitationArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s AcceptResourceShareInvitationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ClientToken != nil {
		v := *s.ClientToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "clientToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ResourceShareInvitationArn != nil {
		v := *s.ResourceShareInvitationArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "resourceShareInvitationArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ram-2018-01-04/AcceptResourceShareInvitationResponse
type AcceptResourceShareInvitationOutput struct {
	_ struct{} `type:"structure"`

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request.
	ClientToken *string `locationName:"clientToken" type:"string"`

	// Information about the invitation.
	ResourceShareInvitation *ResourceShareInvitation `locationName:"resourceShareInvitation" type:"structure"`
}

// String returns the string representation
func (s AcceptResourceShareInvitationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s AcceptResourceShareInvitationOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ClientToken != nil {
		v := *s.ClientToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "clientToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ResourceShareInvitation != nil {
		v := s.ResourceShareInvitation

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "resourceShareInvitation", v, metadata)
	}
	return nil
}

const opAcceptResourceShareInvitation = "AcceptResourceShareInvitation"

// AcceptResourceShareInvitationRequest returns a request value for making API operation for
// AWS Resource Access Manager.
//
// Accepts an invitation to a resource share from another AWS account.
//
//    // Example sending a request using AcceptResourceShareInvitationRequest.
//    req := client.AcceptResourceShareInvitationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ram-2018-01-04/AcceptResourceShareInvitation
func (c *Client) AcceptResourceShareInvitationRequest(input *AcceptResourceShareInvitationInput) AcceptResourceShareInvitationRequest {
	op := &aws.Operation{
		Name:       opAcceptResourceShareInvitation,
		HTTPMethod: "POST",
		HTTPPath:   "/acceptresourceshareinvitation",
	}

	if input == nil {
		input = &AcceptResourceShareInvitationInput{}
	}

	req := c.newRequest(op, input, &AcceptResourceShareInvitationOutput{})
	return AcceptResourceShareInvitationRequest{Request: req, Input: input, Copy: c.AcceptResourceShareInvitationRequest}
}

// AcceptResourceShareInvitationRequest is the request type for the
// AcceptResourceShareInvitation API operation.
type AcceptResourceShareInvitationRequest struct {
	*aws.Request
	Input *AcceptResourceShareInvitationInput
	Copy  func(*AcceptResourceShareInvitationInput) AcceptResourceShareInvitationRequest
}

// Send marshals and sends the AcceptResourceShareInvitation API request.
func (r AcceptResourceShareInvitationRequest) Send(ctx context.Context) (*AcceptResourceShareInvitationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AcceptResourceShareInvitationResponse{
		AcceptResourceShareInvitationOutput: r.Request.Data.(*AcceptResourceShareInvitationOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AcceptResourceShareInvitationResponse is the response type for the
// AcceptResourceShareInvitation API operation.
type AcceptResourceShareInvitationResponse struct {
	*AcceptResourceShareInvitationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AcceptResourceShareInvitation request.
func (r *AcceptResourceShareInvitationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
