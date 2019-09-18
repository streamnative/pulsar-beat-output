// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mturk

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mturk-requester-2017-01-17/AssociateQualificationWithWorkerRequest
type AssociateQualificationWithWorkerInput struct {
	_ struct{} `type:"structure"`

	// The value of the Qualification to assign.
	IntegerValue *int64 `type:"integer"`

	// The ID of the Qualification type to use for the assigned Qualification.
	//
	// QualificationTypeId is a required field
	QualificationTypeId *string `min:"1" type:"string" required:"true"`

	// Specifies whether to send a notification email message to the Worker saying
	// that the qualification was assigned to the Worker. Note: this is true by
	// default.
	SendNotification *bool `type:"boolean"`

	// The ID of the Worker to whom the Qualification is being assigned. Worker
	// IDs are included with submitted HIT assignments and Qualification requests.
	//
	// WorkerId is a required field
	WorkerId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s AssociateQualificationWithWorkerInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AssociateQualificationWithWorkerInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AssociateQualificationWithWorkerInput"}

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

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mturk-requester-2017-01-17/AssociateQualificationWithWorkerResponse
type AssociateQualificationWithWorkerOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s AssociateQualificationWithWorkerOutput) String() string {
	return awsutil.Prettify(s)
}

const opAssociateQualificationWithWorker = "AssociateQualificationWithWorker"

// AssociateQualificationWithWorkerRequest returns a request value for making API operation for
// Amazon Mechanical Turk.
//
// The AssociateQualificationWithWorker operation gives a Worker a Qualification.
// AssociateQualificationWithWorker does not require that the Worker submit
// a Qualification request. It gives the Qualification directly to the Worker.
//
// You can only assign a Qualification of a Qualification type that you created
// (using the CreateQualificationType operation).
//
// Note: AssociateQualificationWithWorker does not affect any pending Qualification
// requests for the Qualification by the Worker. If you assign a Qualification
// to a Worker, then later grant a Qualification request made by the Worker,
// the granting of the request may modify the Qualification score. To resolve
// a pending Qualification request without affecting the Qualification the Worker
// already has, reject the request with the RejectQualificationRequest operation.
//
//    // Example sending a request using AssociateQualificationWithWorkerRequest.
//    req := client.AssociateQualificationWithWorkerRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mturk-requester-2017-01-17/AssociateQualificationWithWorker
func (c *Client) AssociateQualificationWithWorkerRequest(input *AssociateQualificationWithWorkerInput) AssociateQualificationWithWorkerRequest {
	op := &aws.Operation{
		Name:       opAssociateQualificationWithWorker,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AssociateQualificationWithWorkerInput{}
	}

	req := c.newRequest(op, input, &AssociateQualificationWithWorkerOutput{})
	return AssociateQualificationWithWorkerRequest{Request: req, Input: input, Copy: c.AssociateQualificationWithWorkerRequest}
}

// AssociateQualificationWithWorkerRequest is the request type for the
// AssociateQualificationWithWorker API operation.
type AssociateQualificationWithWorkerRequest struct {
	*aws.Request
	Input *AssociateQualificationWithWorkerInput
	Copy  func(*AssociateQualificationWithWorkerInput) AssociateQualificationWithWorkerRequest
}

// Send marshals and sends the AssociateQualificationWithWorker API request.
func (r AssociateQualificationWithWorkerRequest) Send(ctx context.Context) (*AssociateQualificationWithWorkerResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AssociateQualificationWithWorkerResponse{
		AssociateQualificationWithWorkerOutput: r.Request.Data.(*AssociateQualificationWithWorkerOutput),
		response:                               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AssociateQualificationWithWorkerResponse is the response type for the
// AssociateQualificationWithWorker API operation.
type AssociateQualificationWithWorkerResponse struct {
	*AssociateQualificationWithWorkerOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AssociateQualificationWithWorker request.
func (r *AssociateQualificationWithWorkerResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
