// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentity

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

// Input to the SetIdentityPoolRoles action.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-identity-2014-06-30/SetIdentityPoolRolesInput
type SetIdentityPoolRolesInput struct {
	_ struct{} `type:"structure"`

	// An identity pool ID in the format REGION:GUID.
	//
	// IdentityPoolId is a required field
	IdentityPoolId *string `min:"1" type:"string" required:"true"`

	// How users for a specific identity provider are to mapped to roles. This is
	// a string to RoleMapping object map. The string identifies the identity provider,
	// for example, "graph.facebook.com" or "cognito-idp-east-1.amazonaws.com/us-east-1_abcdefghi:app_client_id".
	//
	// Up to 25 rules can be specified per identity provider.
	RoleMappings map[string]RoleMapping `type:"map"`

	// The map of roles associated with this pool. For a given role, the key will
	// be either "authenticated" or "unauthenticated" and the value will be the
	// Role ARN.
	//
	// Roles is a required field
	Roles map[string]string `type:"map" required:"true"`
}

// String returns the string representation
func (s SetIdentityPoolRolesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SetIdentityPoolRolesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SetIdentityPoolRolesInput"}

	if s.IdentityPoolId == nil {
		invalidParams.Add(aws.NewErrParamRequired("IdentityPoolId"))
	}
	if s.IdentityPoolId != nil && len(*s.IdentityPoolId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("IdentityPoolId", 1))
	}

	if s.Roles == nil {
		invalidParams.Add(aws.NewErrParamRequired("Roles"))
	}
	if s.RoleMappings != nil {
		for i, v := range s.RoleMappings {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "RoleMappings", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-identity-2014-06-30/SetIdentityPoolRolesOutput
type SetIdentityPoolRolesOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s SetIdentityPoolRolesOutput) String() string {
	return awsutil.Prettify(s)
}

const opSetIdentityPoolRoles = "SetIdentityPoolRoles"

// SetIdentityPoolRolesRequest returns a request value for making API operation for
// Amazon Cognito Identity.
//
// Sets the roles for an identity pool. These roles are used when making calls
// to GetCredentialsForIdentity action.
//
// You must use AWS Developer credentials to call this API.
//
//    // Example sending a request using SetIdentityPoolRolesRequest.
//    req := client.SetIdentityPoolRolesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-identity-2014-06-30/SetIdentityPoolRoles
func (c *Client) SetIdentityPoolRolesRequest(input *SetIdentityPoolRolesInput) SetIdentityPoolRolesRequest {
	op := &aws.Operation{
		Name:       opSetIdentityPoolRoles,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &SetIdentityPoolRolesInput{}
	}

	req := c.newRequest(op, input, &SetIdentityPoolRolesOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return SetIdentityPoolRolesRequest{Request: req, Input: input, Copy: c.SetIdentityPoolRolesRequest}
}

// SetIdentityPoolRolesRequest is the request type for the
// SetIdentityPoolRoles API operation.
type SetIdentityPoolRolesRequest struct {
	*aws.Request
	Input *SetIdentityPoolRolesInput
	Copy  func(*SetIdentityPoolRolesInput) SetIdentityPoolRolesRequest
}

// Send marshals and sends the SetIdentityPoolRoles API request.
func (r SetIdentityPoolRolesRequest) Send(ctx context.Context) (*SetIdentityPoolRolesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SetIdentityPoolRolesResponse{
		SetIdentityPoolRolesOutput: r.Request.Data.(*SetIdentityPoolRolesOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// SetIdentityPoolRolesResponse is the response type for the
// SetIdentityPoolRoles API operation.
type SetIdentityPoolRolesResponse struct {
	*SetIdentityPoolRolesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SetIdentityPoolRoles request.
func (r *SetIdentityPoolRolesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
