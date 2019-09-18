// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package connect

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/connect-2017-08-08/GetFederationTokenRequest
type GetFederationTokenInput struct {
	_ struct{} `type:"structure"`

	// The identifier for your Amazon Connect instance. To find the ID of your instance,
	// open the AWS console and select Amazon Connect. Select the alias of the instance
	// in the Instance alias column. The instance ID is displayed in the Overview
	// section of your instance settings. For example, the instance ID is the set
	// of characters at the end of the instance ARN, after instance/, such as 10a4c4eb-f57e-4d4c-b602-bf39176ced07.
	//
	// InstanceId is a required field
	InstanceId *string `location:"uri" locationName:"InstanceId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetFederationTokenInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetFederationTokenInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetFederationTokenInput"}

	if s.InstanceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceId"))
	}
	if s.InstanceId != nil && len(*s.InstanceId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("InstanceId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetFederationTokenInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.InstanceId != nil {
		v := *s.InstanceId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "InstanceId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/connect-2017-08-08/GetFederationTokenResponse
type GetFederationTokenOutput struct {
	_ struct{} `type:"structure"`

	// The credentials to use for federation.
	Credentials *Credentials `type:"structure"`
}

// String returns the string representation
func (s GetFederationTokenOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetFederationTokenOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Credentials != nil {
		v := s.Credentials

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Credentials", v, metadata)
	}
	return nil
}

const opGetFederationToken = "GetFederationToken"

// GetFederationTokenRequest returns a request value for making API operation for
// Amazon Connect Service.
//
// Retrieves a token for federation.
//
//    // Example sending a request using GetFederationTokenRequest.
//    req := client.GetFederationTokenRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/connect-2017-08-08/GetFederationToken
func (c *Client) GetFederationTokenRequest(input *GetFederationTokenInput) GetFederationTokenRequest {
	op := &aws.Operation{
		Name:       opGetFederationToken,
		HTTPMethod: "GET",
		HTTPPath:   "/user/federate/{InstanceId}",
	}

	if input == nil {
		input = &GetFederationTokenInput{}
	}

	req := c.newRequest(op, input, &GetFederationTokenOutput{})
	return GetFederationTokenRequest{Request: req, Input: input, Copy: c.GetFederationTokenRequest}
}

// GetFederationTokenRequest is the request type for the
// GetFederationToken API operation.
type GetFederationTokenRequest struct {
	*aws.Request
	Input *GetFederationTokenInput
	Copy  func(*GetFederationTokenInput) GetFederationTokenRequest
}

// Send marshals and sends the GetFederationToken API request.
func (r GetFederationTokenRequest) Send(ctx context.Context) (*GetFederationTokenResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetFederationTokenResponse{
		GetFederationTokenOutput: r.Request.Data.(*GetFederationTokenOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetFederationTokenResponse is the response type for the
// GetFederationToken API operation.
type GetFederationTokenResponse struct {
	*GetFederationTokenOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetFederationToken request.
func (r *GetFederationTokenResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
