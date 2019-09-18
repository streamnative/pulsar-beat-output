// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53resolver

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53resolver-2018-04-01/ListResolverEndpointsRequest
type ListResolverEndpointsInput struct {
	_ struct{} `type:"structure"`

	// An optional specification to return a subset of resolver endpoints, such
	// as all inbound resolver endpoints.
	//
	// If you submit a second or subsequent ListResolverEndpoints request and specify
	// the NextToken parameter, you must use the same values for Filters, if any,
	// as in the previous request.
	Filters []Filter `type:"list"`

	// The maximum number of resolver endpoints that you want to return in the response
	// to a ListResolverEndpoints request. If you don't specify a value for MaxResults,
	// Resolver returns up to 100 resolver endpoints.
	MaxResults *int64 `min:"1" type:"integer"`

	// For the first ListResolverEndpoints request, omit this value.
	//
	// If you have more than MaxResults resolver endpoints, you can submit another
	// ListResolverEndpoints request to get the next group of resolver endpoints.
	// In the next request, specify the value of NextToken from the previous response.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListResolverEndpointsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListResolverEndpointsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListResolverEndpointsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.Filters != nil {
		for i, v := range s.Filters {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Filters", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53resolver-2018-04-01/ListResolverEndpointsResponse
type ListResolverEndpointsOutput struct {
	_ struct{} `type:"structure"`

	// The value that you specified for MaxResults in the request.
	MaxResults *int64 `min:"1" type:"integer"`

	// If more than MaxResults IP addresses match the specified criteria, you can
	// submit another ListResolverEndpoint request to get the next group of results.
	// In the next request, specify the value of NextToken from the previous response.
	NextToken *string `type:"string"`

	// The resolver endpoints that were created by using the current AWS account,
	// and that match the specified filters, if any.
	ResolverEndpoints []ResolverEndpoint `type:"list"`
}

// String returns the string representation
func (s ListResolverEndpointsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListResolverEndpoints = "ListResolverEndpoints"

// ListResolverEndpointsRequest returns a request value for making API operation for
// Amazon Route 53 Resolver.
//
// Lists all the resolver endpoints that were created using the current AWS
// account.
//
//    // Example sending a request using ListResolverEndpointsRequest.
//    req := client.ListResolverEndpointsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53resolver-2018-04-01/ListResolverEndpoints
func (c *Client) ListResolverEndpointsRequest(input *ListResolverEndpointsInput) ListResolverEndpointsRequest {
	op := &aws.Operation{
		Name:       opListResolverEndpoints,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListResolverEndpointsInput{}
	}

	req := c.newRequest(op, input, &ListResolverEndpointsOutput{})
	return ListResolverEndpointsRequest{Request: req, Input: input, Copy: c.ListResolverEndpointsRequest}
}

// ListResolverEndpointsRequest is the request type for the
// ListResolverEndpoints API operation.
type ListResolverEndpointsRequest struct {
	*aws.Request
	Input *ListResolverEndpointsInput
	Copy  func(*ListResolverEndpointsInput) ListResolverEndpointsRequest
}

// Send marshals and sends the ListResolverEndpoints API request.
func (r ListResolverEndpointsRequest) Send(ctx context.Context) (*ListResolverEndpointsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListResolverEndpointsResponse{
		ListResolverEndpointsOutput: r.Request.Data.(*ListResolverEndpointsOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListResolverEndpointsRequestPaginator returns a paginator for ListResolverEndpoints.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListResolverEndpointsRequest(input)
//   p := route53resolver.NewListResolverEndpointsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListResolverEndpointsPaginator(req ListResolverEndpointsRequest) ListResolverEndpointsPaginator {
	return ListResolverEndpointsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListResolverEndpointsInput
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

// ListResolverEndpointsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListResolverEndpointsPaginator struct {
	aws.Pager
}

func (p *ListResolverEndpointsPaginator) CurrentPage() *ListResolverEndpointsOutput {
	return p.Pager.CurrentPage().(*ListResolverEndpointsOutput)
}

// ListResolverEndpointsResponse is the response type for the
// ListResolverEndpoints API operation.
type ListResolverEndpointsResponse struct {
	*ListResolverEndpointsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListResolverEndpoints request.
func (r *ListResolverEndpointsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
