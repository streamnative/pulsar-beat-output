// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigatewayv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/apigatewayv2-2018-11-29/CreateIntegrationRequest
type CreateIntegrationInput struct {
	_ struct{} `type:"structure"`

	// ApiId is a required field
	ApiId *string `location:"uri" locationName:"apiId" type:"string" required:"true"`

	// A string with a length between [1-1024].
	ConnectionId *string `locationName:"connectionId" type:"string"`

	// Represents a connection type.
	ConnectionType ConnectionType `locationName:"connectionType" type:"string" enum:"true"`

	// Specifies how to handle response payload content type conversions.
	ContentHandlingStrategy ContentHandlingStrategy `locationName:"contentHandlingStrategy" type:"string" enum:"true"`

	// Represents an Amazon Resource Name (ARN).
	CredentialsArn *string `locationName:"credentialsArn" type:"string"`

	// A string with a length between [0-1024].
	Description *string `locationName:"description" type:"string"`

	// A string with a length between [1-64].
	IntegrationMethod *string `locationName:"integrationMethod" type:"string"`

	// Represents an API method integration type.
	//
	// IntegrationType is a required field
	IntegrationType IntegrationType `locationName:"integrationType" type:"string" required:"true" enum:"true"`

	// A string representation of a URI with a length between [1-2048].
	IntegrationUri *string `locationName:"integrationUri" type:"string"`

	// Represents passthrough behavior for an integration response.
	PassthroughBehavior PassthroughBehavior `locationName:"passthroughBehavior" type:"string" enum:"true"`

	// A key-value map specifying response parameters that are passed to the method
	// response from the backend. The key is a method response header parameter
	// name and the mapped value is an integration response header value, a static
	// value enclosed within a pair of single quotes, or a JSON expression from
	// the integration response body. The mapping key must match the pattern of
	// method.response.header.{name}, where name is a valid and unique header name.
	// The mapped non-static value must match the pattern of integration.response.header.{name}
	// or integration.response.body.{JSON-expression}, where name is a valid and
	// unique response header name and JSON-expression is a valid JSON expression
	// without the $ prefix.
	RequestParameters map[string]string `locationName:"requestParameters" type:"map"`

	// A mapping of identifier keys to templates. The value is an actual template
	// script. The key is typically a SelectionKey which is chosen based on evaluating
	// a selection expression.
	RequestTemplates map[string]string `locationName:"requestTemplates" type:"map"`

	// An expression used to extract information at runtime. See Selection Expressions
	// (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions)
	// for more information.
	TemplateSelectionExpression *string `locationName:"templateSelectionExpression" type:"string"`

	// An integer with a value between [50-29000].
	TimeoutInMillis *int64 `locationName:"timeoutInMillis" min:"50" type:"integer"`
}

