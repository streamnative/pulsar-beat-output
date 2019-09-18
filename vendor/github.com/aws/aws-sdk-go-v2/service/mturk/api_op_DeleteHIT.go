// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mturk

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mturk-requester-2017-01-17/DeleteHITRequest
type DeleteHITInput struct {
	_ struct{} `type:"structure"`

	// The ID of the HIT to be deleted.
	//
	// HITId is a required field
	HITId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteHITInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteHITInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteHITInput"}

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

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mturk-requester-2017-01-17/DeleteHITResponse
type DeleteHITOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteHITOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteHIT = "DeleteHIT"

// DeleteHITRequest returns a request value for making API operation for
// Amazon Mechanical Turk.
//
// The DeleteHIT operation is used to delete HIT that is no longer needed. Only
// the Requester who created the HIT can delete it.
//
// You can only dispose of HITs that are in the Reviewable state, with all of
// their submitted assignments already either approved or rejected. If you call
// the DeleteHIT operation on a HIT that is not in the Reviewable state (for
// example, that has not expired, or still has active assignments), or on a
// HIT that is Reviewable but without all of its submitted assignments already
// approved or rejected, the service will return an error.
//
//    * HITs are automatically disposed of after 120 days.
//
//    * After you dispose of a HIT, you can no longer approve the HIT's rejected
//    assignments.
//
//    * Disposed HITs are not returned in results for the ListHITs operation.
//
//    * Disposing HITs can improve the performance of operations such as ListReviewableHITs
//    and ListHITs.
//
//    // Example sending a request using DeleteHITRequest.
//    req := client.DeleteHITRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mturk-requester-2017-01-17/DeleteHIT
func (c *Client) DeleteHITRequest(input *DeleteHITInput) DeleteHITRequest {
	op := &aws.Operation{
		Name:       opDeleteHIT,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteHITInput{}
	}

	req := c.newRequest(op, input, &DeleteHITOutput{})
	return DeleteHITRequest{Request: req, Input: input, Copy: c.DeleteHITRequest}
}

// DeleteHITRequest is the request type for the
// DeleteHIT API operation.
type DeleteHITRequest struct {
	*aws.Request
	Input *DeleteHITInput
	Copy  func(*DeleteHITInput) DeleteHITRequest
}

// Send marshals and sends the DeleteHIT API request.
func (r DeleteHITRequest) Send(ctx context.Context) (*DeleteHITResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteHITResponse{
		DeleteHITOutput: r.Request.Data.(*DeleteHITOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteHITResponse is the response type for the
// DeleteHIT API operation.
type DeleteHITResponse struct {
	*DeleteHITOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteHIT request.
func (r *DeleteHITResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
