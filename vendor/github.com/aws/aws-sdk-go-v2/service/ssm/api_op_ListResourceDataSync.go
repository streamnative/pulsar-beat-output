// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/ListResourceDataSyncRequest
type ListResourceDataSyncInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of items to return for this call. The call also returns
	// a token that you can specify in a subsequent call to get the next set of
	// results.
	MaxResults *int64 `min:"1" type:"integer"`

	// A token to start the list. Use this token to get the next set of results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListResourceDataSyncInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListResourceDataSyncInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListResourceDataSyncInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/ListResourceDataSyncResult
type ListResourceDataSyncOutput struct {
	_ struct{} `type:"structure"`

	// The token for the next set of items to return. Use this token to get the
	// next set of results.
	NextToken *string `type:"string"`

	// A list of your current Resource Data Sync configurations and their statuses.
	ResourceDataSyncItems []ResourceDataSyncItem `type:"list"`
}

// String returns the string representation
func (s ListResourceDataSyncOutput) String() string {
	return awsutil.Prettify(s)
}

const opListResourceDataSync = "ListResourceDataSync"

// ListResourceDataSyncRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Lists your resource data sync configurations. Includes information about
// the last time a sync attempted to start, the last sync status, and the last
// time a sync successfully completed.
//
// The number of sync configurations might be too large to return using a single
// call to ListResourceDataSync. You can limit the number of sync configurations
// returned by using the MaxResults parameter. To determine whether there are
// more sync configurations to list, check the value of NextToken in the output.
// If there are more sync configurations to list, you can request them by specifying
// the NextToken returned in the call to the parameter of a subsequent call.
//
//    // Example sending a request using ListResourceDataSyncRequest.
//    req := client.ListResourceDataSyncRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/ListResourceDataSync
func (c *Client) ListResourceDataSyncRequest(input *ListResourceDataSyncInput) ListResourceDataSyncRequest {
	op := &aws.Operation{
		Name:       opListResourceDataSync,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListResourceDataSyncInput{}
	}

	req := c.newRequest(op, input, &ListResourceDataSyncOutput{})
	return ListResourceDataSyncRequest{Request: req, Input: input, Copy: c.ListResourceDataSyncRequest}
}

// ListResourceDataSyncRequest is the request type for the
// ListResourceDataSync API operation.
type ListResourceDataSyncRequest struct {
	*aws.Request
	Input *ListResourceDataSyncInput
	Copy  func(*ListResourceDataSyncInput) ListResourceDataSyncRequest
}

// Send marshals and sends the ListResourceDataSync API request.
func (r ListResourceDataSyncRequest) Send(ctx context.Context) (*ListResourceDataSyncResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListResourceDataSyncResponse{
		ListResourceDataSyncOutput: r.Request.Data.(*ListResourceDataSyncOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListResourceDataSyncResponse is the response type for the
// ListResourceDataSync API operation.
type ListResourceDataSyncResponse struct {
	*ListResourceDataSyncOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListResourceDataSync request.
func (r *ListResourceDataSyncResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
