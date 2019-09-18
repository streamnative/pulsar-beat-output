// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitosync

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// The input for the BulkPublish operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-sync-2014-06-30/BulkPublishRequest
type BulkPublishInput struct {
	_ struct{} `type:"structure"`

	// A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE)
	// created by Amazon Cognito. GUID generation is unique within a region.
	//
	// IdentityPoolId is a required field
	IdentityPoolId *string `location:"uri" locationName:"IdentityPoolId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s BulkPublishInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BulkPublishInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "BulkPublishInput"}

	if s.IdentityPoolId == nil {
		invalidParams.Add(aws.NewErrParamRequired("IdentityPoolId"))
	}
	if s.IdentityPoolId != nil && len(*s.IdentityPoolId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("IdentityPoolId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s BulkPublishInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.IdentityPoolId != nil {
		v := *s.IdentityPoolId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "IdentityPoolId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// The output for the BulkPublish operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-sync-2014-06-30/BulkPublishResponse
type BulkPublishOutput struct {
	_ struct{} `type:"structure"`

	// A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE)
	// created by Amazon Cognito. GUID generation is unique within a region.
	IdentityPoolId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s BulkPublishOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s BulkPublishOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.IdentityPoolId != nil {
		v := *s.IdentityPoolId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "IdentityPoolId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opBulkPublish = "BulkPublish"

// BulkPublishRequest returns a request value for making API operation for
// Amazon Cognito Sync.
//
// Initiates a bulk publish of all existing datasets for an Identity Pool to
// the configured stream. Customers are limited to one successful bulk publish
// per 24 hours. Bulk publish is an asynchronous request, customers can see
// the status of the request via the GetBulkPublishDetails operation.
//
// This API can only be called with developer credentials. You cannot call this
// API with the temporary user credentials provided by Cognito Identity.
//
//    // Example sending a request using BulkPublishRequest.
//    req := client.BulkPublishRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-sync-2014-06-30/BulkPublish
func (c *Client) BulkPublishRequest(input *BulkPublishInput) BulkPublishRequest {
	op := &aws.Operation{
		Name:       opBulkPublish,
		HTTPMethod: "POST",
		HTTPPath:   "/identitypools/{IdentityPoolId}/bulkpublish",
	}

	if input == nil {
		input = &BulkPublishInput{}
	}

	req := c.newRequest(op, input, &BulkPublishOutput{})
	return BulkPublishRequest{Request: req, Input: input, Copy: c.BulkPublishRequest}
}

// BulkPublishRequest is the request type for the
// BulkPublish API operation.
type BulkPublishRequest struct {
	*aws.Request
	Input *BulkPublishInput
	Copy  func(*BulkPublishInput) BulkPublishRequest
}

// Send marshals and sends the BulkPublish API request.
func (r BulkPublishRequest) Send(ctx context.Context) (*BulkPublishResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &BulkPublishResponse{
		BulkPublishOutput: r.Request.Data.(*BulkPublishOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// BulkPublishResponse is the response type for the
// BulkPublish API operation.
type BulkPublishResponse struct {
	*BulkPublishOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// BulkPublish request.
func (r *BulkPublishResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
