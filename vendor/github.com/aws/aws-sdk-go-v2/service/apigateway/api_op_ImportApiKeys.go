// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// The POST request to import API keys from an external source, such as a CSV-formatted
// file.
type ImportApiKeysInput struct {
	_ struct{} `type:"structure" payload:"Body"`

	// The payload of the POST request to import API keys. For the payload format,
	// see API Key File Format (https://docs.aws.amazon.com/apigateway/latest/developerguide/api-key-file-format.html).
	//
	// Body is a required field
	Body []byte `locationName:"body" type:"blob" required:"true"`

	// A query parameter to indicate whether to rollback ApiKey importation (true)
	// or not (false) when error is encountered.
	FailOnWarnings *bool `location:"querystring" locationName:"failonwarnings" type:"boolean"`

	// A query parameter to specify the input format to imported API keys. Currently,
	// only the csv format is supported.
	//
	// Format is a required field
	Format ApiKeysFormat `location:"querystring" locationName:"format" type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s ImportApiKeysInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ImportApiKeysInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ImportApiKeysInput"}

	if s.Body == nil {
		invalidParams.Add(aws.NewErrParamRequired("Body"))
	}
	if len(s.Format) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Format"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ImportApiKeysInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Body != nil {
		v := s.Body

		metadata := protocol.Metadata{}
		e.SetStream(protocol.PayloadTarget, "body", protocol.BytesStream(v), metadata)
	}
	if s.FailOnWarnings != nil {
		v := *s.FailOnWarnings

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "failonwarnings", protocol.BoolValue(v), metadata)
	}
	if len(s.Format) > 0 {
		v := s.Format

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "format", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

// The identifier of an ApiKey used in a UsagePlan.
type ImportApiKeysOutput struct {
	_ struct{} `type:"structure"`

	// A list of all the ApiKey identifiers.
	Ids []string `locationName:"ids" type:"list"`

	// A list of warning messages.
	Warnings []string `locationName:"warnings" type:"list"`
}

// String returns the string representation
func (s ImportApiKeysOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ImportApiKeysOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Ids != nil {
		v := s.Ids

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "ids", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.Warnings != nil {
		v := s.Warnings

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "warnings", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	return nil
}

const opImportApiKeys = "ImportApiKeys"

// ImportApiKeysRequest returns a request value for making API operation for
// Amazon API Gateway.
//
// Import API keys from an external source, such as a CSV-formatted file.
//
//    // Example sending a request using ImportApiKeysRequest.
//    req := client.ImportApiKeysRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) ImportApiKeysRequest(input *ImportApiKeysInput) ImportApiKeysRequest {
	op := &aws.Operation{
		Name:       opImportApiKeys,
		HTTPMethod: "POST",
		HTTPPath:   "/apikeys?mode=import",
	}

	if input == nil {
		input = &ImportApiKeysInput{}
	}

	req := c.newRequest(op, input, &ImportApiKeysOutput{})
	return ImportApiKeysRequest{Request: req, Input: input, Copy: c.ImportApiKeysRequest}
}

// ImportApiKeysRequest is the request type for the
// ImportApiKeys API operation.
type ImportApiKeysRequest struct {
	*aws.Request
	Input *ImportApiKeysInput
	Copy  func(*ImportApiKeysInput) ImportApiKeysRequest
}

// Send marshals and sends the ImportApiKeys API request.
func (r ImportApiKeysRequest) Send(ctx context.Context) (*ImportApiKeysResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ImportApiKeysResponse{
		ImportApiKeysOutput: r.Request.Data.(*ImportApiKeysOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ImportApiKeysResponse is the response type for the
// ImportApiKeys API operation.
type ImportApiKeysResponse struct {
	*ImportApiKeysOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ImportApiKeys request.
func (r *ImportApiKeysResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
