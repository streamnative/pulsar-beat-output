// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package redshift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/ModifyEventSubscriptionMessage
type ModifyEventSubscriptionInput struct {
	_ struct{} `type:"structure"`

	// A Boolean value indicating if the subscription is enabled. true indicates
	// the subscription is enabled
	Enabled *bool `type:"boolean"`

	// Specifies the Amazon Redshift event categories to be published by the event
	// notification subscription.
	//
	// Values: configuration, management, monitoring, security
	EventCategories []string `locationNameList:"EventCategory" type:"list"`

	// Specifies the Amazon Redshift event severity to be published by the event
	// notification subscription.
	//
	// Values: ERROR, INFO
	Severity *string `type:"string"`

	// The Amazon Resource Name (ARN) of the SNS topic to be used by the event notification
	// subscription.
	SnsTopicArn *string `type:"string"`

	// A list of one or more identifiers of Amazon Redshift source objects. All
	// of the objects must be of the same type as was specified in the source type
	// parameter. The event subscription will return only events generated by the
	// specified objects. If not specified, then events are returned for all objects
	// within the source type specified.
	//
	// Example: my-cluster-1, my-cluster-2
	//
	// Example: my-snapshot-20131010
	SourceIds []string `locationNameList:"SourceId" type:"list"`

	// The type of source that will be generating the events. For example, if you
	// want to be notified of events generated by a cluster, you would set this
	// parameter to cluster. If this value is not specified, events are returned
	// for all Amazon Redshift objects in your AWS account. You must specify a source
	// type in order to specify source IDs.
	//
	// Valid values: cluster, cluster-parameter-group, cluster-security-group, and
	// cluster-snapshot.
	SourceType *string `type:"string"`

	// The name of the modified Amazon Redshift event notification subscription.
	//
	// SubscriptionName is a required field
	SubscriptionName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ModifyEventSubscriptionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyEventSubscriptionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ModifyEventSubscriptionInput"}

	if s.SubscriptionName == nil {
		invalidParams.Add(aws.NewErrParamRequired("SubscriptionName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/ModifyEventSubscriptionResult
type ModifyEventSubscriptionOutput struct {
	_ struct{} `type:"structure"`

	// Describes event subscriptions.
	EventSubscription *EventSubscription `type:"structure"`
}

// String returns the string representation
func (s ModifyEventSubscriptionOutput) String() string {
	return awsutil.Prettify(s)
}

const opModifyEventSubscription = "ModifyEventSubscription"

// ModifyEventSubscriptionRequest returns a request value for making API operation for
// Amazon Redshift.
//
// Modifies an existing Amazon Redshift event notification subscription.
//
//    // Example sending a request using ModifyEventSubscriptionRequest.
//    req := client.ModifyEventSubscriptionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/ModifyEventSubscription
func (c *Client) ModifyEventSubscriptionRequest(input *ModifyEventSubscriptionInput) ModifyEventSubscriptionRequest {
	op := &aws.Operation{
		Name:       opModifyEventSubscription,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyEventSubscriptionInput{}
	}

	req := c.newRequest(op, input, &ModifyEventSubscriptionOutput{})
	return ModifyEventSubscriptionRequest{Request: req, Input: input, Copy: c.ModifyEventSubscriptionRequest}
}

// ModifyEventSubscriptionRequest is the request type for the
// ModifyEventSubscription API operation.
type ModifyEventSubscriptionRequest struct {
	*aws.Request
	Input *ModifyEventSubscriptionInput
	Copy  func(*ModifyEventSubscriptionInput) ModifyEventSubscriptionRequest
}

// Send marshals and sends the ModifyEventSubscription API request.
func (r ModifyEventSubscriptionRequest) Send(ctx context.Context) (*ModifyEventSubscriptionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ModifyEventSubscriptionResponse{
		ModifyEventSubscriptionOutput: r.Request.Data.(*ModifyEventSubscriptionOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ModifyEventSubscriptionResponse is the response type for the
// ModifyEventSubscription API operation.
type ModifyEventSubscriptionResponse struct {
	*ModifyEventSubscriptionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ModifyEventSubscription request.
func (r *ModifyEventSubscriptionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
