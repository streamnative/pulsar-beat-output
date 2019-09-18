// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/StopLabelingJobRequest
type StopLabelingJobInput struct {
	_ struct{} `type:"structure"`

	// The name of the labeling job to stop.
	//
	// LabelingJobName is a required field
	LabelingJobName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s StopLabelingJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StopLabelingJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StopLabelingJobInput"}

	if s.LabelingJobName == nil {
		invalidParams.Add(aws.NewErrParamRequired("LabelingJobName"))
	}
	if s.LabelingJobName != nil && len(*s.LabelingJobName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("LabelingJobName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/StopLabelingJobOutput
type StopLabelingJobOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s StopLabelingJobOutput) String() string {
	return awsutil.Prettify(s)
}

const opStopLabelingJob = "StopLabelingJob"

// StopLabelingJobRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// Stops a running labeling job. A job that is stopped cannot be restarted.
// Any results obtained before the job is stopped are placed in the Amazon S3
// output bucket.
//
//    // Example sending a request using StopLabelingJobRequest.
//    req := client.StopLabelingJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/StopLabelingJob
func (c *Client) StopLabelingJobRequest(input *StopLabelingJobInput) StopLabelingJobRequest {
	op := &aws.Operation{
		Name:       opStopLabelingJob,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StopLabelingJobInput{}
	}

	req := c.newRequest(op, input, &StopLabelingJobOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return StopLabelingJobRequest{Request: req, Input: input, Copy: c.StopLabelingJobRequest}
}

// StopLabelingJobRequest is the request type for the
// StopLabelingJob API operation.
type StopLabelingJobRequest struct {
	*aws.Request
	Input *StopLabelingJobInput
	Copy  func(*StopLabelingJobInput) StopLabelingJobRequest
}

// Send marshals and sends the StopLabelingJob API request.
func (r StopLabelingJobRequest) Send(ctx context.Context) (*StopLabelingJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StopLabelingJobResponse{
		StopLabelingJobOutput: r.Request.Data.(*StopLabelingJobOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StopLabelingJobResponse is the response type for the
// StopLabelingJob API operation.
type StopLabelingJobResponse struct {
	*StopLabelingJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StopLabelingJob request.
func (r *StopLabelingJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
