// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package swf

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type RegisterActivityTypeInput struct {
	_ struct{} `type:"structure"`

	// If set, specifies the default maximum time before which a worker processing
	// a task of this type must report progress by calling RecordActivityTaskHeartbeat.
	// If the timeout is exceeded, the activity task is automatically timed out.
	// This default can be overridden when scheduling an activity task using the
	// ScheduleActivityTask Decision. If the activity worker subsequently attempts
	// to record a heartbeat or returns a result, the activity worker receives an
	// UnknownResource fault. In this case, Amazon SWF no longer considers the activity
	// task to be valid; the activity worker should clean up the activity task.
	//
	// The duration is specified in seconds, an integer greater than or equal to
	// 0. You can use NONE to specify unlimited duration.
	DefaultTaskHeartbeatTimeout *string `locationName:"defaultTaskHeartbeatTimeout" type:"string"`

	// If set, specifies the default task list to use for scheduling tasks of this
	// activity type. This default task list is used if a task list isn't provided
	// when a task is scheduled through the ScheduleActivityTask Decision.
	DefaultTaskList *TaskList `locationName:"defaultTaskList" type:"structure"`

	// The default task priority to assign to the activity type. If not assigned,
	// then 0 is used. Valid values are integers that range from Java's Integer.MIN_VALUE
	// (-2147483648) to Integer.MAX_VALUE (2147483647). Higher numbers indicate
	// higher priority.
	//
	// For more information about setting task priority, see Setting Task Priority
	// (https://docs.aws.amazon.com/amazonswf/latest/developerguide/programming-priority.html)
	// in the in the Amazon SWF Developer Guide..
	DefaultTaskPriority *string `locationName:"defaultTaskPriority" type:"string"`

	// If set, specifies the default maximum duration for a task of this activity
	// type. This default can be overridden when scheduling an activity task using
	// the ScheduleActivityTask Decision.
	//
	// The duration is specified in seconds, an integer greater than or equal to
	// 0. You can use NONE to specify unlimited duration.
	DefaultTaskScheduleToCloseTimeout *string `locationName:"defaultTaskScheduleToCloseTimeout" type:"string"`

	// If set, specifies the default maximum duration that a task of this activity
	// type can wait before being assigned to a worker. This default can be overridden
	// when scheduling an activity task using the ScheduleActivityTask Decision.
	//
	// The duration is specified in seconds, an integer greater than or equal to
	// 0. You can use NONE to specify unlimited duration.
	DefaultTaskScheduleToStartTimeout *string `locationName:"defaultTaskScheduleToStartTimeout" type:"string"`

	// If set, specifies the default maximum duration that a worker can take to
	// process tasks of this activity type. This default can be overridden when
	// scheduling an activity task using the ScheduleActivityTask Decision.
	//
	// The duration is specified in seconds, an integer greater than or equal to
	// 0. You can use NONE to specify unlimited duration.
	DefaultTaskStartToCloseTimeout *string `locationName:"defaultTaskStartToCloseTimeout" type:"string"`

	// A textual description of the activity type.
	Description *string `locationName:"description" type:"string"`

	// The name of the domain in which this activity is to be registered.
	//
	// Domain is a required field
	Domain *string `locationName:"domain" min:"1" type:"string" required:"true"`

	// The name of the activity type within the domain.
	//
	// The specified string must not start or end with whitespace. It must not contain
	// a : (colon), / (slash), | (vertical bar), or any control characters (\u0000-\u001f
	// | \u007f-\u009f). Also, it must not be the literal string arn.
	//
	// Name is a required field
	Name *string `locationName:"name" min:"1" type:"string" required:"true"`

	// The version of the activity type.
	//
	// The activity type consists of the name and version, the combination of which
	// must be unique within the domain.
	//
	// The specified string must not start or end with whitespace. It must not contain
	// a : (colon), / (slash), | (vertical bar), or any control characters (\u0000-\u001f
	// | \u007f-\u009f). Also, it must not be the literal string arn.
	//
	// Version is a required field
	Version *string `locationName:"version" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s RegisterActivityTypeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RegisterActivityTypeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RegisterActivityTypeInput"}

	if s.Domain == nil {
		invalidParams.Add(aws.NewErrParamRequired("Domain"))
	}
	if s.Domain != nil && len(*s.Domain) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Domain", 1))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if s.Version == nil {
		invalidParams.Add(aws.NewErrParamRequired("Version"))
	}
	if s.Version != nil && len(*s.Version) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Version", 1))
	}
	if s.DefaultTaskList != nil {
		if err := s.DefaultTaskList.Validate(); err != nil {
			invalidParams.AddNested("DefaultTaskList", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type RegisterActivityTypeOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s RegisterActivityTypeOutput) String() string {
	return awsutil.Prettify(s)
}

const opRegisterActivityType = "RegisterActivityType"

// RegisterActivityTypeRequest returns a request value for making API operation for
// Amazon Simple Workflow Service.
//
// Registers a new activity type along with its configuration settings in the
// specified domain.
//
// A TypeAlreadyExists fault is returned if the type already exists in the domain.
// You cannot change any configuration settings of the type after its registration,
// and it must be registered as a new version.
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
//    * Constrain the following parameters by using a Condition element with
//    the appropriate keys. defaultTaskList.name: String constraint. The key
//    is swf:defaultTaskList.name. name: String constraint. The key is swf:name.
//    version: String constraint. The key is swf:version.
//
// If the caller doesn't have sufficient permissions to invoke the action, or
// the parameter values fall outside the specified constraints, the action fails.
// The associated event attribute's cause parameter is set to OPERATION_NOT_PERMITTED.
// For details and example IAM policies, see Using IAM to Manage Access to Amazon
// SWF Workflows (https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html)
// in the Amazon SWF Developer Guide.
//
//    // Example sending a request using RegisterActivityTypeRequest.
//    req := client.RegisterActivityTypeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) RegisterActivityTypeRequest(input *RegisterActivityTypeInput) RegisterActivityTypeRequest {
	op := &aws.Operation{
		Name:       opRegisterActivityType,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RegisterActivityTypeInput{}
	}

	req := c.newRequest(op, input, &RegisterActivityTypeOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return RegisterActivityTypeRequest{Request: req, Input: input, Copy: c.RegisterActivityTypeRequest}
}

// RegisterActivityTypeRequest is the request type for the
// RegisterActivityType API operation.
type RegisterActivityTypeRequest struct {
	*aws.Request
	Input *RegisterActivityTypeInput
	Copy  func(*RegisterActivityTypeInput) RegisterActivityTypeRequest
}

// Send marshals and sends the RegisterActivityType API request.
func (r RegisterActivityTypeRequest) Send(ctx context.Context) (*RegisterActivityTypeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RegisterActivityTypeResponse{
		RegisterActivityTypeOutput: r.Request.Data.(*RegisterActivityTypeOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RegisterActivityTypeResponse is the response type for the
// RegisterActivityType API operation.
type RegisterActivityTypeResponse struct {
	*RegisterActivityTypeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RegisterActivityType request.
func (r *RegisterActivityTypeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
