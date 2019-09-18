// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package waf

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// A request to update an XssMatchSet.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-2015-08-24/UpdateXssMatchSetRequest
type UpdateXssMatchSetInput struct {
	_ struct{} `type:"structure"`

	// The value returned by the most recent call to GetChangeToken.
	//
	// ChangeToken is a required field
	ChangeToken *string `min:"1" type:"string" required:"true"`

	// An array of XssMatchSetUpdate objects that you want to insert into or delete
	// from an XssMatchSet. For more information, see the applicable data types:
	//
	//    * XssMatchSetUpdate: Contains Action and XssMatchTuple
	//
	//    * XssMatchTuple: Contains FieldToMatch and TextTransformation
	//
	//    * FieldToMatch: Contains Data and Type
	//
	// Updates is a required field
	Updates []XssMatchSetUpdate `min:"1" type:"list" required:"true"`

	// The XssMatchSetId of the XssMatchSet that you want to update. XssMatchSetId
	// is returned by CreateXssMatchSet and by ListXssMatchSets.
	//
	// XssMatchSetId is a required field
	XssMatchSetId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateXssMatchSetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateXssMatchSetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateXssMatchSetInput"}

	if s.ChangeToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("ChangeToken"))
	}
	if s.ChangeToken != nil && len(*s.ChangeToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ChangeToken", 1))
	}

	if s.Updates == nil {
		invalidParams.Add(aws.NewErrParamRequired("Updates"))
	}
	if s.Updates != nil && len(s.Updates) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Updates", 1))
	}

	if s.XssMatchSetId == nil {
		invalidParams.Add(aws.NewErrParamRequired("XssMatchSetId"))
	}
	if s.XssMatchSetId != nil && len(*s.XssMatchSetId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("XssMatchSetId", 1))
	}
	if s.Updates != nil {
		for i, v := range s.Updates {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Updates", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The response to an UpdateXssMatchSets request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-2015-08-24/UpdateXssMatchSetResponse
type UpdateXssMatchSetOutput struct {
	_ struct{} `type:"structure"`

	// The ChangeToken that you used to submit the UpdateXssMatchSet request. You
	// can also use this value to query the status of the request. For more information,
	// see GetChangeTokenStatus.
	ChangeToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s UpdateXssMatchSetOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateXssMatchSet = "UpdateXssMatchSet"

// UpdateXssMatchSetRequest returns a request value for making API operation for
// AWS WAF.
//
// Inserts or deletes XssMatchTuple objects (filters) in an XssMatchSet. For
// each XssMatchTuple object, you specify the following values:
//
//    * Action: Whether to insert the object into or delete the object from
//    the array. To change an XssMatchTuple, you delete the existing object
//    and add a new one.
//
//    * FieldToMatch: The part of web requests that you want AWS WAF to inspect
//    and, if you want AWS WAF to inspect a header or custom query parameter,
//    the name of the header or parameter.
//
//    * TextTransformation: Which text transformation, if any, to perform on
//    the web request before inspecting the request for cross-site scripting
//    attacks. You can only specify a single type of TextTransformation.
//
// You use XssMatchSet objects to specify which CloudFront requests that you
// want to allow, block, or count. For example, if you're receiving requests
// that contain cross-site scripting attacks in the request body and you want
// to block the requests, you can create an XssMatchSet with the applicable
// settings, and then configure AWS WAF to block the requests.
//
// To create and configure an XssMatchSet, perform the following steps:
//
// Submit a CreateXssMatchSet request.
//
// Use GetChangeToken to get the change token that you provide in the ChangeToken
// parameter of an UpdateIPSet request.
//
// Submit an UpdateXssMatchSet request to specify the parts of web requests
// that you want AWS WAF to inspect for cross-site scripting attacks.
//
// For more information about how to use the AWS WAF API to allow or block HTTP
// requests, see the AWS WAF Developer Guide (https://docs.aws.amazon.com/waf/latest/developerguide/).
//
//    // Example sending a request using UpdateXssMatchSetRequest.
//    req := client.UpdateXssMatchSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-2015-08-24/UpdateXssMatchSet
func (c *Client) UpdateXssMatchSetRequest(input *UpdateXssMatchSetInput) UpdateXssMatchSetRequest {
	op := &aws.Operation{
		Name:       opUpdateXssMatchSet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateXssMatchSetInput{}
	}

	req := c.newRequest(op, input, &UpdateXssMatchSetOutput{})
	return UpdateXssMatchSetRequest{Request: req, Input: input, Copy: c.UpdateXssMatchSetRequest}
}

// UpdateXssMatchSetRequest is the request type for the
// UpdateXssMatchSet API operation.
type UpdateXssMatchSetRequest struct {
	*aws.Request
	Input *UpdateXssMatchSetInput
	Copy  func(*UpdateXssMatchSetInput) UpdateXssMatchSetRequest
}

// Send marshals and sends the UpdateXssMatchSet API request.
func (r UpdateXssMatchSetRequest) Send(ctx context.Context) (*UpdateXssMatchSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateXssMatchSetResponse{
		UpdateXssMatchSetOutput: r.Request.Data.(*UpdateXssMatchSetOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateXssMatchSetResponse is the response type for the
// UpdateXssMatchSet API operation.
type UpdateXssMatchSetResponse struct {
	*UpdateXssMatchSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateXssMatchSet request.
func (r *UpdateXssMatchSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
