// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sfn

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/states-2016-11-23/ListStateMachinesInput
type ListStateMachinesInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results that are returned per call. You can use nextToken
	// to obtain further pages of results. The default is 100 and the maximum allowed
	// page size is 1000. A value of 0 uses the default.
	//
	// This is only an upper limit. The actual number of results returned per call
	// might be fewer than the specified maximum.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// If nextToken is returned, there are more results available. The value of
	// nextToken is a unique pagination token for each page. Make the call again
	// using the returned token to retrieve the next page. Keep all other arguments
	// unchanged. Each pagination token expires after 24 hours. Using an expired
	// pagination token will return an HTTP 400 InvalidToken error.
	NextToken *string `locationName:"nextToken" min:"1" type:"string"`
}

// String returns the string representation
func (s ListStateMachinesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListStateMachinesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListStateMachinesInput"}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/states-2016-11-23/ListStateMachinesOutput
type ListStateMachinesOutput struct {
	_ struct{} `type:"structure"`

	// If nextToken is returned, there are more results available. The value of
	// nextToken is a unique pagination token for each page. Make the call again
	// using the returned token to retrieve the next page. Keep all other arguments
	// unchanged. Each pagination token expires after 24 hours. Using an expired
	// pagination token will return an HTTP 400 InvalidToken error.
	NextToken *string `locationName:"nextToken" min:"1" type:"string"`

	// StateMachines is a required field
	StateMachines []StateMachineListItem `locationName:"stateMachines" type:"list" required:"true"`
}

// String returns the string representation
func (s ListStateMachinesOutput) String() string {
	return awsutil.Prettify(s)
}

const opListStateMachines = "ListStateMachines"

// ListStateMachinesRequest returns a request value for making API operation for
// AWS Step Functions.
//
// Lists the existing state machines.
//
// If nextToken is returned, there are more results available. The value of
// nextToken is a unique pagination token for each page. Make the call again
// using the returned token to retrieve the next page. Keep all other arguments
// unchanged. Each pagination token expires after 24 hours. Using an expired
// pagination token will return an HTTP 400 InvalidToken error.
//
// This operation is eventually consistent. The results are best effort and
// may not reflect very recent updates and changes.
//
//    // Example sending a request using ListStateMachinesRequest.
//    req := client.ListStateMachinesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/states-2016-11-23/ListStateMachines
func (c *Client) ListStateMachinesRequest(input *ListStateMachinesInput) ListStateMachinesRequest {
	op := &aws.Operation{
		Name:       opListStateMachines,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListStateMachinesInput{}
	}

	req := c.newRequest(op, input, &ListStateMachinesOutput{})
	return ListStateMachinesRequest{Request: req, Input: input, Copy: c.ListStateMachinesRequest}
}

// ListStateMachinesRequest is the request type for the
// ListStateMachines API operation.
type ListStateMachinesRequest struct {
	*aws.Request
	Input *ListStateMachinesInput
	Copy  func(*ListStateMachinesInput) ListStateMachinesRequest
}

// Send marshals and sends the ListStateMachines API request.
func (r ListStateMachinesRequest) Send(ctx context.Context) (*ListStateMachinesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListStateMachinesResponse{
		ListStateMachinesOutput: r.Request.Data.(*ListStateMachinesOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListStateMachinesRequestPaginator returns a paginator for ListStateMachines.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListStateMachinesRequest(input)
//   p := sfn.NewListStateMachinesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListStateMachinesPaginator(req ListStateMachinesRequest) ListStateMachinesPaginator {
	return ListStateMachinesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListStateMachinesInput
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

// ListStateMachinesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListStateMachinesPaginator struct {
	aws.Pager
}

func (p *ListStateMachinesPaginator) CurrentPage() *ListStateMachinesOutput {
	return p.Pager.CurrentPage().(*ListStateMachinesOutput)
}

// ListStateMachinesResponse is the response type for the
// ListStateMachines API operation.
type ListStateMachinesResponse struct {
	*ListStateMachinesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListStateMachines request.
func (r *ListStateMachinesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
