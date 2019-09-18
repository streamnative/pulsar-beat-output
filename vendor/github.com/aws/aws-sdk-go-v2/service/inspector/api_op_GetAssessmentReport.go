// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package inspector

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/inspector-2016-02-16/GetAssessmentReportRequest
type GetAssessmentReportInput struct {
	_ struct{} `type:"structure"`

	// The ARN that specifies the assessment run for which you want to generate
	// a report.
	//
	// AssessmentRunArn is a required field
	AssessmentRunArn *string `locationName:"assessmentRunArn" min:"1" type:"string" required:"true"`

	// Specifies the file format (html or pdf) of the assessment report that you
	// want to generate.
	//
	// ReportFileFormat is a required field
	ReportFileFormat ReportFileFormat `locationName:"reportFileFormat" type:"string" required:"true" enum:"true"`

	// Specifies the type of the assessment report that you want to generate. There
	// are two types of assessment reports: a finding report and a full report.
	// For more information, see Assessment Reports (https://docs.aws.amazon.com/inspector/latest/userguide/inspector_reports.html).
	//
	// ReportType is a required field
	ReportType ReportType `locationName:"reportType" type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s GetAssessmentReportInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetAssessmentReportInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetAssessmentReportInput"}

	if s.AssessmentRunArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("AssessmentRunArn"))
	}
	if s.AssessmentRunArn != nil && len(*s.AssessmentRunArn) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AssessmentRunArn", 1))
	}
	if len(s.ReportFileFormat) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("ReportFileFormat"))
	}
	if len(s.ReportType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("ReportType"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/inspector-2016-02-16/GetAssessmentReportResponse
type GetAssessmentReportOutput struct {
	_ struct{} `type:"structure"`

	// Specifies the status of the request to generate an assessment report.
	//
	// Status is a required field
	Status ReportStatus `locationName:"status" type:"string" required:"true" enum:"true"`

	// Specifies the URL where you can find the generated assessment report. This
	// parameter is only returned if the report is successfully generated.
	Url *string `locationName:"url" type:"string"`
}

// String returns the string representation
func (s GetAssessmentReportOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetAssessmentReport = "GetAssessmentReport"

// GetAssessmentReportRequest returns a request value for making API operation for
// Amazon Inspector.
//
// Produces an assessment report that includes detailed and comprehensive results
// of a specified assessment run.
//
//    // Example sending a request using GetAssessmentReportRequest.
//    req := client.GetAssessmentReportRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/inspector-2016-02-16/GetAssessmentReport
func (c *Client) GetAssessmentReportRequest(input *GetAssessmentReportInput) GetAssessmentReportRequest {
	op := &aws.Operation{
		Name:       opGetAssessmentReport,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetAssessmentReportInput{}
	}

	req := c.newRequest(op, input, &GetAssessmentReportOutput{})
	return GetAssessmentReportRequest{Request: req, Input: input, Copy: c.GetAssessmentReportRequest}
}

// GetAssessmentReportRequest is the request type for the
// GetAssessmentReport API operation.
type GetAssessmentReportRequest struct {
	*aws.Request
	Input *GetAssessmentReportInput
	Copy  func(*GetAssessmentReportInput) GetAssessmentReportRequest
}

// Send marshals and sends the GetAssessmentReport API request.
func (r GetAssessmentReportRequest) Send(ctx context.Context) (*GetAssessmentReportResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetAssessmentReportResponse{
		GetAssessmentReportOutput: r.Request.Data.(*GetAssessmentReportOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetAssessmentReportResponse is the response type for the
// GetAssessmentReport API operation.
type GetAssessmentReportResponse struct {
	*GetAssessmentReportOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetAssessmentReport request.
func (r *GetAssessmentReportResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
