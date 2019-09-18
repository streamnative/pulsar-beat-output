// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sns

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Input for ListPlatformApplications action.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sns-2010-03-31/ListPlatformApplicationsInput
type ListPlatformApplicationsInput struct {
	_ struct{} `type:"structure"`

	// NextToken string is used when calling ListPlatformApplications action to
	// retrieve additional records that are available after the first page results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListPlatformApplicationsInput) String() string {
	return awsutil.Prettify(s)
}

// Response for ListPlatformApplications action.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sns-2010-03-31/ListPlatformApplicationsResponse
type ListPlatformApplicationsOutput struct {
	_ struct{} `type:"structure"`

	// NextToken string is returned when calling ListPlatformApplications action
	// if additional records are available after the first page results.
	NextToken *string `type:"string"`

	// Platform applications returned when calling ListPlatformApplications action.
	PlatformApplications []PlatformApplication `type:"list"`
}

// String returns the string representation
func (s ListPlatformApplicationsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListPlatformApplications = "ListPlatformApplications"

// ListPlatformApplicationsRequest returns a request value for making API operation for
// Amazon Simple Notification Service.
//
// Lists the platform application objects for the supported push notification
// services, such as APNS and GCM. The results for ListPlatformApplications
// are paginated and return a limited list of applications, up to 100. If additional
// records are available after the first page results, then a NextToken string
// will be returned. To receive the next page, you call ListPlatformApplications
// using the NextToken string received from the previous call. When there are
// no more records to return, NextToken will be null. For more information,
// see Using Amazon SNS Mobile Push Notifications (https://docs.aws.amazon.com/sns/latest/dg/SNSMobilePush.html).
//
// This action is throttled at 15 transactions per second (TPS).
//
//    // Example sending a request using ListPlatformApplicationsRequest.
//    req := client.ListPlatformApplicationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sns-2010-03-31/ListPlatformApplications
func (c *Client) ListPlatformApplicationsRequest(input *ListPlatformApplicationsInput) ListPlatformApplicationsRequest {
	op := &aws.Operation{
		Name:       opListPlatformApplications,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListPlatformApplicationsInput{}
	}

	req := c.newRequest(op, input, &ListPlatformApplicationsOutput{})
	return ListPlatformApplicationsRequest{Request: req, Input: input, Copy: c.ListPlatformApplicationsRequest}
}

// ListPlatformApplicationsRequest is the request type for the
// ListPlatformApplications API operation.
type ListPlatformApplicationsRequest struct {
	*aws.Request
	Input *ListPlatformApplicationsInput
	Copy  func(*ListPlatformApplicationsInput) ListPlatformApplicationsRequest
}

// Send marshals and sends the ListPlatformApplications API request.
func (r ListPlatformApplicationsRequest) Send(ctx context.Context) (*ListPlatformApplicationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListPlatformApplicationsResponse{
		ListPlatformApplicationsOutput: r.Request.Data.(*ListPlatformApplicationsOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListPlatformApplicationsRequestPaginator returns a paginator for ListPlatformApplications.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListPlatformApplicationsRequest(input)
//   p := sns.NewListPlatformApplicationsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListPlatformApplicationsPaginator(req ListPlatformApplicationsRequest) ListPlatformApplicationsPaginator {
	return ListPlatformApplicationsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListPlatformApplicationsInput
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

// ListPlatformApplicationsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListPlatformApplicationsPaginator struct {
	aws.Pager
}

func (p *ListPlatformApplicationsPaginator) CurrentPage() *ListPlatformApplicationsOutput {
	return p.Pager.CurrentPage().(*ListPlatformApplicationsOutput)
}

// ListPlatformApplicationsResponse is the response type for the
// ListPlatformApplications API operation.
type ListPlatformApplicationsResponse struct {
	*ListPlatformApplicationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListPlatformApplications request.
func (r *ListPlatformApplicationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
