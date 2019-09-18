// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesisanalyticsv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesisanalyticsv2-2018-05-23/ListTagsForResourceRequest
type ListTagsForResourceInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the application for which to retrieve tags.
	//
	// ResourceARN is a required field
	ResourceARN *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s ListTagsForResourceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListTagsForResourceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListTagsForResourceInput"}

	if s.ResourceARN == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceARN"))
	}
	if s.ResourceARN != nil && len(*s.ResourceARN) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceARN", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesisanalyticsv2-2018-05-23/ListTagsForResourceResponse
type ListTagsForResourceOutput struct {
	_ struct{} `type:"structure"`

	// The key-value tags assigned to the application.
	Tags []Tag `min:"1" type:"list"`
}

// String returns the string representation
func (s ListTagsForResourceOutput) String() string {
	return awsutil.Prettify(s)
}

const opListTagsForResource = "ListTagsForResource"

// ListTagsForResourceRequest returns a request value for making API operation for
// Amazon Kinesis Analytics.
//
// Retrieves the list of key-value tags assigned to the application.
//
//    // Example sending a request using ListTagsForResourceRequest.
//    req := client.ListTagsForResourceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesisanalyticsv2-2018-05-23/ListTagsForResource
func (c *Client) ListTagsForResourceRequest(input *ListTagsForResourceInput) ListTagsForResourceRequest {
	op := &aws.Operation{
		Name:       opListTagsForResource,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListTagsForResourceInput{}
	}

	req := c.newRequest(op, input, &ListTagsForResourceOutput{})
	return ListTagsForResourceRequest{Request: req, Input: input, Copy: c.ListTagsForResourceRequest}
}

// ListTagsForResourceRequest is the request type for the
// ListTagsForResource API operation.
type ListTagsForResourceRequest struct {
	*aws.Request
	Input *ListTagsForResourceInput
	Copy  func(*ListTagsForResourceInput) ListTagsForResourceRequest
}

// Send marshals and sends the ListTagsForResource API request.
func (r ListTagsForResourceRequest) Send(ctx context.Context) (*ListTagsForResourceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListTagsForResourceResponse{
		ListTagsForResourceOutput: r.Request.Data.(*ListTagsForResourceOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListTagsForResourceResponse is the response type for the
// ListTagsForResource API operation.
type ListTagsForResourceResponse struct {
	*ListTagsForResourceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListTagsForResource request.
func (r *ListTagsForResourceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
