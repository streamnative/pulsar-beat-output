// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package shield

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/shield-2016-06-02/DeleteSubscriptionRequest
type DeleteSubscriptionInput struct {
	_ struct{} `deprecated:"true" type:"structure"`
}

// String returns the string representation
func (s DeleteSubscriptionInput) String() string {
	return awsutil.Prettify(s)
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/shield-2016-06-02/DeleteSubscriptionResponse
type DeleteSubscriptionOutput struct {
	_ struct{} `deprecated:"true" type:"structure"`
}

// String returns the string representation
func (s DeleteSubscriptionOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteSubscription = "DeleteSubscription"

// DeleteSubscriptionRequest returns a request value for making API operation for
// AWS Shield.
//
// Removes AWS Shield Advanced from an account. AWS Shield Advanced requires
// a 1-year subscription commitment. You cannot delete a subscription prior
// to the completion of that commitment.
//
//    // Example sending a request using DeleteSubscriptionRequest.
//    req := client.DeleteSubscriptionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/shield-2016-06-02/DeleteSubscription
func (c *Client) DeleteSubscriptionRequest(input *DeleteSubscriptionInput) DeleteSubscriptionRequest {
	if c.Client.Config.Logger != nil {
		c.Client.Config.Logger.Log("This operation, DeleteSubscription, has been deprecated")
	}
	op := &aws.Operation{
		Name:       opDeleteSubscription,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteSubscriptionInput{}
	}

	req := c.newRequest(op, input, &DeleteSubscriptionOutput{})
	return DeleteSubscriptionRequest{Request: req, Input: input, Copy: c.DeleteSubscriptionRequest}
}

// DeleteSubscriptionRequest is the request type for the
// DeleteSubscription API operation.
type DeleteSubscriptionRequest struct {
	*aws.Request
	Input *DeleteSubscriptionInput
	Copy  func(*DeleteSubscriptionInput) DeleteSubscriptionRequest
}

// Send marshals and sends the DeleteSubscription API request.
func (r DeleteSubscriptionRequest) Send(ctx context.Context) (*DeleteSubscriptionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteSubscriptionResponse{
		DeleteSubscriptionOutput: r.Request.Data.(*DeleteSubscriptionOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteSubscriptionResponse is the response type for the
// DeleteSubscription API operation.
type DeleteSubscriptionResponse struct {
	*DeleteSubscriptionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteSubscription request.
func (r *DeleteSubscriptionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
