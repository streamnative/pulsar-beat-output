// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package migrationhub

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/AWSMigrationHub-2017-05-31/ListDiscoveredResourcesRequest
type ListDiscoveredResourcesInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results returned per page.
	MaxResults *int64 `min:"1" type:"integer"`

	// The name of the MigrationTask.
	//
	// MigrationTaskName is a required field
	MigrationTaskName *string `min:"1" type:"string" required:"true"`

	// If a NextToken was returned by a previous call, there are more results available.
	// To retrieve the next page of results, make the call again using the returned
	// token in NextToken.
	NextToken *string `type:"string"`

	// The name of the ProgressUpdateStream.
	//
	// ProgressUpdateStream is a required field
	ProgressUpdateStream *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s ListDiscoveredResourcesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListDiscoveredResourcesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListDiscoveredResourcesInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if s.MigrationTaskName == nil {
		invalidParams.Add(aws.NewErrParamRequired("MigrationTaskName"))
	}
	if s.MigrationTaskName != nil && len(*s.MigrationTaskName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("MigrationTaskName", 1))
	}

	if s.ProgressUpdateStream == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProgressUpdateStream"))
	}
	if s.ProgressUpdateStream != nil && len(*s.ProgressUpdateStream) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ProgressUpdateStream", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/AWSMigrationHub-2017-05-31/ListDiscoveredResourcesResult
type ListDiscoveredResourcesOutput struct {
	_ struct{} `type:"structure"`

	// Returned list of discovered resources associated with the given MigrationTask.
	DiscoveredResourceList []DiscoveredResource `type:"list"`

	// If there are more discovered resources than the max result, return the next
	// token to be passed to the next call as a bookmark of where to start from.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListDiscoveredResourcesOutput) String() string {
	return awsutil.Prettify(s)
}

const opListDiscoveredResources = "ListDiscoveredResources"

// ListDiscoveredResourcesRequest returns a request value for making API operation for
// AWS Migration Hub.
//
// Lists discovered resources associated with the given MigrationTask.
//
//    // Example sending a request using ListDiscoveredResourcesRequest.
//    req := client.ListDiscoveredResourcesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/AWSMigrationHub-2017-05-31/ListDiscoveredResources
func (c *Client) ListDiscoveredResourcesRequest(input *ListDiscoveredResourcesInput) ListDiscoveredResourcesRequest {
	op := &aws.Operation{
		Name:       opListDiscoveredResources,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListDiscoveredResourcesInput{}
	}

	req := c.newRequest(op, input, &ListDiscoveredResourcesOutput{})
	return ListDiscoveredResourcesRequest{Request: req, Input: input, Copy: c.ListDiscoveredResourcesRequest}
}

// ListDiscoveredResourcesRequest is the request type for the
// ListDiscoveredResources API operation.
type ListDiscoveredResourcesRequest struct {
	*aws.Request
	Input *ListDiscoveredResourcesInput
	Copy  func(*ListDiscoveredResourcesInput) ListDiscoveredResourcesRequest
}

// Send marshals and sends the ListDiscoveredResources API request.
func (r ListDiscoveredResourcesRequest) Send(ctx context.Context) (*ListDiscoveredResourcesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListDiscoveredResourcesResponse{
		ListDiscoveredResourcesOutput: r.Request.Data.(*ListDiscoveredResourcesOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListDiscoveredResourcesResponse is the response type for the
// ListDiscoveredResources API operation.
type ListDiscoveredResourcesResponse struct {
	*ListDiscoveredResourcesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListDiscoveredResources request.
func (r *ListDiscoveredResourcesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
