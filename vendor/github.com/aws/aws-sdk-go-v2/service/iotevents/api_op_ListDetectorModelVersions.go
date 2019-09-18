// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotevents

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotevents-2018-07-27/ListDetectorModelVersionsRequest
type ListDetectorModelVersionsInput struct {
	_ struct{} `type:"structure"`

	// The name of the detector model whose versions are returned.
	//
	// DetectorModelName is a required field
	DetectorModelName *string `location:"uri" locationName:"detectorModelName" min:"1" type:"string" required:"true"`

	// The maximum number of results to return at one time.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// The token for the next set of results.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListDetectorModelVersionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListDetectorModelVersionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListDetectorModelVersionsInput"}

	if s.DetectorModelName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DetectorModelName"))
	}
	if s.DetectorModelName != nil && len(*s.DetectorModelName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DetectorModelName", 1))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListDetectorModelVersionsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.DetectorModelName != nil {
		v := *s.DetectorModelName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "detectorModelName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotevents-2018-07-27/ListDetectorModelVersionsResponse
type ListDetectorModelVersionsOutput struct {
	_ struct{} `type:"structure"`

	// Summary information about the detector model versions.
	DetectorModelVersionSummaries []DetectorModelVersionSummary `locationName:"detectorModelVersionSummaries" type:"list"`

	// A token to retrieve the next set of results, or null if there are no additional
	// results.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListDetectorModelVersionsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListDetectorModelVersionsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.DetectorModelVersionSummaries != nil {
		v := s.DetectorModelVersionSummaries

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "detectorModelVersionSummaries", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListDetectorModelVersions = "ListDetectorModelVersions"

// ListDetectorModelVersionsRequest returns a request value for making API operation for
// AWS IoT Events.
//
// Lists all the versions of a detector model. Only the metadata associated
// with each detector model version is returned.
//
//    // Example sending a request using ListDetectorModelVersionsRequest.
//    req := client.ListDetectorModelVersionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotevents-2018-07-27/ListDetectorModelVersions
func (c *Client) ListDetectorModelVersionsRequest(input *ListDetectorModelVersionsInput) ListDetectorModelVersionsRequest {
	op := &aws.Operation{
		Name:       opListDetectorModelVersions,
		HTTPMethod: "GET",
		HTTPPath:   "/detector-models/{detectorModelName}/versions",
	}

	if input == nil {
		input = &ListDetectorModelVersionsInput{}
	}

	req := c.newRequest(op, input, &ListDetectorModelVersionsOutput{})
	return ListDetectorModelVersionsRequest{Request: req, Input: input, Copy: c.ListDetectorModelVersionsRequest}
}

// ListDetectorModelVersionsRequest is the request type for the
// ListDetectorModelVersions API operation.
type ListDetectorModelVersionsRequest struct {
	*aws.Request
	Input *ListDetectorModelVersionsInput
	Copy  func(*ListDetectorModelVersionsInput) ListDetectorModelVersionsRequest
}

// Send marshals and sends the ListDetectorModelVersions API request.
func (r ListDetectorModelVersionsRequest) Send(ctx context.Context) (*ListDetectorModelVersionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListDetectorModelVersionsResponse{
		ListDetectorModelVersionsOutput: r.Request.Data.(*ListDetectorModelVersionsOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListDetectorModelVersionsResponse is the response type for the
// ListDetectorModelVersions API operation.
type ListDetectorModelVersionsResponse struct {
	*ListDetectorModelVersionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListDetectorModelVersions request.
func (r *ListDetectorModelVersionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
