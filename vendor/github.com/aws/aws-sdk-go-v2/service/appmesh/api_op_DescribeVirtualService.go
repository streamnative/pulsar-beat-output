// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appmesh

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/appmesh-2019-01-25/DescribeVirtualServiceInput
type DescribeVirtualServiceInput struct {
	_ struct{} `type:"structure"`

	// MeshName is a required field
	MeshName *string `location:"uri" locationName:"meshName" min:"1" type:"string" required:"true"`

	// VirtualServiceName is a required field
	VirtualServiceName *string `location:"uri" locationName:"virtualServiceName" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeVirtualServiceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeVirtualServiceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeVirtualServiceInput"}

	if s.MeshName == nil {
		invalidParams.Add(aws.NewErrParamRequired("MeshName"))
	}
	if s.MeshName != nil && len(*s.MeshName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("MeshName", 1))
	}

	if s.VirtualServiceName == nil {
		invalidParams.Add(aws.NewErrParamRequired("VirtualServiceName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeVirtualServiceInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.MeshName != nil {
		v := *s.MeshName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "meshName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VirtualServiceName != nil {
		v := *s.VirtualServiceName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "virtualServiceName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/appmesh-2019-01-25/DescribeVirtualServiceOutput
type DescribeVirtualServiceOutput struct {
	_ struct{} `type:"structure" payload:"VirtualService"`

	// An object representing a virtual service returned by a describe operation.
	//
	// VirtualService is a required field
	VirtualService *VirtualServiceData `locationName:"virtualService" type:"structure" required:"true"`
}

// String returns the string representation
func (s DescribeVirtualServiceOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeVirtualServiceOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.VirtualService != nil {
		v := s.VirtualService

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "virtualService", v, metadata)
	}
	return nil
}

const opDescribeVirtualService = "DescribeVirtualService"

// DescribeVirtualServiceRequest returns a request value for making API operation for
// AWS App Mesh.
//
// Describes an existing virtual service.
//
//    // Example sending a request using DescribeVirtualServiceRequest.
//    req := client.DescribeVirtualServiceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appmesh-2019-01-25/DescribeVirtualService
func (c *Client) DescribeVirtualServiceRequest(input *DescribeVirtualServiceInput) DescribeVirtualServiceRequest {
	op := &aws.Operation{
		Name:       opDescribeVirtualService,
		HTTPMethod: "GET",
		HTTPPath:   "/v20190125/meshes/{meshName}/virtualServices/{virtualServiceName}",
	}

	if input == nil {
		input = &DescribeVirtualServiceInput{}
	}

	req := c.newRequest(op, input, &DescribeVirtualServiceOutput{})
	return DescribeVirtualServiceRequest{Request: req, Input: input, Copy: c.DescribeVirtualServiceRequest}
}

// DescribeVirtualServiceRequest is the request type for the
// DescribeVirtualService API operation.
type DescribeVirtualServiceRequest struct {
	*aws.Request
	Input *DescribeVirtualServiceInput
	Copy  func(*DescribeVirtualServiceInput) DescribeVirtualServiceRequest
}

// Send marshals and sends the DescribeVirtualService API request.
func (r DescribeVirtualServiceRequest) Send(ctx context.Context) (*DescribeVirtualServiceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeVirtualServiceResponse{
		DescribeVirtualServiceOutput: r.Request.Data.(*DescribeVirtualServiceOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeVirtualServiceResponse is the response type for the
// DescribeVirtualService API operation.
type DescribeVirtualServiceResponse struct {
	*DescribeVirtualServiceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeVirtualService request.
func (r *DescribeVirtualServiceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
