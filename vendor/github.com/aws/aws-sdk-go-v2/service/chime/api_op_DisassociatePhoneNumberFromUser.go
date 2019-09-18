// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/DisassociatePhoneNumberFromUserRequest
type DisassociatePhoneNumberFromUserInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Chime account ID.
	//
	// AccountId is a required field
	AccountId *string `location:"uri" locationName:"accountId" type:"string" required:"true"`

	// The user ID.
	//
	// UserId is a required field
	UserId *string `location:"uri" locationName:"userId" type:"string" required:"true"`
}

// String returns the string representation
func (s DisassociatePhoneNumberFromUserInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DisassociatePhoneNumberFromUserInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DisassociatePhoneNumberFromUserInput"}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}

	if s.UserId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DisassociatePhoneNumberFromUserInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AccountId != nil {
		v := *s.AccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "accountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.UserId != nil {
		v := *s.UserId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "userId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/DisassociatePhoneNumberFromUserResponse
type DisassociatePhoneNumberFromUserOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DisassociatePhoneNumberFromUserOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DisassociatePhoneNumberFromUserOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDisassociatePhoneNumberFromUser = "DisassociatePhoneNumberFromUser"

// DisassociatePhoneNumberFromUserRequest returns a request value for making API operation for
// Amazon Chime.
//
// Disassociates the primary provisioned phone number from the specified Amazon
// Chime user.
//
//    // Example sending a request using DisassociatePhoneNumberFromUserRequest.
//    req := client.DisassociatePhoneNumberFromUserRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/DisassociatePhoneNumberFromUser
func (c *Client) DisassociatePhoneNumberFromUserRequest(input *DisassociatePhoneNumberFromUserInput) DisassociatePhoneNumberFromUserRequest {
	op := &aws.Operation{
		Name:       opDisassociatePhoneNumberFromUser,
		HTTPMethod: "POST",
		HTTPPath:   "/accounts/{accountId}/users/{userId}?operation=disassociate-phone-number",
	}

	if input == nil {
		input = &DisassociatePhoneNumberFromUserInput{}
	}

	req := c.newRequest(op, input, &DisassociatePhoneNumberFromUserOutput{})
	return DisassociatePhoneNumberFromUserRequest{Request: req, Input: input, Copy: c.DisassociatePhoneNumberFromUserRequest}
}

// DisassociatePhoneNumberFromUserRequest is the request type for the
// DisassociatePhoneNumberFromUser API operation.
type DisassociatePhoneNumberFromUserRequest struct {
	*aws.Request
	Input *DisassociatePhoneNumberFromUserInput
	Copy  func(*DisassociatePhoneNumberFromUserInput) DisassociatePhoneNumberFromUserRequest
}

// Send marshals and sends the DisassociatePhoneNumberFromUser API request.
func (r DisassociatePhoneNumberFromUserRequest) Send(ctx context.Context) (*DisassociatePhoneNumberFromUserResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DisassociatePhoneNumberFromUserResponse{
		DisassociatePhoneNumberFromUserOutput: r.Request.Data.(*DisassociatePhoneNumberFromUserOutput),
		response:                              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DisassociatePhoneNumberFromUserResponse is the response type for the
// DisassociatePhoneNumberFromUser API operation.
type DisassociatePhoneNumberFromUserResponse struct {
	*DisassociatePhoneNumberFromUserOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DisassociatePhoneNumberFromUser request.
func (r *DisassociatePhoneNumberFromUserResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
