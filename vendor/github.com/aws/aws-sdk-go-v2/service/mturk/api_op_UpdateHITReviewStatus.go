// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mturk

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mturk-requester-2017-01-17/UpdateHITReviewStatusRequest
type UpdateHITReviewStatusInput struct {
	_ struct{} `type:"structure"`

	// The ID of the HIT to update.
	//
	// HITId is a required field
	HITId *string `min:"1" type:"string" required:"true"`

	// Specifies how to update the HIT status. Default is False.
	//
	//    * Setting this to false will only transition a HIT from Reviewable to
	//    Reviewing
	//
	//    * Setting this to true will only transition a HIT from Reviewing to Reviewable
	Revert *bool `type:"boolean"`
}

// String returns the string representation
func (s UpdateHITReviewStatusInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateHITReviewStatusInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateHITReviewStatusInput"}

	if s.HITId == nil {
		invalidParams.Add(aws.NewErrParamRequired("HITId"))
	}
	if s.HITId != nil && len(*s.HITId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("HITId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mturk-requester-2017-01-17/UpdateHITReviewStatusResponse
type UpdateHITReviewStatusOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateHITReviewStatusOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateHITReviewStatus = "UpdateHITReviewStatus"

// UpdateHITReviewStatusRequest returns a request value for making API operation for
// Amazon Mechanical Turk.
//
// The UpdateHITReviewStatus operation updates the status of a HIT. If the status
// is Reviewable, this operation can update the status to Reviewing, or it can
// revert a Reviewing HIT back to the Reviewable status.
//
//    // Example sending a request using UpdateHITReviewStatusRequest.
//    req := client.UpdateHITReviewStatusRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mturk-requester-2017-01-17/UpdateHITReviewStatus
func (c *Client) UpdateHITReviewStatusRequest(input *UpdateHITReviewStatusInput) UpdateHITReviewStatusRequest {
	op := &aws.Operation{
		Name:       opUpdateHITReviewStatus,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateHITReviewStatusInput{}
	}

	req := c.newRequest(op, input, &UpdateHITReviewStatusOutput{})
	return UpdateHITReviewStatusRequest{Request: req, Input: input, Copy: c.UpdateHITReviewStatusRequest}
}

// UpdateHITReviewStatusRequest is the request type for the
// UpdateHITReviewStatus API operation.
type UpdateHITReviewStatusRequest struct {
	*aws.Request
	Input *UpdateHITReviewStatusInput
	Copy  func(*UpdateHITReviewStatusInput) UpdateHITReviewStatusRequest
}

// Send marshals and sends the UpdateHITReviewStatus API request.
func (r UpdateHITReviewStatusRequest) Send(ctx context.Context) (*UpdateHITReviewStatusResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateHITReviewStatusResponse{
		UpdateHITReviewStatusOutput: r.Request.Data.(*UpdateHITReviewStatusOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateHITReviewStatusResponse is the response type for the
// UpdateHITReviewStatus API operation.
type UpdateHITReviewStatusResponse struct {
	*UpdateHITReviewStatusOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateHITReviewStatus request.
func (r *UpdateHITReviewStatusResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