// String returns the string representation
func (s CreateIntegrationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateIntegrationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateIntegrationInput"}

	if s.ApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApiId"))
	}
	if len(s.IntegrationType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("IntegrationType"))
	}
	if s.TimeoutInMillis != nil && *s.TimeoutInMillis < 50 {
		invalidParams.Add(aws.NewErrParamMinValue("TimeoutInMillis", 50))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateIntegrationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ConnectionId != nil {
		v := *s.ConnectionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "connectionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.ConnectionType) > 0 {
		v := s.ConnectionType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "connectionType", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if len(s.ContentHandlingStrategy) > 0 {
		v := s.ContentHandlingStrategy

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "contentHandlingStrategy", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.CredentialsArn != nil {
		v := *s.CredentialsArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "credentialsArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.IntegrationMethod != nil {
		v := *s.IntegrationMethod

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "integrationMethod", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.IntegrationType) > 0 {
		v := s.IntegrationType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "integrationType", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.IntegrationUri != nil {
		v := *s.IntegrationUri

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "integrationUri", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.PassthroughBehavior) > 0 {
		v := s.PassthroughBehavior

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "passthroughBehavior", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.RequestParameters != nil {
		v := s.RequestParameters

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "requestParameters", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.RequestTemplates != nil {
		v := s.RequestTemplates

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "requestTemplates", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.TemplateSelectionExpression != nil {
		v := *s.TemplateSelectionExpression

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "templateSelectionExpression", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.TimeoutInMillis != nil {
		v := *s.TimeoutInMillis

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "timeoutInMillis", protocol.Int64Value(v), metadata)
	}
	if s.ApiId != nil {
		v := *s.ApiId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "apiId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/apigatewayv2-2018-11-29/CreateIntegrationResult
type CreateIntegrationOutput struct {
	_ struct{} `type:"structure"`

	// A string with a length between [1-1024].
	ConnectionId *string `locationName:"connectionId" type:"string"`

	// Represents a connection type.
	ConnectionType ConnectionType `locationName:"connectionType" type:"string" enum:"true"`

	// Specifies how to handle response payload content type conversions.
	ContentHandlingStrategy ContentHandlingStrategy `locationName:"contentHandlingStrategy" type:"string" enum:"true"`

	// Represents an Amazon Resource Name (ARN).
	CredentialsArn *string `locationName:"credentialsArn" type:"string"`

	// A string with a length between [0-1024].
	Description *string `locationName:"description" type:"string"`

	// The identifier.
	IntegrationId *string `locationName:"integrationId" type:"string"`

	// A string with a length between [1-64].
	IntegrationMethod *string `locationName:"integrationMethod" type:"string"`

	// An expression used to extract information at runtime. See Selection Expressions
	// (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions)
	// for more information.
	IntegrationResponseSelectionExpression *string `locationName:"integrationResponseSelectionExpression" type:"string"`

	// Represents an API method integration type.
	IntegrationType IntegrationType `locationName:"integrationType" type:"string" enum:"true"`

	// A string representation of a URI with a length between [1-2048].
	IntegrationUri *string `locationName:"integrationUri" type:"string"`

	// Represents passthrough behavior for an integration response.
	PassthroughBehavior PassthroughBehavior `locationName:"passthroughBehavior" type:"string" enum:"true"`

	// A key-value map specifying response parameters that are passed to the method
	// response from the backend. The key is a method response header parameter
	// name and the mapped value is an integration response header value, a static
	// value enclosed within a pair of single quotes, or a JSON expression from
	// the integration response body. The mapping key must match the pattern of
	// method.response.header.{name}, where name is a valid and unique header name.
	// The mapped non-static value must match the pattern of integration.response.header.{name}
	// or integration.response.body.{JSON-expression}, where name is a valid and
	// unique response header name and JSON-expression is a valid JSON expression
	// without the $ prefix.
	RequestParameters map[string]string `locationName:"requestParameters" type:"map"`

	// A mapping of identifier keys to templates. The value is an actual template
	// script. The key is typically a SelectionKey which is chosen based on evaluating
	// a selection expression.
	RequestTemplates map[string]string `locationName:"requestTemplates" type:"map"`

	// An expression used to extract information at runtime. See Selection Expressions
	// (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions)
	// for more information.
	TemplateSelectionExpression *string `locationName:"templateSelectionExpression" type:"string"`

	// An integer with a value between [50-29000].
	TimeoutInMillis *int64 `locationName:"timeoutInMillis" min:"50" type:"integer"`
}

// String returns the string representation
func (s CreateIntegrationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateIntegrationOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ConnectionId != nil {
		v := *s.ConnectionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "connectionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.ConnectionType) > 0 {
		v := s.ConnectionType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "connectionType", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if len(s.ContentHandlingStrategy) > 0 {
		v := s.ContentHandlingStrategy

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "contentHandlingStrategy", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.CredentialsArn != nil {
		v := *s.CredentialsArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "credentialsArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.IntegrationId != nil {
		v := *s.IntegrationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "integrationId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.IntegrationMethod != nil {
		v := *s.IntegrationMethod

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "integrationMethod", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.IntegrationResponseSelectionExpression != nil {
		v := *s.IntegrationResponseSelectionExpression

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "integrationResponseSelectionExpression", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.IntegrationType) > 0 {
		v := s.IntegrationType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "integrationType", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.IntegrationUri != nil {
		v := *s.IntegrationUri

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "integrationUri", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.PassthroughBehavior) > 0 {
		v := s.PassthroughBehavior

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "passthroughBehavior", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.RequestParameters != nil {
		v := s.RequestParameters

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "requestParameters", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.RequestTemplates != nil {
		v := s.RequestTemplates

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "requestTemplates", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.TemplateSelectionExpression != nil {
		v := *s.TemplateSelectionExpression

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "templateSelectionExpression", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.TimeoutInMillis != nil {
		v := *s.TimeoutInMillis

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "timeoutInMillis", protocol.Int64Value(v), metadata)
	}
	return nil
}

const opCreateIntegration = "CreateIntegration"

// CreateIntegrationRequest returns a request value for making API operation for
// AmazonApiGatewayV2.
//
// Creates an Integration.
//
//    // Example sending a request using CreateIntegrationRequest.
//    req := client.CreateIntegrationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/apigatewayv2-2018-11-29/CreateIntegration
func (c *Client) CreateIntegrationRequest(input *CreateIntegrationInput) CreateIntegrationRequest {
	op := &aws.Operation{
		Name:       opCreateIntegration,
		HTTPMethod: "POST",
		HTTPPath:   "/v2/apis/{apiId}/integrations",
	}

	if input == nil {
		input = &CreateIntegrationInput{}
	}

	req := c.newRequest(op, input, &CreateIntegrationOutput{})
	return CreateIntegrationRequest{Request: req, Input: input, Copy: c.CreateIntegrationRequest}
}

// CreateIntegrationRequest is the request type for the
// CreateIntegration API operation.
type CreateIntegrationRequest struct {
	*aws.Request
	Input *CreateIntegrationInput
	Copy  func(*CreateIntegrationInput) CreateIntegrationRequest
}

// Send marshals and sends the CreateIntegration API request.
func (r CreateIntegrationRequest) Send(ctx context.Context) (*CreateIntegrationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateIntegrationResponse{
		CreateIntegrationOutput: r.Request.Data.(*CreateIntegrationOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateIntegrationResponse is the response type for the
// CreateIntegration API operation.
type CreateIntegrationResponse struct {
	*CreateIntegrationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateIntegration request.
func (r *CreateIntegrationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
