// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/GetPatchBaselineForPatchGroupRequest
type GetPatchBaselineForPatchGroupInput struct {
	_ struct{} `type:"structure"`

	// Returns he operating system rule specified for patch groups using the patch
	// baseline.
	OperatingSystem OperatingSystem `type:"string" enum:"true"`

	// The name of the patch group whose patch baseline should be retrieved.
	//
	// PatchGroup is a required field
	PatchGroup *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetPatchBaselineForPatchGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetPatchBaselineForPatchGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetPatchBaselineForPatchGroupInput"}

	if s.PatchGroup == nil {
		invalidParams.Add(aws.NewErrParamRequired("PatchGroup"))
	}
	if s.PatchGroup != nil && len(*s.PatchGroup) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PatchGroup", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/GetPatchBaselineForPatchGroupResult
type GetPatchBaselineForPatchGroupOutput struct {
	_ struct{} `type:"structure"`

	// The ID of the patch baseline that should be used for the patch group.
	BaselineId *string `min:"20" type:"string"`

	// The operating system rule specified for patch groups using the patch baseline.
	OperatingSystem OperatingSystem `type:"string" enum:"true"`

	// The name of the patch group.
	PatchGroup *string `min:"1" type:"string"`
}

// String returns the string representation
func (s GetPatchBaselineForPatchGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetPatchBaselineForPatchGroup = "GetPatchBaselineForPatchGroup"

// GetPatchBaselineForPatchGroupRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Retrieves the patch baseline that should be used for the specified patch
// group.
//
//    // Example sending a request using GetPatchBaselineForPatchGroupRequest.
//    req := client.GetPatchBaselineForPatchGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/GetPatchBaselineForPatchGroup
func (c *Client) GetPatchBaselineForPatchGroupRequest(input *GetPatchBaselineForPatchGroupInput) GetPatchBaselineForPatchGroupRequest {
	op := &aws.Operation{
		Name:       opGetPatchBaselineForPatchGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetPatchBaselineForPatchGroupInput{}
	}

	req := c.newRequest(op, input, &GetPatchBaselineForPatchGroupOutput{})
	return GetPatchBaselineForPatchGroupRequest{Request: req, Input: input, Copy: c.GetPatchBaselineForPatchGroupRequest}
}

// GetPatchBaselineForPatchGroupRequest is the request type for the
// GetPatchBaselineForPatchGroup API operation.
type GetPatchBaselineForPatchGroupRequest struct {
	*aws.Request
	Input *GetPatchBaselineForPatchGroupInput
	Copy  func(*GetPatchBaselineForPatchGroupInput) GetPatchBaselineForPatchGroupRequest
}

// Send marshals and sends the GetPatchBaselineForPatchGroup API request.
func (r GetPatchBaselineForPatchGroupRequest) Send(ctx context.Context) (*GetPatchBaselineForPatchGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetPatchBaselineForPatchGroupResponse{
		GetPatchBaselineForPatchGroupOutput: r.Request.Data.(*GetPatchBaselineForPatchGroupOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetPatchBaselineForPatchGroupResponse is the response type for the
// GetPatchBaselineForPatchGroup API operation.
type GetPatchBaselineForPatchGroupResponse struct {
	*GetPatchBaselineForPatchGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetPatchBaselineForPatchGroup request.
func (r *GetPatchBaselineForPatchGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
