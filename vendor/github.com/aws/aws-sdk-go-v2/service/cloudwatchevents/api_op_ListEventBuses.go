// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudwatchevents

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/events-2015-10-07/ListEventBusesRequest
type ListEventBusesInput struct {
	_ struct{} `type:"structure"`

	// Specifying this limits the number of results returned by this operation.
	// The operation also returns a NextToken that you can use in a subsequent operation
	// to retrieve the next set of results.
	Limit *int64 `min:"1" type:"integer"`

	// Specifying this limits the results to only those event buses with names that
	// start with the specified prefix.
	NamePrefix *string `min:"1" type:"string"`

	// The token returned by a previous call to retrieve the next set of results.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListEventBusesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListEventBusesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListEventBusesInput"}
	if s.Limit != nil && *s.Limit < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Limit", 1))
	}
	if s.NamePrefix != nil && len(*s.NamePrefix) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NamePrefix", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/events-2015-10-07/ListEventBusesResponse
type ListEventBusesOutput struct {
	_ struct{} `type:"structure"`

	// This list of event buses.
	EventBuses []EventBus `type:"list"`

	// A token you can use in a subsequent operation to retrieve the next set of
	// results.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListEventBusesOutput) String() string {
	return awsutil.Prettify(s)
}

const opListEventBuses = "ListEventBuses"

// ListEventBusesRequest returns a request value for making API operation for
// Amazon CloudWatch Events.
//
// Lists all the event buses in your account, including the default event bus,
// custom event buses, and partner event buses.
//
// This operation is run by AWS customers, not by SaaS partners.
//
//    // Example sending a request using ListEventBusesRequest.
//    req := client.ListEventBusesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/events-2015-10-07/ListEventBuses
func (c *Client) ListEventBusesRequest(input *ListEventBusesInput) ListEventBusesRequest {
	op := &aws.Operation{
		Name:       opListEventBuses,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListEventBusesInput{}
	}

	req := c.newRequest(op, input, &ListEventBusesOutput{})
	return ListEventBusesRequest{Request: req, Input: input, Copy: c.ListEventBusesRequest}
}

// ListEventBusesRequest is the request type for the
// ListEventBuses API operation.
type ListEventBusesRequest struct {
	*aws.Request
	Input *ListEventBusesInput
	Copy  func(*ListEventBusesInput) ListEventBusesRequest
}

// Send marshals and sends the ListEventBuses API request.
func (r ListEventBusesRequest) Send(ctx context.Context) (*ListEventBusesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListEventBusesResponse{
		ListEventBusesOutput: r.Request.Data.(*ListEventBusesOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListEventBusesResponse is the response type for the
// ListEventBuses API operation.
type ListEventBusesResponse struct {
	*ListEventBusesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListEventBuses request.
func (r *ListEventBusesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
