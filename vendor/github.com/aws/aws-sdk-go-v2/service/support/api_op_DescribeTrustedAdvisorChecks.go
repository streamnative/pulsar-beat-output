// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package support

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/support-2013-04-15/DescribeTrustedAdvisorChecksRequest
type DescribeTrustedAdvisorChecksInput struct {
	_ struct{} `type:"structure"`

	// The ISO 639-1 code for the language in which AWS provides support. AWS Support
	// currently supports English ("en") and Japanese ("ja"). Language parameters
	// must be passed explicitly for operations that take them.
	//
	// Language is a required field
	Language *string `locationName:"language" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeTrustedAdvisorChecksInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeTrustedAdvisorChecksInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeTrustedAdvisorChecksInput"}

	if s.Language == nil {
		invalidParams.Add(aws.NewErrParamRequired("Language"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Information about the Trusted Advisor checks returned by the DescribeTrustedAdvisorChecks
// operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/support-2013-04-15/DescribeTrustedAdvisorChecksResponse
type DescribeTrustedAdvisorChecksOutput struct {
	_ struct{} `type:"structure"`

	// Information about all available Trusted Advisor checks.
	//
	// Checks is a required field
	Checks []TrustedAdvisorCheckDescription `locationName:"checks" type:"list" required:"true"`
}

// String returns the string representation
func (s DescribeTrustedAdvisorChecksOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeTrustedAdvisorChecks = "DescribeTrustedAdvisorChecks"

// DescribeTrustedAdvisorChecksRequest returns a request value for making API operation for
// AWS Support.
//
// Returns information about all available Trusted Advisor checks, including
// name, ID, category, description, and metadata. You must specify a language
// code; English ("en") and Japanese ("ja") are currently supported. The response
// contains a TrustedAdvisorCheckDescription for each check.
//
//    // Example sending a request using DescribeTrustedAdvisorChecksRequest.
//    req := client.DescribeTrustedAdvisorChecksRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/support-2013-04-15/DescribeTrustedAdvisorChecks
func (c *Client) DescribeTrustedAdvisorChecksRequest(input *DescribeTrustedAdvisorChecksInput) DescribeTrustedAdvisorChecksRequest {
	op := &aws.Operation{
		Name:       opDescribeTrustedAdvisorChecks,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeTrustedAdvisorChecksInput{}
	}

	req := c.newRequest(op, input, &DescribeTrustedAdvisorChecksOutput{})
	return DescribeTrustedAdvisorChecksRequest{Request: req, Input: input, Copy: c.DescribeTrustedAdvisorChecksRequest}
}

// DescribeTrustedAdvisorChecksRequest is the request type for the
// DescribeTrustedAdvisorChecks API operation.
type DescribeTrustedAdvisorChecksRequest struct {
	*aws.Request
	Input *DescribeTrustedAdvisorChecksInput
	Copy  func(*DescribeTrustedAdvisorChecksInput) DescribeTrustedAdvisorChecksRequest
}

// Send marshals and sends the DescribeTrustedAdvisorChecks API request.
func (r DescribeTrustedAdvisorChecksRequest) Send(ctx context.Context) (*DescribeTrustedAdvisorChecksResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeTrustedAdvisorChecksResponse{
		DescribeTrustedAdvisorChecksOutput: r.Request.Data.(*DescribeTrustedAdvisorChecksOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeTrustedAdvisorChecksResponse is the response type for the
// DescribeTrustedAdvisorChecks API operation.
type DescribeTrustedAdvisorChecksResponse struct {
	*DescribeTrustedAdvisorChecksOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeTrustedAdvisorChecks request.
func (r *DescribeTrustedAdvisorChecksResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
