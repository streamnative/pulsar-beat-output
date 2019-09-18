// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mturk

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mturk-requester-2017-01-17/DisassociateQualificationFromWorkerRequest
type DisassociateQualificationFromWorkerInput struct {
	_ struct{} `type:"structure"`

	// The ID of the Qualification type of the Qualification to be revoked.
	//
	// QualificationTypeId is a required field
	QualificationTypeId *string `min:"1" type:"string" required:"true"`

	// A text message that explains why the Qualification was revoked. The user
	// who had the Qualification sees this message.
	Reason *string `type:"string"`

	// The ID of the Worker who possesses the Qualification to be revoked.
	//
	// WorkerId is a required field
	WorkerId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DisassociateQualificationFromWorkerInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DisassociateQualificationFromWorkerInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DisassociateQualificationFromWorkerInput"}

	if s.QualificationTypeId == nil {
		invalidParams.Add(aws.NewErrParamRequired("QualificationTypeId"))
	}
	if s.QualificationTypeId != nil && len(*s.QualificationTypeId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("QualificationTypeId", 1))
	}

	if s.WorkerId == nil {
		invalidParams.Add(aws.NewErrParamRequired("WorkerId"))
	}
	if s.WorkerId != nil && len(*s.WorkerId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("WorkerId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mturk-requester-2017-01-17/DisassociateQualificationFromWorkerResponse
type DisassociateQualificationFromWorkerOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DisassociateQualificationFromWorkerOutput) String() string {
	return awsutil.Prettify(s)
}

const opDisassociateQualificationFromWorker = "DisassociateQualificationFromWorker"

// DisassociateQualificationFromWorkerRequest returns a request value for making API operation for
// Amazon Mechanical Turk.
//
// The DisassociateQualificationFromWorker revokes a previously granted Qualification
// from a user.
//
// You can provide a text message explaining why the Qualification was revoked.
// The user who had the Qualification can see this message.
//
//    // Example sending a request using DisassociateQualificationFromWorkerRequest.
//    req := client.DisassociateQualificationFromWorkerRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mturk-requester-2017-01-17/DisassociateQualificationFromWorker
func (c *Client) DisassociateQualificationFromWorkerRequest(input *DisassociateQualificationFromWorkerInput) DisassociateQualificationFromWorkerRequest {
	op := &aws.Operation{
		Name:       opDisassociateQualificationFromWorker,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DisassociateQualificationFromWorkerInput{}
	}

	req := c.newRequest(op, input, &DisassociateQualificationFromWorkerOutput{})
	return DisassociateQualificationFromWorkerRequest{Request: req, Input: input, Copy: c.DisassociateQualificationFromWorkerRequest}
}

// DisassociateQualificationFromWorkerRequest is the request type for the
// DisassociateQualificationFromWorker API operation.
type DisassociateQualificationFromWorkerRequest struct {
	*aws.Request
	Input *DisassociateQualificationFromWorkerInput
	Copy  func(*DisassociateQualificationFromWorkerInput) DisassociateQualificationFromWorkerRequest
}

// Send marshals and sends the DisassociateQualificationFromWorker API request.
func (r DisassociateQualificationFromWorkerRequest) Send(ctx context.Context) (*DisassociateQualificationFromWorkerResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DisassociateQualificationFromWorkerResponse{
		DisassociateQualificationFromWorkerOutput: r.Request.Data.(*DisassociateQualificationFromWorkerOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DisassociateQualificationFromWorkerResponse is the response type for the
// DisassociateQualificationFromWorker API operation.
type DisassociateQualificationFromWorkerResponse struct {
	*DisassociateQualificationFromWorkerOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DisassociateQualificationFromWorker request.
func (r *DisassociateQualificationFromWorkerResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
