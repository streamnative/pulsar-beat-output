// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentity

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Input to the GetIdentityPoolRoles action.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-identity-2014-06-30/GetIdentityPoolRolesInput
type GetIdentityPoolRolesInput struct {
	_ struct{} `type:"structure"`

	// An identity pool ID in the format REGION:GUID.
	//
	// IdentityPoolId is a required field
	IdentityPoolId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetIdentityPoolRolesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetIdentityPoolRolesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetIdentityPoolRolesInput"}

	if s.IdentityPoolId == nil {
		invalidParams.Add(aws.NewErrParamRequired("IdentityPoolId"))
	}
	if s.IdentityPoolId != nil && len(*s.IdentityPoolId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("IdentityPoolId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Returned in response to a successful GetIdentityPoolRoles operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-identity-2014-06-30/GetIdentityPoolRolesResponse
type GetIdentityPoolRolesOutput struct {
	_ struct{} `type:"structure"`

	// An identity pool ID in the format REGION:GUID.
	IdentityPoolId *string `min:"1" type:"string"`

	// How users for a specific identity provider are to mapped to roles. This is
	// a String-to-RoleMapping object map. The string identifies the identity provider,
	// for example, "graph.facebook.com" or "cognito-idp.us-east-1.amazonaws.com/us-east-1_abcdefghi:app_client_id".
	RoleMappings map[string]RoleMapping `type:"map"`

	// The map of roles associated with this pool. Currently only authenticated
	// and unauthenticated roles are supported.
	Roles map[string]string `type:"map"`
}

// String returns the string representation
func (s GetIdentityPoolRolesOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetIdentityPoolRoles = "GetIdentityPoolRoles"

// GetIdentityPoolRolesRequest returns a request value for making API operation for
// Amazon Cognito Identity.
//
// Gets the roles for an identity pool.
//
// You must use AWS Developer credentials to call this API.
//
//    // Example sending a request using GetIdentityPoolRolesRequest.
//    req := client.GetIdentityPoolRolesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-identity-2014-06-30/GetIdentityPoolRoles
func (c *Client) GetIdentityPoolRolesRequest(input *GetIdentityPoolRolesInput) GetIdentityPoolRolesRequest {
	op := &aws.Operation{
		Name:       opGetIdentityPoolRoles,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetIdentityPoolRolesInput{}
	}

	req := c.newRequest(op, input, &GetIdentityPoolRolesOutput{})
	return GetIdentityPoolRolesRequest{Request: req, Input: input, Copy: c.GetIdentityPoolRolesRequest}
}

// GetIdentityPoolRolesRequest is the request type for the
// GetIdentityPoolRoles API operation.
type GetIdentityPoolRolesRequest struct {
	*aws.Request
	Input *GetIdentityPoolRolesInput
	Copy  func(*GetIdentityPoolRolesInput) GetIdentityPoolRolesRequest
}

// Send marshals and sends the GetIdentityPoolRoles API request.
func (r GetIdentityPoolRolesRequest) Send(ctx context.Context) (*GetIdentityPoolRolesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetIdentityPoolRolesResponse{
		GetIdentityPoolRolesOutput: r.Request.Data.(*GetIdentityPoolRolesOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetIdentityPoolRolesResponse is the response type for the
// GetIdentityPoolRoles API operation.
type GetIdentityPoolRolesResponse struct {
	*GetIdentityPoolRolesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetIdentityPoolRoles request.
func (r *GetIdentityPoolRolesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
