// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the request to start the user import job.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/StartUserImportJobRequest
type StartUserImportJobInput struct {
	_ struct{} `type:"structure"`

	// The job ID for the user import job.
	//
	// JobId is a required field
	JobId *string `min:"1" type:"string" required:"true"`

	// The user pool ID for the user pool that the users are being imported into.
	//
	// UserPoolId is a required field
	UserPoolId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s StartUserImportJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartUserImportJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartUserImportJobInput"}

	if s.JobId == nil {
		invalidParams.Add(aws.NewErrParamRequired("JobId"))
	}
	if s.JobId != nil && len(*s.JobId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("JobId", 1))
	}

	if s.UserPoolId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserPoolId"))
	}
	if s.UserPoolId != nil && len(*s.UserPoolId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserPoolId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the response from the server to the request to start the user
// import job.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/StartUserImportJobResponse
type StartUserImportJobOutput struct {
	_ struct{} `type:"structure"`

	// The job object that represents the user import job.
	UserImportJob *UserImportJobType `type:"structure"`
}

// String returns the string representation
func (s StartUserImportJobOutput) String() string {
	return awsutil.Prettify(s)
}

const opStartUserImportJob = "StartUserImportJob"

// StartUserImportJobRequest returns a request value for making API operation for
// Amazon Cognito Identity Provider.
//
// Starts the user import.
//
//    // Example sending a request using StartUserImportJobRequest.
//    req := client.StartUserImportJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/StartUserImportJob
func (c *Client) StartUserImportJobRequest(input *StartUserImportJobInput) StartUserImportJobRequest {
	op := &aws.Operation{
		Name:       opStartUserImportJob,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StartUserImportJobInput{}
	}

	req := c.newRequest(op, input, &StartUserImportJobOutput{})
	return StartUserImportJobRequest{Request: req, Input: input, Copy: c.StartUserImportJobRequest}
}

// StartUserImportJobRequest is the request type for the
// StartUserImportJob API operation.
type StartUserImportJobRequest struct {
	*aws.Request
	Input *StartUserImportJobInput
	Copy  func(*StartUserImportJobInput) StartUserImportJobRequest
}

// Send marshals and sends the StartUserImportJob API request.
func (r StartUserImportJobRequest) Send(ctx context.Context) (*StartUserImportJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartUserImportJobResponse{
		StartUserImportJobOutput: r.Request.Data.(*StartUserImportJobOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartUserImportJobResponse is the response type for the
// StartUserImportJob API operation.
type StartUserImportJobResponse struct {
	*StartUserImportJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartUserImportJob request.
func (r *StartUserImportJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
