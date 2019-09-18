// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediaconnect

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Adds outputs to an existing flow. You can create up to 20 outputs per flow.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediaconnect-2018-11-14/AddFlowOutputsRequest
type AddFlowOutputsInput struct {
	_ struct{} `type:"structure"`

	// FlowArn is a required field
	FlowArn *string `location:"uri" locationName:"flowArn" type:"string" required:"true"`

	// A list of outputs that you want to add.
	//
	// Outputs is a required field
	Outputs []AddOutputRequest `locationName:"outputs" type:"list" required:"true"`
}

// String returns the string representation
func (s AddFlowOutputsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AddFlowOutputsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AddFlowOutputsInput"}

	if s.FlowArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("FlowArn"))
	}

	if s.Outputs == nil {
		invalidParams.Add(aws.NewErrParamRequired("Outputs"))
	}
	if s.Outputs != nil {
		for i, v := range s.Outputs {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Outputs", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s AddFlowOutputsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Outputs != nil {
		v := s.Outputs

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "outputs", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.FlowArn != nil {
		v := *s.FlowArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "flowArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// The result of a successful AddOutput request. The response includes the details
// of the newly added outputs.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediaconnect-2018-11-14/AddFlowOutputsResponse
type AddFlowOutputsOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the flow that these outputs were added to.
	FlowArn *string `locationName:"flowArn" type:"string"`

	// The details of the newly added outputs.
	Outputs []Output `locationName:"outputs" type:"list"`
}

// String returns the string representation
func (s AddFlowOutputsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s AddFlowOutputsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.FlowArn != nil {
		v := *s.FlowArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "flowArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Outputs != nil {
		v := s.Outputs

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "outputs", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opAddFlowOutputs = "AddFlowOutputs"

// AddFlowOutputsRequest returns a request value for making API operation for
// AWS MediaConnect.
//
// Adds outputs to an existing flow. You can create up to 20 outputs per flow.
//
//    // Example sending a request using AddFlowOutputsRequest.
//    req := client.AddFlowOutputsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediaconnect-2018-11-14/AddFlowOutputs
func (c *Client) AddFlowOutputsRequest(input *AddFlowOutputsInput) AddFlowOutputsRequest {
	op := &aws.Operation{
		Name:       opAddFlowOutputs,
		HTTPMethod: "POST",
		HTTPPath:   "/v1/flows/{flowArn}/outputs",
	}

	if input == nil {
		input = &AddFlowOutputsInput{}
	}

	req := c.newRequest(op, input, &AddFlowOutputsOutput{})
	return AddFlowOutputsRequest{Request: req, Input: input, Copy: c.AddFlowOutputsRequest}
}

// AddFlowOutputsRequest is the request type for the
// AddFlowOutputs API operation.
type AddFlowOutputsRequest struct {
	*aws.Request
	Input *AddFlowOutputsInput
	Copy  func(*AddFlowOutputsInput) AddFlowOutputsRequest
}

// Send marshals and sends the AddFlowOutputs API request.
func (r AddFlowOutputsRequest) Send(ctx context.Context) (*AddFlowOutputsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AddFlowOutputsResponse{
		AddFlowOutputsOutput: r.Request.Data.(*AddFlowOutputsOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AddFlowOutputsResponse is the response type for the
// AddFlowOutputs API operation.
type AddFlowOutputsResponse struct {
	*AddFlowOutputsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AddFlowOutputs request.
func (r *AddFlowOutputsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
