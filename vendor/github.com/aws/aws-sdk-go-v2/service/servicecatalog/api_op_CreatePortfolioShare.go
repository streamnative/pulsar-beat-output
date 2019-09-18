// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicecatalog

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/CreatePortfolioShareInput
type CreatePortfolioShareInput struct {
	_ struct{} `type:"structure"`

	// The language code.
	//
	//    * en - English (default)
	//
	//    * jp - Japanese
	//
	//    * zh - Chinese
	AcceptLanguage *string `type:"string"`

	// The AWS account ID. For example, 123456789012.
	AccountId *string `type:"string"`

	// The organization node to whom you are going to share. If OrganizationNode
	// is passed in, PortfolioShare will be created for the node and its children
	// (when applies), and a PortfolioShareToken will be returned in the output
	// in order for the administrator to monitor the status of the PortfolioShare
	// creation process.
	OrganizationNode *OrganizationNode `type:"structure"`

	// The portfolio identifier.
	//
	// PortfolioId is a required field
	PortfolioId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreatePortfolioShareInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreatePortfolioShareInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreatePortfolioShareInput"}

	if s.PortfolioId == nil {
		invalidParams.Add(aws.NewErrParamRequired("PortfolioId"))
	}
	if s.PortfolioId != nil && len(*s.PortfolioId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PortfolioId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/CreatePortfolioShareOutput
type CreatePortfolioShareOutput struct {
	_ struct{} `type:"structure"`

	// The portfolio share unique identifier. This will only be returned if portfolio
	// is shared to an organization node.
	PortfolioShareToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s CreatePortfolioShareOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreatePortfolioShare = "CreatePortfolioShare"

// CreatePortfolioShareRequest returns a request value for making API operation for
// AWS Service Catalog.
//
// Shares the specified portfolio with the specified account or organization
// node. Shares to an organization node can only be created by the master account
// of an Organization. AWSOrganizationsAccess must be enabled in order to create
// a portfolio share to an organization node.
//
//    // Example sending a request using CreatePortfolioShareRequest.
//    req := client.CreatePortfolioShareRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/CreatePortfolioShare
func (c *Client) CreatePortfolioShareRequest(input *CreatePortfolioShareInput) CreatePortfolioShareRequest {
	op := &aws.Operation{
		Name:       opCreatePortfolioShare,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreatePortfolioShareInput{}
	}

	req := c.newRequest(op, input, &CreatePortfolioShareOutput{})
	return CreatePortfolioShareRequest{Request: req, Input: input, Copy: c.CreatePortfolioShareRequest}
}

// CreatePortfolioShareRequest is the request type for the
// CreatePortfolioShare API operation.
type CreatePortfolioShareRequest struct {
	*aws.Request
	Input *CreatePortfolioShareInput
	Copy  func(*CreatePortfolioShareInput) CreatePortfolioShareRequest
}

// Send marshals and sends the CreatePortfolioShare API request.
func (r CreatePortfolioShareRequest) Send(ctx context.Context) (*CreatePortfolioShareResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreatePortfolioShareResponse{
		CreatePortfolioShareOutput: r.Request.Data.(*CreatePortfolioShareOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreatePortfolioShareResponse is the response type for the
// CreatePortfolioShare API operation.
type CreatePortfolioShareResponse struct {
	*CreatePortfolioShareOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreatePortfolioShare request.
func (r *CreatePortfolioShareResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
