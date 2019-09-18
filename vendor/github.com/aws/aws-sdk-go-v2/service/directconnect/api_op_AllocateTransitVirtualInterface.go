// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package directconnect

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/directconnect-2012-10-25/AllocateTransitVirtualInterfaceRequest
type AllocateTransitVirtualInterfaceInput struct {
	_ struct{} `type:"structure"`

	// The ID of the connection on which the transit virtual interface is provisioned.
	//
	// ConnectionId is a required field
	ConnectionId *string `locationName:"connectionId" type:"string" required:"true"`

	// Information about the transit virtual interface.
	//
	// NewTransitVirtualInterfaceAllocation is a required field
	NewTransitVirtualInterfaceAllocation *NewTransitVirtualInterfaceAllocation `locationName:"newTransitVirtualInterfaceAllocation" type:"structure" required:"true"`

	// The ID of the AWS account that owns the transit virtual interface.
	//
	// OwnerAccount is a required field
	OwnerAccount *string `locationName:"ownerAccount" type:"string" required:"true"`
}

// String returns the string representation
func (s AllocateTransitVirtualInterfaceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AllocateTransitVirtualInterfaceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AllocateTransitVirtualInterfaceInput"}

	if s.ConnectionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConnectionId"))
	}

	if s.NewTransitVirtualInterfaceAllocation == nil {
		invalidParams.Add(aws.NewErrParamRequired("NewTransitVirtualInterfaceAllocation"))
	}

	if s.OwnerAccount == nil {
		invalidParams.Add(aws.NewErrParamRequired("OwnerAccount"))
	}
	if s.NewTransitVirtualInterfaceAllocation != nil {
		if err := s.NewTransitVirtualInterfaceAllocation.Validate(); err != nil {
			invalidParams.AddNested("NewTransitVirtualInterfaceAllocation", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/directconnect-2012-10-25/AllocateTransitVirtualInterfaceResult
type AllocateTransitVirtualInterfaceOutput struct {
	_ struct{} `type:"structure"`

	// Information about a virtual interface.
	VirtualInterface *VirtualInterface `locationName:"virtualInterface" type:"structure"`
}

// String returns the string representation
func (s AllocateTransitVirtualInterfaceOutput) String() string {
	return awsutil.Prettify(s)
}

const opAllocateTransitVirtualInterface = "AllocateTransitVirtualInterface"

// AllocateTransitVirtualInterfaceRequest returns a request value for making API operation for
// AWS Direct Connect.
//
// Provisions a transit virtual interface to be owned by the specified AWS account.
// Use this type of interface to connect a transit gateway to your Direct Connect
// gateway.
//
// The owner of a connection provisions a transit virtual interface to be owned
// by the specified AWS account.
//
// After you create a transit virtual interface, it must be confirmed by the
// owner using ConfirmTransitVirtualInterface. Until this step has been completed,
// the transit virtual interface is in the requested state and is not available
// to handle traffic.
//
//    // Example sending a request using AllocateTransitVirtualInterfaceRequest.
//    req := client.AllocateTransitVirtualInterfaceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/directconnect-2012-10-25/AllocateTransitVirtualInterface
func (c *Client) AllocateTransitVirtualInterfaceRequest(input *AllocateTransitVirtualInterfaceInput) AllocateTransitVirtualInterfaceRequest {
	op := &aws.Operation{
		Name:       opAllocateTransitVirtualInterface,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AllocateTransitVirtualInterfaceInput{}
	}

	req := c.newRequest(op, input, &AllocateTransitVirtualInterfaceOutput{})
	return AllocateTransitVirtualInterfaceRequest{Request: req, Input: input, Copy: c.AllocateTransitVirtualInterfaceRequest}
}

// AllocateTransitVirtualInterfaceRequest is the request type for the
// AllocateTransitVirtualInterface API operation.
type AllocateTransitVirtualInterfaceRequest struct {
	*aws.Request
	Input *AllocateTransitVirtualInterfaceInput
	Copy  func(*AllocateTransitVirtualInterfaceInput) AllocateTransitVirtualInterfaceRequest
}

// Send marshals and sends the AllocateTransitVirtualInterface API request.
func (r AllocateTransitVirtualInterfaceRequest) Send(ctx context.Context) (*AllocateTransitVirtualInterfaceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AllocateTransitVirtualInterfaceResponse{
		AllocateTransitVirtualInterfaceOutput: r.Request.Data.(*AllocateTransitVirtualInterfaceOutput),
		response:                              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AllocateTransitVirtualInterfaceResponse is the response type for the
// AllocateTransitVirtualInterface API operation.
type AllocateTransitVirtualInterfaceResponse struct {
	*AllocateTransitVirtualInterfaceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AllocateTransitVirtualInterface request.
func (r *AllocateTransitVirtualInterfaceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
