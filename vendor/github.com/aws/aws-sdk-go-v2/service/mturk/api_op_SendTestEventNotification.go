// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mturk

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mturk-requester-2017-01-17/SendTestEventNotificationRequest
type SendTestEventNotificationInput struct {
	_ struct{} `type:"structure"`

	// The notification specification to test. This value is identical to the value
	// you would provide to the UpdateNotificationSettings operation when you establish
	// the notification specification for a HIT type.
	//
	// Notification is a required field
	Notification *NotificationSpecification `type:"structure" required:"true"`

	// The event to simulate to test the notification specification. This event
	// is included in the test message even if the notification specification does
	// not include the event type. The notification specification does not filter
	// out the test event.
	//
	// TestEventType is a required field
	TestEventType EventType `type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s SendTestEventNotificationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SendTestEventNotificationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SendTestEventNotificationInput"}

	if s.Notification == nil {
		invalidParams.Add(aws.NewErrParamRequired("Notification"))
	}
	if len(s.TestEventType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("TestEventType"))
	}
	if s.Notification != nil {
		if err := s.Notification.Validate(); err != nil {
			invalidParams.AddNested("Notification", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mturk-requester-2017-01-17/SendTestEventNotificationResponse
type SendTestEventNotificationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s SendTestEventNotificationOutput) String() string {
	return awsutil.Prettify(s)
}

const opSendTestEventNotification = "SendTestEventNotification"

// SendTestEventNotificationRequest returns a request value for making API operation for
// Amazon Mechanical Turk.
//
// The SendTestEventNotification operation causes Amazon Mechanical Turk to
// send a notification message as if a HIT event occurred, according to the
// provided notification specification. This allows you to test notifications
// without setting up notifications for a real HIT type and trying to trigger
// them using the website. When you call this operation, the service attempts
// to send the test notification immediately.
//
//    // Example sending a request using SendTestEventNotificationRequest.
//    req := client.SendTestEventNotificationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mturk-requester-2017-01-17/SendTestEventNotification
func (c *Client) SendTestEventNotificationRequest(input *SendTestEventNotificationInput) SendTestEventNotificationRequest {
	op := &aws.Operation{
		Name:       opSendTestEventNotification,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &SendTestEventNotificationInput{}
	}

	req := c.newRequest(op, input, &SendTestEventNotificationOutput{})
	return SendTestEventNotificationRequest{Request: req, Input: input, Copy: c.SendTestEventNotificationRequest}
}

// SendTestEventNotificationRequest is the request type for the
// SendTestEventNotification API operation.
type SendTestEventNotificationRequest struct {
	*aws.Request
	Input *SendTestEventNotificationInput
	Copy  func(*SendTestEventNotificationInput) SendTestEventNotificationRequest
}

// Send marshals and sends the SendTestEventNotification API request.
func (r SendTestEventNotificationRequest) Send(ctx context.Context) (*SendTestEventNotificationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SendTestEventNotificationResponse{
		SendTestEventNotificationOutput: r.Request.Data.(*SendTestEventNotificationOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// SendTestEventNotificationResponse is the response type for the
// SendTestEventNotification API operation.
type SendTestEventNotificationResponse struct {
	*SendTestEventNotificationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SendTestEventNotification request.
func (r *SendTestEventNotificationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
