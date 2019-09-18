// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package opsworks

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/UpdateInstanceRequest
type UpdateInstanceInput struct {
	_ struct{} `type:"structure"`

	// The default AWS OpsWorks Stacks agent version. You have the following options:
	//
	//    * INHERIT - Use the stack's default agent version setting.
	//
	//    * version_number - Use the specified agent version. This value overrides
	//    the stack's default setting. To update the agent version, you must edit
	//    the instance configuration and specify a new version. AWS OpsWorks Stacks
	//    then automatically installs that version on the instance.
	//
	// The default setting is INHERIT. To specify an agent version, you must use
	// the complete version number, not the abbreviated number shown on the console.
	// For a list of available agent version numbers, call DescribeAgentVersions.
	//
	// AgentVersion cannot be set to Chef 12.2.
	AgentVersion *string `type:"string"`

	// The ID of the AMI that was used to create the instance. The value of this
	// parameter must be the same AMI ID that the instance is already using. You
	// cannot apply a new AMI to an instance by running UpdateInstance. UpdateInstance
	// does not work on instances that are using custom AMIs.
	AmiId *string `type:"string"`

	// The instance architecture. Instance types do not necessarily support both
	// architectures. For a list of the architectures that are supported by the
	// different instance types, see Instance Families and Types (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html).
	Architecture Architecture `type:"string" enum:"true"`

	// For load-based or time-based instances, the type. Windows stacks can use
	// only time-based instances.
	AutoScalingType AutoScalingType `type:"string" enum:"true"`

	// This property cannot be updated.
	EbsOptimized *bool `type:"boolean"`

	// The instance host name.
	Hostname *string `type:"string"`

	// Whether to install operating system and package updates when the instance
	// boots. The default value is true. To control when updates are installed,
	// set this value to false. You must then update your instances manually by
	// using CreateDeployment to run the update_dependencies stack command or by
	// manually running yum (Amazon Linux) or apt-get (Ubuntu) on the instances.
	//
	// We strongly recommend using the default value of true, to ensure that your
	// instances have the latest security updates.
	InstallUpdatesOnBoot *bool `type:"boolean"`

	// The instance ID.
	//
	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`

	// The instance type, such as t2.micro. For a list of supported instance types,
	// open the stack in the console, choose Instances, and choose + Instance. The
	// Size list contains the currently supported types. For more information, see
	// Instance Families and Types (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html).
	// The parameter values that you use to specify the various types are in the
	// API Name column of the Available Instance Types table.
	InstanceType *string `type:"string"`

	// The instance's layer IDs.
	LayerIds []string `type:"list"`

	// The instance's operating system, which must be set to one of the following.
	// You cannot update an instance that is using a custom AMI.
	//
	//    * A supported Linux operating system: An Amazon Linux version, such as
	//    Amazon Linux 2018.03, Amazon Linux 2017.09, Amazon Linux 2017.03, Amazon
	//    Linux 2016.09, Amazon Linux 2016.03, Amazon Linux 2015.09, or Amazon Linux
	//    2015.03.
	//
	//    * A supported Ubuntu operating system, such as Ubuntu 16.04 LTS, Ubuntu
	//    14.04 LTS, or Ubuntu 12.04 LTS.
	//
	//    * CentOS Linux 7
	//
	//    * Red Hat Enterprise Linux 7
	//
	//    * A supported Windows operating system, such as Microsoft Windows Server
	//    2012 R2 Base, Microsoft Windows Server 2012 R2 with SQL Server Express,
	//    Microsoft Windows Server 2012 R2 with SQL Server Standard, or Microsoft
	//    Windows Server 2012 R2 with SQL Server Web.
	//
	// For more information about supported operating systems, see AWS OpsWorks
	// Stacks Operating Systems (https://docs.aws.amazon.com/opsworks/latest/userguide/workinginstances-os.html).
	//
	// The default option is the current Amazon Linux version. If you set this parameter
	// to Custom, you must use the AmiId parameter to specify the custom AMI that
	// you want to use. For more information about supported operating systems,
	// see Operating Systems (https://docs.aws.amazon.com/opsworks/latest/userguide/workinginstances-os.html).
	// For more information about how to use custom AMIs with OpsWorks, see Using
	// Custom AMIs (https://docs.aws.amazon.com/opsworks/latest/userguide/workinginstances-custom-ami.html).
	//
	// You can specify a different Linux operating system for the updated stack,
	// but you cannot change from Linux to Windows or Windows to Linux.
	Os *string `type:"string"`

	// The instance's Amazon EC2 key name.
	SshKeyName *string `type:"string"`
}

// String returns the string representation
func (s UpdateInstanceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateInstanceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateInstanceInput"}

	if s.InstanceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/UpdateInstanceOutput
type UpdateInstanceOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateInstanceOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateInstance = "UpdateInstance"

// UpdateInstanceRequest returns a request value for making API operation for
// AWS OpsWorks.
//
// Updates a specified instance.
//
// Required Permissions: To use this action, an IAM user must have a Manage
// permissions level for the stack, or an attached policy that explicitly grants
// permissions. For more information on user permissions, see Managing User
// Permissions (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
//
//    // Example sending a request using UpdateInstanceRequest.
//    req := client.UpdateInstanceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/UpdateInstance
func (c *Client) UpdateInstanceRequest(input *UpdateInstanceInput) UpdateInstanceRequest {
	op := &aws.Operation{
		Name:       opUpdateInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateInstanceInput{}
	}

	req := c.newRequest(op, input, &UpdateInstanceOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return UpdateInstanceRequest{Request: req, Input: input, Copy: c.UpdateInstanceRequest}
}

// UpdateInstanceRequest is the request type for the
// UpdateInstance API operation.
type UpdateInstanceRequest struct {
	*aws.Request
	Input *UpdateInstanceInput
	Copy  func(*UpdateInstanceInput) UpdateInstanceRequest
}

// Send marshals and sends the UpdateInstance API request.
func (r UpdateInstanceRequest) Send(ctx context.Context) (*UpdateInstanceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateInstanceResponse{
		UpdateInstanceOutput: r.Request.Data.(*UpdateInstanceOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateInstanceResponse is the response type for the
// UpdateInstance API operation.
type UpdateInstanceResponse struct {
	*UpdateInstanceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateInstance request.
func (r *UpdateInstanceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
