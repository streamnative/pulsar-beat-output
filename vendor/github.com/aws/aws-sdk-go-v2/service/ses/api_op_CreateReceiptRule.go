// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ses

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents a request to create a receipt rule. You use receipt rules to receive
// email with Amazon SES. For more information, see the Amazon SES Developer
// Guide (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-concepts.html).
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/CreateReceiptRuleRequest
type CreateReceiptRuleInput struct {
	_ struct{} `type:"structure"`

	// The name of an existing rule after which the new rule will be placed. If
	// this parameter is null, the new rule will be inserted at the beginning of
	// the rule list.
	After *string `type:"string"`

	// A data structure that contains the specified rule's name, actions, recipients,
	// domains, enabled status, scan status, and TLS policy.
	//
	// Rule is a required field
	Rule *ReceiptRule `type:"structure" required:"true"`

	// The name of the rule set that the receipt rule will be added to.
	//
	// RuleSetName is a required field
	RuleSetName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateReceiptRuleInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateReceiptRuleInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateReceiptRuleInput"}

	if s.Rule == nil {
		invalidParams.Add(aws.NewErrParamRequired("Rule"))
	}

	if s.RuleSetName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RuleSetName"))
	}
	if s.Rule != nil {
		if err := s.Rule.Validate(); err != nil {
			invalidParams.AddNested("Rule", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// An empty element returned on a successful request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/CreateReceiptRuleResponse
type CreateReceiptRuleOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CreateReceiptRuleOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateReceiptRule = "CreateReceiptRule"

// CreateReceiptRuleRequest returns a request value for making API operation for
// Amazon Simple Email Service.
//
// Creates a receipt rule.
//
// For information about setting up receipt rules, see the Amazon SES Developer
// Guide (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-receipt-rules.html).
//
// You can execute this operation no more than once per second.
//
//    // Example sending a request using CreateReceiptRuleRequest.
//    req := client.CreateReceiptRuleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/CreateReceiptRule
func (c *Client) CreateReceiptRuleRequest(input *CreateReceiptRuleInput) CreateReceiptRuleRequest {
	op := &aws.Operation{
		Name:       opCreateReceiptRule,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateReceiptRuleInput{}
	}

	req := c.newRequest(op, input, &CreateReceiptRuleOutput{})
	return CreateReceiptRuleRequest{Request: req, Input: input, Copy: c.CreateReceiptRuleRequest}
}

// CreateReceiptRuleRequest is the request type for the
// CreateReceiptRule API operation.
type CreateReceiptRuleRequest struct {
	*aws.Request
	Input *CreateReceiptRuleInput
	Copy  func(*CreateReceiptRuleInput) CreateReceiptRuleRequest
}

// Send marshals and sends the CreateReceiptRule API request.
func (r CreateReceiptRuleRequest) Send(ctx context.Context) (*CreateReceiptRuleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateReceiptRuleResponse{
		CreateReceiptRuleOutput: r.Request.Data.(*CreateReceiptRuleOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateReceiptRuleResponse is the response type for the
// CreateReceiptRule API operation.
type CreateReceiptRuleResponse struct {
	*CreateReceiptRuleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateReceiptRule request.
func (r *CreateReceiptRuleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
