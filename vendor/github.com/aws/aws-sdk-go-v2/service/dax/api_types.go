// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package dax

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

var _ aws.Config
var _ = awsutil.Prettify

// Contains all of the attributes of a specific DAX cluster.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dax-2017-04-19/Cluster
type Cluster struct {
	_ struct{} `type:"structure"`

	// The number of nodes in the cluster that are active (i.e., capable of serving
	// requests).
	ActiveNodes *int64 `type:"integer"`

	// The Amazon Resource Name (ARN) that uniquely identifies the cluster.
	ClusterArn *string `type:"string"`

	// The configuration endpoint for this DAX cluster, consisting of a DNS name
	// and a port number. Client applications can specify this endpoint, rather
	// than an individual node endpoint, and allow the DAX client software to intelligently
	// route requests and responses to nodes in the DAX cluster.
	ClusterDiscoveryEndpoint *Endpoint `type:"structure"`

	// The name of the DAX cluster.
	ClusterName *string `type:"string"`

	// The description of the cluster.
	Description *string `type:"string"`

	// A valid Amazon Resource Name (ARN) that identifies an IAM role. At runtime,
	// DAX will assume this role and use the role's permissions to access DynamoDB
	// on your behalf.
	IamRoleArn *string `type:"string"`

	// A list of nodes to be removed from the cluster.
	NodeIdsToRemove []string `type:"list"`

	// The node type for the nodes in the cluster. (All nodes in a DAX cluster are
	// of the same type.)
	NodeType *string `type:"string"`

	// A list of nodes that are currently in the cluster.
	Nodes []Node `type:"list"`

	// Describes a notification topic and its status. Notification topics are used
	// for publishing DAX events to subscribers using Amazon Simple Notification
	// Service (SNS).
	NotificationConfiguration *NotificationConfiguration `type:"structure"`

	// The parameter group being used by nodes in the cluster.
	ParameterGroup *ParameterGroupStatus `type:"structure"`

	// A range of time when maintenance of DAX cluster software will be performed.
	// For example: sun:01:00-sun:09:00. Cluster maintenance normally takes less
	// than 30 minutes, and is performed automatically within the maintenance window.
	PreferredMaintenanceWindow *string `type:"string"`

	// The description of the server-side encryption status on the specified DAX
	// cluster.
	SSEDescription *SSEDescription `type:"structure"`

	// A list of security groups, and the status of each, for the nodes in the cluster.
	SecurityGroups []SecurityGroupMembership `type:"list"`

	// The current status of the cluster.
	Status *string `type:"string"`

	// The subnet group where the DAX cluster is running.
	SubnetGroup *string `type:"string"`

	// The total number of nodes in the cluster.
	TotalNodes *int64 `type:"integer"`
}

// String returns the string representation
func (s Cluster) String() string {
	return awsutil.Prettify(s)
}

// Represents the information required for client programs to connect to the
// configuration endpoint for a DAX cluster, or to an individual node within
// the cluster.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dax-2017-04-19/Endpoint
type Endpoint struct {
	_ struct{} `type:"structure"`

	// The DNS hostname of the endpoint.
	Address *string `type:"string"`

	// The port number that applications should use to connect to the endpoint.
	Port *int64 `type:"integer"`
}

// String returns the string representation
func (s Endpoint) String() string {
	return awsutil.Prettify(s)
}

// Represents a single occurrence of something interesting within the system.
// Some examples of events are creating a DAX cluster, adding or removing a
// node, or rebooting a node.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dax-2017-04-19/Event
type Event struct {
	_ struct{} `type:"structure"`

	// The date and time when the event occurred.
	Date *time.Time `type:"timestamp"`

	// A user-defined message associated with the event.
	Message *string `type:"string"`

	// The source of the event. For example, if the event occurred at the node level,
	// the source would be the node ID.
	SourceName *string `type:"string"`

	// Specifies the origin of this event - a cluster, a parameter group, a node
	// ID, etc.
	SourceType SourceType `type:"string" enum:"true"`
}

// String returns the string representation
func (s Event) String() string {
	return awsutil.Prettify(s)
}

