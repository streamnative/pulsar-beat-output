// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotjobsdataplane

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

var _ aws.Config
var _ = awsutil.Prettify

// Contains data about a job execution.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iot-jobs-data-2017-09-29/JobExecution
type JobExecution struct {
	_ struct{} `type:"structure"`

	// The estimated number of seconds that remain before the job execution status
	// will be changed to TIMED_OUT.
	ApproximateSecondsBeforeTimedOut *int64 `locationName:"approximateSecondsBeforeTimedOut" type:"long"`

	// A number that identifies a particular job execution on a particular device.
	// It can be used later in commands that return or update job execution information.
	ExecutionNumber *int64 `locationName:"executionNumber" type:"long"`

	// The content of the job document.
	JobDocument *string `locationName:"jobDocument" type:"string"`

	// The unique identifier you assigned to this job when it was created.
	JobId *string `locationName:"jobId" min:"1" type:"string"`

	// The time, in milliseconds since the epoch, when the job execution was last
	// updated.
	LastUpdatedAt *int64 `locationName:"lastUpdatedAt" type:"long"`

	// The time, in milliseconds since the epoch, when the job execution was enqueued.
	QueuedAt *int64 `locationName:"queuedAt" type:"long"`

	// The time, in milliseconds since the epoch, when the job execution was started.
	StartedAt *int64 `locationName:"startedAt" type:"long"`

	// The status of the job execution. Can be one of: "QUEUED", "IN_PROGRESS",
	// "FAILED", "SUCCESS", "CANCELED", "REJECTED", or "REMOVED".
	Status JobExecutionStatus `locationName:"status" type:"string" enum:"true"`

	// A collection of name/value pairs that describe the status of the job execution.
	StatusDetails map[string]string `locationName:"statusDetails" type:"map"`

	// The name of the thing that is executing the job.
	ThingName *string `locationName:"thingName" min:"1" type:"string"`

	// The version of the job execution. Job execution versions are incremented
	// each time they are updated by a device.
	VersionNumber *int64 `locationName:"versionNumber" type:"long"`
}

// String returns the string representation
func (s JobExecution) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s JobExecution) MarshalFields(e protocol.FieldEncoder) error {
	if s.ApproximateSecondsBeforeTimedOut != nil {
		v := *s.ApproximateSecondsBeforeTimedOut

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "approximateSecondsBeforeTimedOut", protocol.Int64Value(v), metadata)
	}
	if s.ExecutionNumber != nil {
		v := *s.ExecutionNumber

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "executionNumber", protocol.Int64Value(v), metadata)
	}
	if s.JobDocument != nil {
		v := *s.JobDocument

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "jobDocument", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.JobId != nil {
		v := *s.JobId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "jobId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.LastUpdatedAt != nil {
		v := *s.LastUpdatedAt

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "lastUpdatedAt", protocol.Int64Value(v), metadata)
	}
	if s.QueuedAt != nil {
		v := *s.QueuedAt

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "queuedAt", protocol.Int64Value(v), metadata)
	}
	if s.StartedAt != nil {
		v := *s.StartedAt

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "startedAt", protocol.Int64Value(v), metadata)
	}
	if len(s.Status) > 0 {
		v := s.Status

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "status", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.StatusDetails != nil {
		v := s.StatusDetails

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "statusDetails", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.ThingName != nil {
		v := *s.ThingName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "thingName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VersionNumber != nil {
		v := *s.VersionNumber

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "versionNumber", protocol.Int64Value(v), metadata)
	}
	return nil
}

// Contains data about the state of a job execution.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iot-jobs-data-2017-09-29/JobExecutionState
type JobExecutionState struct {
	_ struct{} `type:"structure"`

	// The status of the job execution. Can be one of: "QUEUED", "IN_PROGRESS",
	// "FAILED", "SUCCESS", "CANCELED", "REJECTED", or "REMOVED".
	Status JobExecutionStatus `locationName:"status" type:"string" enum:"true"`

	// A collection of name/value pairs that describe the status of the job execution.
	StatusDetails map[string]string `locationName:"statusDetails" type:"map"`

	// The version of the job execution. Job execution versions are incremented
	// each time they are updated by a device.
	VersionNumber *int64 `locationName:"versionNumber" type:"long"`
}

// String returns the string representation
func (s JobExecutionState) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s JobExecutionState) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.Status) > 0 {
		v := s.Status

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "status", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.StatusDetails != nil {
		v := s.StatusDetails

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "statusDetails", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.VersionNumber != nil {
		v := *s.VersionNumber

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "versionNumber", protocol.Int64Value(v), metadata)
	}
	return nil
}

// Contains a subset of information about a job execution.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iot-jobs-data-2017-09-29/JobExecutionSummary
type JobExecutionSummary struct {
	_ struct{} `type:"structure"`

	// A number that identifies a particular job execution on a particular device.
	ExecutionNumber *int64 `locationName:"executionNumber" type:"long"`

	// The unique identifier you assigned to this job when it was created.
	JobId *string `locationName:"jobId" min:"1" type:"string"`

	// The time, in milliseconds since the epoch, when the job execution was last
	// updated.
	LastUpdatedAt *int64 `locationName:"lastUpdatedAt" type:"long"`

	// The time, in milliseconds since the epoch, when the job execution was enqueued.
	QueuedAt *int64 `locationName:"queuedAt" type:"long"`

	// The time, in milliseconds since the epoch, when the job execution started.
	StartedAt *int64 `locationName:"startedAt" type:"long"`

	// The version of the job execution. Job execution versions are incremented
	// each time AWS IoT Jobs receives an update from a device.
	VersionNumber *int64 `locationName:"versionNumber" type:"long"`
}

// String returns the string representation
func (s JobExecutionSummary) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s JobExecutionSummary) MarshalFields(e protocol.FieldEncoder) error {
	if s.ExecutionNumber != nil {
		v := *s.ExecutionNumber

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "executionNumber", protocol.Int64Value(v), metadata)
	}
	if s.JobId != nil {
		v := *s.JobId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "jobId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.LastUpdatedAt != nil {
		v := *s.LastUpdatedAt

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "lastUpdatedAt", protocol.Int64Value(v), metadata)
	}
	if s.QueuedAt != nil {
		v := *s.QueuedAt

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "queuedAt", protocol.Int64Value(v), metadata)
	}
	if s.StartedAt != nil {
		v := *s.StartedAt

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "startedAt", protocol.Int64Value(v), metadata)
	}
	if s.VersionNumber != nil {
		v := *s.VersionNumber

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "versionNumber", protocol.Int64Value(v), metadata)
	}
	return nil
}
