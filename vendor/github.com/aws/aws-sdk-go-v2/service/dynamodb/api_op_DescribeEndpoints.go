// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package dynamodb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/dynamodb-2012-08-10/DescribeEndpointsRequest
type DescribeEndpointsInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DescribeEndpointsInput) String() string {
	return awsutil.Prettify(s)
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/dynamodb-2012-08-10/DescribeEndpointsResponse
type DescribeEndpointsOutput struct {
	_ struct{} `type:"structure"`

	// List of endpoints.
	//
	// Endpoints is a required field
	Endpoints []Endpoint `type:"list" required:"true"`
}

// String returns the string representation
func (s DescribeEndpointsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeEndpoints = "DescribeEndpoints"

// DescribeEndpointsRequest returns a request value for making API operation for
// Amazon DynamoDB.
//
// Returns the regional endpoint information.
//
//    // Example sending a request using DescribeEndpointsRequest.
//    req := client.DescribeEndpointsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dynamodb-2012-08-10/DescribeEndpoints
func (c *Client) DescribeEndpointsRequest(input *DescribeEndpointsInput) DescribeEndpointsRequest {
	op := &aws.Operation{
		Name:       opDescribeEndpoints,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeEndpointsInput{}
	}

	req := c.newRequest(op, input, &DescribeEndpointsOutput{})
	return DescribeEndpointsRequest{Request: req, Input: input, Copy: c.DescribeEndpointsRequest}
}

// DescribeEndpointsRequest is the request type for the
// DescribeEndpoints API operation.
type DescribeEndpointsRequest struct {
	*aws.Request
	Input *DescribeEndpointsInput
	Copy  func(*DescribeEndpointsInput) DescribeEndpointsRequest
}

// Send marshals and sends the DescribeEndpoints API request.
func (r DescribeEndpointsRequest) Send(ctx context.Context) (*DescribeEndpointsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeEndpointsResponse{
		DescribeEndpointsOutput: r.Request.Data.(*DescribeEndpointsOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeEndpointsResponse is the response type for the
// DescribeEndpoints API operation.
type DescribeEndpointsResponse struct {
	*DescribeEndpointsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeEndpoints request.
func (r *DescribeEndpointsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
