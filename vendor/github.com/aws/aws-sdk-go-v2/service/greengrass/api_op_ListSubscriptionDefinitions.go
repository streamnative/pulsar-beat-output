// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package greengrass

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/ListSubscriptionDefinitionsRequest
type ListSubscriptionDefinitionsInput struct {
	_ struct{} `type:"structure"`

	MaxResults *string `location:"querystring" locationName:"MaxResults" type:"string"`

	NextToken *string `location:"querystring" locationName:"NextToken" type:"string"`
}

// String returns the string representation
func (s ListSubscriptionDefinitionsInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListSubscriptionDefinitionsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "MaxResults", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/ListSubscriptionDefinitionsResponse
type ListSubscriptionDefinitionsOutput struct {
	_ struct{} `type:"structure"`

	Definitions []DefinitionInformation `type:"list"`

	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListSubscriptionDefinitionsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListSubscriptionDefinitionsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Definitions != nil {
		v := s.Definitions

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Definitions", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListSubscriptionDefinitions = "ListSubscriptionDefinitions"

// ListSubscriptionDefinitionsRequest returns a request value for making API operation for
// AWS Greengrass.
//
// Retrieves a list of subscription definitions.
//
//    // Example sending a request using ListSubscriptionDefinitionsRequest.
//    req := client.ListSubscriptionDefinitionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/ListSubscriptionDefinitions
func (c *Client) ListSubscriptionDefinitionsRequest(input *ListSubscriptionDefinitionsInput) ListSubscriptionDefinitionsRequest {
	op := &aws.Operation{
		Name:       opListSubscriptionDefinitions,
		HTTPMethod: "GET",
		HTTPPath:   "/greengrass/definition/subscriptions",
	}

	if input == nil {
		input = &ListSubscriptionDefinitionsInput{}
	}

	req := c.newRequest(op, input, &ListSubscriptionDefinitionsOutput{})
	return ListSubscriptionDefinitionsRequest{Request: req, Input: input, Copy: c.ListSubscriptionDefinitionsRequest}
}

// ListSubscriptionDefinitionsRequest is the request type for the
// ListSubscriptionDefinitions API operation.
type ListSubscriptionDefinitionsRequest struct {
	*aws.Request
	Input *ListSubscriptionDefinitionsInput
	Copy  func(*ListSubscriptionDefinitionsInput) ListSubscriptionDefinitionsRequest
}

// Send marshals and sends the ListSubscriptionDefinitions API request.
func (r ListSubscriptionDefinitionsRequest) Send(ctx context.Context) (*ListSubscriptionDefinitionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListSubscriptionDefinitionsResponse{
		ListSubscriptionDefinitionsOutput: r.Request.Data.(*ListSubscriptionDefinitionsOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListSubscriptionDefinitionsResponse is the response type for the
// ListSubscriptionDefinitions API operation.
type ListSubscriptionDefinitionsResponse struct {
	*ListSubscriptionDefinitionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListSubscriptionDefinitions request.
func (r *ListSubscriptionDefinitionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
