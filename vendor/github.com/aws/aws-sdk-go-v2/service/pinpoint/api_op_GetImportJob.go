// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/GetImportJobRequest
type GetImportJobInput struct {
	_ struct{} `type:"structure"`

	// ApplicationId is a required field
	ApplicationId *string `location:"uri" locationName:"application-id" type:"string" required:"true"`

	// JobId is a required field
	JobId *string `location:"uri" locationName:"job-id" type:"string" required:"true"`
}

// String returns the string representation
func (s GetImportJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetImportJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetImportJobInput"}

	if s.ApplicationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationId"))
	}

	if s.JobId == nil {
		invalidParams.Add(aws.NewErrParamRequired("JobId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetImportJobInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ApplicationId != nil {
		v := *s.ApplicationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "application-id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.JobId != nil {
		v := *s.JobId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "job-id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/GetImportJobResponse
type GetImportJobOutput struct {
	_ struct{} `type:"structure" payload:"ImportJobResponse"`

	// Provides information about the status and settings of a job that imports
	// endpoint definitions from one or more files. The files can be stored in an
	// Amazon Simple Storage Service (Amazon S3) bucket or uploaded directly from
	// a computer by using the Amazon Pinpoint console.
	//
	// ImportJobResponse is a required field
	ImportJobResponse *ImportJobResponse `type:"structure" required:"true"`
}

// String returns the string representation
func (s GetImportJobOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetImportJobOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ImportJobResponse != nil {
		v := s.ImportJobResponse

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "ImportJobResponse", v, metadata)
	}
	return nil
}

const opGetImportJob = "GetImportJob"

// GetImportJobRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Retrieves information about the status and settings of a specific import
// job for an application.
//
//    // Example sending a request using GetImportJobRequest.
//    req := client.GetImportJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/GetImportJob
func (c *Client) GetImportJobRequest(input *GetImportJobInput) GetImportJobRequest {
	op := &aws.Operation{
		Name:       opGetImportJob,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/apps/{application-id}/jobs/import/{job-id}",
	}

	if input == nil {
		input = &GetImportJobInput{}
	}

	req := c.newRequest(op, input, &GetImportJobOutput{})
	return GetImportJobRequest{Request: req, Input: input, Copy: c.GetImportJobRequest}
}

// GetImportJobRequest is the request type for the
// GetImportJob API operation.
type GetImportJobRequest struct {
	*aws.Request
	Input *GetImportJobInput
	Copy  func(*GetImportJobInput) GetImportJobRequest
}

// Send marshals and sends the GetImportJob API request.
func (r GetImportJobRequest) Send(ctx context.Context) (*GetImportJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetImportJobResponse{
		GetImportJobOutput: r.Request.Data.(*GetImportJobOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetImportJobResponse is the response type for the
// GetImportJob API operation.
type GetImportJobResponse struct {
	*GetImportJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetImportJob request.
func (r *GetImportJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
