// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpointemail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// A request to enable or disable the automatic IP address warm-up feature.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-email-2018-07-26/PutAccountDedicatedIpWarmupAttributesRequest
type PutAccountDedicatedIpWarmupAttributesInput struct {
	_ struct{} `type:"structure"`

	// Enables or disables the automatic warm-up feature for dedicated IP addresses
	// that are associated with your Amazon Pinpoint account in the current AWS
	// Region. Set to true to enable the automatic warm-up feature, or set to false
	// to disable it.
	AutoWarmupEnabled *bool `type:"boolean"`
}

// String returns the string representation
func (s PutAccountDedicatedIpWarmupAttributesInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutAccountDedicatedIpWarmupAttributesInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AutoWarmupEnabled != nil {
		v := *s.AutoWarmupEnabled

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "AutoWarmupEnabled", protocol.BoolValue(v), metadata)
	}
	return nil
}

// An HTTP 200 response if the request succeeds, or an error message if the
// request fails.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-email-2018-07-26/PutAccountDedicatedIpWarmupAttributesResponse
type PutAccountDedicatedIpWarmupAttributesOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutAccountDedicatedIpWarmupAttributesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutAccountDedicatedIpWarmupAttributesOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opPutAccountDedicatedIpWarmupAttributes = "PutAccountDedicatedIpWarmupAttributes"

// PutAccountDedicatedIpWarmupAttributesRequest returns a request value for making API operation for
// Amazon Pinpoint Email Service.
//
// Enable or disable the automatic warm-up feature for dedicated IP addresses.
//
//    // Example sending a request using PutAccountDedicatedIpWarmupAttributesRequest.
//    req := client.PutAccountDedicatedIpWarmupAttributesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-email-2018-07-26/PutAccountDedicatedIpWarmupAttributes
func (c *Client) PutAccountDedicatedIpWarmupAttributesRequest(input *PutAccountDedicatedIpWarmupAttributesInput) PutAccountDedicatedIpWarmupAttributesRequest {
	op := &aws.Operation{
		Name:       opPutAccountDedicatedIpWarmupAttributes,
		HTTPMethod: "PUT",
		HTTPPath:   "/v1/email/account/dedicated-ips/warmup",
	}

	if input == nil {
		input = &PutAccountDedicatedIpWarmupAttributesInput{}
	}

	req := c.newRequest(op, input, &PutAccountDedicatedIpWarmupAttributesOutput{})
	return PutAccountDedicatedIpWarmupAttributesRequest{Request: req, Input: input, Copy: c.PutAccountDedicatedIpWarmupAttributesRequest}
}

// PutAccountDedicatedIpWarmupAttributesRequest is the request type for the
// PutAccountDedicatedIpWarmupAttributes API operation.
type PutAccountDedicatedIpWarmupAttributesRequest struct {
	*aws.Request
	Input *PutAccountDedicatedIpWarmupAttributesInput
	Copy  func(*PutAccountDedicatedIpWarmupAttributesInput) PutAccountDedicatedIpWarmupAttributesRequest
}

// Send marshals and sends the PutAccountDedicatedIpWarmupAttributes API request.
func (r PutAccountDedicatedIpWarmupAttributesRequest) Send(ctx context.Context) (*PutAccountDedicatedIpWarmupAttributesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutAccountDedicatedIpWarmupAttributesResponse{
		PutAccountDedicatedIpWarmupAttributesOutput: r.Request.Data.(*PutAccountDedicatedIpWarmupAttributesOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutAccountDedicatedIpWarmupAttributesResponse is the response type for the
// PutAccountDedicatedIpWarmupAttributes API operation.
type PutAccountDedicatedIpWarmupAttributesResponse struct {
	*PutAccountDedicatedIpWarmupAttributesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutAccountDedicatedIpWarmupAttributes request.
func (r *PutAccountDedicatedIpWarmupAttributesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
