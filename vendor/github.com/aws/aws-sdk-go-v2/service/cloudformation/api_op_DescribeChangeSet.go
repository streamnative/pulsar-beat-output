// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudformation

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// The input for the DescribeChangeSet action.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/DescribeChangeSetInput
type DescribeChangeSetInput struct {
	_ struct{} `type:"structure"`

	// The name or Amazon Resource Name (ARN) of the change set that you want to
	// describe.
	//
	// ChangeSetName is a required field
	ChangeSetName *string `min:"1" type:"string" required:"true"`

	// A string (provided by the DescribeChangeSet response output) that identifies
	// the next page of information that you want to retrieve.
	NextToken *string `min:"1" type:"string"`

	// If you specified the name of a change set, specify the stack name or ID (ARN)
	// of the change set you want to describe.
	StackName *string `min:"1" type:"string"`
}

// String returns the string representation
func (s DescribeChangeSetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeChangeSetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeChangeSetInput"}

	if s.ChangeSetName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ChangeSetName"))
	}
	if s.ChangeSetName != nil && len(*s.ChangeSetName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ChangeSetName", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}
	if s.StackName != nil && len(*s.StackName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("StackName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The output for the DescribeChangeSet action.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/DescribeChangeSetOutput
type DescribeChangeSetOutput struct {
	_ struct{} `type:"structure"`

	// If you execute the change set, the list of capabilities that were explicitly
	// acknowledged when the change set was created.
	Capabilities []Capability `type:"list"`

	// The ARN of the change set.
	ChangeSetId *string `min:"1" type:"string"`

	// The name of the change set.
	ChangeSetName *string `min:"1" type:"string"`

	// A list of Change structures that describes the resources AWS CloudFormation
	// changes if you execute the change set.
	Changes []Change `type:"list"`

	// The start time when the change set was created, in UTC.
	CreationTime *time.Time `type:"timestamp"`

	// Information about the change set.
	Description *string `min:"1" type:"string"`

	// If the change set execution status is AVAILABLE, you can execute the change
	// set. If you can’t execute the change set, the status indicates why. For
	// example, a change set might be in an UNAVAILABLE state because AWS CloudFormation
	// is still creating it or in an OBSOLETE state because the stack was already
	// updated.
	ExecutionStatus ExecutionStatus `type:"string" enum:"true"`

	// If the output exceeds 1 MB, a string that identifies the next page of changes.
	// If there is no additional page, this value is null.
	NextToken *string `min:"1" type:"string"`

	// The ARNs of the Amazon Simple Notification Service (Amazon SNS) topics that
	// will be associated with the stack if you execute the change set.
	NotificationARNs []string `type:"list"`

	// A list of Parameter structures that describes the input parameters and their
	// values used to create the change set. For more information, see the Parameter
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_Parameter.html)
	// data type.
	Parameters []Parameter `type:"list"`

	// The rollback triggers for AWS CloudFormation to monitor during stack creation
	// and updating operations, and for the specified monitoring period afterwards.
	RollbackConfiguration *RollbackConfiguration `type:"structure"`

	// The ARN of the stack that is associated with the change set.
	StackId *string `type:"string"`

	// The name of the stack that is associated with the change set.
	StackName *string `type:"string"`

	// The current status of the change set, such as CREATE_IN_PROGRESS, CREATE_COMPLETE,
	// or FAILED.
	Status ChangeSetStatus `type:"string" enum:"true"`

	// A description of the change set's status. For example, if your attempt to
	// create a change set failed, AWS CloudFormation shows the error message.
	StatusReason *string `type:"string"`

	// If you execute the change set, the tags that will be associated with the
	// stack.
	Tags []Tag `type:"list"`
}

// String returns the string representation
func (s DescribeChangeSetOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeChangeSet = "DescribeChangeSet"

// DescribeChangeSetRequest returns a request value for making API operation for
// AWS CloudFormation.
//
// Returns the inputs for the change set and a list of changes that AWS CloudFormation
// will make if you execute the change set. For more information, see Updating
// Stacks Using Change Sets (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-changesets.html)
// in the AWS CloudFormation User Guide.
//
//    // Example sending a request using DescribeChangeSetRequest.
//    req := client.DescribeChangeSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/DescribeChangeSet
func (c *Client) DescribeChangeSetRequest(input *DescribeChangeSetInput) DescribeChangeSetRequest {
	op := &aws.Operation{
		Name:       opDescribeChangeSet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeChangeSetInput{}
	}

	req := c.newRequest(op, input, &DescribeChangeSetOutput{})
	return DescribeChangeSetRequest{Request: req, Input: input, Copy: c.DescribeChangeSetRequest}
}

// DescribeChangeSetRequest is the request type for the
// DescribeChangeSet API operation.
type DescribeChangeSetRequest struct {
	*aws.Request
	Input *DescribeChangeSetInput
	Copy  func(*DescribeChangeSetInput) DescribeChangeSetRequest
}

// Send marshals and sends the DescribeChangeSet API request.
func (r DescribeChangeSetRequest) Send(ctx context.Context) (*DescribeChangeSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeChangeSetResponse{
		DescribeChangeSetOutput: r.Request.Data.(*DescribeChangeSetOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeChangeSetResponse is the response type for the
// DescribeChangeSet API operation.
type DescribeChangeSetResponse struct {
	*DescribeChangeSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeChangeSet request.
func (r *DescribeChangeSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
