// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CreateThingGroupInput struct {
	_ struct{} `type:"structure"`

	// The name of the parent thing group.
	ParentGroupName *string `locationName:"parentGroupName" min:"1" type:"string"`

	// Metadata which can be used to manage the thing group.
	Tags []Tag `locationName:"tags" type:"list"`

	// The thing group name to create.
	//
	// ThingGroupName is a required field
	ThingGroupName *string `location:"uri" locationName:"thingGroupName" min:"1" type:"string" required:"true"`

	// The thing group properties.
	ThingGroupProperties *ThingGroupProperties `locationName:"thingGroupProperties" type:"structure"`
}

// String returns the string representation
func (s CreateThingGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateThingGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateThingGroupInput"}
	if s.ParentGroupName != nil && len(*s.ParentGroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ParentGroupName", 1))
	}

	if s.ThingGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ThingGroupName"))
	}
	if s.ThingGroupName != nil && len(*s.ThingGroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ThingGroupName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateThingGroupInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ParentGroupName != nil {
		v := *s.ParentGroupName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "parentGroupName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "tags", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.ThingGroupProperties != nil {
		v := s.ThingGroupProperties

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "thingGroupProperties", v, metadata)
	}
	if s.ThingGroupName != nil {
		v := *s.ThingGroupName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "thingGroupName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type CreateThingGroupOutput struct {
	_ struct{} `type:"structure"`

	// The thing group ARN.
	ThingGroupArn *string `locationName:"thingGroupArn" type:"string"`

	// The thing group ID.
	ThingGroupId *string `locationName:"thingGroupId" min:"1" type:"string"`

	// The thing group name.
	ThingGroupName *string `locationName:"thingGroupName" min:"1" type:"string"`
}

// String returns the string representation
func (s CreateThingGroupOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateThingGroupOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ThingGroupArn != nil {
		v := *s.ThingGroupArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "thingGroupArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ThingGroupId != nil {
		v := *s.ThingGroupId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "thingGroupId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ThingGroupName != nil {
		v := *s.ThingGroupName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "thingGroupName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opCreateThingGroup = "CreateThingGroup"

// CreateThingGroupRequest returns a request value for making API operation for
// AWS IoT.
//
// Create a thing group.
//
// This is a control plane operation. See Authorization (https://docs.aws.amazon.com/iot/latest/developerguide/authorization.html)
// for information about authorizing control plane actions.
//
//    // Example sending a request using CreateThingGroupRequest.
//    req := client.CreateThingGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) CreateThingGroupRequest(input *CreateThingGroupInput) CreateThingGroupRequest {
	op := &aws.Operation{
		Name:       opCreateThingGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/thing-groups/{thingGroupName}",
	}

	if input == nil {
		input = &CreateThingGroupInput{}
	}

	req := c.newRequest(op, input, &CreateThingGroupOutput{})
	return CreateThingGroupRequest{Request: req, Input: input, Copy: c.CreateThingGroupRequest}
}

// CreateThingGroupRequest is the request type for the
// CreateThingGroup API operation.
type CreateThingGroupRequest struct {
	*aws.Request
	Input *CreateThingGroupInput
	Copy  func(*CreateThingGroupInput) CreateThingGroupRequest
}

// Send marshals and sends the CreateThingGroup API request.
func (r CreateThingGroupRequest) Send(ctx context.Context) (*CreateThingGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateThingGroupResponse{
		CreateThingGroupOutput: r.Request.Data.(*CreateThingGroupOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateThingGroupResponse is the response type for the
// CreateThingGroup API operation.
type CreateThingGroupResponse struct {
	*CreateThingGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateThingGroup request.
func (r *CreateThingGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
