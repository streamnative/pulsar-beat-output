// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package dynamodb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/dynamodb-2012-08-10/ListTagsOfResourceInput
type ListTagsOfResourceInput struct {
	_ struct{} `type:"structure"`

	// An optional string that, if supplied, must be copied from the output of a
	// previous call to ListTagOfResource. When provided in this manner, this API
	// fetches the next page of results.
	NextToken *string `type:"string"`

	// The Amazon DynamoDB resource with tags to be listed. This value is an Amazon
	// Resource Name (ARN).
	//
	// ResourceArn is a required field
	ResourceArn *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s ListTagsOfResourceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListTagsOfResourceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListTagsOfResourceInput"}

	if s.ResourceArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceArn"))
	}
	if s.ResourceArn != nil && len(*s.ResourceArn) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceArn", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/dynamodb-2012-08-10/ListTagsOfResourceOutput
type ListTagsOfResourceOutput struct {
	_ struct{} `type:"structure"`

	// If this value is returned, there are additional results to be displayed.
	// To retrieve them, call ListTagsOfResource again, with NextToken set to this
	// value.
	NextToken *string `type:"string"`

	// The tags currently associated with the Amazon DynamoDB resource.
	Tags []Tag `type:"list"`
}

// String returns the string representation
func (s ListTagsOfResourceOutput) String() string {
	return awsutil.Prettify(s)
}

const opListTagsOfResource = "ListTagsOfResource"

// ListTagsOfResourceRequest returns a request value for making API operation for
// Amazon DynamoDB.
//
// List all tags on an Amazon DynamoDB resource. You can call ListTagsOfResource
// up to 10 times per second, per account.
//
// For an overview on tagging DynamoDB resources, see Tagging for DynamoDB (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Tagging.html)
// in the Amazon DynamoDB Developer Guide.
//
//    // Example sending a request using ListTagsOfResourceRequest.
//    req := client.ListTagsOfResourceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dynamodb-2012-08-10/ListTagsOfResource
func (c *Client) ListTagsOfResourceRequest(input *ListTagsOfResourceInput) ListTagsOfResourceRequest {
	op := &aws.Operation{
		Name:       opListTagsOfResource,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListTagsOfResourceInput{}
	}

	req := c.newRequest(op, input, &ListTagsOfResourceOutput{})
	return ListTagsOfResourceRequest{Request: req, Input: input, Copy: c.ListTagsOfResourceRequest}
}

// ListTagsOfResourceRequest is the request type for the
// ListTagsOfResource API operation.
type ListTagsOfResourceRequest struct {
	*aws.Request
	Input *ListTagsOfResourceInput
	Copy  func(*ListTagsOfResourceInput) ListTagsOfResourceRequest
}

// Send marshals and sends the ListTagsOfResource API request.
func (r ListTagsOfResourceRequest) Send(ctx context.Context) (*ListTagsOfResourceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListTagsOfResourceResponse{
		ListTagsOfResourceOutput: r.Request.Data.(*ListTagsOfResourceOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListTagsOfResourceResponse is the response type for the
// ListTagsOfResource API operation.
type ListTagsOfResourceResponse struct {
	*ListTagsOfResourceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListTagsOfResource request.
func (r *ListTagsOfResourceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
