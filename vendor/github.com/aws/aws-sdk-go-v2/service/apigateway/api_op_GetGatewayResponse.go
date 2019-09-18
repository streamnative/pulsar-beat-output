// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Gets a GatewayResponse of a specified response type on the given RestApi.
type GetGatewayResponseInput struct {
	_ struct{} `type:"structure"`

	// [Required]
	// The response type of the associated GatewayResponse. Valid values are
	//    * ACCESS_DENIED
	//
	//    * API_CONFIGURATION_ERROR
	//
	//    * AUTHORIZER_FAILURE
	//
	//    * AUTHORIZER_CONFIGURATION_ERROR
	//
	//    * BAD_REQUEST_PARAMETERS
	//
	//    * BAD_REQUEST_BODY
	//
	//    * DEFAULT_4XX
	//
	//    * DEFAULT_5XX
	//
	//    * EXPIRED_TOKEN
	//
	//    * INVALID_SIGNATURE
	//
	//    * INTEGRATION_FAILURE
	//
	//    * INTEGRATION_TIMEOUT
	//
	//    * INVALID_API_KEY
	//
	//    * MISSING_AUTHENTICATION_TOKEN
	//
	//    * QUOTA_EXCEEDED
	//
	//    * REQUEST_TOO_LARGE
	//
	//    * RESOURCE_NOT_FOUND
	//
	//    * THROTTLED
	//
	//    * UNAUTHORIZED
	//
	//    * UNSUPPORTED_MEDIA_TYPE
	//
	// ResponseType is a required field
	ResponseType GatewayResponseType `location:"uri" locationName:"response_type" type:"string" required:"true" enum:"true"`

	// [Required] The string identifier of the associated RestApi.
	//
	// RestApiId is a required field
	RestApiId *string `location:"uri" locationName:"restapi_id" type:"string" required:"true"`
}

