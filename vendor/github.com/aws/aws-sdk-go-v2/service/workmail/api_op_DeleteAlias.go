// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workmail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/workmail-2017-10-01/DeleteAliasRequest
type DeleteAliasInput struct {
	_ struct{} `type:"structure"`

	// The aliases to be removed from the user's set of aliases. Duplicate entries
	// in the list are collapsed into single entries (the list is transformed into
	// a set).
	//
	// Alias is a required field
	Alias *string `min:"1" type:"string" required:"true"`

	// The identifier for the member (user or group) from which to have the aliases
	// removed.
	//
	// EntityId is a required field
	EntityId *string `min:"12" type:"string" required:"true"`

	// The identifier for the organization under which the user exists.
	//
	// OrganizationId is a required field
	OrganizationId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteAliasInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteAliasInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteAliasInput"}

	if s.Alias == nil {
		invalidParams.Add(aws.NewErrParamRequired("Alias"))
	}
	if s.Alias != nil && len(*s.Alias) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Alias", 1))
	}

	if s.EntityId == nil {
		invalidParams.Add(aws.NewErrParamRequired("EntityId"))
	}
	if s.EntityId != nil && len(*s.EntityId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("EntityId", 12))
	}

	if s.OrganizationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("OrganizationId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/workmail-2017-10-01/DeleteAliasResponse
type DeleteAliasOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteAliasOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteAlias = "DeleteAlias"

// DeleteAliasRequest returns a request value for making API operation for
// Amazon WorkMail.
//
// Remove one or more specified aliases from a set of aliases for a given user.
//
//    // Example sending a request using DeleteAliasRequest.
//    req := client.DeleteAliasRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workmail-2017-10-01/DeleteAlias
func (c *Client) DeleteAliasRequest(input *DeleteAliasInput) DeleteAliasRequest {
	op := &aws.Operation{
		Name:       opDeleteAlias,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteAliasInput{}
	}

	req := c.newRequest(op, input, &DeleteAliasOutput{})
	return DeleteAliasRequest{Request: req, Input: input, Copy: c.DeleteAliasRequest}
}

// DeleteAliasRequest is the request type for the
// DeleteAlias API operation.
type DeleteAliasRequest struct {
	*aws.Request
	Input *DeleteAliasInput
	Copy  func(*DeleteAliasInput) DeleteAliasRequest
}

// Send marshals and sends the DeleteAlias API request.
func (r DeleteAliasRequest) Send(ctx context.Context) (*DeleteAliasResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteAliasResponse{
		DeleteAliasOutput: r.Request.Data.(*DeleteAliasOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteAliasResponse is the response type for the
// DeleteAlias API operation.
type DeleteAliasResponse struct {
	*DeleteAliasOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteAlias request.
func (r *DeleteAliasResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
