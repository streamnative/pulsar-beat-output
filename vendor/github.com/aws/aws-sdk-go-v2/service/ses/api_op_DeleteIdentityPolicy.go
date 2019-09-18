// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ses

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents a request to delete a sending authorization policy for an identity.
// Sending authorization is an Amazon SES feature that enables you to authorize
// other senders to use your identities. For information, see the Amazon SES
// Developer Guide (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/sending-authorization.html).
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/DeleteIdentityPolicyRequest
type DeleteIdentityPolicyInput struct {
	_ struct{} `type:"structure"`

	// The identity that is associated with the policy that you want to delete.
	// You can specify the identity by using its name or by using its Amazon Resource
	// Name (ARN). Examples: user@example.com, example.com, arn:aws:ses:us-east-1:123456789012:identity/example.com.
	//
	// To successfully call this API, you must own the identity.
	//
	// Identity is a required field
	Identity *string `type:"string" required:"true"`

	// The name of the policy to be deleted.
	//
	// PolicyName is a required field
	PolicyName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteIdentityPolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteIdentityPolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteIdentityPolicyInput"}

	if s.Identity == nil {
		invalidParams.Add(aws.NewErrParamRequired("Identity"))
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

// An empty element returned on a successful request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/DeleteIdentityPolicyResponse
type DeleteIdentityPolicyOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteIdentityPolicyOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteIdentityPolicy = "DeleteIdentityPolicy"

// DeleteIdentityPolicyRequest returns a request value for making API operation for
// Amazon Simple Email Service.
//
// Deletes the specified sending authorization policy for the given identity
// (an email address or a domain). This API returns successfully even if a policy
// with the specified name does not exist.
//
// This API is for the identity owner only. If you have not verified the identity,
// this API will return an error.
//
// Sending authorization is a feature that enables an identity owner to authorize
// other senders to use its identities. For information about using sending
// authorization, see the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/sending-authorization.html).
//
// You can execute this operation no more than once per second.
//
//    // Example sending a request using DeleteIdentityPolicyRequest.
//    req := client.DeleteIdentityPolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/DeleteIdentityPolicy
func (c *Client) DeleteIdentityPolicyRequest(input *DeleteIdentityPolicyInput) DeleteIdentityPolicyRequest {
	op := &aws.Operation{
		Name:       opDeleteIdentityPolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteIdentityPolicyInput{}
	}

	req := c.newRequest(op, input, &DeleteIdentityPolicyOutput{})
	return DeleteIdentityPolicyRequest{Request: req, Input: input, Copy: c.DeleteIdentityPolicyRequest}
}

// DeleteIdentityPolicyRequest is the request type for the
// DeleteIdentityPolicy API operation.
type DeleteIdentityPolicyRequest struct {
	*aws.Request
	Input *DeleteIdentityPolicyInput
	Copy  func(*DeleteIdentityPolicyInput) DeleteIdentityPolicyRequest
}

// Send marshals and sends the DeleteIdentityPolicy API request.
func (r DeleteIdentityPolicyRequest) Send(ctx context.Context) (*DeleteIdentityPolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteIdentityPolicyResponse{
		DeleteIdentityPolicyOutput: r.Request.Data.(*DeleteIdentityPolicyOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteIdentityPolicyResponse is the response type for the
// DeleteIdentityPolicy API operation.
type DeleteIdentityPolicyResponse struct {
	*DeleteIdentityPolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteIdentityPolicy request.
func (r *DeleteIdentityPolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
