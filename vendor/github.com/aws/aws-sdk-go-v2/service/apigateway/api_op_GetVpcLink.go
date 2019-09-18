// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Gets a specified VPC link under the caller's account in a region.
type GetVpcLinkInput struct {
	_ struct{} `type:"structure"`

	// [Required] The identifier of the VpcLink. It is used in an Integration to
	// reference this VpcLink.
	//
	// VpcLinkId is a required field
	VpcLinkId *string `location:"uri" locationName:"vpclink_id" type:"string" required:"true"`
}

// String returns the string representation
func (s GetVpcLinkInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetVpcLinkInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetVpcLinkInput"}

	if s.VpcLinkId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VpcLinkId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetVpcLinkInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.VpcLinkId != nil {
		v := *s.VpcLinkId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "vpclink_id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// A API Gateway VPC link for a RestApi to access resources in an Amazon Virtual
// Private Cloud (VPC).
//
// To enable access to a resource in an Amazon Virtual Private Cloud through
// Amazon API Gateway, you, as an API developer, create a VpcLink resource targeted
// for one or more network load balancers of the VPC and then integrate an API
// method with a private integration that uses the VpcLink. The private integration
// has an integration type of HTTP or HTTP_PROXY and has a connection type of
// VPC_LINK. The integration uses the connectionId property to identify the
// VpcLink used.
type GetVpcLinkOutput struct {
	_ struct{} `type:"structure"`

	// The description of the VPC link.
	Description *string `locationName:"description" type:"string"`

	// The identifier of the VpcLink. It is used in an Integration to reference
	// this VpcLink.
	Id *string `locationName:"id" type:"string"`

	// The name used to label and identify the VPC link.
	Name *string `locationName:"name" type:"string"`

	// The status of the VPC link. The valid values are AVAILABLE, PENDING, DELETING,
	// or FAILED. Deploying an API will wait if the status is PENDING and will fail
	// if the status is DELETING.
	Status VpcLinkStatus `locationName:"status" type:"string" enum:"true"`

	// A description about the VPC link status.
	StatusMessage *string `locationName:"statusMessage" type:"string"`

	// The collection of tags. Each tag element is associated with a given resource.
	Tags map[string]string `locationName:"tags" type:"map"`

	// The ARNs of network load balancers of the VPC targeted by the VPC link. The
	// network load balancers must be owned by the same AWS account of the API owner.
	TargetArns []string `locationName:"targetArns" type:"list"`
}

// String returns the string representation
func (s GetVpcLinkOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetVpcLinkOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Status) > 0 {
		v := s.Status

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "status", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.StatusMessage != nil {
		v := *s.StatusMessage

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "statusMessage", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
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
	if s.TargetArns != nil {
		v := s.TargetArns

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "targetArns", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	return nil
}

const opGetVpcLink = "GetVpcLink"

// GetVpcLinkRequest returns a request value for making API operation for
// Amazon API Gateway.
//
// Gets a specified VPC link under the caller's account in a region.
//
//    // Example sending a request using GetVpcLinkRequest.
//    req := client.GetVpcLinkRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) GetVpcLinkRequest(input *GetVpcLinkInput) GetVpcLinkRequest {
	op := &aws.Operation{
		Name:       opGetVpcLink,
		HTTPMethod: "GET",
		HTTPPath:   "/vpclinks/{vpclink_id}",
	}

	if input == nil {
		input = &GetVpcLinkInput{}
	}

	req := c.newRequest(op, input, &GetVpcLinkOutput{})
	return GetVpcLinkRequest{Request: req, Input: input, Copy: c.GetVpcLinkRequest}
}

// GetVpcLinkRequest is the request type for the
// GetVpcLink API operation.
type GetVpcLinkRequest struct {
	*aws.Request
	Input *GetVpcLinkInput
	Copy  func(*GetVpcLinkInput) GetVpcLinkRequest
}

// Send marshals and sends the GetVpcLink API request.
func (r GetVpcLinkRequest) Send(ctx context.Context) (*GetVpcLinkResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetVpcLinkResponse{
		GetVpcLinkOutput: r.Request.Data.(*GetVpcLinkOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetVpcLinkResponse is the response type for the
// GetVpcLink API operation.
type GetVpcLinkResponse struct {
	*GetVpcLinkOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetVpcLink request.
func (r *GetVpcLinkResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
