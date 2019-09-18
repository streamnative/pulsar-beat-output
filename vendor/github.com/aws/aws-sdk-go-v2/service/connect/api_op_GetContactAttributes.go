// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package connect

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/connect-2017-08-08/GetContactAttributesRequest
type GetContactAttributesInput struct {
	_ struct{} `type:"structure"`

	// The ID for the initial contact in Amazon Connect associated with the attributes
	// to update.
	//
	// InitialContactId is a required field
	InitialContactId *string `location:"uri" locationName:"InitialContactId" min:"1" type:"string" required:"true"`

	// The instance ID for the instance from which to retrieve contact attributes.
	//
	// InstanceId is a required field
	InstanceId *string `location:"uri" locationName:"InstanceId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetContactAttributesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetContactAttributesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetContactAttributesInput"}

	if s.InitialContactId == nil {
		invalidParams.Add(aws.NewErrParamRequired("InitialContactId"))
	}
	if s.InitialContactId != nil && len(*s.InitialContactId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("InitialContactId", 1))
	}

	if s.InstanceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceId"))
	}
	if s.InstanceId != nil && len(*s.InstanceId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("InstanceId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetContactAttributesInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.InitialContactId != nil {
		v := *s.InitialContactId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "InitialContactId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.InstanceId != nil {
		v := *s.InstanceId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "InstanceId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/connect-2017-08-08/GetContactAttributesResponse
type GetContactAttributesOutput struct {
	_ struct{} `type:"structure"`

	// The attributes to update.
	Attributes map[string]string `type:"map"`
}

// String returns the string representation
func (s GetContactAttributesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetContactAttributesOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Attributes != nil {
		v := s.Attributes

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "Attributes", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	return nil
}

const opGetContactAttributes = "GetContactAttributes"

// GetContactAttributesRequest returns a request value for making API operation for
// Amazon Connect Service.
//
// Retrieves the contact attributes associated with a contact.
//
//    // Example sending a request using GetContactAttributesRequest.
//    req := client.GetContactAttributesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/connect-2017-08-08/GetContactAttributes
func (c *Client) GetContactAttributesRequest(input *GetContactAttributesInput) GetContactAttributesRequest {
	op := &aws.Operation{
		Name:       opGetContactAttributes,
		HTTPMethod: "GET",
		HTTPPath:   "/contact/attributes/{InstanceId}/{InitialContactId}",
	}

	if input == nil {
		input = &GetContactAttributesInput{}
	}

	req := c.newRequest(op, input, &GetContactAttributesOutput{})
	return GetContactAttributesRequest{Request: req, Input: input, Copy: c.GetContactAttributesRequest}
}

// GetContactAttributesRequest is the request type for the
// GetContactAttributes API operation.
type GetContactAttributesRequest struct {
	*aws.Request
	Input *GetContactAttributesInput
	Copy  func(*GetContactAttributesInput) GetContactAttributesRequest
}

// Send marshals and sends the GetContactAttributes API request.
func (r GetContactAttributesRequest) Send(ctx context.Context) (*GetContactAttributesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetContactAttributesResponse{
		GetContactAttributesOutput: r.Request.Data.(*GetContactAttributesOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetContactAttributesResponse is the response type for the
// GetContactAttributes API operation.
type GetContactAttributesResponse struct {
	*GetContactAttributesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetContactAttributes request.
func (r *GetContactAttributesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
