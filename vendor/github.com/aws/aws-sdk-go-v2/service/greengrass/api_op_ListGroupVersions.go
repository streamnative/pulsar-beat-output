// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package greengrass

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/ListGroupVersionsRequest
type ListGroupVersionsInput struct {
	_ struct{} `type:"structure"`

	// GroupId is a required field
	GroupId *string `location:"uri" locationName:"GroupId" type:"string" required:"true"`

	MaxResults *string `location:"querystring" locationName:"MaxResults" type:"string"`

	NextToken *string `location:"querystring" locationName:"NextToken" type:"string"`
}

// String returns the string representation
func (s ListGroupVersionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListGroupVersionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListGroupVersionsInput"}

	if s.GroupId == nil {
		invalidParams.Add(aws.NewErrParamRequired("GroupId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListGroupVersionsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.GroupId != nil {
		v := *s.GroupId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "GroupId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "MaxResults", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/ListGroupVersionsResponse
type ListGroupVersionsOutput struct {
	_ struct{} `type:"structure"`

	NextToken *string `type:"string"`

	Versions []VersionInformation `type:"list"`
}

// String returns the string representation
func (s ListGroupVersionsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListGroupVersionsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Versions != nil {
		v := s.Versions

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Versions", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opListGroupVersions = "ListGroupVersions"

// ListGroupVersionsRequest returns a request value for making API operation for
// AWS Greengrass.
//
// Lists the versions of a group.
//
//    // Example sending a request using ListGroupVersionsRequest.
//    req := client.ListGroupVersionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/ListGroupVersions
func (c *Client) ListGroupVersionsRequest(input *ListGroupVersionsInput) ListGroupVersionsRequest {
	op := &aws.Operation{
		Name:       opListGroupVersions,
		HTTPMethod: "GET",
		HTTPPath:   "/greengrass/groups/{GroupId}/versions",
	}

	if input == nil {
		input = &ListGroupVersionsInput{}
	}

	req := c.newRequest(op, input, &ListGroupVersionsOutput{})
	return ListGroupVersionsRequest{Request: req, Input: input, Copy: c.ListGroupVersionsRequest}
}

// ListGroupVersionsRequest is the request type for the
// ListGroupVersions API operation.
type ListGroupVersionsRequest struct {
	*aws.Request
	Input *ListGroupVersionsInput
	Copy  func(*ListGroupVersionsInput) ListGroupVersionsRequest
}

// Send marshals and sends the ListGroupVersions API request.
func (r ListGroupVersionsRequest) Send(ctx context.Context) (*ListGroupVersionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListGroupVersionsResponse{
		ListGroupVersionsOutput: r.Request.Data.(*ListGroupVersionsOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListGroupVersionsResponse is the response type for the
// ListGroupVersions API operation.
type ListGroupVersionsResponse struct {
	*ListGroupVersionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListGroupVersions request.
func (r *ListGroupVersionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
