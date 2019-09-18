// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53resolver

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53resolver-2018-04-01/DeleteResolverRuleRequest
type DeleteResolverRuleInput struct {
	_ struct{} `type:"structure"`

	// The ID of the resolver rule that you want to delete.
	//
	// ResolverRuleId is a required field
	ResolverRuleId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteResolverRuleInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteResolverRuleInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteResolverRuleInput"}

	if s.ResolverRuleId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResolverRuleId"))
	}
	if s.ResolverRuleId != nil && len(*s.ResolverRuleId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResolverRuleId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53resolver-2018-04-01/DeleteResolverRuleResponse
type DeleteResolverRuleOutput struct {
	_ struct{} `type:"structure"`

	// Information about the DeleteResolverRule request, including the status of
	// the request.
	ResolverRule *ResolverRule `type:"structure"`
}

// String returns the string representation
func (s DeleteResolverRuleOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteResolverRule = "DeleteResolverRule"

// DeleteResolverRuleRequest returns a request value for making API operation for
// Amazon Route 53 Resolver.
//
// Deletes a resolver rule. Before you can delete a resolver rule, you must
// disassociate it from all the VPCs that you associated the resolver rule with.
// For more infomation, see DisassociateResolverRule.
//
//    // Example sending a request using DeleteResolverRuleRequest.
//    req := client.DeleteResolverRuleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53resolver-2018-04-01/DeleteResolverRule
func (c *Client) DeleteResolverRuleRequest(input *DeleteResolverRuleInput) DeleteResolverRuleRequest {
	op := &aws.Operation{
		Name:       opDeleteResolverRule,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteResolverRuleInput{}
	}

	req := c.newRequest(op, input, &DeleteResolverRuleOutput{})
	return DeleteResolverRuleRequest{Request: req, Input: input, Copy: c.DeleteResolverRuleRequest}
}

// DeleteResolverRuleRequest is the request type for the
// DeleteResolverRule API operation.
type DeleteResolverRuleRequest struct {
	*aws.Request
	Input *DeleteResolverRuleInput
	Copy  func(*DeleteResolverRuleInput) DeleteResolverRuleRequest
}

// Send marshals and sends the DeleteResolverRule API request.
func (r DeleteResolverRuleRequest) Send(ctx context.Context) (*DeleteResolverRuleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteResolverRuleResponse{
		DeleteResolverRuleOutput: r.Request.Data.(*DeleteResolverRuleOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteResolverRuleResponse is the response type for the
// DeleteResolverRule API operation.
type DeleteResolverRuleResponse struct {
	*DeleteResolverRuleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteResolverRule request.
func (r *DeleteResolverRuleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
