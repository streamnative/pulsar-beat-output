// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/BatchStopJobRunRequest
type BatchStopJobRunInput struct {
	_ struct{} `type:"structure"`

	// The name of the job definition for which to stop job runs.
	//
	// JobName is a required field
	JobName *string `min:"1" type:"string" required:"true"`

	// A list of the JobRunIds that should be stopped for that job definition.
	//
	// JobRunIds is a required field
	JobRunIds []string `min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s BatchStopJobRunInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BatchStopJobRunInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "BatchStopJobRunInput"}

	if s.JobName == nil {
		invalidParams.Add(aws.NewErrParamRequired("JobName"))
	}
	if s.JobName != nil && len(*s.JobName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("JobName", 1))
	}

	if s.JobRunIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("JobRunIds"))
	}
	if s.JobRunIds != nil && len(s.JobRunIds) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("JobRunIds", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/BatchStopJobRunResponse
type BatchStopJobRunOutput struct {
	_ struct{} `type:"structure"`

	// A list of the errors that were encountered in trying to stop JobRuns, including
	// the JobRunId for which each error was encountered and details about the error.
	Errors []BatchStopJobRunError `type:"list"`

	// A list of the JobRuns that were successfully submitted for stopping.
	SuccessfulSubmissions []BatchStopJobRunSuccessfulSubmission `type:"list"`
}

// String returns the string representation
func (s BatchStopJobRunOutput) String() string {
	return awsutil.Prettify(s)
}

const opBatchStopJobRun = "BatchStopJobRun"

// BatchStopJobRunRequest returns a request value for making API operation for
// AWS Glue.
//
// Stops one or more job runs for a specified job definition.
//
//    // Example sending a request using BatchStopJobRunRequest.
//    req := client.BatchStopJobRunRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/BatchStopJobRun
func (c *Client) BatchStopJobRunRequest(input *BatchStopJobRunInput) BatchStopJobRunRequest {
	op := &aws.Operation{
		Name:       opBatchStopJobRun,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &BatchStopJobRunInput{}
	}

	req := c.newRequest(op, input, &BatchStopJobRunOutput{})
	return BatchStopJobRunRequest{Request: req, Input: input, Copy: c.BatchStopJobRunRequest}
}

// BatchStopJobRunRequest is the request type for the
// BatchStopJobRun API operation.
type BatchStopJobRunRequest struct {
	*aws.Request
	Input *BatchStopJobRunInput
	Copy  func(*BatchStopJobRunInput) BatchStopJobRunRequest
}

// Send marshals and sends the BatchStopJobRun API request.
func (r BatchStopJobRunRequest) Send(ctx context.Context) (*BatchStopJobRunResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &BatchStopJobRunResponse{
		BatchStopJobRunOutput: r.Request.Data.(*BatchStopJobRunOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// BatchStopJobRunResponse is the response type for the
// BatchStopJobRun API operation.
type BatchStopJobRunResponse struct {
	*BatchStopJobRunOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// BatchStopJobRun request.
func (r *BatchStopJobRunResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
