// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticbeanstalk

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Results message indicating whether a CNAME is available.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticbeanstalk-2010-12-01/CheckDNSAvailabilityMessage
type CheckDNSAvailabilityInput struct {
	_ struct{} `type:"structure"`

	// The prefix used when this CNAME is reserved.
	//
	// CNAMEPrefix is a required field
	CNAMEPrefix *string `min:"4" type:"string" required:"true"`
}

// String returns the string representation
func (s CheckDNSAvailabilityInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CheckDNSAvailabilityInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CheckDNSAvailabilityInput"}

	if s.CNAMEPrefix == nil {
		invalidParams.Add(aws.NewErrParamRequired("CNAMEPrefix"))
	}
	if s.CNAMEPrefix != nil && len(*s.CNAMEPrefix) < 4 {
		invalidParams.Add(aws.NewErrParamMinLen("CNAMEPrefix", 4))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Indicates if the specified CNAME is available.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticbeanstalk-2010-12-01/CheckDNSAvailabilityResultMessage
type CheckDNSAvailabilityOutput struct {
	_ struct{} `type:"structure"`

	// Indicates if the specified CNAME is available:
	//
	//    * true : The CNAME is available.
	//
	//    * false : The CNAME is not available.
	Available *bool `type:"boolean"`

	// The fully qualified CNAME to reserve when CreateEnvironment is called with
	// the provided prefix.
	FullyQualifiedCNAME *string `min:"1" type:"string"`
}

// String returns the string representation
func (s CheckDNSAvailabilityOutput) String() string {
	return awsutil.Prettify(s)
}

const opCheckDNSAvailability = "CheckDNSAvailability"

// CheckDNSAvailabilityRequest returns a request value for making API operation for
// AWS Elastic Beanstalk.
//
// Checks if the specified CNAME is available.
//
//    // Example sending a request using CheckDNSAvailabilityRequest.
//    req := client.CheckDNSAvailabilityRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticbeanstalk-2010-12-01/CheckDNSAvailability
func (c *Client) CheckDNSAvailabilityRequest(input *CheckDNSAvailabilityInput) CheckDNSAvailabilityRequest {
	op := &aws.Operation{
		Name:       opCheckDNSAvailability,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CheckDNSAvailabilityInput{}
	}

	req := c.newRequest(op, input, &CheckDNSAvailabilityOutput{})
	return CheckDNSAvailabilityRequest{Request: req, Input: input, Copy: c.CheckDNSAvailabilityRequest}
}

// CheckDNSAvailabilityRequest is the request type for the
// CheckDNSAvailability API operation.
type CheckDNSAvailabilityRequest struct {
	*aws.Request
	Input *CheckDNSAvailabilityInput
	Copy  func(*CheckDNSAvailabilityInput) CheckDNSAvailabilityRequest
}

// Send marshals and sends the CheckDNSAvailability API request.
func (r CheckDNSAvailabilityRequest) Send(ctx context.Context) (*CheckDNSAvailabilityResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CheckDNSAvailabilityResponse{
		CheckDNSAvailabilityOutput: r.Request.Data.(*CheckDNSAvailabilityOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CheckDNSAvailabilityResponse is the response type for the
// CheckDNSAvailability API operation.
type CheckDNSAvailabilityResponse struct {
	*CheckDNSAvailabilityOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CheckDNSAvailability request.
func (r *CheckDNSAvailabilityResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
