// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package fms

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/fms-2018-01-01/ListPoliciesRequest
type ListPoliciesInput struct {
	_ struct{} `type:"structure"`

	// Specifies the number of PolicySummary objects that you want AWS Firewall
	// Manager to return for this request. If you have more PolicySummary objects
	// than the number that you specify for MaxResults, the response includes a
	// NextToken value that you can use to get another batch of PolicySummary objects.
	MaxResults *int64 `min:"1" type:"integer"`

	// If you specify a value for MaxResults and you have more PolicySummary objects
	// than the number that you specify for MaxResults, AWS Firewall Manager returns
	// a NextToken value in the response that allows you to list another group of
	// PolicySummary objects. For the second and subsequent ListPolicies requests,
	// specify the value of NextToken from the previous response to get information
	// about another batch of PolicySummary objects.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListPoliciesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListPoliciesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListPoliciesInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/fms-2018-01-01/ListPoliciesResponse
type ListPoliciesOutput struct {
	_ struct{} `type:"structure"`

	// If you have more PolicySummary objects than the number that you specified
	// for MaxResults in the request, the response includes a NextToken value. To
	// list more PolicySummary objects, submit another ListPolicies request, and
	// specify the NextToken value from the response in the NextToken value in the
	// next request.
	NextToken *string `min:"1" type:"string"`

	// An array of PolicySummary objects.
	PolicyList []PolicySummary `type:"list"`
}

// String returns the string representation
func (s ListPoliciesOutput) String() string {
	return awsutil.Prettify(s)
}

const opListPolicies = "ListPolicies"

// ListPoliciesRequest returns a request value for making API operation for
// Firewall Management Service.
//
// Returns an array of PolicySummary objects in the response.
//
//    // Example sending a request using ListPoliciesRequest.
//    req := client.ListPoliciesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/fms-2018-01-01/ListPolicies
func (c *Client) ListPoliciesRequest(input *ListPoliciesInput) ListPoliciesRequest {
	op := &aws.Operation{
		Name:       opListPolicies,
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
		input = &ListPoliciesInput{}
	}

	req := c.newRequest(op, input, &ListPoliciesOutput{})
	return ListPoliciesRequest{Request: req, Input: input, Copy: c.ListPoliciesRequest}
}

// ListPoliciesRequest is the request type for the
// ListPolicies API operation.
type ListPoliciesRequest struct {
	*aws.Request
	Input *ListPoliciesInput
	Copy  func(*ListPoliciesInput) ListPoliciesRequest
}

// Send marshals and sends the ListPolicies API request.
func (r ListPoliciesRequest) Send(ctx context.Context) (*ListPoliciesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListPoliciesResponse{
		ListPoliciesOutput: r.Request.Data.(*ListPoliciesOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListPoliciesRequestPaginator returns a paginator for ListPolicies.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListPoliciesRequest(input)
//   p := fms.NewListPoliciesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListPoliciesPaginator(req ListPoliciesRequest) ListPoliciesPaginator {
	return ListPoliciesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListPoliciesInput
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

// ListPoliciesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListPoliciesPaginator struct {
	aws.Pager
}

func (p *ListPoliciesPaginator) CurrentPage() *ListPoliciesOutput {
	return p.Pager.CurrentPage().(*ListPoliciesOutput)
}

// ListPoliciesResponse is the response type for the
// ListPolicies API operation.
type ListPoliciesResponse struct {
	*ListPoliciesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListPolicies request.
func (r *ListPoliciesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
