// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudtrail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Passes the request to CloudTrail to stop logging AWS API calls for the specified
// account.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudtrail-2013-11-01/StopLoggingRequest
type StopLoggingInput struct {
	_ struct{} `type:"structure"`

	// Specifies the name or the CloudTrail ARN of the trail for which CloudTrail
	// will stop logging AWS API calls. The format of a trail ARN is:
	//
	// arn:aws:cloudtrail:us-east-2:123456789012:trail/MyTrail
	//
	// Name is a required field
	Name *string `type:"string" required:"true"`
}

// String returns the string representation
func (s StopLoggingInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StopLoggingInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StopLoggingInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Returns the objects or data listed below if successful. Otherwise, returns
// an error.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudtrail-2013-11-01/StopLoggingResponse
type StopLoggingOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s StopLoggingOutput) String() string {
	return awsutil.Prettify(s)
}

const opStopLogging = "StopLogging"

// StopLoggingRequest returns a request value for making API operation for
// AWS CloudTrail.
//
// Suspends the recording of AWS API calls and log file delivery for the specified
// trail. Under most circumstances, there is no need to use this action. You
// can update a trail without stopping it first. This action is the only way
// to stop recording. For a trail enabled in all regions, this operation must
// be called from the region in which the trail was created, or an InvalidHomeRegionException
// will occur. This operation cannot be called on the shadow trails (replicated
// trails in other regions) of a trail enabled in all regions.
//
//    // Example sending a request using StopLoggingRequest.
//    req := client.StopLoggingRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudtrail-2013-11-01/StopLogging
func (c *Client) StopLoggingRequest(input *StopLoggingInput) StopLoggingRequest {
	op := &aws.Operation{
		Name:       opStopLogging,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StopLoggingInput{}
	}

	req := c.newRequest(op, input, &StopLoggingOutput{})
	return StopLoggingRequest{Request: req, Input: input, Copy: c.StopLoggingRequest}
}

// StopLoggingRequest is the request type for the
// StopLogging API operation.
type StopLoggingRequest struct {
	*aws.Request
	Input *StopLoggingInput
	Copy  func(*StopLoggingInput) StopLoggingRequest
}

// Send marshals and sends the StopLogging API request.
func (r StopLoggingRequest) Send(ctx context.Context) (*StopLoggingResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StopLoggingResponse{
		StopLoggingOutput: r.Request.Data.(*StopLoggingOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StopLoggingResponse is the response type for the
// StopLogging API operation.
type StopLoggingResponse struct {
	*StopLoggingOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StopLogging request.
func (r *StopLoggingResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
