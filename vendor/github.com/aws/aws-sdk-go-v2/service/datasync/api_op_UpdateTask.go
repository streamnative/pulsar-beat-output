// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package datasync

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// UpdateTaskResponse
// Please also see https://docs.aws.amazon.com/goto/WebAPI/datasync-2018-11-09/UpdateTaskRequest
type UpdateTaskInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the resource name of the CloudWatch LogGroup.
	CloudWatchLogGroupArn *string `type:"string"`

	// A list of filter rules that determines which files to exclude from a task.
	// The list should contain a single filter string that consists of the patterns
	// to exclude. The patterns are delimited by "|" (that is, a pipe), for example:
	// "/folder1|/folder2"
	Excludes []FilterRule `type:"list"`

	// The name of the task to update.
	Name *string `min:"1" type:"string"`

	// Represents the options that are available to control the behavior of a StartTaskExecution
	// operation. Behavior includes preserving metadata such as user ID (UID), group
	// ID (GID), and file permissions, and also overwriting files in the destination,
	// data integrity verification, and so on.
	//
	// A task has a set of default options associated with it. If you don't specify
	// an option in StartTaskExecution, the default value is used. You can override
	// the defaults options on each task execution by specifying an overriding Options
	// value to StartTaskExecution.
	Options *Options `type:"structure"`

	// The Amazon Resource Name (ARN) of the resource name of the task to update.
	//
	// TaskArn is a required field
	TaskArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateTaskInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateTaskInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateTaskInput"}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if s.TaskArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("TaskArn"))
	}
	if s.Options != nil {
		if err := s.Options.Validate(); err != nil {
			invalidParams.AddNested("Options", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/datasync-2018-11-09/UpdateTaskResponse
type UpdateTaskOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateTaskOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateTask = "UpdateTask"

// UpdateTaskRequest returns a request value for making API operation for
// AWS DataSync.
//
// Updates the metadata associated with a task.
//
//    // Example sending a request using UpdateTaskRequest.
//    req := client.UpdateTaskRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/datasync-2018-11-09/UpdateTask
func (c *Client) UpdateTaskRequest(input *UpdateTaskInput) UpdateTaskRequest {
	op := &aws.Operation{
		Name:       opUpdateTask,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateTaskInput{}
	}

	req := c.newRequest(op, input, &UpdateTaskOutput{})
	return UpdateTaskRequest{Request: req, Input: input, Copy: c.UpdateTaskRequest}
}

// UpdateTaskRequest is the request type for the
// UpdateTask API operation.
type UpdateTaskRequest struct {
	*aws.Request
	Input *UpdateTaskInput
	Copy  func(*UpdateTaskInput) UpdateTaskRequest
}

// Send marshals and sends the UpdateTask API request.
func (r UpdateTaskRequest) Send(ctx context.Context) (*UpdateTaskResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateTaskResponse{
		UpdateTaskOutput: r.Request.Data.(*UpdateTaskOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateTaskResponse is the response type for the
// UpdateTask API operation.
type UpdateTaskResponse struct {
	*UpdateTaskOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateTask request.
func (r *UpdateTaskResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
