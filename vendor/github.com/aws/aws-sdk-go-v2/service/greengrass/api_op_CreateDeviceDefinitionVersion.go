// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package greengrass

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/CreateDeviceDefinitionVersionRequest
type CreateDeviceDefinitionVersionInput struct {
	_ struct{} `type:"structure"`

	AmznClientToken *string `location:"header" locationName:"X-Amzn-Client-Token" type:"string"`

	// DeviceDefinitionId is a required field
	DeviceDefinitionId *string `location:"uri" locationName:"DeviceDefinitionId" type:"string" required:"true"`

	Devices []Device `type:"list"`
}

// String returns the string representation
func (s CreateDeviceDefinitionVersionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDeviceDefinitionVersionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateDeviceDefinitionVersionInput"}

	if s.DeviceDefinitionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DeviceDefinitionId"))
	}
	if s.Devices != nil {
		for i, v := range s.Devices {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Devices", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateDeviceDefinitionVersionInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Devices != nil {
		v := s.Devices

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Devices", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.AmznClientToken != nil {
		v := *s.AmznClientToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "X-Amzn-Client-Token", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DeviceDefinitionId != nil {
		v := *s.DeviceDefinitionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "DeviceDefinitionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/CreateDeviceDefinitionVersionResponse
type CreateDeviceDefinitionVersionOutput struct {
	_ struct{} `type:"structure"`

	Arn *string `type:"string"`

	CreationTimestamp *string `type:"string"`

	Id *string `type:"string"`

	Version *string `type:"string"`
}

// String returns the string representation
func (s CreateDeviceDefinitionVersionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateDeviceDefinitionVersionOutput) MarshalFields(e protocol.FieldEncoder) error {
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
	if s.Version != nil {
		v := *s.Version

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Version", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opCreateDeviceDefinitionVersion = "CreateDeviceDefinitionVersion"

// CreateDeviceDefinitionVersionRequest returns a request value for making API operation for
// AWS Greengrass.
//
// Creates a version of a device definition that has already been defined.
//
//    // Example sending a request using CreateDeviceDefinitionVersionRequest.
//    req := client.CreateDeviceDefinitionVersionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/CreateDeviceDefinitionVersion
func (c *Client) CreateDeviceDefinitionVersionRequest(input *CreateDeviceDefinitionVersionInput) CreateDeviceDefinitionVersionRequest {
	op := &aws.Operation{
		Name:       opCreateDeviceDefinitionVersion,
		HTTPMethod: "POST",
		HTTPPath:   "/greengrass/definition/devices/{DeviceDefinitionId}/versions",
	}

	if input == nil {
		input = &CreateDeviceDefinitionVersionInput{}
	}

	req := c.newRequest(op, input, &CreateDeviceDefinitionVersionOutput{})
	return CreateDeviceDefinitionVersionRequest{Request: req, Input: input, Copy: c.CreateDeviceDefinitionVersionRequest}
}

// CreateDeviceDefinitionVersionRequest is the request type for the
// CreateDeviceDefinitionVersion API operation.
type CreateDeviceDefinitionVersionRequest struct {
	*aws.Request
	Input *CreateDeviceDefinitionVersionInput
	Copy  func(*CreateDeviceDefinitionVersionInput) CreateDeviceDefinitionVersionRequest
}

// Send marshals and sends the CreateDeviceDefinitionVersion API request.
func (r CreateDeviceDefinitionVersionRequest) Send(ctx context.Context) (*CreateDeviceDefinitionVersionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateDeviceDefinitionVersionResponse{
		CreateDeviceDefinitionVersionOutput: r.Request.Data.(*CreateDeviceDefinitionVersionOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateDeviceDefinitionVersionResponse is the response type for the
// CreateDeviceDefinitionVersion API operation.
type CreateDeviceDefinitionVersionResponse struct {
	*CreateDeviceDefinitionVersionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateDeviceDefinitionVersion request.
func (r *CreateDeviceDefinitionVersionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
