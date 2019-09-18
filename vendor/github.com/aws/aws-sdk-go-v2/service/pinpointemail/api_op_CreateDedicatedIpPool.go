// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpointemail

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// A request to create a new dedicated IP pool.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-email-2018-07-26/CreateDedicatedIpPoolRequest
type CreateDedicatedIpPoolInput struct {
	_ struct{} `type:"structure"`

	// The name of the dedicated IP pool.
	//
	// PoolName is a required field
	PoolName *string `type:"string" required:"true"`

	// An object that defines the tags (keys and values) that you want to associate
	// with the pool.
	Tags []Tag `type:"list"`
}

// String returns the string representation
func (s CreateDedicatedIpPoolInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDedicatedIpPoolInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateDedicatedIpPoolInput"}

	if s.PoolName == nil {
		invalidParams.Add(aws.NewErrParamRequired("PoolName"))
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateDedicatedIpPoolInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.PoolName != nil {
		v := *s.PoolName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "PoolName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Tags", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

// An HTTP 200 response if the request succeeds, or an error message if the
// request fails.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-email-2018-07-26/CreateDedicatedIpPoolResponse
type CreateDedicatedIpPoolOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CreateDedicatedIpPoolOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateDedicatedIpPoolOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opCreateDedicatedIpPool = "CreateDedicatedIpPool"

// CreateDedicatedIpPoolRequest returns a request value for making API operation for
// Amazon Pinpoint Email Service.
//
// Create a new pool of dedicated IP addresses. A pool can include one or more
// dedicated IP addresses that are associated with your Amazon Pinpoint account.
// You can associate a pool with a configuration set. When you send an email
// that uses that configuration set, Amazon Pinpoint sends it using only the
// IP addresses in the associated pool.
//
//    // Example sending a request using CreateDedicatedIpPoolRequest.
//    req := client.CreateDedicatedIpPoolRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-email-2018-07-26/CreateDedicatedIpPool
func (c *Client) CreateDedicatedIpPoolRequest(input *CreateDedicatedIpPoolInput) CreateDedicatedIpPoolRequest {
	op := &aws.Operation{
		Name:       opCreateDedicatedIpPool,
		HTTPMethod: "POST",
		HTTPPath:   "/v1/email/dedicated-ip-pools",
	}

	if input == nil {
		input = &CreateDedicatedIpPoolInput{}
	}

	req := c.newRequest(op, input, &CreateDedicatedIpPoolOutput{})
	return CreateDedicatedIpPoolRequest{Request: req, Input: input, Copy: c.CreateDedicatedIpPoolRequest}
}

// CreateDedicatedIpPoolRequest is the request type for the
// CreateDedicatedIpPool API operation.
type CreateDedicatedIpPoolRequest struct {
	*aws.Request
	Input *CreateDedicatedIpPoolInput
	Copy  func(*CreateDedicatedIpPoolInput) CreateDedicatedIpPoolRequest
}

// Send marshals and sends the CreateDedicatedIpPool API request.
func (r CreateDedicatedIpPoolRequest) Send(ctx context.Context) (*CreateDedicatedIpPoolResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateDedicatedIpPoolResponse{
		CreateDedicatedIpPoolOutput: r.Request.Data.(*CreateDedicatedIpPoolOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateDedicatedIpPoolResponse is the response type for the
// CreateDedicatedIpPool API operation.
type CreateDedicatedIpPoolResponse struct {
	*CreateDedicatedIpPoolOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateDedicatedIpPool request.
func (r *CreateDedicatedIpPoolResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
