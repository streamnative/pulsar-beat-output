// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package fms

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/fms-2018-01-01/GetProtectionStatusRequest
type GetProtectionStatusInput struct {
	_ struct{} `type:"structure"`

	// The end of the time period to query for the attacks. This is a timestamp
	// type. The sample request above indicates a number type because the default
	// used by AWS Firewall Manager is Unix time in seconds. However, any valid
	// timestamp format is allowed.
	EndTime *time.Time `type:"timestamp"`

	// Specifies the number of objects that you want AWS Firewall Manager to return
	// for this request. If you have more objects than the number that you specify
	// for MaxResults, the response includes a NextToken value that you can use
	// to get another batch of objects.
	MaxResults *int64 `min:"1" type:"integer"`

	// The AWS account that is in scope of the policy that you want to get the details
	// for.
	MemberAccountId *string `min:"1" type:"string"`

	// If you specify a value for MaxResults and you have more objects than the
	// number that you specify for MaxResults, AWS Firewall Manager returns a NextToken
	// value in the response that allows you to list another group of objects. For
	// the second and subsequent GetProtectionStatus requests, specify the value
	// of NextToken from the previous response to get information about another
	// batch of objects.
	NextToken *string `min:"1" type:"string"`

	// The ID of the policy for which you want to get the attack information.
	//
	// PolicyId is a required field
	PolicyId *string `min:"36" type:"string" required:"true"`

	// The start of the time period to query for the attacks. This is a timestamp
	// type. The sample request above indicates a number type because the default
	// used by AWS Firewall Manager is Unix time in seconds. However, any valid
	// timestamp format is allowed.
	StartTime *time.Time `type:"timestamp"`
}

// String returns the string representation
func (s GetProtectionStatusInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetProtectionStatusInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetProtectionStatusInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.MemberAccountId != nil && len(*s.MemberAccountId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("MemberAccountId", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if s.PolicyId == nil {
		invalidParams.Add(aws.NewErrParamRequired("PolicyId"))
	}
	if s.PolicyId != nil && len(*s.PolicyId) < 36 {
		invalidParams.Add(aws.NewErrParamMinLen("PolicyId", 36))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/fms-2018-01-01/GetProtectionStatusResponse
type GetProtectionStatusOutput struct {
	_ struct{} `type:"structure"`

	// The ID of the AWS Firewall administrator account for this policy.
	AdminAccountId *string `min:"1" type:"string"`

	// Details about the attack, including the following:
	//
	//    * Attack type
	//
	//    * Account ID
	//
	//    * ARN of the resource attacked
	//
	//    * Start time of the attack
	//
	//    * End time of the attack (ongoing attacks will not have an end time)
	//
	// The details are in JSON format. An example is shown in the Examples section
	// below.
	Data *string `type:"string"`

	// If you have more objects than the number that you specified for MaxResults
	// in the request, the response includes a NextToken value. To list more objects,
	// submit another GetProtectionStatus request, and specify the NextToken value
	// from the response in the NextToken value in the next request.
	//
	// AWS SDKs provide auto-pagination that identify NextToken in a response and
	// make subsequent request calls automatically on your behalf. However, this
	// feature is not supported by GetProtectionStatus. You must submit subsequent
	// requests with NextToken using your own processes.
	NextToken *string `min:"1" type:"string"`

	// The service type that is protected by the policy. Currently, this is always
	// SHIELD_ADVANCED.
	ServiceType SecurityServiceType `type:"string" enum:"true"`
}

// String returns the string representation
func (s GetProtectionStatusOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetProtectionStatus = "GetProtectionStatus"

// GetProtectionStatusRequest returns a request value for making API operation for
// Firewall Management Service.
//
// If you created a Shield Advanced policy, returns policy-level attack summary
// information in the event of a potential DDoS attack.
//
//    // Example sending a request using GetProtectionStatusRequest.
//    req := client.GetProtectionStatusRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/fms-2018-01-01/GetProtectionStatus
func (c *Client) GetProtectionStatusRequest(input *GetProtectionStatusInput) GetProtectionStatusRequest {
	op := &aws.Operation{
		Name:       opGetProtectionStatus,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetProtectionStatusInput{}
	}

	req := c.newRequest(op, input, &GetProtectionStatusOutput{})
	return GetProtectionStatusRequest{Request: req, Input: input, Copy: c.GetProtectionStatusRequest}
}

// GetProtectionStatusRequest is the request type for the
// GetProtectionStatus API operation.
type GetProtectionStatusRequest struct {
	*aws.Request
	Input *GetProtectionStatusInput
	Copy  func(*GetProtectionStatusInput) GetProtectionStatusRequest
}

// Send marshals and sends the GetProtectionStatus API request.
func (r GetProtectionStatusRequest) Send(ctx context.Context) (*GetProtectionStatusResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetProtectionStatusResponse{
		GetProtectionStatusOutput: r.Request.Data.(*GetProtectionStatusOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetProtectionStatusResponse is the response type for the
// GetProtectionStatus API operation.
type GetProtectionStatusResponse struct {
	*GetProtectionStatusOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetProtectionStatus request.
func (r *GetProtectionStatusResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