// Represents an individual node within a DAX cluster.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dax-2017-04-19/Node
type Node struct {
	_ struct{} `type:"structure"`

	// The Availability Zone (AZ) in which the node has been deployed.
	AvailabilityZone *string `type:"string"`

	// The endpoint for the node, consisting of a DNS name and a port number. Client
	// applications can connect directly to a node endpoint, if desired (as an alternative
	// to allowing DAX client software to intelligently route requests and responses
	// to nodes in the DAX cluster.
	Endpoint *Endpoint `type:"structure"`

	// The date and time (in UNIX epoch format) when the node was launched.
	NodeCreateTime *time.Time `type:"timestamp"`

	// A system-generated identifier for the node.
	NodeId *string `type:"string"`

	// The current status of the node. For example: available.
	NodeStatus *string `type:"string"`

	// The status of the parameter group associated with this node. For example,
	// in-sync.
	ParameterGroupStatus *string `type:"string"`
}

// String returns the string representation
func (s Node) String() string {
	return awsutil.Prettify(s)
}

// Represents a parameter value that is applicable to a particular node type.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dax-2017-04-19/NodeTypeSpecificValue
type NodeTypeSpecificValue struct {
	_ struct{} `type:"structure"`

	// A node type to which the parameter value applies.
	NodeType *string `type:"string"`

	// The parameter value for this node type.
	Value *string `type:"string"`
}

// String returns the string representation
func (s NodeTypeSpecificValue) String() string {
	return awsutil.Prettify(s)
}

// Describes a notification topic and its status. Notification topics are used
// for publishing DAX events to subscribers using Amazon Simple Notification
// Service (SNS).
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dax-2017-04-19/NotificationConfiguration
type NotificationConfiguration struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) that identifies the topic.
	TopicArn *string `type:"string"`

	// The current state of the topic.
	TopicStatus *string `type:"string"`
}

// String returns the string representation
func (s NotificationConfiguration) String() string {
	return awsutil.Prettify(s)
}

// Describes an individual setting that controls some aspect of DAX behavior.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dax-2017-04-19/Parameter
type Parameter struct {
	_ struct{} `type:"structure"`

	// A range of values within which the parameter can be set.
	AllowedValues *string `type:"string"`

	// The conditions under which changes to this parameter can be applied. For
	// example, requires-reboot indicates that a new value for this parameter will
	// only take effect if a node is rebooted.
	ChangeType ChangeType `type:"string" enum:"true"`

	// The data type of the parameter. For example, integer:
	DataType *string `type:"string"`

	// A description of the parameter
	Description *string `type:"string"`

	// Whether the customer is allowed to modify the parameter.
	IsModifiable IsModifiable `type:"string" enum:"true"`

	// A list of node types, and specific parameter values for each node.
	NodeTypeSpecificValues []NodeTypeSpecificValue `type:"list"`

	// The name of the parameter.
	ParameterName *string `type:"string"`

	// Determines whether the parameter can be applied to any nodes, or only nodes
	// of a particular type.
	ParameterType ParameterType `type:"string" enum:"true"`

	// The value for the parameter.
	ParameterValue *string `type:"string"`

	// How the parameter is defined. For example, system denotes a system-defined
	// parameter.
	Source *string `type:"string"`
}

// String returns the string representation
func (s Parameter) String() string {
	return awsutil.Prettify(s)
}

// A named set of parameters that are applied to all of the nodes in a DAX cluster.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dax-2017-04-19/ParameterGroup
type ParameterGroup struct {
	_ struct{} `type:"structure"`

	// A description of the parameter group.
	Description *string `type:"string"`

	// The name of the parameter group.
	ParameterGroupName *string `type:"string"`
}

// String returns the string representation
func (s ParameterGroup) String() string {
	return awsutil.Prettify(s)
}

// The status of a parameter group.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dax-2017-04-19/ParameterGroupStatus
type ParameterGroupStatus struct {
	_ struct{} `type:"structure"`

	// The node IDs of one or more nodes to be rebooted.
	NodeIdsToReboot []string `type:"list"`

	// The status of parameter updates.
	ParameterApplyStatus *string `type:"string"`

	// The name of the parameter group.
	ParameterGroupName *string `type:"string"`
}

// String returns the string representation
func (s ParameterGroupStatus) String() string {
	return awsutil.Prettify(s)
}

