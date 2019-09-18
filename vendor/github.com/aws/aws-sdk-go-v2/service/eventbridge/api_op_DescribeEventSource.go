// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package eventbridge

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/eventbridge-2015-10-07/DescribeEventSourceRequest
type DescribeEventSourceInput struct {
	_ struct{} `type:"structure"`

	// The name of the partner event source to display the details of.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeEventSourceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeEventSourceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeEventSourceInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/eventbridge-2015-10-07/DescribeEventSourceResponse
type DescribeEventSourceOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the partner event source.
	Arn *string `type:"string"`

	// The name of the SaaS partner that created the event source.
	CreatedBy *string `type:"string"`

	// The date and time that the event source was created.
	CreationTime *time.Time `type:"timestamp"`

	// The date and time that the event source will expire if you don't create a
	// matching event bus.
	ExpirationTime *time.Time `type:"timestamp"`

	// The name of the partner event source.
	Name *string `type:"string"`

	// The state of the event source. If it's ACTIVE, you have already created a
	// matching event bus for this event source, and that event bus is active. If
	// it's PENDING, either you haven't yet created a matching event bus, or that
	// event bus is deactivated. If it's DELETED, you have created a matching event
	// bus, but the event source has since been deleted.
	State EventSourceState `type:"string" enum:"true"`
}

// String returns the string representation
func (s DescribeEventSourceOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeEventSource = "DescribeEventSource"

// DescribeEventSourceRequest returns a request value for making API operation for
// Amazon EventBridge.
//
// This operation lists details about a partner event source that is shared
// with your account.
//
// This operation is run by AWS customers, not by SaaS partners.
//
//    // Example sending a request using DescribeEventSourceRequest.
//    req := client.DescribeEventSourceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/eventbridge-2015-10-07/DescribeEventSource
func (c *Client) DescribeEventSourceRequest(input *DescribeEventSourceInput) DescribeEventSourceRequest {
	op := &aws.Operation{
		Name:       opDescribeEventSource,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeEventSourceInput{}
	}

	req := c.newRequest(op, input, &DescribeEventSourceOutput{})
	return DescribeEventSourceRequest{Request: req, Input: input, Copy: c.DescribeEventSourceRequest}
}

// DescribeEventSourceRequest is the request type for the
// DescribeEventSource API operation.
type DescribeEventSourceRequest struct {
	*aws.Request
	Input *DescribeEventSourceInput
	Copy  func(*DescribeEventSourceInput) DescribeEventSourceRequest
}

// Send marshals and sends the DescribeEventSource API request.
func (r DescribeEventSourceRequest) Send(ctx context.Context) (*DescribeEventSourceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeEventSourceResponse{
		DescribeEventSourceOutput: r.Request.Data.(*DescribeEventSourceOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeEventSourceResponse is the response type for the
// DescribeEventSource API operation.
type DescribeEventSourceResponse struct {
	*DescribeEventSourceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeEventSource request.
func (r *DescribeEventSourceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
