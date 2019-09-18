// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/ListNotebookInstanceLifecycleConfigsInput
type ListNotebookInstanceLifecycleConfigsInput struct {
	_ struct{} `type:"structure"`

	// A filter that returns only lifecycle configurations that were created after
	// the specified time (timestamp).
	CreationTimeAfter *time.Time `type:"timestamp"`

	// A filter that returns only lifecycle configurations that were created before
	// the specified time (timestamp).
	CreationTimeBefore *time.Time `type:"timestamp"`

	// A filter that returns only lifecycle configurations that were modified after
	// the specified time (timestamp).
	LastModifiedTimeAfter *time.Time `type:"timestamp"`

	// A filter that returns only lifecycle configurations that were modified before
	// the specified time (timestamp).
	LastModifiedTimeBefore *time.Time `type:"timestamp"`

	// The maximum number of lifecycle configurations to return in the response.
	MaxResults *int64 `min:"1" type:"integer"`

	// A string in the lifecycle configuration name. This filter returns only lifecycle
	// configurations whose name contains the specified string.
	NameContains *string `type:"string"`

	// If the result of a ListNotebookInstanceLifecycleConfigs request was truncated,
	// the response includes a NextToken. To get the next set of lifecycle configurations,
	// use the token in the next request.
	NextToken *string `type:"string"`

	// Sorts the list of results. The default is CreationTime.
	SortBy NotebookInstanceLifecycleConfigSortKey `type:"string" enum:"true"`

	// The sort order for results.
	SortOrder NotebookInstanceLifecycleConfigSortOrder `type:"string" enum:"true"`
}

// String returns the string representation
func (s ListNotebookInstanceLifecycleConfigsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListNotebookInstanceLifecycleConfigsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListNotebookInstanceLifecycleConfigsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/ListNotebookInstanceLifecycleConfigsOutput
type ListNotebookInstanceLifecycleConfigsOutput struct {
	_ struct{} `type:"structure"`

	// If the response is truncated, Amazon SageMaker returns this token. To get
	// the next set of lifecycle configurations, use it in the next request.
	NextToken *string `type:"string"`

	// An array of NotebookInstanceLifecycleConfiguration objects, each listing
	// a lifecycle configuration.
	NotebookInstanceLifecycleConfigs []NotebookInstanceLifecycleConfigSummary `type:"list"`
}

// String returns the string representation
func (s ListNotebookInstanceLifecycleConfigsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListNotebookInstanceLifecycleConfigs = "ListNotebookInstanceLifecycleConfigs"

// ListNotebookInstanceLifecycleConfigsRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// Lists notebook instance lifestyle configurations created with the CreateNotebookInstanceLifecycleConfig
// API.
//
//    // Example sending a request using ListNotebookInstanceLifecycleConfigsRequest.
//    req := client.ListNotebookInstanceLifecycleConfigsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/ListNotebookInstanceLifecycleConfigs
func (c *Client) ListNotebookInstanceLifecycleConfigsRequest(input *ListNotebookInstanceLifecycleConfigsInput) ListNotebookInstanceLifecycleConfigsRequest {
	op := &aws.Operation{
		Name:       opListNotebookInstanceLifecycleConfigs,
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
		input = &ListNotebookInstanceLifecycleConfigsInput{}
	}

	req := c.newRequest(op, input, &ListNotebookInstanceLifecycleConfigsOutput{})
	return ListNotebookInstanceLifecycleConfigsRequest{Request: req, Input: input, Copy: c.ListNotebookInstanceLifecycleConfigsRequest}
}

// ListNotebookInstanceLifecycleConfigsRequest is the request type for the
// ListNotebookInstanceLifecycleConfigs API operation.
type ListNotebookInstanceLifecycleConfigsRequest struct {
	*aws.Request
	Input *ListNotebookInstanceLifecycleConfigsInput
	Copy  func(*ListNotebookInstanceLifecycleConfigsInput) ListNotebookInstanceLifecycleConfigsRequest
}

// Send marshals and sends the ListNotebookInstanceLifecycleConfigs API request.
func (r ListNotebookInstanceLifecycleConfigsRequest) Send(ctx context.Context) (*ListNotebookInstanceLifecycleConfigsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListNotebookInstanceLifecycleConfigsResponse{
		ListNotebookInstanceLifecycleConfigsOutput: r.Request.Data.(*ListNotebookInstanceLifecycleConfigsOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListNotebookInstanceLifecycleConfigsRequestPaginator returns a paginator for ListNotebookInstanceLifecycleConfigs.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListNotebookInstanceLifecycleConfigsRequest(input)
//   p := sagemaker.NewListNotebookInstanceLifecycleConfigsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListNotebookInstanceLifecycleConfigsPaginator(req ListNotebookInstanceLifecycleConfigsRequest) ListNotebookInstanceLifecycleConfigsPaginator {
	return ListNotebookInstanceLifecycleConfigsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListNotebookInstanceLifecycleConfigsInput
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

// ListNotebookInstanceLifecycleConfigsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListNotebookInstanceLifecycleConfigsPaginator struct {
	aws.Pager
}

func (p *ListNotebookInstanceLifecycleConfigsPaginator) CurrentPage() *ListNotebookInstanceLifecycleConfigsOutput {
	return p.Pager.CurrentPage().(*ListNotebookInstanceLifecycleConfigsOutput)
}

// ListNotebookInstanceLifecycleConfigsResponse is the response type for the
// ListNotebookInstanceLifecycleConfigs API operation.
type ListNotebookInstanceLifecycleConfigsResponse struct {
	*ListNotebookInstanceLifecycleConfigsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListNotebookInstanceLifecycleConfigs request.
func (r *ListNotebookInstanceLifecycleConfigsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