// String returns the string representation
func (s GetGatewayResponseInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetGatewayResponseInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetGatewayResponseInput"}
	if len(s.ResponseType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("ResponseType"))
	}

	if s.RestApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RestApiId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetGatewayResponseInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if len(s.ResponseType) > 0 {
		v := s.ResponseType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "response_type", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.RestApiId != nil {
		v := *s.RestApiId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "restapi_id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// A gateway response of a given response type and status code, with optional
// response parameters and mapping templates.
//
// For more information about valid gateway response types, see Gateway Response
// Types Supported by API Gateway (https://docs.aws.amazon.com/apigateway/latest/developerguide/supported-gateway-response-types.html)
//
// Example: Get a Gateway Response of a given response type
//
// Request
//
// This example shows how to get a gateway response of the MISSING_AUTHENTICATION_TOKEN
// type.
//  GET /restapis/o81lxisefl/gatewayresponses/MISSING_AUTHENTICATION_TOKEN HTTP/1.1
//  Host: beta-apigateway.us-east-1.amazonaws.com Content-Type: application/json
//  X-Amz-Date: 20170503T202516Z Authorization: AWS4-HMAC-SHA256 Credential={access-key-id}/20170503/us-east-1/apigateway/aws4_request,
//  SignedHeaders=content-type;host;x-amz-date, Signature=1b52460e3159c1a26cff29093855d50ea141c1c5b937528fecaf60f51129697a
//  Cache-Control: no-cache Postman-Token: 3b2a1ce9-c848-2e26-2e2f-9c2caefbed45
// The response type is specified as a URL path.
//
// Response
//
// The successful operation returns the 200 OK status code and a payload similar
// to the following:
//  { "_links": { "curies": { "href": "http://docs.aws.amazon.com/apigateway/latest/developerguide/restapi-gatewayresponse-{rel}.html",
//  "name": "gatewayresponse", "templated": true }, "self": { "href": "/restapis/o81lxisefl/gatewayresponses/MISSING_AUTHENTICATION_TOKEN"
//  }, "gatewayresponse:delete": { "href": "/restapis/o81lxisefl/gatewayresponses/MISSING_AUTHENTICATION_TOKEN"
//  }, "gatewayresponse:put": { "href": "/restapis/o81lxisefl/gatewayresponses/{response_type}",
//  "templated": true }, "gatewayresponse:update": { "href": "/restapis/o81lxisefl/gatewayresponses/MISSING_AUTHENTICATION_TOKEN"
//  } }, "defaultResponse": false, "responseParameters": { "gatewayresponse.header.x-request-path":
//  "method.request.path.petId", "gatewayresponse.header.Access-Control-Allow-Origin":
//  "'a.b.c'", "gatewayresponse.header.x-request-query": "method.request.querystring.q",
//  "gatewayresponse.header.x-request-header": "method.request.header.Accept"
//  }, "responseTemplates": { "application/json": "{\n \"message\": $context.error.messageString,\n
//  \"type\": \"$context.error.responseType\",\n \"stage\": \"$context.stage\",\n
//  \"resourcePath\": \"$context.resourcePath\",\n \"stageVariables.a\": \"$stageVariables.a\",\n
//  \"statusCode\": \"'404'\"\n}" }, "responseType": "MISSING_AUTHENTICATION_TOKEN",
//  "statusCode": "404" }
//
// Customize Gateway Responses (https://docs.aws.amazon.com/apigateway/latest/developerguide/customize-gateway-responses.html)
type GetGatewayResponseOutput struct {
	_ struct{} `type:"structure"`

	// A Boolean flag to indicate whether this GatewayResponse is the default gateway
	// response (true) or not (false). A default gateway response is one generated
	// by API Gateway without any customization by an API developer.
	DefaultResponse *bool `locationName:"defaultResponse" type:"boolean"`

	// Response parameters (paths, query strings and headers) of the GatewayResponse
	// as a string-to-string map of key-value pairs.
	ResponseParameters map[string]string `locationName:"responseParameters" type:"map"`

	// Response templates of the GatewayResponse as a string-to-string map of key-value
	// pairs.
	ResponseTemplates map[string]string `locationName:"responseTemplates" type:"map"`

	// The response type of the associated GatewayResponse. Valid values are
	//    * ACCESS_DENIED
	//
	//    * API_CONFIGURATION_ERROR
	//
	//    * AUTHORIZER_FAILURE
	//
	//    * AUTHORIZER_CONFIGURATION_ERROR
	//
	//    * BAD_REQUEST_PARAMETERS
	//
	//    * BAD_REQUEST_BODY
	//
	//    * DEFAULT_4XX
	//
	//    * DEFAULT_5XX
	//
	//    * EXPIRED_TOKEN
	//
	//    * INVALID_SIGNATURE
	//
	//    * INTEGRATION_FAILURE
	//
	//    * INTEGRATION_TIMEOUT
	//
	//    * INVALID_API_KEY
	//
	//    * MISSING_AUTHENTICATION_TOKEN
	//
	//    * QUOTA_EXCEEDED
	//
	//    * REQUEST_TOO_LARGE
	//
	//    * RESOURCE_NOT_FOUND
	//
	//    * THROTTLED
	//
	//    * UNAUTHORIZED
	//
	//    * UNSUPPORTED_MEDIA_TYPE
	ResponseType GatewayResponseType `locationName:"responseType" type:"string" enum:"true"`

	// The HTTP status code for this GatewayResponse.
	StatusCode *string `locationName:"statusCode" type:"string"`
}

// String returns the string representation
func (s GetGatewayResponseOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetGatewayResponseOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.DefaultResponse != nil {
		v := *s.DefaultResponse

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "defaultResponse", protocol.BoolValue(v), metadata)
	}
	if s.ResponseParameters != nil {
		v := s.ResponseParameters

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "responseParameters", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.ResponseTemplates != nil {
		v := s.ResponseTemplates

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "responseTemplates", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if len(s.ResponseType) > 0 {
		v := s.ResponseType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "responseType", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.StatusCode != nil {
		v := *s.StatusCode

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "statusCode", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opGetGatewayResponse = "GetGatewayResponse"

// GetGatewayResponseRequest returns a request value for making API operation for
// Amazon API Gateway.
//
// Gets a GatewayResponse of a specified response type on the given RestApi.
//
//    // Example sending a request using GetGatewayResponseRequest.
//    req := client.GetGatewayResponseRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) GetGatewayResponseRequest(input *GetGatewayResponseInput) GetGatewayResponseRequest {
	op := &aws.Operation{
		Name:       opGetGatewayResponse,
		HTTPMethod: "GET",
		HTTPPath:   "/restapis/{restapi_id}/gatewayresponses/{response_type}",
	}

	if input == nil {
		input = &GetGatewayResponseInput{}
	}

	req := c.newRequest(op, input, &GetGatewayResponseOutput{})
	return GetGatewayResponseRequest{Request: req, Input: input, Copy: c.GetGatewayResponseRequest}
}

// GetGatewayResponseRequest is the request type for the
// GetGatewayResponse API operation.
type GetGatewayResponseRequest struct {
	*aws.Request
	Input *GetGatewayResponseInput
	Copy  func(*GetGatewayResponseInput) GetGatewayResponseRequest
}

// Send marshals and sends the GetGatewayResponse API request.
func (r GetGatewayResponseRequest) Send(ctx context.Context) (*GetGatewayResponseResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetGatewayResponseResponse{
		GetGatewayResponseOutput: r.Request.Data.(*GetGatewayResponseOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetGatewayResponseResponse is the response type for the
// GetGatewayResponse API operation.
type GetGatewayResponseResponse struct {
	*GetGatewayResponseOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetGatewayResponse request.
func (r *GetGatewayResponseResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
