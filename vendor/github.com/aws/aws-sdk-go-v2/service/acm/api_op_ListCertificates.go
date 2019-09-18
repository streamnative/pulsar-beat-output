// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package acm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/acm-2015-12-08/ListCertificatesRequest
type ListCertificatesInput struct {
	_ struct{} `type:"structure"`

	// Filter the certificate list by status value.
	CertificateStatuses []CertificateStatus `type:"list"`

	// Filter the certificate list. For more information, see the Filters structure.
	Includes *Filters `type:"structure"`

	// Use this parameter when paginating results to specify the maximum number
	// of items to return in the response. If additional items exist beyond the
	// number you specify, the NextToken element is sent in the response. Use this
	// NextToken value in a subsequent request to retrieve additional items.
	MaxItems *int64 `min:"1" type:"integer"`

	// Use this parameter only when paginating results and only in a subsequent
	// request after you receive a response with truncated results. Set it to the
	// value of NextToken from the response you just received.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListCertificatesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListCertificatesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListCertificatesInput"}
	if s.MaxItems != nil && *s.MaxItems < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxItems", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/acm-2015-12-08/ListCertificatesResponse
type ListCertificatesOutput struct {
	_ struct{} `type:"structure"`

	// A list of ACM certificates.
	CertificateSummaryList []CertificateSummary `type:"list"`

	// When the list is truncated, this value is present and contains the value
	// to use for the NextToken parameter in a subsequent pagination request.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListCertificatesOutput) String() string {
	return awsutil.Prettify(s)
}

const opListCertificates = "ListCertificates"

// ListCertificatesRequest returns a request value for making API operation for
// AWS Certificate Manager.
//
// Retrieves a list of certificate ARNs and domain names. You can request that
// only certificates that match a specific status be listed. You can also filter
// by specific attributes of the certificate.
//
//    // Example sending a request using ListCertificatesRequest.
//    req := client.ListCertificatesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/acm-2015-12-08/ListCertificates
func (c *Client) ListCertificatesRequest(input *ListCertificatesInput) ListCertificatesRequest {
	op := &aws.Operation{
		Name:       opListCertificates,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxItems",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListCertificatesInput{}
	}

	req := c.newRequest(op, input, &ListCertificatesOutput{})
	return ListCertificatesRequest{Request: req, Input: input, Copy: c.ListCertificatesRequest}
}

// ListCertificatesRequest is the request type for the
// ListCertificates API operation.
type ListCertificatesRequest struct {
	*aws.Request
	Input *ListCertificatesInput
	Copy  func(*ListCertificatesInput) ListCertificatesRequest
}

// Send marshals and sends the ListCertificates API request.
func (r ListCertificatesRequest) Send(ctx context.Context) (*ListCertificatesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListCertificatesResponse{
		ListCertificatesOutput: r.Request.Data.(*ListCertificatesOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListCertificatesRequestPaginator returns a paginator for ListCertificates.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListCertificatesRequest(input)
//   p := acm.NewListCertificatesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListCertificatesPaginator(req ListCertificatesRequest) ListCertificatesPaginator {
	return ListCertificatesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListCertificatesInput
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

// ListCertificatesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListCertificatesPaginator struct {
	aws.Pager
}

func (p *ListCertificatesPaginator) CurrentPage() *ListCertificatesOutput {
	return p.Pager.CurrentPage().(*ListCertificatesOutput)
}

// ListCertificatesResponse is the response type for the
// ListCertificates API operation.
type ListCertificatesResponse struct {
	*ListCertificatesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListCertificates request.
func (r *ListCertificatesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
