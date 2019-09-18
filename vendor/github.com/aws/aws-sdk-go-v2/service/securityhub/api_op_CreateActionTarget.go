// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package securityhub

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/securityhub-2018-10-26/CreateActionTargetRequest
type CreateActionTargetInput struct {
	_ struct{} `type:"structure"`

	// The description for the custom action target.
	//
	// Description is a required field
	Description *string `type:"string" required:"true"`

	// The ID for the custom action target.
	//
	// Id is a required field
	Id *string `type:"string" required:"true"`

	// The name of the custom action target.
	//
	// Name is a required field
	Name *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateActionTargetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateActionTargetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateActionTargetInput"}

	if s.Description == nil {
		invalidParams.Add(aws.NewErrParamRequired("Description"))
	}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateActionTargetInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/securityhub-2018-10-26/CreateActionTargetResponse
type CreateActionTargetOutput struct {
	_ struct{} `type:"structure"`

	// The ARN for the custom action target.
	//
	// ActionTargetArn is a required field
	ActionTargetArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateActionTargetOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateActionTargetOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ActionTargetArn != nil {
		v := *s.ActionTargetArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ActionTargetArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opCreateActionTarget = "CreateActionTarget"

// CreateActionTargetRequest returns a request value for making API operation for
// AWS SecurityHub.
//
// Creates a custom action target in Security Hub. You can use custom actions
// on findings and insights in Security Hub to trigger target actions in Amazon
// CloudWatch Events.
//
//    // Example sending a request using CreateActionTargetRequest.
//    req := client.CreateActionTargetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/securityhub-2018-10-26/CreateActionTarget
func (c *Client) CreateActionTargetRequest(input *CreateActionTargetInput) CreateActionTargetRequest {
	op := &aws.Operation{
		Name:       opCreateActionTarget,
		HTTPMethod: "POST",
		HTTPPath:   "/actionTargets",
	}

	if input == nil {
		input = &CreateActionTargetInput{}
	}

	req := c.newRequest(op, input, &CreateActionTargetOutput{})
	return CreateActionTargetRequest{Request: req, Input: input, Copy: c.CreateActionTargetRequest}
}

// CreateActionTargetRequest is the request type for the
// CreateActionTarget API operation.
type CreateActionTargetRequest struct {
	*aws.Request
	Input *CreateActionTargetInput
	Copy  func(*CreateActionTargetInput) CreateActionTargetRequest
}

// Send marshals and sends the CreateActionTarget API request.
func (r CreateActionTargetRequest) Send(ctx context.Context) (*CreateActionTargetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateActionTargetResponse{
		CreateActionTargetOutput: r.Request.Data.(*CreateActionTargetOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateActionTargetResponse is the response type for the
// CreateActionTarget API operation.
type CreateActionTargetResponse struct {
	*CreateActionTargetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateActionTarget request.
func (r *CreateActionTargetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
