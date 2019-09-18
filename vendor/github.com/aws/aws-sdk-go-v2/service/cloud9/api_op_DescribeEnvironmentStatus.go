// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloud9

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloud9-2017-09-23/DescribeEnvironmentStatusRequest
type DescribeEnvironmentStatusInput struct {
	_ struct{} `type:"structure"`

	// The ID of the environment to get status information about.
	//
	// EnvironmentId is a required field
	EnvironmentId *string `locationName:"environmentId" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeEnvironmentStatusInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeEnvironmentStatusInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeEnvironmentStatusInput"}

	if s.EnvironmentId == nil {
		invalidParams.Add(aws.NewErrParamRequired("EnvironmentId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloud9-2017-09-23/DescribeEnvironmentStatusResult
type DescribeEnvironmentStatusOutput struct {
	_ struct{} `type:"structure"`

	// Any informational message about the status of the environment.
	Message *string `locationName:"message" type:"string"`

	// The status of the environment. Available values include:
	//
	//    * connecting: The environment is connecting.
	//
	//    * creating: The environment is being created.
	//
	//    * deleting: The environment is being deleted.
	//
	//    * error: The environment is in an error state.
	//
	//    * ready: The environment is ready.
	//
	//    * stopped: The environment is stopped.
	//
	//    * stopping: The environment is stopping.
	Status EnvironmentStatus `locationName:"status" type:"string" enum:"true"`
}

// String returns the string representation
func (s DescribeEnvironmentStatusOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeEnvironmentStatus = "DescribeEnvironmentStatus"

// DescribeEnvironmentStatusRequest returns a request value for making API operation for
// AWS Cloud9.
//
// Gets status information for an AWS Cloud9 development environment.
//
//    // Example sending a request using DescribeEnvironmentStatusRequest.
//    req := client.DescribeEnvironmentStatusRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloud9-2017-09-23/DescribeEnvironmentStatus
func (c *Client) DescribeEnvironmentStatusRequest(input *DescribeEnvironmentStatusInput) DescribeEnvironmentStatusRequest {
	op := &aws.Operation{
		Name:       opDescribeEnvironmentStatus,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeEnvironmentStatusInput{}
	}

	req := c.newRequest(op, input, &DescribeEnvironmentStatusOutput{})
	return DescribeEnvironmentStatusRequest{Request: req, Input: input, Copy: c.DescribeEnvironmentStatusRequest}
}

// DescribeEnvironmentStatusRequest is the request type for the
// DescribeEnvironmentStatus API operation.
type DescribeEnvironmentStatusRequest struct {
	*aws.Request
	Input *DescribeEnvironmentStatusInput
	Copy  func(*DescribeEnvironmentStatusInput) DescribeEnvironmentStatusRequest
}

// Send marshals and sends the DescribeEnvironmentStatus API request.
func (r DescribeEnvironmentStatusRequest) Send(ctx context.Context) (*DescribeEnvironmentStatusResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeEnvironmentStatusResponse{
		DescribeEnvironmentStatusOutput: r.Request.Data.(*DescribeEnvironmentStatusOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeEnvironmentStatusResponse is the response type for the
// DescribeEnvironmentStatus API operation.
type DescribeEnvironmentStatusResponse struct {
	*DescribeEnvironmentStatusOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeEnvironmentStatus request.
func (r *DescribeEnvironmentStatusResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
