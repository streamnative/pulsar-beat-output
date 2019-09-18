// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package waf

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-2015-08-24/GetSizeConstraintSetRequest
type GetSizeConstraintSetInput struct {
	_ struct{} `type:"structure"`

	// The SizeConstraintSetId of the SizeConstraintSet that you want to get. SizeConstraintSetId
	// is returned by CreateSizeConstraintSet and by ListSizeConstraintSets.
	//
	// SizeConstraintSetId is a required field
	SizeConstraintSetId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetSizeConstraintSetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetSizeConstraintSetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetSizeConstraintSetInput"}

	if s.SizeConstraintSetId == nil {
		invalidParams.Add(aws.NewErrParamRequired("SizeConstraintSetId"))
	}
	if s.SizeConstraintSetId != nil && len(*s.SizeConstraintSetId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("SizeConstraintSetId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-2015-08-24/GetSizeConstraintSetResponse
type GetSizeConstraintSetOutput struct {
	_ struct{} `type:"structure"`

	// Information about the SizeConstraintSet that you specified in the GetSizeConstraintSet
	// request. For more information, see the following topics:
	//
	//    * SizeConstraintSet: Contains SizeConstraintSetId, SizeConstraints, and
	//    Name
	//
	//    * SizeConstraints: Contains an array of SizeConstraint objects. Each SizeConstraint
	//    object contains FieldToMatch, TextTransformation, ComparisonOperator,
	//    and Size
	//
	//    * FieldToMatch: Contains Data and Type
	SizeConstraintSet *SizeConstraintSet `type:"structure"`
}

// String returns the string representation
func (s GetSizeConstraintSetOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetSizeConstraintSet = "GetSizeConstraintSet"

// GetSizeConstraintSetRequest returns a request value for making API operation for
// AWS WAF.
//
// Returns the SizeConstraintSet specified by SizeConstraintSetId.
//
//    // Example sending a request using GetSizeConstraintSetRequest.
//    req := client.GetSizeConstraintSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-2015-08-24/GetSizeConstraintSet
func (c *Client) GetSizeConstraintSetRequest(input *GetSizeConstraintSetInput) GetSizeConstraintSetRequest {
	op := &aws.Operation{
		Name:       opGetSizeConstraintSet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetSizeConstraintSetInput{}
	}

	req := c.newRequest(op, input, &GetSizeConstraintSetOutput{})
	return GetSizeConstraintSetRequest{Request: req, Input: input, Copy: c.GetSizeConstraintSetRequest}
}

// GetSizeConstraintSetRequest is the request type for the
// GetSizeConstraintSet API operation.
type GetSizeConstraintSetRequest struct {
	*aws.Request
	Input *GetSizeConstraintSetInput
	Copy  func(*GetSizeConstraintSetInput) GetSizeConstraintSetRequest
}

// Send marshals and sends the GetSizeConstraintSet API request.
func (r GetSizeConstraintSetRequest) Send(ctx context.Context) (*GetSizeConstraintSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetSizeConstraintSetResponse{
		GetSizeConstraintSetOutput: r.Request.Data.(*GetSizeConstraintSetOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetSizeConstraintSetResponse is the response type for the
// GetSizeConstraintSet API operation.
type GetSizeConstraintSetResponse struct {
	*GetSizeConstraintSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetSizeConstraintSet request.
func (r *GetSizeConstraintSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
