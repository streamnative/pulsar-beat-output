// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appsync

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/appsync-2017-07-25/CreateResolverRequest
type CreateResolverInput struct {
	_ struct{} `type:"structure"`

	// The ID for the GraphQL API for which the resolver is being created.
	//
	// ApiId is a required field
	ApiId *string `location:"uri" locationName:"apiId" type:"string" required:"true"`

	// The name of the data source for which the resolver is being created.
	DataSourceName *string `locationName:"dataSourceName" type:"string"`

	// The name of the field to attach the resolver to.
	//
	// FieldName is a required field
	FieldName *string `locationName:"fieldName" type:"string" required:"true"`

	// The resolver type.
	//
	//    * UNIT: A UNIT resolver type. A UNIT resolver is the default resolver
	//    type. A UNIT resolver enables you to execute a GraphQL query against a
	//    single data source.
	//
	//    * PIPELINE: A PIPELINE resolver type. A PIPELINE resolver enables you
	//    to execute a series of Function in a serial manner. You can use a pipeline
	//    resolver to execute a GraphQL query against multiple data sources.
	Kind ResolverKind `locationName:"kind" type:"string" enum:"true"`

	// The PipelineConfig.
	PipelineConfig *PipelineConfig `locationName:"pipelineConfig" type:"structure"`

	// The mapping template to be used for requests.
	//
	// A resolver uses a request mapping template to convert a GraphQL expression
	// into a format that a data source can understand. Mapping templates are written
	// in Apache Velocity Template Language (VTL).
	//
	// RequestMappingTemplate is a required field
	RequestMappingTemplate *string `locationName:"requestMappingTemplate" min:"1" type:"string" required:"true"`

	// The mapping template to be used for responses from the data source.
	ResponseMappingTemplate *string `locationName:"responseMappingTemplate" min:"1" type:"string"`

	// The name of the Type.
	//
	// TypeName is a required field
	TypeName *string `location:"uri" locationName:"typeName" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateResolverInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateResolverInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateResolverInput"}

	if s.ApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApiId"))
	}

	if s.FieldName == nil {
		invalidParams.Add(aws.NewErrParamRequired("FieldName"))
	}

	if s.RequestMappingTemplate == nil {
		invalidParams.Add(aws.NewErrParamRequired("RequestMappingTemplate"))
	}
	if s.RequestMappingTemplate != nil && len(*s.RequestMappingTemplate) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RequestMappingTemplate", 1))
	}
	if s.ResponseMappingTemplate != nil && len(*s.ResponseMappingTemplate) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResponseMappingTemplate", 1))
	}

	if s.TypeName == nil {
		invalidParams.Add(aws.NewErrParamRequired("TypeName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateResolverInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.DataSourceName != nil {
		v := *s.DataSourceName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "dataSourceName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.FieldName != nil {
		v := *s.FieldName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "fieldName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Kind) > 0 {
		v := s.Kind

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "kind", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.PipelineConfig != nil {
		v := s.PipelineConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "pipelineConfig", v, metadata)
	}
	if s.RequestMappingTemplate != nil {
		v := *s.RequestMappingTemplate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "requestMappingTemplate", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ResponseMappingTemplate != nil {
		v := *s.ResponseMappingTemplate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "responseMappingTemplate", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ApiId != nil {
		v := *s.ApiId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "apiId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.TypeName != nil {
		v := *s.TypeName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "typeName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/appsync-2017-07-25/CreateResolverResponse
type CreateResolverOutput struct {
	_ struct{} `type:"structure"`

	// The Resolver object.
	Resolver *Resolver `locationName:"resolver" type:"structure"`
}

// String returns the string representation
func (s CreateResolverOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateResolverOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Resolver != nil {
		v := s.Resolver

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "resolver", v, metadata)
	}
	return nil
}

const opCreateResolver = "CreateResolver"

// CreateResolverRequest returns a request value for making API operation for
// AWS AppSync.
//
// Creates a Resolver object.
//
// A resolver converts incoming requests into a format that a data source can
// understand and converts the data source's responses into GraphQL.
//
//    // Example sending a request using CreateResolverRequest.
//    req := client.CreateResolverRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appsync-2017-07-25/CreateResolver
func (c *Client) CreateResolverRequest(input *CreateResolverInput) CreateResolverRequest {
	op := &aws.Operation{
		Name:       opCreateResolver,
		HTTPMethod: "POST",
		HTTPPath:   "/v1/apis/{apiId}/types/{typeName}/resolvers",
	}

	if input == nil {
		input = &CreateResolverInput{}
	}

	req := c.newRequest(op, input, &CreateResolverOutput{})
	return CreateResolverRequest{Request: req, Input: input, Copy: c.CreateResolverRequest}
}

// CreateResolverRequest is the request type for the
// CreateResolver API operation.
type CreateResolverRequest struct {
	*aws.Request
	Input *CreateResolverInput
	Copy  func(*CreateResolverInput) CreateResolverRequest
}

// Send marshals and sends the CreateResolver API request.
func (r CreateResolverRequest) Send(ctx context.Context) (*CreateResolverResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateResolverResponse{
		CreateResolverOutput: r.Request.Data.(*CreateResolverOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateResolverResponse is the response type for the
// CreateResolver API operation.
type CreateResolverResponse struct {
	*CreateResolverOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateResolver request.
func (r *CreateResolverResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
