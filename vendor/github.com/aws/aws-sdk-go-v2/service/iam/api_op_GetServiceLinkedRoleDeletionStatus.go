// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/GetServiceLinkedRoleDeletionStatusRequest
type GetServiceLinkedRoleDeletionStatusInput struct {
	_ struct{} `type:"structure"`

	// The deletion task identifier. This identifier is returned by the DeleteServiceLinkedRole
	// operation in the format task/aws-service-role/<service-principal-name>/<role-name>/<task-uuid>.
	//
	// DeletionTaskId is a required field
	DeletionTaskId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetServiceLinkedRoleDeletionStatusInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetServiceLinkedRoleDeletionStatusInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetServiceLinkedRoleDeletionStatusInput"}

	if s.DeletionTaskId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DeletionTaskId"))
	}
	if s.DeletionTaskId != nil && len(*s.DeletionTaskId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DeletionTaskId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/GetServiceLinkedRoleDeletionStatusResponse
type GetServiceLinkedRoleDeletionStatusOutput struct {
	_ struct{} `type:"structure"`

	// An object that contains details about the reason the deletion failed.
	Reason *DeletionTaskFailureReasonType `type:"structure"`

	// The status of the deletion.
	//
	// Status is a required field
	Status DeletionTaskStatusType `type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s GetServiceLinkedRoleDeletionStatusOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetServiceLinkedRoleDeletionStatus = "GetServiceLinkedRoleDeletionStatus"

// GetServiceLinkedRoleDeletionStatusRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Retrieves the status of your service-linked role deletion. After you use
// the DeleteServiceLinkedRole API operation to submit a service-linked role
// for deletion, you can use the DeletionTaskId parameter in GetServiceLinkedRoleDeletionStatus
// to check the status of the deletion. If the deletion fails, this operation
// returns the reason that it failed, if that information is returned by the
// service.
//
//    // Example sending a request using GetServiceLinkedRoleDeletionStatusRequest.
//    req := client.GetServiceLinkedRoleDeletionStatusRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/GetServiceLinkedRoleDeletionStatus
func (c *Client) GetServiceLinkedRoleDeletionStatusRequest(input *GetServiceLinkedRoleDeletionStatusInput) GetServiceLinkedRoleDeletionStatusRequest {
	op := &aws.Operation{
		Name:       opGetServiceLinkedRoleDeletionStatus,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetServiceLinkedRoleDeletionStatusInput{}
	}

	req := c.newRequest(op, input, &GetServiceLinkedRoleDeletionStatusOutput{})
	return GetServiceLinkedRoleDeletionStatusRequest{Request: req, Input: input, Copy: c.GetServiceLinkedRoleDeletionStatusRequest}
}

// GetServiceLinkedRoleDeletionStatusRequest is the request type for the
// GetServiceLinkedRoleDeletionStatus API operation.
type GetServiceLinkedRoleDeletionStatusRequest struct {
	*aws.Request
	Input *GetServiceLinkedRoleDeletionStatusInput
	Copy  func(*GetServiceLinkedRoleDeletionStatusInput) GetServiceLinkedRoleDeletionStatusRequest
}

// Send marshals and sends the GetServiceLinkedRoleDeletionStatus API request.
func (r GetServiceLinkedRoleDeletionStatusRequest) Send(ctx context.Context) (*GetServiceLinkedRoleDeletionStatusResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetServiceLinkedRoleDeletionStatusResponse{
		GetServiceLinkedRoleDeletionStatusOutput: r.Request.Data.(*GetServiceLinkedRoleDeletionStatusOutput),
		response:                                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetServiceLinkedRoleDeletionStatusResponse is the response type for the
// GetServiceLinkedRoleDeletionStatus API operation.
type GetServiceLinkedRoleDeletionStatusResponse struct {
	*GetServiceLinkedRoleDeletionStatusOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetServiceLinkedRoleDeletionStatus request.
func (r *GetServiceLinkedRoleDeletionStatusResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
