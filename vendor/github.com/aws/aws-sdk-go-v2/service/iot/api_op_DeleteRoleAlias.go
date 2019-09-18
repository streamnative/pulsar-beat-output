// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DeleteRoleAliasInput struct {
	_ struct{} `type:"structure"`

	// The role alias to delete.
	//
	// RoleAlias is a required field
	RoleAlias *string `location:"uri" locationName:"roleAlias" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteRoleAliasInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteRoleAliasInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteRoleAliasInput"}

	if s.RoleAlias == nil {
		invalidParams.Add(aws.NewErrParamRequired("RoleAlias"))
	}
	if s.RoleAlias != nil && len(*s.RoleAlias) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RoleAlias", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteRoleAliasInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.RoleAlias != nil {
		v := *s.RoleAlias

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "roleAlias", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteRoleAliasOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteRoleAliasOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteRoleAliasOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteRoleAlias = "DeleteRoleAlias"

// DeleteRoleAliasRequest returns a request value for making API operation for
// AWS IoT.
//
// Deletes a role alias
//
//    // Example sending a request using DeleteRoleAliasRequest.
//    req := client.DeleteRoleAliasRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DeleteRoleAliasRequest(input *DeleteRoleAliasInput) DeleteRoleAliasRequest {
	op := &aws.Operation{
		Name:       opDeleteRoleAlias,
		HTTPMethod: "DELETE",
		HTTPPath:   "/role-aliases/{roleAlias}",
	}

	if input == nil {
		input = &DeleteRoleAliasInput{}
	}

	req := c.newRequest(op, input, &DeleteRoleAliasOutput{})
	return DeleteRoleAliasRequest{Request: req, Input: input, Copy: c.DeleteRoleAliasRequest}
}

// DeleteRoleAliasRequest is the request type for the
// DeleteRoleAlias API operation.
type DeleteRoleAliasRequest struct {
	*aws.Request
	Input *DeleteRoleAliasInput
	Copy  func(*DeleteRoleAliasInput) DeleteRoleAliasRequest
}

// Send marshals and sends the DeleteRoleAlias API request.
func (r DeleteRoleAliasRequest) Send(ctx context.Context) (*DeleteRoleAliasResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteRoleAliasResponse{
		DeleteRoleAliasOutput: r.Request.Data.(*DeleteRoleAliasOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteRoleAliasResponse is the response type for the
// DeleteRoleAlias API operation.
type DeleteRoleAliasResponse struct {
	*DeleteRoleAliasOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteRoleAlias request.
func (r *DeleteRoleAliasResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
