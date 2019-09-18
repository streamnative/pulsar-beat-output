// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package wafregional

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/waf"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-regional-2016-11-28/CreateIPSetRequest
type CreateIPSetInput struct {
	_ struct{} `type:"structure"`

	// The value returned by the most recent call to GetChangeToken.
	//
	// ChangeToken is a required field
	ChangeToken *string `min:"1" type:"string" required:"true"`

	// A friendly name or description of the IPSet. You can't change Name after
	// you create the IPSet.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateIPSetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateIPSetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateIPSetInput"}

	if s.ChangeToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("ChangeToken"))
	}
	if s.ChangeToken != nil && len(*s.ChangeToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ChangeToken", 1))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-regional-2016-11-28/CreateIPSetResponse
type CreateIPSetOutput struct {
	_ struct{} `type:"structure"`

	// The ChangeToken that you used to submit the CreateIPSet request. You can
	// also use this value to query the status of the request. For more information,
	// see GetChangeTokenStatus.
	ChangeToken *string `min:"1" type:"string"`

	// The IPSet returned in the CreateIPSet response.
	IPSet *waf.IPSet `type:"structure"`
}

// String returns the string representation
func (s CreateIPSetOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateIPSet = "CreateIPSet"

// CreateIPSetRequest returns a request value for making API operation for
// AWS WAF Regional.
//
// Creates an IPSet, which you use to specify which web requests that you want
// to allow or block based on the IP addresses that the requests originate from.
// For example, if you're receiving a lot of requests from one or more individual
// IP addresses or one or more ranges of IP addresses and you want to block
// the requests, you can create an IPSet that contains those IP addresses and
// then configure AWS WAF to block the requests.
//
// To create and configure an IPSet, perform the following steps:
//
// Use GetChangeToken to get the change token that you provide in the ChangeToken
// parameter of a CreateIPSet request.
//
// Submit a CreateIPSet request.
//
// Use GetChangeToken to get the change token that you provide in the ChangeToken
// parameter of an UpdateIPSet request.
//
// Submit an UpdateIPSet request to specify the IP addresses that you want AWS
// WAF to watch for.
//
// For more information about how to use the AWS WAF API to allow or block HTTP
// requests, see the AWS WAF Developer Guide (https://docs.aws.amazon.com/waf/latest/developerguide/).
//
//    // Example sending a request using CreateIPSetRequest.
//    req := client.CreateIPSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-regional-2016-11-28/CreateIPSet
func (c *Client) CreateIPSetRequest(input *CreateIPSetInput) CreateIPSetRequest {
	op := &aws.Operation{
		Name:       opCreateIPSet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateIPSetInput{}
	}

	req := c.newRequest(op, input, &CreateIPSetOutput{})
	return CreateIPSetRequest{Request: req, Input: input, Copy: c.CreateIPSetRequest}
}

// CreateIPSetRequest is the request type for the
// CreateIPSet API operation.
type CreateIPSetRequest struct {
	*aws.Request
	Input *CreateIPSetInput
	Copy  func(*CreateIPSetInput) CreateIPSetRequest
}

// Send marshals and sends the CreateIPSet API request.
func (r CreateIPSetRequest) Send(ctx context.Context) (*CreateIPSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateIPSetResponse{
		CreateIPSetOutput: r.Request.Data.(*CreateIPSetOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateIPSetResponse is the response type for the
// CreateIPSet API operation.
type CreateIPSetResponse struct {
	*CreateIPSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateIPSet request.
func (r *CreateIPSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
