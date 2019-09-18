// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package swf

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type StartWorkflowExecutionInput struct {
	_ struct{} `type:"structure"`

	// If set, specifies the policy to use for the child workflow executions of
	// this workflow execution if it is terminated, by calling the TerminateWorkflowExecution
	// action explicitly or due to an expired timeout. This policy overrides the
	// default child policy specified when registering the workflow type using RegisterWorkflowType.
	//
	// The supported child policies are:
	//
	//    * TERMINATE – The child executions are terminated.
	//
	//    * REQUEST_CANCEL – A request to cancel is attempted for each child execution
	//    by recording a WorkflowExecutionCancelRequested event in its history.
	//    It is up to the decider to take appropriate actions when it receives an
	//    execution history with this event.
	//
	//    * ABANDON – No action is taken. The child executions continue to run.
	//
	// A child policy for this workflow execution must be specified either as a
	// default for the workflow type or through this parameter. If neither this
	// parameter is set nor a default child policy was specified at registration
	// time then a fault is returned.
	ChildPolicy ChildPolicy `locationName:"childPolicy" type:"string" enum:"true"`

	// The name of the domain in which the workflow execution is created.
	//
	// Domain is a required field
	Domain *string `locationName:"domain" min:"1" type:"string" required:"true"`

	// The total duration for this workflow execution. This overrides the defaultExecutionStartToCloseTimeout
	// specified when registering the workflow type.
	//
	// The duration is specified in seconds; an integer greater than or equal to
	// 0. Exceeding this limit causes the workflow execution to time out. Unlike
	// some of the other timeout parameters in Amazon SWF, you cannot specify a
	// value of "NONE" for this timeout; there is a one-year max limit on the time
	// that a workflow execution can run.
	//
	// An execution start-to-close timeout must be specified either through this
	// parameter or as a default when the workflow type is registered. If neither
	// this parameter nor a default execution start-to-close timeout is specified,
	// a fault is returned.
	ExecutionStartToCloseTimeout *string `locationName:"executionStartToCloseTimeout" type:"string"`

	// The input for the workflow execution. This is a free form string which should
	// be meaningful to the workflow you are starting. This input is made available
	// to the new workflow execution in the WorkflowExecutionStarted history event.
	Input *string `locationName:"input" type:"string"`

	// The IAM role to attach to this workflow execution.
	//
	// Executions of this workflow type need IAM roles to invoke Lambda functions.
	// If you don't attach an IAM role, any attempt to schedule a Lambda task fails.
	// This results in a ScheduleLambdaFunctionFailed history event. For more information,
	// see https://docs.aws.amazon.com/amazonswf/latest/developerguide/lambda-task.html
	// (https://docs.aws.amazon.com/amazonswf/latest/developerguide/lambda-task.html)
	// in the Amazon SWF Developer Guide.
	LambdaRole *string `locationName:"lambdaRole" min:"1" type:"string"`

	// The list of tags to associate with the workflow execution. You can specify
	// a maximum of 5 tags. You can list workflow executions with a specific tag
	// by calling ListOpenWorkflowExecutions or ListClosedWorkflowExecutions and
	// specifying a TagFilter.
	TagList []string `locationName:"tagList" type:"list"`

	// The task list to use for the decision tasks generated for this workflow execution.
	// This overrides the defaultTaskList specified when registering the workflow
	// type.
	//
	// A task list for this workflow execution must be specified either as a default
	// for the workflow type or through this parameter. If neither this parameter
	// is set nor a default task list was specified at registration time then a
	// fault is returned.
	//
	// The specified string must not start or end with whitespace. It must not contain
	// a : (colon), / (slash), | (vertical bar), or any control characters (\u0000-\u001f
	// | \u007f-\u009f). Also, it must not be the literal string arn.
	TaskList *TaskList `locationName:"taskList" type:"structure"`

	// The task priority to use for this workflow execution. This overrides any
	// default priority that was assigned when the workflow type was registered.
	// If not set, then the default task priority for the workflow type is used.
	// Valid values are integers that range from Java's Integer.MIN_VALUE (-2147483648)
	// to Integer.MAX_VALUE (2147483647). Higher numbers indicate higher priority.
	//
	// For more information about setting task priority, see Setting Task Priority
	// (https://docs.aws.amazon.com/amazonswf/latest/developerguide/programming-priority.html)
	// in the Amazon SWF Developer Guide.
	TaskPriority *string `locationName:"taskPriority" type:"string"`

	// Specifies the maximum duration of decision tasks for this workflow execution.
	// This parameter overrides the defaultTaskStartToCloseTimout specified when
	// registering the workflow type using RegisterWorkflowType.
	//
	// The duration is specified in seconds, an integer greater than or equal to
	// 0. You can use NONE to specify unlimited duration.
	//
	// A task start-to-close timeout for this workflow execution must be specified
	// either as a default for the workflow type or through this parameter. If neither
	// this parameter is set nor a default task start-to-close timeout was specified
	// at registration time then a fault is returned.
	TaskStartToCloseTimeout *string `locationName:"taskStartToCloseTimeout" type:"string"`

	// The user defined identifier associated with the workflow execution. You can
	// use this to associate a custom identifier with the workflow execution. You
	// may specify the same identifier if a workflow execution is logically a restart
	// of a previous execution. You cannot have two open workflow executions with
	// the same workflowId at the same time within the same domain.
	//
	// The specified string must not start or end with whitespace. It must not contain
	// a : (colon), / (slash), | (vertical bar), or any control characters (\u0000-\u001f
	// | \u007f-\u009f). Also, it must not be the literal string arn.
	//
	// WorkflowId is a required field
	WorkflowId *string `locationName:"workflowId" min:"1" type:"string" required:"true"`

	// The type of the workflow to start.
	//
	// WorkflowType is a required field
	WorkflowType *WorkflowType `locationName:"workflowType" type:"structure" required:"true"`
}

