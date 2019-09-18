// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicequotas

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

var _ aws.Config
var _ = awsutil.Prettify

// Returns an error that explains why the action did not succeed.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/service-quotas-2019-06-24/ErrorReason
type ErrorReason struct {
	_ struct{} `type:"structure"`

	// Service Quotas returns the following error values.
	//
	// DEPENDENCY_ACCESS_DENIED_ERROR is returned when the caller does not have
	// permission to call the service or service quota. To resolve the error, you
	// need permission to access the service or service quota.
	//
	// DEPENDENCY_THROTTLING_ERROR is returned when the service being called is
	// throttling Service Quotas.
	//
	// DEPENDENCY_SERVICE_ERROR is returned when the service being called has availability
	// issues.
	//
	// SERVICE_QUOTA_NOT_AVAILABLE_ERROR is returned when there was an error in
	// Service Quotas.
	ErrorCode ErrorCode `type:"string" enum:"true"`

	// The error message that provides more detail.
	ErrorMessage *string `type:"string"`
}

// String returns the string representation
func (s ErrorReason) String() string {
	return awsutil.Prettify(s)
}

// A structure that uses CloudWatch metrics to gather data about the service
// quota.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/service-quotas-2019-06-24/MetricInfo
type MetricInfo struct {
	_ struct{} `type:"structure"`

	// A dimension is a name/value pair that is part of the identity of a metric.
	// Every metric has specific characteristics that describe it, and you can think
	// of dimensions as categories for those characteristics. These dimensions are
	// part of the CloudWatch Metric Identity that measures usage against a particular
	// service quota.
	MetricDimensions map[string]string `type:"map"`

	// The name of the CloudWatch metric that measures usage of a service quota.
	// This is a required field.
	MetricName *string `type:"string"`

	// The namespace of the metric. The namespace is a container for CloudWatch
	// metrics. You can specify a name for the namespace when you create a metric.
	MetricNamespace *string `type:"string"`

	// Statistics are metric data aggregations over specified periods of time. This
	// is the recommended statistic to use when comparing usage in the CloudWatch
	// Metric against your Service Quota.
	MetricStatisticRecommendation *string `min:"1" type:"string"`
}

// String returns the string representation
func (s MetricInfo) String() string {
	return awsutil.Prettify(s)
}

// A structure that contains information about the quota period.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/service-quotas-2019-06-24/QuotaPeriod
type QuotaPeriod struct {
	_ struct{} `type:"structure"`

	// The time unit of a period.
	PeriodUnit PeriodUnit `type:"string" enum:"true"`

	// The value of a period.
	PeriodValue *int64 `type:"integer"`
}

// String returns the string representation
func (s QuotaPeriod) String() string {
	return awsutil.Prettify(s)
}

// A structure that contains information about a requested change for a quota.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/service-quotas-2019-06-24/RequestedServiceQuotaChange
type RequestedServiceQuotaChange struct {
	_ struct{} `type:"structure"`

	// The case Id for the service quota increase request.
	CaseId *string `type:"string"`

	// The date and time when the service quota increase request was received and
	// the case Id was created.
	Created *time.Time `type:"timestamp"`

	// New increased value for the service quota.
	DesiredValue *float64 `type:"double"`

	// Identifies if the quota is global.
	GlobalQuota *bool `type:"boolean"`

	// The unique identifier of a requested service quota change.
	Id *string `min:"1" type:"string"`

	// The date and time of the most recent change in the service quota increase
	// request.
	LastUpdated *time.Time `type:"timestamp"`

	// The Amazon Resource Name (ARN) of the service quota.
	QuotaArn *string `type:"string"`

	// Specifies the service quota that you want to use.
	QuotaCode *string `min:"1" type:"string"`

	// Name of the service quota.
	QuotaName *string `type:"string"`

	// The IAM identity who submitted the service quota increase request.
	Requester *string `type:"string"`

	// Specifies the service that you want to use.
	ServiceCode *string `min:"1" type:"string"`

	// The name of the AWS service specified in the increase request.
	ServiceName *string `type:"string"`

	// State of the service quota increase request.
	Status RequestStatus `type:"string" enum:"true"`

	// Specifies the unit used for the quota.
	Unit *string `type:"string"`
}

// String returns the string representation
func (s RequestedServiceQuotaChange) String() string {
	return awsutil.Prettify(s)
}

// A structure that contains the ServiceName and ServiceCode. It does not include
// all details of the service quota. To get those values, use the ListServiceQuotas
// operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/service-quotas-2019-06-24/ServiceInfo
type ServiceInfo struct {
	_ struct{} `type:"structure"`

	// Specifies the service that you want to use.
	ServiceCode *string `min:"1" type:"string"`

	// The name of the AWS service specified in the increase request.
	ServiceName *string `type:"string"`
}

// String returns the string representation
func (s ServiceInfo) String() string {
	return awsutil.Prettify(s)
}

// A structure that contains the full set of details that define the service
// quota.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/service-quotas-2019-06-24/ServiceQuota
type ServiceQuota struct {
	_ struct{} `type:"structure"`

	// Specifies if the quota value can be increased.
	Adjustable *bool `type:"boolean"`

	// Specifies the ErrorCode and ErrorMessage when success isn't achieved.
	ErrorReason *ErrorReason `type:"structure"`

	// Specifies if the quota is global.
	GlobalQuota *bool `type:"boolean"`

	// Identifies the unit and value of how time is measured.
	Period *QuotaPeriod `type:"structure"`

	// The Amazon Resource Name (ARN) of the service quota.
	QuotaArn *string `type:"string"`

	// The code identifier for the service quota specified.
	QuotaCode *string `min:"1" type:"string"`

	// The name identifier of the service quota.
	QuotaName *string `type:"string"`

	// Specifies the service that you want to use.
	ServiceCode *string `min:"1" type:"string"`

	// The name of the AWS service specified in the increase request.
	ServiceName *string `type:"string"`

	// The unit of measurement for the value of the service quota.
	Unit *string `type:"string"`

	// Specifies the details about the measurement.
	UsageMetric *MetricInfo `type:"structure"`

	// The value of service quota.
	Value *float64 `type:"double"`
}

// String returns the string representation
func (s ServiceQuota) String() string {
	return awsutil.Prettify(s)
}

// A structure that contains information about one service quota increase request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/service-quotas-2019-06-24/ServiceQuotaIncreaseRequestInTemplate
type ServiceQuotaIncreaseRequestInTemplate struct {
	_ struct{} `type:"structure"`

	// The AWS Region where the increase request occurs.
	AwsRegion *string `min:"1" type:"string"`

	// Identifies the new, increased value of the service quota in the increase
	// request.
	DesiredValue *float64 `type:"double"`

	// Specifies if the quota is a global quota.
	GlobalQuota *bool `type:"boolean"`

	// The code identifier for the service quota specified in the increase request.
	QuotaCode *string `min:"1" type:"string"`

	// The name of the service quota in the increase request.
	QuotaName *string `type:"string"`

	// The code identifier for the AWS service specified in the increase request.
	ServiceCode *string `min:"1" type:"string"`

	// The name of the AWS service specified in the increase request.
	ServiceName *string `type:"string"`

	// The unit of measure for the increase request.
	Unit *string `type:"string"`
}

// String returns the string representation
func (s ServiceQuotaIncreaseRequestInTemplate) String() string {
	return awsutil.Prettify(s)
}
