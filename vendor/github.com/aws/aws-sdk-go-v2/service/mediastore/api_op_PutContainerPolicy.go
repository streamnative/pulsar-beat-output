// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediastore

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediastore-2017-09-01/PutContainerPolicyInput
type PutContainerPolicyInput struct {
	_ struct{} `type:"structure"`

	// The name of the container.
	//
	// ContainerName is a required field
	ContainerName *string `min:"1" type:"string" required:"true"`

	// The contents of the policy, which includes the following:
	//
	//    * One Version tag
	//
	//    * One Statement tag that contains the standard tags for the policy.
	//
	// Policy is a required field
	Policy *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s PutContainerPolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutContainerPolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutContainerPolicyInput"}

	if s.ContainerName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ContainerName"))
	}
	if s.ContainerName != nil && len(*s.ContainerName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ContainerName", 1))
	}

	if s.Policy == nil {
		invalidParams.Add(aws.NewErrParamRequired("Policy"))
	}
	if s.Policy != nil && len(*s.Policy) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Policy", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediastore-2017-09-01/PutContainerPolicyOutput
type PutContainerPolicyOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutContainerPolicyOutput) String() string {
	return awsutil.Prettify(s)
}

const opPutContainerPolicy = "PutContainerPolicy"

// PutContainerPolicyRequest returns a request value for making API operation for
// AWS Elemental MediaStore.
//
// Creates an access policy for the specified container to restrict the users
// and clients that can access it. For information about the data that is included
// in an access policy, see the AWS Identity and Access Management User Guide
// (https://aws.amazon.com/documentation/iam/).
//
// For this release of the REST API, you can create only one policy for a container.
// If you enter PutContainerPolicy twice, the second command modifies the existing
// policy.
//
//    // Example sending a request using PutContainerPolicyRequest.
//    req := client.PutContainerPolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediastore-2017-09-01/PutContainerPolicy
func (c *Client) PutContainerPolicyRequest(input *PutContainerPolicyInput) PutContainerPolicyRequest {
	op := &aws.Operation{
		Name:       opPutContainerPolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PutContainerPolicyInput{}
	}

	req := c.newRequest(op, input, &PutContainerPolicyOutput{})
	return PutContainerPolicyRequest{Request: req, Input: input, Copy: c.PutContainerPolicyRequest}
}

// PutContainerPolicyRequest is the request type for the
// PutContainerPolicy API operation.
type PutContainerPolicyRequest struct {
	*aws.Request
	Input *PutContainerPolicyInput
	Copy  func(*PutContainerPolicyInput) PutContainerPolicyRequest
}

// Send marshals and sends the PutContainerPolicy API request.
func (r PutContainerPolicyRequest) Send(ctx context.Context) (*PutContainerPolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutContainerPolicyResponse{
		PutContainerPolicyOutput: r.Request.Data.(*PutContainerPolicyOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutContainerPolicyResponse is the response type for the
// PutContainerPolicy API operation.
type PutContainerPolicyResponse struct {
	*PutContainerPolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutContainerPolicy request.
func (r *PutContainerPolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
