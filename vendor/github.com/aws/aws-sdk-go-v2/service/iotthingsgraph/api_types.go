// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotthingsgraph

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

var _ aws.Config
var _ = awsutil.Prettify

// A document that defines an entity.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotthingsgraph-2018-09-06/DefinitionDocument
type DefinitionDocument struct {
	_ struct{} `type:"structure"`

	// The language used to define the entity. GRAPHQL is the only valid value.
	//
	// Language is a required field
	Language DefinitionLanguage `locationName:"language" type:"string" required:"true" enum:"true"`

	// The GraphQL text that defines the entity.
	//
	// Text is a required field
	Text *string `locationName:"text" type:"string" required:"true"`
}

// String returns the string representation
func (s DefinitionDocument) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DefinitionDocument) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DefinitionDocument"}
	if len(s.Language) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Language"))
	}

	if s.Text == nil {
		invalidParams.Add(aws.NewErrParamRequired("Text"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// An object that contains the ID and revision number of a workflow or system
// that is part of a deployment.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotthingsgraph-2018-09-06/DependencyRevision
type DependencyRevision struct {
	_ struct{} `type:"structure"`

	// The ID of the workflow or system.
	Id *string `locationName:"id" type:"string"`

	// The revision number of the workflow or system.
	RevisionNumber *int64 `locationName:"revisionNumber" type:"long"`
}

// String returns the string representation
func (s DependencyRevision) String() string {
	return awsutil.Prettify(s)
}

// Describes the properties of an entity.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotthingsgraph-2018-09-06/EntityDescription
type EntityDescription struct {
	_ struct{} `type:"structure"`

	// The entity ARN.
	Arn *string `locationName:"arn" type:"string"`

	// The time at which the entity was created.
	CreatedAt *time.Time `locationName:"createdAt" type:"timestamp"`

	// The definition document of the entity.
	Definition *DefinitionDocument `locationName:"definition" type:"structure"`

	// The entity ID.
	Id *string `locationName:"id" type:"string"`

	// The entity type.
	Type EntityType `locationName:"type" type:"string" enum:"true"`
}

// String returns the string representation
func (s EntityDescription) String() string {
	return awsutil.Prettify(s)
}

// An object that filters an entity search. Multiple filters function as OR
// criteria in the search. For example a search that includes a NAMESPACE and
// a REFERENCED_ENTITY_ID filter searches for entities in the specified namespace
// that use the entity specified by the value of REFERENCED_ENTITY_ID.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotthingsgraph-2018-09-06/EntityFilter
type EntityFilter struct {
	_ struct{} `type:"structure"`

	// The name of the entity search filter field. REFERENCED_ENTITY_ID filters
	// on entities that are used by the entity in the result set. For example, you
	// can filter on the ID of a property that is used in a state.
	Name EntityFilterName `locationName:"name" type:"string" enum:"true"`

	// An array of string values for the search filter field. Multiple values function
	// as AND criteria in the search.
	Value []string `locationName:"value" type:"list"`
}

// String returns the string representation
func (s EntityFilter) String() string {
	return awsutil.Prettify(s)
}

// An object that contains information about a flow event.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotthingsgraph-2018-09-06/FlowExecutionMessage
type FlowExecutionMessage struct {
	_ struct{} `type:"structure"`

	// The type of flow event .
	EventType FlowExecutionEventType `locationName:"eventType" type:"string" enum:"true"`

	// The unique identifier of the message.
	MessageId *string `locationName:"messageId" type:"string"`

	// A string containing information about the flow event.
	Payload *string `locationName:"payload" type:"string"`

	// The date and time when the message was last updated.
	Timestamp *time.Time `locationName:"timestamp" type:"timestamp"`
}

// String returns the string representation
func (s FlowExecutionMessage) String() string {
	return awsutil.Prettify(s)
}

// An object that contains summary information about a flow execution.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotthingsgraph-2018-09-06/FlowExecutionSummary
type FlowExecutionSummary struct {
	_ struct{} `type:"structure"`

	// The date and time when the flow execution summary was created.
	CreatedAt *time.Time `locationName:"createdAt" type:"timestamp"`

	// The ID of the flow execution.
	FlowExecutionId *string `locationName:"flowExecutionId" type:"string"`

	// The ID of the flow.
	FlowTemplateId *string `locationName:"flowTemplateId" type:"string"`

	// The current status of the flow execution.
	Status FlowExecutionStatus `locationName:"status" type:"string" enum:"true"`

	// The ID of the system instance that contains the flow.
	SystemInstanceId *string `locationName:"systemInstanceId" type:"string"`

	// The date and time when the flow execution summary was last updated.
	UpdatedAt *time.Time `locationName:"updatedAt" type:"timestamp"`
}

// String returns the string representation
func (s FlowExecutionSummary) String() string {
	return awsutil.Prettify(s)
}

// An object that contains a workflow's definition and summary information.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotthingsgraph-2018-09-06/FlowTemplateDescription
type FlowTemplateDescription struct {
	_ struct{} `type:"structure"`

	// A workflow's definition document.
	Definition *DefinitionDocument `locationName:"definition" type:"structure"`

	// An object that contains summary information about a workflow.
	Summary *FlowTemplateSummary `locationName:"summary" type:"structure"`

	// The version of the user's namespace against which the workflow was validated.
	// Use this value in your system instance.
	ValidatedNamespaceVersion *int64 `locationName:"validatedNamespaceVersion" type:"long"`
}

// String returns the string representation
func (s FlowTemplateDescription) String() string {
	return awsutil.Prettify(s)
}

// An object that filters a workflow search.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotthingsgraph-2018-09-06/FlowTemplateFilter
type FlowTemplateFilter struct {
	_ struct{} `type:"structure"`

	// The name of the search filter field.
	//
	// Name is a required field
	Name FlowTemplateFilterName `locationName:"name" type:"string" required:"true" enum:"true"`

	// An array of string values for the search filter field. Multiple values function
	// as AND criteria in the search.
	//
	// Value is a required field
	Value []string `locationName:"value" type:"list" required:"true"`
}

// String returns the string representation
func (s FlowTemplateFilter) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *FlowTemplateFilter) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "FlowTemplateFilter"}
	if len(s.Name) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if s.Value == nil {
		invalidParams.Add(aws.NewErrParamRequired("Value"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// An object that contains summary information about a workflow.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotthingsgraph-2018-09-06/FlowTemplateSummary
type FlowTemplateSummary struct {
	_ struct{} `type:"structure"`

	// The ARN of the workflow.
	Arn *string `locationName:"arn" type:"string"`

	// The date when the workflow was created.
	CreatedAt *time.Time `locationName:"createdAt" type:"timestamp"`

	// The ID of the workflow.
	Id *string `locationName:"id" type:"string"`

	// The revision number of the workflow.
	RevisionNumber *int64 `locationName:"revisionNumber" type:"long"`
}

// String returns the string representation
func (s FlowTemplateSummary) String() string {
	return awsutil.Prettify(s)
}

// An object that specifies whether cloud metrics are collected in a deployment
// and, if so, what role is used to collect metrics.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotthingsgraph-2018-09-06/MetricsConfiguration
type MetricsConfiguration struct {
	_ struct{} `type:"structure"`

	// A Boolean that specifies whether cloud metrics are collected.
	CloudMetricEnabled *bool `locationName:"cloudMetricEnabled" type:"boolean"`

	// The ARN of the role that is used to collect cloud metrics.
	MetricRuleRoleArn *string `locationName:"metricRuleRoleArn" min:"20" type:"string"`
}

// String returns the string representation
func (s MetricsConfiguration) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *MetricsConfiguration) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "MetricsConfiguration"}
	if s.MetricRuleRoleArn != nil && len(*s.MetricRuleRoleArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("MetricRuleRoleArn", 20))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// An object that contains a system instance definition and summary information.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotthingsgraph-2018-09-06/SystemInstanceDescription
type SystemInstanceDescription struct {
	_ struct{} `type:"structure"`

	// A document that defines an entity.
	Definition *DefinitionDocument `locationName:"definition" type:"structure"`

	// The AWS Identity and Access Management (IAM) role that AWS IoT Things Graph
	// assumes during flow execution in a cloud deployment. This role must have
	// read and write permissionss to AWS Lambda and AWS IoT and to any other AWS
	// services that the flow uses.
	FlowActionsRoleArn *string `locationName:"flowActionsRoleArn" min:"20" type:"string"`

	// An object that specifies whether cloud metrics are collected in a deployment
	// and, if so, what role is used to collect metrics.
	MetricsConfiguration *MetricsConfiguration `locationName:"metricsConfiguration" type:"structure"`

	// The Amazon Simple Storage Service bucket where information about a system
	// instance is stored.
	S3BucketName *string `locationName:"s3BucketName" type:"string"`

	// An object that contains summary information about a system instance.
	Summary *SystemInstanceSummary `locationName:"summary" type:"structure"`

	// A list of objects that contain all of the IDs and revision numbers of workflows
	// and systems that are used in a system instance.
	ValidatedDependencyRevisions []DependencyRevision `locationName:"validatedDependencyRevisions" type:"list"`

	// The version of the user's namespace against which the system instance was
	// validated.
	ValidatedNamespaceVersion *int64 `locationName:"validatedNamespaceVersion" type:"long"`
}

// String returns the string representation
func (s SystemInstanceDescription) String() string {
	return awsutil.Prettify(s)
}

// An object that filters a system instance search. Multiple filters function
// as OR criteria in the search. For example a search that includes a GREENGRASS_GROUP_NAME
// and a STATUS filter searches for system instances in the specified Greengrass
// group that have the specified status.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotthingsgraph-2018-09-06/SystemInstanceFilter
type SystemInstanceFilter struct {
	_ struct{} `type:"structure"`

	// The name of the search filter field.
	Name SystemInstanceFilterName `locationName:"name" type:"string" enum:"true"`

	// An array of string values for the search filter field. Multiple values function
	// as AND criteria in the search.
	Value []string `locationName:"value" type:"list"`
}

// String returns the string representation
func (s SystemInstanceFilter) String() string {
	return awsutil.Prettify(s)
}

// An object that contains summary information about a system instance.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotthingsgraph-2018-09-06/SystemInstanceSummary
type SystemInstanceSummary struct {
	_ struct{} `type:"structure"`

	// The ARN of the system instance.
	Arn *string `locationName:"arn" type:"string"`

	// The date when the system instance was created.
	CreatedAt *time.Time `locationName:"createdAt" type:"timestamp"`

	// The ID of the Greengrass group where the system instance is deployed.
	GreengrassGroupId *string `locationName:"greengrassGroupId" type:"string"`

	// The ID of the Greengrass group where the system instance is deployed.
	GreengrassGroupName *string `locationName:"greengrassGroupName" type:"string"`

	// The version of the Greengrass group where the system instance is deployed.
	GreengrassGroupVersionId *string `locationName:"greengrassGroupVersionId" type:"string"`

	// The ID of the system instance.
	Id *string `locationName:"id" type:"string"`

	// The status of the system instance.
	Status SystemInstanceDeploymentStatus `locationName:"status" type:"string" enum:"true"`

	// The target of the system instance.
	Target DeploymentTarget `locationName:"target" type:"string" enum:"true"`

	// The date and time when the system instance was last updated.
	UpdatedAt *time.Time `locationName:"updatedAt" type:"timestamp"`
}

// String returns the string representation
func (s SystemInstanceSummary) String() string {
	return awsutil.Prettify(s)
}

// An object that contains a system's definition document and summary information.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotthingsgraph-2018-09-06/SystemTemplateDescription
type SystemTemplateDescription struct {
	_ struct{} `type:"structure"`

	// The definition document of a system.
	Definition *DefinitionDocument `locationName:"definition" type:"structure"`

	// An object that contains summary information about a system.
	Summary *SystemTemplateSummary `locationName:"summary" type:"structure"`

	// The namespace version against which the system was validated. Use this value
	// in your system instance.
	ValidatedNamespaceVersion *int64 `locationName:"validatedNamespaceVersion" type:"long"`
}

// String returns the string representation
func (s SystemTemplateDescription) String() string {
	return awsutil.Prettify(s)
}

// An object that filters a system search.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotthingsgraph-2018-09-06/SystemTemplateFilter
type SystemTemplateFilter struct {
	_ struct{} `type:"structure"`

	// The name of the system search filter field.
	//
	// Name is a required field
	Name SystemTemplateFilterName `locationName:"name" type:"string" required:"true" enum:"true"`

	// An array of string values for the search filter field. Multiple values function
	// as AND criteria in the search.
	//
	// Value is a required field
	Value []string `locationName:"value" type:"list" required:"true"`
}

// String returns the string representation
func (s SystemTemplateFilter) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SystemTemplateFilter) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SystemTemplateFilter"}
	if len(s.Name) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if s.Value == nil {
		invalidParams.Add(aws.NewErrParamRequired("Value"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// An object that contains information about a system.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotthingsgraph-2018-09-06/SystemTemplateSummary
type SystemTemplateSummary struct {
	_ struct{} `type:"structure"`

	// The ARN of the system.
	Arn *string `locationName:"arn" type:"string"`

	// The date when the system was created.
	CreatedAt *time.Time `locationName:"createdAt" type:"timestamp"`

	// The ID of the system.
	Id *string `locationName:"id" type:"string"`

	// The revision number of the system.
	RevisionNumber *int64 `locationName:"revisionNumber" type:"long"`
}

// String returns the string representation
func (s SystemTemplateSummary) String() string {
	return awsutil.Prettify(s)
}

// Metadata assigned to an AWS IoT Things Graph resource consisting of a key-value
// pair.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotthingsgraph-2018-09-06/Tag
type Tag struct {
	_ struct{} `type:"structure"`

	// The required name of the tag. The string value can be from 1 to 128 Unicode
	// characters in length.
	//
	// Key is a required field
	Key *string `locationName:"key" min:"1" type:"string" required:"true"`

	// The optional value of the tag. The string value can be from 1 to 256 Unicode
	// characters in length.
	//
	// Value is a required field
	Value *string `locationName:"value" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s Tag) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Tag) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "Tag"}

	if s.Key == nil {
		invalidParams.Add(aws.NewErrParamRequired("Key"))
	}
	if s.Key != nil && len(*s.Key) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Key", 1))
	}

	if s.Value == nil {
		invalidParams.Add(aws.NewErrParamRequired("Value"))
	}
	if s.Value != nil && len(*s.Value) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Value", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// An AWS IoT thing.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotthingsgraph-2018-09-06/Thing
type Thing struct {
	_ struct{} `type:"structure"`

	// The ARN of the thing.
	ThingArn *string `locationName:"thingArn" type:"string"`

	// The name of the thing.
	ThingName *string `locationName:"thingName" min:"1" type:"string"`
}

// String returns the string representation
func (s Thing) String() string {
	return awsutil.Prettify(s)
}
