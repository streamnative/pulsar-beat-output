// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package amplify

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Request structure for stop job request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/amplify-2017-07-25/StopJobRequest
type StopJobInput struct {
	_ struct{} `type:"structure"`

	// Unique Id for an Amplify App.
	//
	// AppId is a required field
	AppId *string `location:"uri" locationName:"appId" min:"1" type:"string" required:"true"`

	// Name for the branch, for the Job.
	//
	// BranchName is a required field
	BranchName *string `location:"uri" locationName:"branchName" min:"1" type:"string" required:"true"`

	// Unique Id for the Job.
	//
	// JobId is a required field
	JobId *string `location:"uri" locationName:"jobId" type:"string" required:"true"`
}

// String returns the string representation
func (s StopJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StopJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StopJobInput"}

	if s.AppId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AppId"))
	}
	if s.AppId != nil && len(*s.AppId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AppId", 1))
	}

	if s.BranchName == nil {
		invalidParams.Add(aws.NewErrParamRequired("BranchName"))
	}
	if s.BranchName != nil && len(*s.BranchName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("BranchName", 1))
	}

	if s.JobId == nil {
		invalidParams.Add(aws.NewErrParamRequired("JobId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s StopJobInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AppId != nil {
		v := *s.AppId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "appId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.BranchName != nil {
		v := *s.BranchName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "branchName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.JobId != nil {
		v := *s.JobId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "jobId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Result structure for the stop job request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/amplify-2017-07-25/StopJobResult
type StopJobOutput struct {
	_ struct{} `type:"structure"`

	// Summary for the Job.
	//
	// JobSummary is a required field
	JobSummary *JobSummary `locationName:"jobSummary" type:"structure" required:"true"`
}

// String returns the string representation
func (s StopJobOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s StopJobOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.JobSummary != nil {
		v := s.JobSummary

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "jobSummary", v, metadata)
	}
	return nil
}

const opStopJob = "StopJob"

// StopJobRequest returns a request value for making API operation for
// AWS Amplify.
//
// Stop a job that is in progress, for an Amplify branch, part of Amplify App.
//
//    // Example sending a request using StopJobRequest.
//    req := client.StopJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/amplify-2017-07-25/StopJob
func (c *Client) StopJobRequest(input *StopJobInput) StopJobRequest {
	op := &aws.Operation{
		Name:       opStopJob,
		HTTPMethod: "DELETE",
		HTTPPath:   "/apps/{appId}/branches/{branchName}/jobs/{jobId}/stop",
	}

	if input == nil {
		input = &StopJobInput{}
	}

	req := c.newRequest(op, input, &StopJobOutput{})
	return StopJobRequest{Request: req, Input: input, Copy: c.StopJobRequest}
}

// StopJobRequest is the request type for the
// StopJob API operation.
type StopJobRequest struct {
	*aws.Request
	Input *StopJobInput
	Copy  func(*StopJobInput) StopJobRequest
}

// Send marshals and sends the StopJob API request.
func (r StopJobRequest) Send(ctx context.Context) (*StopJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StopJobResponse{
		StopJobOutput: r.Request.Data.(*StopJobOutput),
		response:      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StopJobResponse is the response type for the
// StopJob API operation.
type StopJobResponse struct {
	*StopJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StopJob request.
func (r *StopJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
