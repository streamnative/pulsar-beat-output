// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package worklink

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/worklink-2018-09-25/DescribeWebsiteCertificateAuthorityRequest
type DescribeWebsiteCertificateAuthorityInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the fleet.
	//
	// FleetArn is a required field
	FleetArn *string `min:"20" type:"string" required:"true"`

	// A unique identifier for the certificate authority.
	//
	// WebsiteCaId is a required field
	WebsiteCaId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeWebsiteCertificateAuthorityInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeWebsiteCertificateAuthorityInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeWebsiteCertificateAuthorityInput"}

	if s.FleetArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("FleetArn"))
	}
	if s.FleetArn != nil && len(*s.FleetArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("FleetArn", 20))
	}

	if s.WebsiteCaId == nil {
		invalidParams.Add(aws.NewErrParamRequired("WebsiteCaId"))
	}
	if s.WebsiteCaId != nil && len(*s.WebsiteCaId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("WebsiteCaId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeWebsiteCertificateAuthorityInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.FleetArn != nil {
		v := *s.FleetArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "FleetArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.WebsiteCaId != nil {
		v := *s.WebsiteCaId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "WebsiteCaId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/worklink-2018-09-25/DescribeWebsiteCertificateAuthorityResponse
type DescribeWebsiteCertificateAuthorityOutput struct {
	_ struct{} `type:"structure"`

	// The root certificate of the certificate authority.
	Certificate *string `min:"1" type:"string"`

	// The time that the certificate authority was added.
	CreatedTime *time.Time `type:"timestamp"`

	// The certificate name to display.
	DisplayName *string `type:"string"`
}

// String returns the string representation
func (s DescribeWebsiteCertificateAuthorityOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeWebsiteCertificateAuthorityOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Certificate != nil {
		v := *s.Certificate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Certificate", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CreatedTime != nil {
		v := *s.CreatedTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "CreatedTime",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.DisplayName != nil {
		v := *s.DisplayName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DisplayName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opDescribeWebsiteCertificateAuthority = "DescribeWebsiteCertificateAuthority"

// DescribeWebsiteCertificateAuthorityRequest returns a request value for making API operation for
// Amazon WorkLink.
//
// Provides information about the certificate authority.
//
//    // Example sending a request using DescribeWebsiteCertificateAuthorityRequest.
//    req := client.DescribeWebsiteCertificateAuthorityRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/worklink-2018-09-25/DescribeWebsiteCertificateAuthority
func (c *Client) DescribeWebsiteCertificateAuthorityRequest(input *DescribeWebsiteCertificateAuthorityInput) DescribeWebsiteCertificateAuthorityRequest {
	op := &aws.Operation{
		Name:       opDescribeWebsiteCertificateAuthority,
		HTTPMethod: "POST",
		HTTPPath:   "/describeWebsiteCertificateAuthority",
	}

	if input == nil {
		input = &DescribeWebsiteCertificateAuthorityInput{}
	}

	req := c.newRequest(op, input, &DescribeWebsiteCertificateAuthorityOutput{})
	return DescribeWebsiteCertificateAuthorityRequest{Request: req, Input: input, Copy: c.DescribeWebsiteCertificateAuthorityRequest}
}

// DescribeWebsiteCertificateAuthorityRequest is the request type for the
// DescribeWebsiteCertificateAuthority API operation.
type DescribeWebsiteCertificateAuthorityRequest struct {
	*aws.Request
	Input *DescribeWebsiteCertificateAuthorityInput
	Copy  func(*DescribeWebsiteCertificateAuthorityInput) DescribeWebsiteCertificateAuthorityRequest
}

// Send marshals and sends the DescribeWebsiteCertificateAuthority API request.
func (r DescribeWebsiteCertificateAuthorityRequest) Send(ctx context.Context) (*DescribeWebsiteCertificateAuthorityResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeWebsiteCertificateAuthorityResponse{
		DescribeWebsiteCertificateAuthorityOutput: r.Request.Data.(*DescribeWebsiteCertificateAuthorityOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeWebsiteCertificateAuthorityResponse is the response type for the
// DescribeWebsiteCertificateAuthority API operation.
type DescribeWebsiteCertificateAuthorityResponse struct {
	*DescribeWebsiteCertificateAuthorityOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeWebsiteCertificateAuthority request.
func (r *DescribeWebsiteCertificateAuthorityResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
