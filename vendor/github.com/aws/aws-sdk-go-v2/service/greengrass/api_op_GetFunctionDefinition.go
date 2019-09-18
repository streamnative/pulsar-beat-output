// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package greengrass

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/GetFunctionDefinitionRequest
type GetFunctionDefinitionInput struct {
	_ struct{} `type:"structure"`

	// FunctionDefinitionId is a required field
	FunctionDefinitionId *string `location:"uri" locationName:"FunctionDefinitionId" type:"string" required:"true"`
}

// String returns the string representation
func (s GetFunctionDefinitionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetFunctionDefinitionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetFunctionDefinitionInput"}

	if s.FunctionDefinitionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("FunctionDefinitionId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetFunctionDefinitionInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.FunctionDefinitionId != nil {
		v := *s.FunctionDefinitionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "FunctionDefinitionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/GetFunctionDefinitionResponse
type GetFunctionDefinitionOutput struct {
	_ struct{} `type:"structure"`

	Arn *string `type:"string"`

	CreationTimestamp *string `type:"string"`

	Id *string `type:"string"`

	LastUpdatedTimestamp *string `type:"string"`

	LatestVersion *string `type:"string"`

	LatestVersionArn *string `type:"string"`

	Name *string `type:"string"`

	// The key-value pair for the resource tag.
	Tags map[string]string `locationName:"tags" type:"map"`
}

// String returns the string representation
func (s GetFunctionDefinitionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetFunctionDefinitionOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CreationTimestamp != nil {
		v := *s.CreationTimestamp

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "CreationTimestamp", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.LastUpdatedTimestamp != nil {
		v := *s.LastUpdatedTimestamp

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "LastUpdatedTimestamp", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.LatestVersion != nil {
		v := *s.LatestVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "LatestVersion", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.LatestVersionArn != nil {
		v := *s.LatestVersionArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "LatestVersionArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	return nil
}

const opGetFunctionDefinition = "GetFunctionDefinition"

// GetFunctionDefinitionRequest returns a request value for making API operation for
// AWS Greengrass.
//
// Retrieves information about a Lambda function definition, including its creation
// time and latest version.
//
//    // Example sending a request using GetFunctionDefinitionRequest.
//    req := client.GetFunctionDefinitionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/GetFunctionDefinition
func (c *Client) GetFunctionDefinitionRequest(input *GetFunctionDefinitionInput) GetFunctionDefinitionRequest {
	op := &aws.Operation{
		Name:       opGetFunctionDefinition,
		HTTPMethod: "GET",
		HTTPPath:   "/greengrass/definition/functions/{FunctionDefinitionId}",
	}

	if input == nil {
		input = &GetFunctionDefinitionInput{}
	}

	req := c.newRequest(op, input, &GetFunctionDefinitionOutput{})
	return GetFunctionDefinitionRequest{Request: req, Input: input, Copy: c.GetFunctionDefinitionRequest}
}

// GetFunctionDefinitionRequest is the request type for the
// GetFunctionDefinition API operation.
type GetFunctionDefinitionRequest struct {
	*aws.Request
	Input *GetFunctionDefinitionInput
	Copy  func(*GetFunctionDefinitionInput) GetFunctionDefinitionRequest
}

// Send marshals and sends the GetFunctionDefinition API request.
func (r GetFunctionDefinitionRequest) Send(ctx context.Context) (*GetFunctionDefinitionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetFunctionDefinitionResponse{
		GetFunctionDefinitionOutput: r.Request.Data.(*GetFunctionDefinitionOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetFunctionDefinitionResponse is the response type for the
// GetFunctionDefinition API operation.
type GetFunctionDefinitionResponse struct {
	*GetFunctionDefinitionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetFunctionDefinition request.
func (r *GetFunctionDefinitionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
