// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ses

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/GetSendQuotaInput
type GetSendQuotaInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s GetSendQuotaInput) String() string {
	return awsutil.Prettify(s)
}

// Represents your Amazon SES daily sending quota, maximum send rate, and the
// number of emails you have sent in the last 24 hours.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/GetSendQuotaResponse
type GetSendQuotaOutput struct {
	_ struct{} `type:"structure"`

	// The maximum number of emails the user is allowed to send in a 24-hour interval.
	// A value of -1 signifies an unlimited quota.
	Max24HourSend *float64 `type:"double"`

	// The maximum number of emails that Amazon SES can accept from the user's account
	// per second.
	//
	// The rate at which Amazon SES accepts the user's messages might be less than
	// the maximum send rate.
	MaxSendRate *float64 `type:"double"`

	// The number of emails sent during the previous 24 hours.
	SentLast24Hours *float64 `type:"double"`
}

// String returns the string representation
func (s GetSendQuotaOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetSendQuota = "GetSendQuota"

// GetSendQuotaRequest returns a request value for making API operation for
// Amazon Simple Email Service.
//
// Provides the sending limits for the Amazon SES account.
//
// You can execute this operation no more than once per second.
//
//    // Example sending a request using GetSendQuotaRequest.
//    req := client.GetSendQuotaRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/GetSendQuota
func (c *Client) GetSendQuotaRequest(input *GetSendQuotaInput) GetSendQuotaRequest {
	op := &aws.Operation{
		Name:       opGetSendQuota,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetSendQuotaInput{}
	}

	req := c.newRequest(op, input, &GetSendQuotaOutput{})
	return GetSendQuotaRequest{Request: req, Input: input, Copy: c.GetSendQuotaRequest}
}

// GetSendQuotaRequest is the request type for the
// GetSendQuota API operation.
type GetSendQuotaRequest struct {
	*aws.Request
	Input *GetSendQuotaInput
	Copy  func(*GetSendQuotaInput) GetSendQuotaRequest
}

// Send marshals and sends the GetSendQuota API request.
func (r GetSendQuotaRequest) Send(ctx context.Context) (*GetSendQuotaResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetSendQuotaResponse{
		GetSendQuotaOutput: r.Request.Data.(*GetSendQuotaOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetSendQuotaResponse is the response type for the
// GetSendQuota API operation.
type GetSendQuotaResponse struct {
	*GetSendQuotaOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetSendQuota request.
func (r *GetSendQuotaResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
