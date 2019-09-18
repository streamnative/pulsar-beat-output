// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lexmodelbuildingservice

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lex-models-2017-04-19/GetImportRequest
type GetImportInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the import job information to return.
	//
	// ImportId is a required field
	ImportId *string `location:"uri" locationName:"importId" type:"string" required:"true"`
}

// String returns the string representation
func (s GetImportInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetImportInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetImportInput"}

	if s.ImportId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ImportId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetImportInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ImportId != nil {
		v := *s.ImportId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "importId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lex-models-2017-04-19/GetImportResponse
type GetImportOutput struct {
	_ struct{} `type:"structure"`

	// A timestamp for the date and time that the import job was created.
	CreatedDate *time.Time `locationName:"createdDate" type:"timestamp"`

	// A string that describes why an import job failed to complete.
	FailureReason []string `locationName:"failureReason" type:"list"`

	// The identifier for the specific import job.
	ImportId *string `locationName:"importId" type:"string"`

	// The status of the import job. If the status is FAILED, you can get the reason
	// for the failure from the failureReason field.
	ImportStatus ImportStatus `locationName:"importStatus" type:"string" enum:"true"`

	// The action taken when there was a conflict between an existing resource and
	// a resource in the import file.
	MergeStrategy MergeStrategy `locationName:"mergeStrategy" type:"string" enum:"true"`

	// The name given to the import job.
	Name *string `locationName:"name" min:"1" type:"string"`

	// The type of resource imported.
	ResourceType ResourceType `locationName:"resourceType" type:"string" enum:"true"`
}

// String returns the string representation
func (s GetImportOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetImportOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.CreatedDate != nil {
		v := *s.CreatedDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "createdDate",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.FailureReason != nil {
		v := s.FailureReason

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "failureReason", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.ImportId != nil {
		v := *s.ImportId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "importId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.ImportStatus) > 0 {
		v := s.ImportStatus

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "importStatus", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if len(s.MergeStrategy) > 0 {
		v := s.MergeStrategy

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "mergeStrategy", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.ResourceType) > 0 {
		v := s.ResourceType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "resourceType", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

const opGetImport = "GetImport"

// GetImportRequest returns a request value for making API operation for
// Amazon Lex Model Building Service.
//
// Gets information about an import job started with the StartImport operation.
//
//    // Example sending a request using GetImportRequest.
//    req := client.GetImportRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lex-models-2017-04-19/GetImport
func (c *Client) GetImportRequest(input *GetImportInput) GetImportRequest {
	op := &aws.Operation{
		Name:       opGetImport,
		HTTPMethod: "GET",
		HTTPPath:   "/imports/{importId}",
	}

	if input == nil {
		input = &GetImportInput{}
	}

	req := c.newRequest(op, input, &GetImportOutput{})
	return GetImportRequest{Request: req, Input: input, Copy: c.GetImportRequest}
}

// GetImportRequest is the request type for the
// GetImport API operation.
type GetImportRequest struct {
	*aws.Request
	Input *GetImportInput
	Copy  func(*GetImportInput) GetImportRequest
}

// Send marshals and sends the GetImport API request.
func (r GetImportRequest) Send(ctx context.Context) (*GetImportResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetImportResponse{
		GetImportOutput: r.Request.Data.(*GetImportOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetImportResponse is the response type for the
// GetImport API operation.
type GetImportResponse struct {
	*GetImportOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetImport request.
func (r *GetImportResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
