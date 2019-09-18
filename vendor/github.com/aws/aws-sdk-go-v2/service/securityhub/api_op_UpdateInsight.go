// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package securityhub

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/securityhub-2018-10-26/UpdateInsightRequest
type UpdateInsightInput struct {
	_ struct{} `type:"structure"`

	// The updated filters that define this insight.
	Filters *AwsSecurityFindingFilters `type:"structure"`

	// The updated GroupBy attribute that defines this insight.
	GroupByAttribute *string `type:"string"`

	// The ARN of the insight that you want to update.
	//
	// InsightArn is a required field
	InsightArn *string `location:"uri" locationName:"InsightArn" type:"string" required:"true"`

	// The updated name for the insight.
	Name *string `type:"string"`
}

// String returns the string representation
func (s UpdateInsightInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateInsightInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateInsightInput"}

	if s.InsightArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("InsightArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateInsightInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Filters != nil {
		v := s.Filters

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Filters", v, metadata)
	}
	if s.GroupByAttribute != nil {
		v := *s.GroupByAttribute

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "GroupByAttribute", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.InsightArn != nil {
		v := *s.InsightArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "InsightArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/securityhub-2018-10-26/UpdateInsightResponse
type UpdateInsightOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateInsightOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateInsightOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opUpdateInsight = "UpdateInsight"

// UpdateInsightRequest returns a request value for making API operation for
// AWS SecurityHub.
//
// Updates the Security Hub insight that the insight ARN specifies.
//
//    // Example sending a request using UpdateInsightRequest.
//    req := client.UpdateInsightRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/securityhub-2018-10-26/UpdateInsight
func (c *Client) UpdateInsightRequest(input *UpdateInsightInput) UpdateInsightRequest {
	op := &aws.Operation{
		Name:       opUpdateInsight,
		HTTPMethod: "PATCH",
		HTTPPath:   "/insights/{InsightArn+}",
	}

	if input == nil {
		input = &UpdateInsightInput{}
	}

	req := c.newRequest(op, input, &UpdateInsightOutput{})
	return UpdateInsightRequest{Request: req, Input: input, Copy: c.UpdateInsightRequest}
}

// UpdateInsightRequest is the request type for the
// UpdateInsight API operation.
type UpdateInsightRequest struct {
	*aws.Request
	Input *UpdateInsightInput
	Copy  func(*UpdateInsightInput) UpdateInsightRequest
}

// Send marshals and sends the UpdateInsight API request.
func (r UpdateInsightRequest) Send(ctx context.Context) (*UpdateInsightResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateInsightResponse{
		UpdateInsightOutput: r.Request.Data.(*UpdateInsightOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateInsightResponse is the response type for the
// UpdateInsight API operation.
type UpdateInsightResponse struct {
	*UpdateInsightOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateInsight request.
func (r *UpdateInsightResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
