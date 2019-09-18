// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workmail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/workmail-2017-10-01/DisassociateDelegateFromResourceRequest
type DisassociateDelegateFromResourceInput struct {
	_ struct{} `type:"structure"`

	// The identifier for the member (user, group) to be removed from the resource's
	// delegates.
	//
	// EntityId is a required field
	EntityId *string `min:"12" type:"string" required:"true"`

	// The identifier for the organization under which the resource exists.
	//
	// OrganizationId is a required field
	OrganizationId *string `type:"string" required:"true"`

	// The identifier of the resource from which delegates' set members are removed.
	//
	// ResourceId is a required field
	ResourceId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DisassociateDelegateFromResourceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DisassociateDelegateFromResourceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DisassociateDelegateFromResourceInput"}

	if s.EntityId == nil {
		invalidParams.Add(aws.NewErrParamRequired("EntityId"))
	}
	if s.EntityId != nil && len(*s.EntityId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("EntityId", 12))
	}

	if s.OrganizationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("OrganizationId"))
	}

	if s.ResourceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/workmail-2017-10-01/DisassociateDelegateFromResourceResponse
type DisassociateDelegateFromResourceOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DisassociateDelegateFromResourceOutput) String() string {
	return awsutil.Prettify(s)
}

const opDisassociateDelegateFromResource = "DisassociateDelegateFromResource"

// DisassociateDelegateFromResourceRequest returns a request value for making API operation for
// Amazon WorkMail.
//
// Removes a member from the resource's set of delegates.
//
//    // Example sending a request using DisassociateDelegateFromResourceRequest.
//    req := client.DisassociateDelegateFromResourceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workmail-2017-10-01/DisassociateDelegateFromResource
func (c *Client) DisassociateDelegateFromResourceRequest(input *DisassociateDelegateFromResourceInput) DisassociateDelegateFromResourceRequest {
	op := &aws.Operation{
		Name:       opDisassociateDelegateFromResource,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DisassociateDelegateFromResourceInput{}
	}

	req := c.newRequest(op, input, &DisassociateDelegateFromResourceOutput{})
	return DisassociateDelegateFromResourceRequest{Request: req, Input: input, Copy: c.DisassociateDelegateFromResourceRequest}
}

// DisassociateDelegateFromResourceRequest is the request type for the
// DisassociateDelegateFromResource API operation.
type DisassociateDelegateFromResourceRequest struct {
	*aws.Request
	Input *DisassociateDelegateFromResourceInput
	Copy  func(*DisassociateDelegateFromResourceInput) DisassociateDelegateFromResourceRequest
}

// Send marshals and sends the DisassociateDelegateFromResource API request.
func (r DisassociateDelegateFromResourceRequest) Send(ctx context.Context) (*DisassociateDelegateFromResourceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DisassociateDelegateFromResourceResponse{
		DisassociateDelegateFromResourceOutput: r.Request.Data.(*DisassociateDelegateFromResourceOutput),
		response:                               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DisassociateDelegateFromResourceResponse is the response type for the
// DisassociateDelegateFromResource API operation.
type DisassociateDelegateFromResourceResponse struct {
	*DisassociateDelegateFromResourceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DisassociateDelegateFromResource request.
func (r *DisassociateDelegateFromResourceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
