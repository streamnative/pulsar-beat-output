// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotanalytics

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotanalytics-2017-11-27/GetDatasetContentRequest
type GetDatasetContentInput struct {
	_ struct{} `type:"structure"`

	// The name of the data set whose contents are retrieved.
	//
	// DatasetName is a required field
	DatasetName *string `location:"uri" locationName:"datasetName" min:"1" type:"string" required:"true"`

	// The version of the data set whose contents are retrieved. You can also use
	// the strings "$LATEST" or "$LATEST_SUCCEEDED" to retrieve the contents of
	// the latest or latest successfully completed data set. If not specified, "$LATEST_SUCCEEDED"
	// is the default.
	VersionId *string `location:"querystring" locationName:"versionId" min:"7" type:"string"`
}

// String returns the string representation
func (s GetDatasetContentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetDatasetContentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetDatasetContentInput"}

	if s.DatasetName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DatasetName"))
	}
	if s.DatasetName != nil && len(*s.DatasetName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DatasetName", 1))
	}
	if s.VersionId != nil && len(*s.VersionId) < 7 {
		invalidParams.Add(aws.NewErrParamMinLen("VersionId", 7))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetDatasetContentInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.DatasetName != nil {
		v := *s.DatasetName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "datasetName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VersionId != nil {
		v := *s.VersionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "versionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotanalytics-2017-11-27/GetDatasetContentResponse
type GetDatasetContentOutput struct {
	_ struct{} `type:"structure"`

	// A list of "DatasetEntry" objects.
	Entries []DatasetEntry `locationName:"entries" type:"list"`

	// The status of the data set content.
	Status *DatasetContentStatus `locationName:"status" type:"structure"`

	// The time when the request was made.
	Timestamp *time.Time `locationName:"timestamp" type:"timestamp"`
}

// String returns the string representation
func (s GetDatasetContentOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetDatasetContentOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Entries != nil {
		v := s.Entries

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "entries", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.Status != nil {
		v := s.Status

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "status", v, metadata)
	}
	if s.Timestamp != nil {
		v := *s.Timestamp

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "timestamp",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	return nil
}

const opGetDatasetContent = "GetDatasetContent"

// GetDatasetContentRequest returns a request value for making API operation for
// AWS IoT Analytics.
//
// Retrieves the contents of a data set as pre-signed URIs.
//
//    // Example sending a request using GetDatasetContentRequest.
//    req := client.GetDatasetContentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotanalytics-2017-11-27/GetDatasetContent
func (c *Client) GetDatasetContentRequest(input *GetDatasetContentInput) GetDatasetContentRequest {
	op := &aws.Operation{
		Name:       opGetDatasetContent,
		HTTPMethod: "GET",
		HTTPPath:   "/datasets/{datasetName}/content",
	}

	if input == nil {
		input = &GetDatasetContentInput{}
	}

	req := c.newRequest(op, input, &GetDatasetContentOutput{})
	return GetDatasetContentRequest{Request: req, Input: input, Copy: c.GetDatasetContentRequest}
}

// GetDatasetContentRequest is the request type for the
// GetDatasetContent API operation.
type GetDatasetContentRequest struct {
	*aws.Request
	Input *GetDatasetContentInput
	Copy  func(*GetDatasetContentInput) GetDatasetContentRequest
}

// Send marshals and sends the GetDatasetContent API request.
func (r GetDatasetContentRequest) Send(ctx context.Context) (*GetDatasetContentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetDatasetContentResponse{
		GetDatasetContentOutput: r.Request.Data.(*GetDatasetContentOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetDatasetContentResponse is the response type for the
// GetDatasetContent API operation.
type GetDatasetContentResponse struct {
	*GetDatasetContentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetDatasetContent request.
func (r *GetDatasetContentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
