// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package redshift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/DescribeHsmClientCertificatesMessage
type DescribeHsmClientCertificatesInput struct {
	_ struct{} `type:"structure"`

	// The identifier of a specific HSM client certificate for which you want information.
	// If no identifier is specified, information is returned for all HSM client
	// certificates owned by your AWS customer account.
	HsmClientCertificateIdentifier *string `type:"string"`

	// An optional parameter that specifies the starting point to return a set of
	// response records. When the results of a DescribeHsmClientCertificates request
	// exceed the value specified in MaxRecords, AWS returns a value in the Marker
	// field of the response. You can retrieve the next set of response records
	// by providing the returned marker value in the Marker parameter and retrying
	// the request.
	Marker *string `type:"string"`

	// The maximum number of response records to return in each call. If the number
	// of remaining response records exceeds the specified MaxRecords value, a value
	// is returned in a marker field of the response. You can retrieve the next
	// set of records by retrying the command with the returned marker value.
	//
	// Default: 100
	//
	// Constraints: minimum 20, maximum 100.
	MaxRecords *int64 `type:"integer"`

	// A tag key or keys for which you want to return all matching HSM client certificates
	// that are associated with the specified key or keys. For example, suppose
	// that you have HSM client certificates that are tagged with keys called owner
	// and environment. If you specify both of these tag keys in the request, Amazon
	// Redshift returns a response with the HSM client certificates that have either
	// or both of these tag keys associated with them.
	TagKeys []string `locationNameList:"TagKey" type:"list"`

	// A tag value or values for which you want to return all matching HSM client
	// certificates that are associated with the specified tag value or values.
	// For example, suppose that you have HSM client certificates that are tagged
	// with values called admin and test. If you specify both of these tag values
	// in the request, Amazon Redshift returns a response with the HSM client certificates
	// that have either or both of these tag values associated with them.
	TagValues []string `locationNameList:"TagValue" type:"list"`
}

// String returns the string representation
func (s DescribeHsmClientCertificatesInput) String() string {
	return awsutil.Prettify(s)
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/HsmClientCertificateMessage
type DescribeHsmClientCertificatesOutput struct {
	_ struct{} `type:"structure"`

	// A list of the identifiers for one or more HSM client certificates used by
	// Amazon Redshift clusters to store and retrieve database encryption keys in
	// an HSM.
	HsmClientCertificates []HsmClientCertificate `locationNameList:"HsmClientCertificate" type:"list"`

	// A value that indicates the starting point for the next set of response records
	// in a subsequent request. If a value is returned in a response, you can retrieve
	// the next set of records by providing this returned marker value in the Marker
	// parameter and retrying the command. If the Marker field is empty, all response
	// records have been retrieved for the request.
	Marker *string `type:"string"`
}

// String returns the string representation
func (s DescribeHsmClientCertificatesOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeHsmClientCertificates = "DescribeHsmClientCertificates"

// DescribeHsmClientCertificatesRequest returns a request value for making API operation for
// Amazon Redshift.
//
// Returns information about the specified HSM client certificate. If no certificate
// ID is specified, returns information about all the HSM certificates owned
// by your AWS customer account.
//
// If you specify both tag keys and tag values in the same request, Amazon Redshift
// returns all HSM client certificates that match any combination of the specified
// keys and values. For example, if you have owner and environment for tag keys,
// and admin and test for tag values, all HSM client certificates that have
// any combination of those values are returned.
//
// If both tag keys and values are omitted from the request, HSM client certificates
// are returned regardless of whether they have tag keys or values associated
// with them.
//
//    // Example sending a request using DescribeHsmClientCertificatesRequest.
//    req := client.DescribeHsmClientCertificatesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/DescribeHsmClientCertificates
func (c *Client) DescribeHsmClientCertificatesRequest(input *DescribeHsmClientCertificatesInput) DescribeHsmClientCertificatesRequest {
	op := &aws.Operation{
		Name:       opDescribeHsmClientCertificates,
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
		input = &DescribeHsmClientCertificatesInput{}
	}

	req := c.newRequest(op, input, &DescribeHsmClientCertificatesOutput{})
	return DescribeHsmClientCertificatesRequest{Request: req, Input: input, Copy: c.DescribeHsmClientCertificatesRequest}
}

// DescribeHsmClientCertificatesRequest is the request type for the
// DescribeHsmClientCertificates API operation.
type DescribeHsmClientCertificatesRequest struct {
	*aws.Request
	Input *DescribeHsmClientCertificatesInput
	Copy  func(*DescribeHsmClientCertificatesInput) DescribeHsmClientCertificatesRequest
}

// Send marshals and sends the DescribeHsmClientCertificates API request.
func (r DescribeHsmClientCertificatesRequest) Send(ctx context.Context) (*DescribeHsmClientCertificatesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeHsmClientCertificatesResponse{
		DescribeHsmClientCertificatesOutput: r.Request.Data.(*DescribeHsmClientCertificatesOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeHsmClientCertificatesRequestPaginator returns a paginator for DescribeHsmClientCertificates.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeHsmClientCertificatesRequest(input)
//   p := redshift.NewDescribeHsmClientCertificatesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeHsmClientCertificatesPaginator(req DescribeHsmClientCertificatesRequest) DescribeHsmClientCertificatesPaginator {
	return DescribeHsmClientCertificatesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeHsmClientCertificatesInput
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

// DescribeHsmClientCertificatesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeHsmClientCertificatesPaginator struct {
	aws.Pager
}

func (p *DescribeHsmClientCertificatesPaginator) CurrentPage() *DescribeHsmClientCertificatesOutput {
	return p.Pager.CurrentPage().(*DescribeHsmClientCertificatesOutput)
}

// DescribeHsmClientCertificatesResponse is the response type for the
// DescribeHsmClientCertificates API operation.
type DescribeHsmClientCertificatesResponse struct {
	*DescribeHsmClientCertificatesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeHsmClientCertificates request.
func (r *DescribeHsmClientCertificatesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
