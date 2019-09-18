// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kms

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/kms-2014-11-01/GetKeyPolicyRequest
type GetKeyPolicyInput struct {
	_ struct{} `type:"structure"`

	// A unique identifier for the customer master key (CMK).
	//
	// Specify the key ID or the Amazon Resource Name (ARN) of the CMK.
	//
	// For example:
	//
	//    * Key ID: 1234abcd-12ab-34cd-56ef-1234567890ab
	//
	//    * Key ARN: arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab
	//
	// To get the key ID and key ARN for a CMK, use ListKeys or DescribeKey.
	//
	// KeyId is a required field
	KeyId *string `min:"1" type:"string" required:"true"`

	// Specifies the name of the key policy. The only valid name is default. To
	// get the names of key policies, use ListKeyPolicies.
	//
	// PolicyName is a required field
	PolicyName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetKeyPolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetKeyPolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetKeyPolicyInput"}

	if s.KeyId == nil {
		invalidParams.Add(aws.NewErrParamRequired("KeyId"))
	}
	if s.KeyId != nil && len(*s.KeyId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("KeyId", 1))
	}

	if s.PolicyName == nil {
		invalidParams.Add(aws.NewErrParamRequired("PolicyName"))
	}
	if s.PolicyName != nil && len(*s.PolicyName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PolicyName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/kms-2014-11-01/GetKeyPolicyResponse
type GetKeyPolicyOutput struct {
	_ struct{} `type:"structure"`

	// A key policy document in JSON format.
	Policy *string `min:"1" type:"string"`
}

// String returns the string representation
func (s GetKeyPolicyOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetKeyPolicy = "GetKeyPolicy"

// GetKeyPolicyRequest returns a request value for making API operation for
// AWS Key Management Service.
//
// Gets a key policy attached to the specified customer master key (CMK). You
// cannot perform this operation on a CMK in a different AWS account.
//
//    // Example sending a request using GetKeyPolicyRequest.
//    req := client.GetKeyPolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kms-2014-11-01/GetKeyPolicy
func (c *Client) GetKeyPolicyRequest(input *GetKeyPolicyInput) GetKeyPolicyRequest {
	op := &aws.Operation{
		Name:       opGetKeyPolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetKeyPolicyInput{}
	}

	req := c.newRequest(op, input, &GetKeyPolicyOutput{})
	return GetKeyPolicyRequest{Request: req, Input: input, Copy: c.GetKeyPolicyRequest}
}

// GetKeyPolicyRequest is the request type for the
// GetKeyPolicy API operation.
type GetKeyPolicyRequest struct {
	*aws.Request
	Input *GetKeyPolicyInput
	Copy  func(*GetKeyPolicyInput) GetKeyPolicyRequest
}

// Send marshals and sends the GetKeyPolicy API request.
func (r GetKeyPolicyRequest) Send(ctx context.Context) (*GetKeyPolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetKeyPolicyResponse{
		GetKeyPolicyOutput: r.Request.Data.(*GetKeyPolicyOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetKeyPolicyResponse is the response type for the
// GetKeyPolicy API operation.
type GetKeyPolicyResponse struct {
	*GetKeyPolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetKeyPolicy request.
func (r *GetKeyPolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
