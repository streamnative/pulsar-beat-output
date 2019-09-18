// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package securityhub

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/securityhub-2018-10-26/BatchImportFindingsRequest
type BatchImportFindingsInput struct {
	_ struct{} `type:"structure"`

	// A list of findings to import. To successfully import a finding, it must follow
	// the AWS Security Finding Format (https://docs.aws.amazon.com/securityhub/latest/userguide/securityhub-findings-format.html).
	//
	// Findings is a required field
	Findings []AwsSecurityFinding `type:"list" required:"true"`
}

// String returns the string representation
func (s BatchImportFindingsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BatchImportFindingsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "BatchImportFindingsInput"}

	if s.Findings == nil {
		invalidParams.Add(aws.NewErrParamRequired("Findings"))
	}
	if s.Findings != nil {
		for i, v := range s.Findings {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Findings", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s BatchImportFindingsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Findings != nil {
		v := s.Findings

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Findings", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/securityhub-2018-10-26/BatchImportFindingsResponse
type BatchImportFindingsOutput struct {
	_ struct{} `type:"structure"`

	// The number of findings that failed to import.
	//
	// FailedCount is a required field
	FailedCount *int64 `type:"integer" required:"true"`

	// The list of the findings that failed to import.
	FailedFindings []ImportFindingsError `type:"list"`

	// The number of findings that were successfully imported.
	//
	// SuccessCount is a required field
	SuccessCount *int64 `type:"integer" required:"true"`
}

// String returns the string representation
func (s BatchImportFindingsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s BatchImportFindingsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.FailedCount != nil {
		v := *s.FailedCount

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "FailedCount", protocol.Int64Value(v), metadata)
	}
	if s.FailedFindings != nil {
		v := s.FailedFindings

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "FailedFindings", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.SuccessCount != nil {
		v := *s.SuccessCount

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "SuccessCount", protocol.Int64Value(v), metadata)
	}
	return nil
}

const opBatchImportFindings = "BatchImportFindings"

// BatchImportFindingsRequest returns a request value for making API operation for
// AWS SecurityHub.
//
// Imports security findings generated from an integrated third-party product
// into Security Hub. This action is requested by the integrated product to
// import its findings into Security Hub. The maximum allowed size for a finding
// is 240 Kb. An error is returned for any finding larger than 240 Kb.
//
//    // Example sending a request using BatchImportFindingsRequest.
//    req := client.BatchImportFindingsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/securityhub-2018-10-26/BatchImportFindings
func (c *Client) BatchImportFindingsRequest(input *BatchImportFindingsInput) BatchImportFindingsRequest {
	op := &aws.Operation{
		Name:       opBatchImportFindings,
		HTTPMethod: "POST",
		HTTPPath:   "/findings/import",
	}

	if input == nil {
		input = &BatchImportFindingsInput{}
	}

	req := c.newRequest(op, input, &BatchImportFindingsOutput{})
	return BatchImportFindingsRequest{Request: req, Input: input, Copy: c.BatchImportFindingsRequest}
}

// BatchImportFindingsRequest is the request type for the
// BatchImportFindings API operation.
type BatchImportFindingsRequest struct {
	*aws.Request
	Input *BatchImportFindingsInput
	Copy  func(*BatchImportFindingsInput) BatchImportFindingsRequest
}

// Send marshals and sends the BatchImportFindings API request.
func (r BatchImportFindingsRequest) Send(ctx context.Context) (*BatchImportFindingsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &BatchImportFindingsResponse{
		BatchImportFindingsOutput: r.Request.Data.(*BatchImportFindingsOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// BatchImportFindingsResponse is the response type for the
// BatchImportFindings API operation.
type BatchImportFindingsResponse struct {
	*BatchImportFindingsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// BatchImportFindings request.
func (r *BatchImportFindingsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
