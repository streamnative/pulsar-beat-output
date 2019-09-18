// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mturk

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mturk-requester-2017-01-17/RejectAssignmentRequest
type RejectAssignmentInput struct {
	_ struct{} `type:"structure"`

	// The ID of the assignment. The assignment must correspond to a HIT created
	// by the Requester.
	//
	// AssignmentId is a required field
	AssignmentId *string `min:"1" type:"string" required:"true"`

	// A message for the Worker, which the Worker can see in the Status section
	// of the web site.
	//
	// RequesterFeedback is a required field
	RequesterFeedback *string `type:"string" required:"true"`
}

// String returns the string representation
func (s RejectAssignmentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RejectAssignmentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RejectAssignmentInput"}

	if s.AssignmentId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AssignmentId"))
	}
	if s.AssignmentId != nil && len(*s.AssignmentId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AssignmentId", 1))
	}

	if s.RequesterFeedback == nil {
		invalidParams.Add(aws.NewErrParamRequired("RequesterFeedback"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mturk-requester-2017-01-17/RejectAssignmentResponse
type RejectAssignmentOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s RejectAssignmentOutput) String() string {
	return awsutil.Prettify(s)
}

const opRejectAssignment = "RejectAssignment"

// RejectAssignmentRequest returns a request value for making API operation for
// Amazon Mechanical Turk.
//
// The RejectAssignment operation rejects the results of a completed assignment.
//
// You can include an optional feedback message with the rejection, which the
// Worker can see in the Status section of the web site. When you include a
// feedback message with the rejection, it helps the Worker understand why the
// assignment was rejected, and can improve the quality of the results the Worker
// submits in the future.
//
// Only the Requester who created the HIT can reject an assignment for the HIT.
//
//    // Example sending a request using RejectAssignmentRequest.
//    req := client.RejectAssignmentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mturk-requester-2017-01-17/RejectAssignment
func (c *Client) RejectAssignmentRequest(input *RejectAssignmentInput) RejectAssignmentRequest {
	op := &aws.Operation{
		Name:       opRejectAssignment,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RejectAssignmentInput{}
	}

	req := c.newRequest(op, input, &RejectAssignmentOutput{})
	return RejectAssignmentRequest{Request: req, Input: input, Copy: c.RejectAssignmentRequest}
}

// RejectAssignmentRequest is the request type for the
// RejectAssignment API operation.
type RejectAssignmentRequest struct {
	*aws.Request
	Input *RejectAssignmentInput
	Copy  func(*RejectAssignmentInput) RejectAssignmentRequest
}

// Send marshals and sends the RejectAssignment API request.
func (r RejectAssignmentRequest) Send(ctx context.Context) (*RejectAssignmentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RejectAssignmentResponse{
		RejectAssignmentOutput: r.Request.Data.(*RejectAssignmentOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RejectAssignmentResponse is the response type for the
// RejectAssignment API operation.
type RejectAssignmentResponse struct {
	*RejectAssignmentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RejectAssignment request.
func (r *RejectAssignmentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