// An individual DAX parameter.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dax-2017-04-19/ParameterNameValue
type ParameterNameValue struct {
	_ struct{} `type:"structure"`

	// The name of the parameter.
	ParameterName *string `type:"string"`

	// The value of the parameter.
	ParameterValue *string `type:"string"`
}

// String returns the string representation
func (s ParameterNameValue) String() string {
	return awsutil.Prettify(s)
}

// The description of the server-side encryption status on the specified DAX
// cluster.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dax-2017-04-19/SSEDescription
type SSEDescription struct {
	_ struct{} `type:"structure"`

	// The current state of server-side encryption:
	//
	//    * ENABLING - Server-side encryption is being enabled.
	//
	//    * ENABLED - Server-side encryption is enabled.
	//
	//    * DISABLING - Server-side encryption is being disabled.
	//
	//    * DISABLED - Server-side encryption is disabled.
	Status SSEStatus `type:"string" enum:"true"`
}

// String returns the string representation
func (s SSEDescription) String() string {
	return awsutil.Prettify(s)
}

// Represents the settings used to enable server-side encryption.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dax-2017-04-19/SSESpecification
type SSESpecification struct {
	_ struct{} `type:"structure"`

	// Indicates whether server-side encryption is enabled (true) or disabled (false)
	// on the cluster.
	//
	// Enabled is a required field
	Enabled *bool `type:"boolean" required:"true"`
}

// String returns the string representation
func (s SSESpecification) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SSESpecification) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SSESpecification"}

	if s.Enabled == nil {
		invalidParams.Add(aws.NewErrParamRequired("Enabled"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// An individual VPC security group and its status.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dax-2017-04-19/SecurityGroupMembership
type SecurityGroupMembership struct {
	_ struct{} `type:"structure"`

	// The unique ID for this security group.
	SecurityGroupIdentifier *string `type:"string"`

	// The status of this security group.
	Status *string `type:"string"`
}

// String returns the string representation
func (s SecurityGroupMembership) String() string {
	return awsutil.Prettify(s)
}

// Represents the subnet associated with a DAX cluster. This parameter refers
// to subnets defined in Amazon Virtual Private Cloud (Amazon VPC) and used
// with DAX.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dax-2017-04-19/Subnet
type Subnet struct {
	_ struct{} `type:"structure"`

	// The Availability Zone (AZ) for subnet subnet.
	SubnetAvailabilityZone *string `type:"string"`

	// The system-assigned identifier for the subnet.
	SubnetIdentifier *string `type:"string"`
}

// String returns the string representation
func (s Subnet) String() string {
	return awsutil.Prettify(s)
}

// Represents the output of one of the following actions:
//
//    * CreateSubnetGroup
//
//    * ModifySubnetGroup
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dax-2017-04-19/SubnetGroup
type SubnetGroup struct {
	_ struct{} `type:"structure"`

	// The description of the subnet group.
	Description *string `type:"string"`

	// The name of the subnet group.
	SubnetGroupName *string `type:"string"`

	// A list of subnets associated with the subnet group.
	Subnets []Subnet `type:"list"`

	// The Amazon Virtual Private Cloud identifier (VPC ID) of the subnet group.
	VpcId *string `type:"string"`
}

// String returns the string representation
func (s SubnetGroup) String() string {
	return awsutil.Prettify(s)
}

// A description of a tag. Every tag is a key-value pair. You can add up to
// 50 tags to a single DAX cluster.
//
// AWS-assigned tag names and values are automatically assigned the aws: prefix,
// which the user cannot assign. AWS-assigned tag names do not count towards
// the tag limit of 50. User-assigned tag names have the prefix user:.
//
// You cannot backdate the application of a tag.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dax-2017-04-19/Tag
type Tag struct {
	_ struct{} `type:"structure"`

	// The key for the tag. Tag keys are case sensitive. Every DAX cluster can only
	// have one tag with the same key. If you try to add an existing tag (same key),
	// the existing tag value will be updated to the new value.
	Key *string `type:"string"`

	// The value of the tag. Tag values are case-sensitive and can be null.
	Value *string `type:"string"`
}

// String returns the string representation
func (s Tag) String() string {
	return awsutil.Prettify(s)
}