// String returns the string representation
func (s StartWorkflowExecutionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartWorkflowExecutionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartWorkflowExecutionInput"}

	if s.Domain == nil {
		invalidParams.Add(aws.NewErrParamRequired("Domain"))
	}
	if s.Domain != nil && len(*s.Domain) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Domain", 1))
	}
	if s.LambdaRole != nil && len(*s.LambdaRole) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("LambdaRole", 1))
	}

	if s.WorkflowId == nil {
		invalidParams.Add(aws.NewErrParamRequired("WorkflowId"))
	}
	if s.WorkflowId != nil && len(*s.WorkflowId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("WorkflowId", 1))
	}

	if s.WorkflowType == nil {
		invalidParams.Add(aws.NewErrParamRequired("WorkflowType"))
	}
	if s.TaskList != nil {
		if err := s.TaskList.Validate(); err != nil {
			invalidParams.AddNested("TaskList", err.(aws.ErrInvalidParams))
		}
	}
	if s.WorkflowType != nil {
		if err := s.WorkflowType.Validate(); err != nil {
			invalidParams.AddNested("WorkflowType", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Specifies the runId of a workflow execution.
type StartWorkflowExecutionOutput struct {
	_ struct{} `type:"structure"`

	// The runId of a workflow execution. This ID is generated by the service and
	// can be used to uniquely identify the workflow execution within a domain.
	RunId *string `locationName:"runId" min:"1" type:"string"`
}

// String returns the string representation
func (s StartWorkflowExecutionOutput) String() string {
	return awsutil.Prettify(s)
}

const opStartWorkflowExecution = "StartWorkflowExecution"

// StartWorkflowExecutionRequest returns a request value for making API operation for
// Amazon Simple Workflow Service.
//
// Starts an execution of the workflow type in the specified domain using the
// provided workflowId and input data.
//
// This action returns the newly started workflow execution.
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
//    the appropriate keys. tagList.member.0: The key is swf:tagList.member.0.
//    tagList.member.1: The key is swf:tagList.member.1. tagList.member.2: The
//    key is swf:tagList.member.2. tagList.member.3: The key is swf:tagList.member.3.
//    tagList.member.4: The key is swf:tagList.member.4. taskList: String constraint.
//    The key is swf:taskList.name. workflowType.name: String constraint. The
//    key is swf:workflowType.name. workflowType.version: String constraint.
//    The key is swf:workflowType.version.
//
// If the caller doesn't have sufficient permissions to invoke the action, or
// the parameter values fall outside the specified constraints, the action fails.
// The associated event attribute's cause parameter is set to OPERATION_NOT_PERMITTED.
// For details and example IAM policies, see Using IAM to Manage Access to Amazon
// SWF Workflows (https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html)
// in the Amazon SWF Developer Guide.
//
//    // Example sending a request using StartWorkflowExecutionRequest.
//    req := client.StartWorkflowExecutionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) StartWorkflowExecutionRequest(input *StartWorkflowExecutionInput) StartWorkflowExecutionRequest {
	op := &aws.Operation{
		Name:       opStartWorkflowExecution,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StartWorkflowExecutionInput{}
	}

	req := c.newRequest(op, input, &StartWorkflowExecutionOutput{})
	return StartWorkflowExecutionRequest{Request: req, Input: input, Copy: c.StartWorkflowExecutionRequest}
}

// StartWorkflowExecutionRequest is the request type for the
// StartWorkflowExecution API operation.
type StartWorkflowExecutionRequest struct {
	*aws.Request
	Input *StartWorkflowExecutionInput
	Copy  func(*StartWorkflowExecutionInput) StartWorkflowExecutionRequest
}

// Send marshals and sends the StartWorkflowExecution API request.
func (r StartWorkflowExecutionRequest) Send(ctx context.Context) (*StartWorkflowExecutionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartWorkflowExecutionResponse{
		StartWorkflowExecutionOutput: r.Request.Data.(*StartWorkflowExecutionOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartWorkflowExecutionResponse is the response type for the
// StartWorkflowExecution API operation.
type StartWorkflowExecutionResponse struct {
	*StartWorkflowExecutionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartWorkflowExecution request.
func (r *StartWorkflowExecutionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
