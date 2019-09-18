// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appstream

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/appstream-2016-12-01/DeleteUsageReportSubscriptionRequest
type DeleteUsageReportSubscriptionInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteUsageReportSubscriptionInput) String() string {
	return awsutil.Prettify(s)
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/appstream-2016-12-01/DeleteUsageReportSubscriptionResult
type DeleteUsageReportSubscriptionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteUsageReportSubscriptionOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteUsageReportSubscription = "DeleteUsageReportSubscription"

// DeleteUsageReportSubscriptionRequest returns a request value for making API operation for
// Amazon AppStream.
//
// Disables usage report generation.
//
//    // Example sending a request using DeleteUsageReportSubscriptionRequest.
//    req := client.DeleteUsageReportSubscriptionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appstream-2016-12-01/DeleteUsageReportSubscription
func (c *Client) DeleteUsageReportSubscriptionRequest(input *DeleteUsageReportSubscriptionInput) DeleteUsageReportSubscriptionRequest {
	op := &aws.Operation{
		Name:       opDeleteUsageReportSubscription,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteUsageReportSubscriptionInput{}
	}

	req := c.newRequest(op, input, &DeleteUsageReportSubscriptionOutput{})
	return DeleteUsageReportSubscriptionRequest{Request: req, Input: input, Copy: c.DeleteUsageReportSubscriptionRequest}
}

// DeleteUsageReportSubscriptionRequest is the request type for the
// DeleteUsageReportSubscription API operation.
type DeleteUsageReportSubscriptionRequest struct {
	*aws.Request
	Input *DeleteUsageReportSubscriptionInput
	Copy  func(*DeleteUsageReportSubscriptionInput) DeleteUsageReportSubscriptionRequest
}

// Send marshals and sends the DeleteUsageReportSubscription API request.
func (r DeleteUsageReportSubscriptionRequest) Send(ctx context.Context) (*DeleteUsageReportSubscriptionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteUsageReportSubscriptionResponse{
		DeleteUsageReportSubscriptionOutput: r.Request.Data.(*DeleteUsageReportSubscriptionOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteUsageReportSubscriptionResponse is the response type for the
// DeleteUsageReportSubscription API operation.
type DeleteUsageReportSubscriptionResponse struct {
	*DeleteUsageReportSubscriptionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteUsageReportSubscription request.
func (r *DeleteUsageReportSubscriptionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
