// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package wafregional

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/waf"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-regional-2016-11-28/CreateSizeConstraintSetRequest
type CreateSizeConstraintSetInput struct {
	_ struct{} `type:"structure"`

	// The value returned by the most recent call to GetChangeToken.
	//
	// ChangeToken is a required field
	ChangeToken *string `min:"1" type:"string" required:"true"`

	// A friendly name or description of the SizeConstraintSet. You can't change
	// Name after you create a SizeConstraintSet.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateSizeConstraintSetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateSizeConstraintSetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateSizeConstraintSetInput"}

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

// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-regional-2016-11-28/CreateSizeConstraintSetResponse
type CreateSizeConstraintSetOutput struct {
	_ struct{} `type:"structure"`

	// The ChangeToken that you used to submit the CreateSizeConstraintSet request.
	// You can also use this value to query the status of the request. For more
	// information, see GetChangeTokenStatus.
	ChangeToken *string `min:"1" type:"string"`

	// A SizeConstraintSet that contains no SizeConstraint objects.
	SizeConstraintSet *waf.SizeConstraintSet `type:"structure"`
}

// String returns the string representation
func (s CreateSizeConstraintSetOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateSizeConstraintSet = "CreateSizeConstraintSet"

// CreateSizeConstraintSetRequest returns a request value for making API operation for
// AWS WAF Regional.
//
// Creates a SizeConstraintSet. You then use UpdateSizeConstraintSet to identify
// the part of a web request that you want AWS WAF to check for length, such
// as the length of the User-Agent header or the length of the query string.
// For example, you can create a SizeConstraintSet that matches any requests
// that have a query string that is longer than 100 bytes. You can then configure
// AWS WAF to reject those requests.
//
// To create and configure a SizeConstraintSet, perform the following steps:
//
// Use GetChangeToken to get the change token that you provide in the ChangeToken
// parameter of a CreateSizeConstraintSet request.
//
// Submit a CreateSizeConstraintSet request.
//
// Use GetChangeToken to get the change token that you provide in the ChangeToken
// parameter of an UpdateSizeConstraintSet request.
//
// Submit an UpdateSizeConstraintSet request to specify the part of the request
// that you want AWS WAF to inspect (for example, the header or the URI) and
// the value that you want AWS WAF to watch for.
//
// For more information about how to use the AWS WAF API to allow or block HTTP
// requests, see the AWS WAF Developer Guide (https://docs.aws.amazon.com/waf/latest/developerguide/).
//
//    // Example sending a request using CreateSizeConstraintSetRequest.
//    req := client.CreateSizeConstraintSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-regional-2016-11-28/CreateSizeConstraintSet
func (c *Client) CreateSizeConstraintSetRequest(input *CreateSizeConstraintSetInput) CreateSizeConstraintSetRequest {
	op := &aws.Operation{
		Name:       opCreateSizeConstraintSet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateSizeConstraintSetInput{}
	}

	req := c.newRequest(op, input, &CreateSizeConstraintSetOutput{})
	return CreateSizeConstraintSetRequest{Request: req, Input: input, Copy: c.CreateSizeConstraintSetRequest}
}

// CreateSizeConstraintSetRequest is the request type for the
// CreateSizeConstraintSet API operation.
type CreateSizeConstraintSetRequest struct {
	*aws.Request
	Input *CreateSizeConstraintSetInput
	Copy  func(*CreateSizeConstraintSetInput) CreateSizeConstraintSetRequest
}

// Send marshals and sends the CreateSizeConstraintSet API request.
func (r CreateSizeConstraintSetRequest) Send(ctx context.Context) (*CreateSizeConstraintSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateSizeConstraintSetResponse{
		CreateSizeConstraintSetOutput: r.Request.Data.(*CreateSizeConstraintSetOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateSizeConstraintSetResponse is the response type for the
// CreateSizeConstraintSet API operation.
type CreateSizeConstraintSetResponse struct {
	*CreateSizeConstraintSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateSizeConstraintSet request.
func (r *CreateSizeConstraintSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
