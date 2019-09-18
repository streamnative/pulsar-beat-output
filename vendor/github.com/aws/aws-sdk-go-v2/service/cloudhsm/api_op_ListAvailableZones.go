// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudhsm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the inputs for the ListAvailableZones action.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudhsm-2014-05-30/ListAvailableZonesRequest
type ListAvailableZonesInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s ListAvailableZonesInput) String() string {
	return awsutil.Prettify(s)
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudhsm-2014-05-30/ListAvailableZonesResponse
type ListAvailableZonesOutput struct {
	_ struct{} `type:"structure"`

	// The list of Availability Zones that have available AWS CloudHSM capacity.
	AZList []string `type:"list"`
}

// String returns the string representation
func (s ListAvailableZonesOutput) String() string {
	return awsutil.Prettify(s)
}

const opListAvailableZones = "ListAvailableZones"

// ListAvailableZonesRequest returns a request value for making API operation for
// Amazon CloudHSM.
//
// This is documentation for AWS CloudHSM Classic. For more information, see
// AWS CloudHSM Classic FAQs (http://aws.amazon.com/cloudhsm/faqs-classic/),
// the AWS CloudHSM Classic User Guide (http://docs.aws.amazon.com/cloudhsm/classic/userguide/),
// and the AWS CloudHSM Classic API Reference (http://docs.aws.amazon.com/cloudhsm/classic/APIReference/).
//
// For information about the current version of AWS CloudHSM, see AWS CloudHSM
// (http://aws.amazon.com/cloudhsm/), the AWS CloudHSM User Guide (http://docs.aws.amazon.com/cloudhsm/latest/userguide/),
// and the AWS CloudHSM API Reference (http://docs.aws.amazon.com/cloudhsm/latest/APIReference/).
//
// Lists the Availability Zones that have available AWS CloudHSM capacity.
//
//    // Example sending a request using ListAvailableZonesRequest.
//    req := client.ListAvailableZonesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudhsm-2014-05-30/ListAvailableZones
func (c *Client) ListAvailableZonesRequest(input *ListAvailableZonesInput) ListAvailableZonesRequest {
	op := &aws.Operation{
		Name:       opListAvailableZones,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListAvailableZonesInput{}
	}

	req := c.newRequest(op, input, &ListAvailableZonesOutput{})
	return ListAvailableZonesRequest{Request: req, Input: input, Copy: c.ListAvailableZonesRequest}
}

// ListAvailableZonesRequest is the request type for the
// ListAvailableZones API operation.
type ListAvailableZonesRequest struct {
	*aws.Request
	Input *ListAvailableZonesInput
	Copy  func(*ListAvailableZonesInput) ListAvailableZonesRequest
}

// Send marshals and sends the ListAvailableZones API request.
func (r ListAvailableZonesRequest) Send(ctx context.Context) (*ListAvailableZonesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListAvailableZonesResponse{
		ListAvailableZonesOutput: r.Request.Data.(*ListAvailableZonesOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListAvailableZonesResponse is the response type for the
// ListAvailableZones API operation.
type ListAvailableZonesResponse struct {
	*ListAvailableZonesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListAvailableZones request.
func (r *ListAvailableZonesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
