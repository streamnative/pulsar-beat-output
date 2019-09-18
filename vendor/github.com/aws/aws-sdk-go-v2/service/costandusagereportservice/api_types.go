// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package costandusagereportservice

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

var _ aws.Config
var _ = awsutil.Prettify

// The definition of AWS Cost and Usage Report. You can specify the report name,
// time unit, report format, compression format, S3 bucket, additional artifacts,
// and schema elements in the definition.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cur-2017-01-06/ReportDefinition
type ReportDefinition struct {
	_ struct{} `type:"structure"`

	// A list of manifests that you want Amazon Web Services to create for this
	// report.
	AdditionalArtifacts []AdditionalArtifact `type:"list"`

	// A list of strings that indicate additional content that Amazon Web Services
	// includes in the report, such as individual resource IDs.
	//
	// AdditionalSchemaElements is a required field
	AdditionalSchemaElements []SchemaElement `type:"list" required:"true"`

	// The compression format that AWS uses for the report.
	//
	// Compression is a required field
	Compression CompressionFormat `type:"string" required:"true" enum:"true"`

	// The format that AWS saves the report in.
	//
	// Format is a required field
	Format ReportFormat `type:"string" required:"true" enum:"true"`

	// Whether you want Amazon Web Services to update your reports after they have
	// been finalized if Amazon Web Services detects charges related to previous
	// months. These charges can include refunds, credits, or support fees.
	RefreshClosedReports *bool `type:"boolean"`

	// The name of the report that you want to create. The name must be unique,
	// is case sensitive, and can't include spaces.
	//
	// ReportName is a required field
	ReportName *string `type:"string" required:"true"`

	// Whether you want Amazon Web Services to overwrite the previous version of
	// each report or to deliver the report in addition to the previous versions.
	ReportVersioning ReportVersioning `type:"string" enum:"true"`

	// The S3 bucket where AWS delivers the report.
	//
	// S3Bucket is a required field
	S3Bucket *string `type:"string" required:"true"`

	// The prefix that AWS adds to the report name when AWS delivers the report.
	// Your prefix can't include spaces.
	//
	// S3Prefix is a required field
	S3Prefix *string `type:"string" required:"true"`

	// The region of the S3 bucket that AWS delivers the report into.
	//
	// S3Region is a required field
	S3Region AWSRegion `type:"string" required:"true" enum:"true"`

	// The length of time covered by the report.
	//
	// TimeUnit is a required field
	TimeUnit TimeUnit `type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s ReportDefinition) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ReportDefinition) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ReportDefinition"}

	if s.AdditionalSchemaElements == nil {
		invalidParams.Add(aws.NewErrParamRequired("AdditionalSchemaElements"))
	}
	if len(s.Compression) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Compression"))
	}
	if len(s.Format) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Format"))
	}

	if s.ReportName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ReportName"))
	}

	if s.S3Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("S3Bucket"))
	}

	if s.S3Prefix == nil {
		invalidParams.Add(aws.NewErrParamRequired("S3Prefix"))
	}
	if len(s.S3Region) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("S3Region"))
	}
	if len(s.TimeUnit) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("TimeUnit"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}
