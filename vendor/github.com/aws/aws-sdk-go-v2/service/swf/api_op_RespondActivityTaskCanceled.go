// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package swf

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type RespondActivityTaskCanceledInput struct {
	_ struct{} `type:"structure"`

	// Information about the cancellation.
	Details *string `locationName:"details" type:"string"`

	// The taskToken of the ActivityTask.
	//
	// taskToken is generated by the service and should be treated as an opaque
	// value. If the task is passed to another process, its taskToken must also
	// be passed. This enables it to provide its progress and respond with results.
	//
	// TaskToken is a required field
	TaskToken *string `locationName:"taskToken" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s RespondActivityTaskCanceledInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RespondActivityTaskCanceledInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RespondActivityTaskCanceledInput"}

	if s.TaskToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("TaskToken"))
	}
	if s.TaskToken != nil && len(*s.TaskToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TaskToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type RespondActivityTaskCanceledOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s RespondActivityTaskCanceledOutput) String() string {
	return awsutil.Prettify(s)
}

const opRespondActivityTaskCanceled = "RespondActivityTaskCanceled"

// RespondActivityTaskCanceledRequest returns a request value for making API operation for
// Amazon Simple Workflow Service.
//
// Used by workers to tell the service that the ActivityTask identified by the
// taskToken was successfully canceled. Additional details can be provided using
// the details argument.
//
// These details (if provided) appear in the ActivityTaskCanceled event added
// to the workflow history.
//
// Only use this operation if the canceled flag of a RecordActivityTaskHeartbeat
// request returns true and if the activity can be safely undone or abandoned.
//
// A task is considered open from the time that it is scheduled until it is
// closed. Therefore a task is reported as open while a worker is processing
// it. A task is closed after it has been specified in a call to RespondActivityTaskCompleted,
// RespondActivityTaskCanceled, RespondActivityTaskFailed, or the task has timed
// out (https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dg-basic.html#swf-dev-timeout-types).
//
// Access Control
//
// You can use IAM policies to control this action's access to Amazon SWF resources
// as follows:
//
//    * Use a Resource element with the domain name to limit the action to only
//    specified domains.
//
//    * Use an Action element to allow or deny permission to call this action.
//
//    * You cannot use an IAM policy to constrain this action's parameters.
//
// If the caller doesn't have sufficient permissions to invoke the action, or
// the parameter values fall outside the specified constraints, the action fails.
// The associated event attribute's cause parameter is set to OPERATION_NOT_PERMITTED.
// For details and example IAM policies, see Using IAM to Manage Access to Amazon
// SWF Workflows (https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html)
// in the Amazon SWF Developer Guide.
//
//    // Example sending a request using RespondActivityTaskCanceledRequest.
//    req := client.RespondActivityTaskCanceledRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) RespondActivityTaskCanceledRequest(input *RespondActivityTaskCanceledInput) RespondActivityTaskCanceledRequest {
	op := &aws.Operation{
		Name:       opRespondActivityTaskCanceled,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RespondActivityTaskCanceledInput{}
	}

	req := c.newRequest(op, input, &RespondActivityTaskCanceledOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return RespondActivityTaskCanceledRequest{Request: req, Input: input, Copy: c.RespondActivityTaskCanceledRequest}
}

// RespondActivityTaskCanceledRequest is the request type for the
// RespondActivityTaskCanceled API operation.
type RespondActivityTaskCanceledRequest struct {
	*aws.Request
	Input *RespondActivityTaskCanceledInput
	Copy  func(*RespondActivityTaskCanceledInput) RespondActivityTaskCanceledRequest
}

// Send marshals and sends the RespondActivityTaskCanceled API request.
func (r RespondActivityTaskCanceledRequest) Send(ctx context.Context) (*RespondActivityTaskCanceledResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RespondActivityTaskCanceledResponse{
		RespondActivityTaskCanceledOutput: r.Request.Data.(*RespondActivityTaskCanceledOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RespondActivityTaskCanceledResponse is the response type for the
// RespondActivityTaskCanceled API operation.
type RespondActivityTaskCanceledResponse struct {
	*RespondActivityTaskCanceledOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RespondActivityTaskCanceled request.
func (r *RespondActivityTaskCanceledResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
