// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package managedblockchain

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/managedblockchain-2018-09-24/GetNetworkInput
type GetNetworkInput struct {
	_ struct{} `type:"structure"`

	// The unique identifier of the network to get information about.
	//
	// NetworkId is a required field
	NetworkId *string `location:"uri" locationName:"networkId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetNetworkInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetNetworkInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetNetworkInput"}

	if s.NetworkId == nil {
		invalidParams.Add(aws.NewErrParamRequired("NetworkId"))
	}
	if s.NetworkId != nil && len(*s.NetworkId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NetworkId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetNetworkInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.NetworkId != nil {
		v := *s.NetworkId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "networkId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/managedblockchain-2018-09-24/GetNetworkOutput
type GetNetworkOutput struct {
	_ struct{} `type:"structure"`

	// An object containing network configuration parameters.
	Network *Network `type:"structure"`
}

// String returns the string representation
func (s GetNetworkOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetNetworkOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Network != nil {
		v := s.Network

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Network", v, metadata)
	}
	return nil
}

const opGetNetwork = "GetNetwork"

// GetNetworkRequest returns a request value for making API operation for
// Amazon Managed Blockchain.
//
// Returns detailed information about a network.
//
//    // Example sending a request using GetNetworkRequest.
//    req := client.GetNetworkRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/managedblockchain-2018-09-24/GetNetwork
func (c *Client) GetNetworkRequest(input *GetNetworkInput) GetNetworkRequest {
	op := &aws.Operation{
		Name:       opGetNetwork,
		HTTPMethod: "GET",
		HTTPPath:   "/networks/{networkId}",
	}

	if input == nil {
		input = &GetNetworkInput{}
	}

	req := c.newRequest(op, input, &GetNetworkOutput{})
	return GetNetworkRequest{Request: req, Input: input, Copy: c.GetNetworkRequest}
}

// GetNetworkRequest is the request type for the
// GetNetwork API operation.
type GetNetworkRequest struct {
	*aws.Request
	Input *GetNetworkInput
	Copy  func(*GetNetworkInput) GetNetworkRequest
}

// Send marshals and sends the GetNetwork API request.
func (r GetNetworkRequest) Send(ctx context.Context) (*GetNetworkResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetNetworkResponse{
		GetNetworkOutput: r.Request.Data.(*GetNetworkOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetNetworkResponse is the response type for the
// GetNetwork API operation.
type GetNetworkResponse struct {
	*GetNetworkOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetNetwork request.
func (r *GetNetworkResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
