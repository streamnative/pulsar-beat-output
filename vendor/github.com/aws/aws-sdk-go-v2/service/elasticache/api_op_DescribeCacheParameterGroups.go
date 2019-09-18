// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticache

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a DescribeCacheParameterGroups operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticache-2015-02-02/DescribeCacheParameterGroupsMessage
type DescribeCacheParameterGroupsInput struct {
	_ struct{} `type:"structure"`

	// The name of a specific cache parameter group to return details for.
	CacheParameterGroupName *string `type:"string"`

	// An optional marker returned from a prior request. Use this marker for pagination
	// of results from this operation. If this parameter is specified, the response
	// includes only records beyond the marker, up to the value specified by MaxRecords.
	Marker *string `type:"string"`

	// The maximum number of records to include in the response. If more records
	// exist than the specified MaxRecords value, a marker is included in the response
	// so that the remaining results can be retrieved.
	//
	// Default: 100
	//
	// Constraints: minimum 20; maximum 100.
	MaxRecords *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribeCacheParameterGroupsInput) String() string {
	return awsutil.Prettify(s)
}

// Represents the output of a DescribeCacheParameterGroups operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticache-2015-02-02/CacheParameterGroupsMessage
type DescribeCacheParameterGroupsOutput struct {
	_ struct{} `type:"structure"`

	// A list of cache parameter groups. Each element in the list contains detailed
	// information about one cache parameter group.
	CacheParameterGroups []CacheParameterGroup `locationNameList:"CacheParameterGroup" type:"list"`

	// Provides an identifier to allow retrieval of paginated results.
	Marker *string `type:"string"`
}

// String returns the string representation
func (s DescribeCacheParameterGroupsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeCacheParameterGroups = "DescribeCacheParameterGroups"

// DescribeCacheParameterGroupsRequest returns a request value for making API operation for
// Amazon ElastiCache.
//
// Returns a list of cache parameter group descriptions. If a cache parameter
// group name is specified, the list contains only the descriptions for that
// group.
//
//    // Example sending a request using DescribeCacheParameterGroupsRequest.
//    req := client.DescribeCacheParameterGroupsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticache-2015-02-02/DescribeCacheParameterGroups
func (c *Client) DescribeCacheParameterGroupsRequest(input *DescribeCacheParameterGroupsInput) DescribeCacheParameterGroupsRequest {
	op := &aws.Operation{
		Name:       opDescribeCacheParameterGroups,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"Marker"},
			LimitToken:      "MaxRecords",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeCacheParameterGroupsInput{}
	}

	req := c.newRequest(op, input, &DescribeCacheParameterGroupsOutput{})
	return DescribeCacheParameterGroupsRequest{Request: req, Input: input, Copy: c.DescribeCacheParameterGroupsRequest}
}

// DescribeCacheParameterGroupsRequest is the request type for the
// DescribeCacheParameterGroups API operation.
type DescribeCacheParameterGroupsRequest struct {
	*aws.Request
	Input *DescribeCacheParameterGroupsInput
	Copy  func(*DescribeCacheParameterGroupsInput) DescribeCacheParameterGroupsRequest
}

// Send marshals and sends the DescribeCacheParameterGroups API request.
func (r DescribeCacheParameterGroupsRequest) Send(ctx context.Context) (*DescribeCacheParameterGroupsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeCacheParameterGroupsResponse{
		DescribeCacheParameterGroupsOutput: r.Request.Data.(*DescribeCacheParameterGroupsOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeCacheParameterGroupsRequestPaginator returns a paginator for DescribeCacheParameterGroups.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeCacheParameterGroupsRequest(input)
//   p := elasticache.NewDescribeCacheParameterGroupsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeCacheParameterGroupsPaginator(req DescribeCacheParameterGroupsRequest) DescribeCacheParameterGroupsPaginator {
	return DescribeCacheParameterGroupsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeCacheParameterGroupsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// DescribeCacheParameterGroupsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeCacheParameterGroupsPaginator struct {
	aws.Pager
}

func (p *DescribeCacheParameterGroupsPaginator) CurrentPage() *DescribeCacheParameterGroupsOutput {
	return p.Pager.CurrentPage().(*DescribeCacheParameterGroupsOutput)
}

// DescribeCacheParameterGroupsResponse is the response type for the
// DescribeCacheParameterGroups API operation.
type DescribeCacheParameterGroupsResponse struct {
	*DescribeCacheParameterGroupsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeCacheParameterGroups request.
func (r *DescribeCacheParameterGroupsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
