// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigatewayv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/apigatewayv2-2018-11-29/GetIntegrationRequest
type GetIntegrationInput struct {
	_ struct{} `type:"structure"`

	// ApiId is a required field
	ApiId *string `location:"uri" locationName:"apiId" type:"string" required:"true"`

	// IntegrationId is a required field
	IntegrationId *string `location:"uri" locationName:"integrationId" type:"string" required:"true"`
}

// String returns the string representation
func (s GetIntegrationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetIntegrationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetIntegrationInput"}

	if s.ApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApiId"))
	}

	if s.IntegrationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("IntegrationId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetIntegrationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ApiId != nil {
		v := *s.ApiId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "apiId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.IntegrationId != nil {
		v := *s.IntegrationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "integrationId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/apigatewayv2-2018-11-29/GetIntegrationResult
type GetIntegrationOutput struct {
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
func (s GetIntegrationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetIntegrationOutput) MarshalFields(e protocol.FieldEncoder) error {
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

const opGetIntegration = "GetIntegration"

// GetIntegrationRequest returns a request value for making API operation for
// AmazonApiGatewayV2.
//
// Gets an Integration.
//
//    // Example sending a request using GetIntegrationRequest.
//    req := client.GetIntegrationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/apigatewayv2-2018-11-29/GetIntegration
func (c *Client) GetIntegrationRequest(input *GetIntegrationInput) GetIntegrationRequest {
	op := &aws.Operation{
		Name:       opGetIntegration,
		HTTPMethod: "GET",
		HTTPPath:   "/v2/apis/{apiId}/integrations/{integrationId}",
	}

	if input == nil {
		input = &GetIntegrationInput{}
	}

	req := c.newRequest(op, input, &GetIntegrationOutput{})
	return GetIntegrationRequest{Request: req, Input: input, Copy: c.GetIntegrationRequest}
}

// GetIntegrationRequest is the request type for the
// GetIntegration API operation.
type GetIntegrationRequest struct {
	*aws.Request
	Input *GetIntegrationInput
	Copy  func(*GetIntegrationInput) GetIntegrationRequest
}

// Send marshals and sends the GetIntegration API request.
func (r GetIntegrationRequest) Send(ctx context.Context) (*GetIntegrationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetIntegrationResponse{
		GetIntegrationOutput: r.Request.Data.(*GetIntegrationOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetIntegrationResponse is the response type for the
// GetIntegration API operation.
type GetIntegrationResponse struct {
	*GetIntegrationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetIntegration request.
func (r *GetIntegrationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
