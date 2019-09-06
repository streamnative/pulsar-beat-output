// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/GetPhoneNumberRequest
type GetPhoneNumberInput struct {
	_ struct{} `type:"structure"`

	// The phone number ID.
	//
	// PhoneNumberId is a required field
	PhoneNumberId *string `location:"uri" locationName:"phoneNumberId" type:"string" required:"true"`
}

// String returns the string representation
func (s GetPhoneNumberInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetPhoneNumberInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetPhoneNumberInput"}

	if s.PhoneNumberId == nil {
		invalidParams.Add(aws.NewErrParamRequired("PhoneNumberId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetPhoneNumberInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.PhoneNumberId != nil {
		v := *s.PhoneNumberId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "phoneNumberId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/GetPhoneNumberResponse
type GetPhoneNumberOutput struct {
	_ struct{} `type:"structure"`

	// The phone number details.
	PhoneNumber *PhoneNumber `type:"structure"`
}

// String returns the string representation
func (s GetPhoneNumberOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetPhoneNumberOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.PhoneNumber != nil {
		v := s.PhoneNumber

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "PhoneNumber", v, metadata)
	}
	return nil
}

const opGetPhoneNumber = "GetPhoneNumber"

// GetPhoneNumberRequest returns a request value for making API operation for
// Amazon Chime.
//
// Retrieves details for the specified phone number ID, such as associations,
// capabilities, and product type.
//
//    // Example sending a request using GetPhoneNumberRequest.
//    req := client.GetPhoneNumberRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/GetPhoneNumber
func (c *Client) GetPhoneNumberRequest(input *GetPhoneNumberInput) GetPhoneNumberRequest {
	op := &aws.Operation{
		Name:       opGetPhoneNumber,
		HTTPMethod: "GET",
		HTTPPath:   "/phone-numbers/{phoneNumberId}",
	}

	if input == nil {
		input = &GetPhoneNumberInput{}
	}

	req := c.newRequest(op, input, &GetPhoneNumberOutput{})
	return GetPhoneNumberRequest{Request: req, Input: input, Copy: c.GetPhoneNumberRequest}
}

// GetPhoneNumberRequest is the request type for the
// GetPhoneNumber API operation.
type GetPhoneNumberRequest struct {
	*aws.Request
	Input *GetPhoneNumberInput
	Copy  func(*GetPhoneNumberInput) GetPhoneNumberRequest
}

// Send marshals and sends the GetPhoneNumber API request.
func (r GetPhoneNumberRequest) Send(ctx context.Context) (*GetPhoneNumberResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetPhoneNumberResponse{
		GetPhoneNumberOutput: r.Request.Data.(*GetPhoneNumberOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetPhoneNumberResponse is the response type for the
// GetPhoneNumber API operation.
type GetPhoneNumberResponse struct {
	*GetPhoneNumberOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetPhoneNumber request.
func (r *GetPhoneNumberResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}