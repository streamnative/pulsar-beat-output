// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workdocs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/workdocs-2016-05-01/DescribeDocumentVersionsRequest
type DescribeDocumentVersionsInput struct {
	_ struct{} `type:"structure"`

	// Amazon WorkDocs authentication token. Do not set this field when using administrative
	// API actions, as in accessing the API using AWS credentials.
	AuthenticationToken *string `location:"header" locationName:"Authentication" min:"1" type:"string"`

	// The ID of the document.
	//
	// DocumentId is a required field
	DocumentId *string `location:"uri" locationName:"DocumentId" min:"1" type:"string" required:"true"`

	// Specify "SOURCE" to include initialized versions and a URL for the source
	// document.
	Fields *string `location:"querystring" locationName:"fields" min:"1" type:"string"`

	// A comma-separated list of values. Specify "INITIALIZED" to include incomplete
	// versions.
	Include *string `location:"querystring" locationName:"include" min:"1" type:"string"`

	// The maximum number of versions to return with this call.
	Limit *int64 `location:"querystring" locationName:"limit" min:"1" type:"integer"`

	// The marker for the next set of results. (You received this marker from a
	// previous call.)
	Marker *string `location:"querystring" locationName:"marker" min:"1" type:"string"`
}

// String returns the string representation
func (s DescribeDocumentVersionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeDocumentVersionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeDocumentVersionsInput"}
	if s.AuthenticationToken != nil && len(*s.AuthenticationToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AuthenticationToken", 1))
	}

	if s.DocumentId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DocumentId"))
	}
	if s.DocumentId != nil && len(*s.DocumentId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DocumentId", 1))
	}
	if s.Fields != nil && len(*s.Fields) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Fields", 1))
	}
	if s.Include != nil && len(*s.Include) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Include", 1))
	}
	if s.Limit != nil && *s.Limit < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Limit", 1))
	}
	if s.Marker != nil && len(*s.Marker) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Marker", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeDocumentVersionsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AuthenticationToken != nil {
		v := *s.AuthenticationToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Authentication", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DocumentId != nil {
		v := *s.DocumentId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "DocumentId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Fields != nil {
		v := *s.Fields

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "fields", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Include != nil {
		v := *s.Include

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "include", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Limit != nil {
		v := *s.Limit

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "limit", protocol.Int64Value(v), metadata)
	}
	if s.Marker != nil {
		v := *s.Marker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "marker", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/workdocs-2016-05-01/DescribeDocumentVersionsResponse
type DescribeDocumentVersionsOutput struct {
	_ struct{} `type:"structure"`

	// The document versions.
	DocumentVersions []DocumentVersionMetadata `type:"list"`

	// The marker to use when requesting the next set of results. If there are no
	// additional results, the string is empty.
	Marker *string `min:"1" type:"string"`
}

// String returns the string representation
func (s DescribeDocumentVersionsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeDocumentVersionsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.DocumentVersions != nil {
		v := s.DocumentVersions

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "DocumentVersions", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.Marker != nil {
		v := *s.Marker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Marker", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opDescribeDocumentVersions = "DescribeDocumentVersions"

// DescribeDocumentVersionsRequest returns a request value for making API operation for
// Amazon WorkDocs.
//
// Retrieves the document versions for the specified document.
//
// By default, only active versions are returned.
//
//    // Example sending a request using DescribeDocumentVersionsRequest.
//    req := client.DescribeDocumentVersionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workdocs-2016-05-01/DescribeDocumentVersions
func (c *Client) DescribeDocumentVersionsRequest(input *DescribeDocumentVersionsInput) DescribeDocumentVersionsRequest {
	op := &aws.Operation{
		Name:       opDescribeDocumentVersions,
		HTTPMethod: "GET",
		HTTPPath:   "/api/v1/documents/{DocumentId}/versions",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"Marker"},
			LimitToken:      "Limit",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeDocumentVersionsInput{}
	}

	req := c.newRequest(op, input, &DescribeDocumentVersionsOutput{})
	return DescribeDocumentVersionsRequest{Request: req, Input: input, Copy: c.DescribeDocumentVersionsRequest}
}

// DescribeDocumentVersionsRequest is the request type for the
// DescribeDocumentVersions API operation.
type DescribeDocumentVersionsRequest struct {
	*aws.Request
	Input *DescribeDocumentVersionsInput
	Copy  func(*DescribeDocumentVersionsInput) DescribeDocumentVersionsRequest
}

// Send marshals and sends the DescribeDocumentVersions API request.
func (r DescribeDocumentVersionsRequest) Send(ctx context.Context) (*DescribeDocumentVersionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeDocumentVersionsResponse{
		DescribeDocumentVersionsOutput: r.Request.Data.(*DescribeDocumentVersionsOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeDocumentVersionsRequestPaginator returns a paginator for DescribeDocumentVersions.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeDocumentVersionsRequest(input)
//   p := workdocs.NewDescribeDocumentVersionsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeDocumentVersionsPaginator(req DescribeDocumentVersionsRequest) DescribeDocumentVersionsPaginator {
	return DescribeDocumentVersionsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeDocumentVersionsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// DescribeDocumentVersionsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeDocumentVersionsPaginator struct {
	aws.Pager
}

func (p *DescribeDocumentVersionsPaginator) CurrentPage() *DescribeDocumentVersionsOutput {
	return p.Pager.CurrentPage().(*DescribeDocumentVersionsOutput)
}

// DescribeDocumentVersionsResponse is the response type for the
// DescribeDocumentVersions API operation.
type DescribeDocumentVersionsResponse struct {
	*DescribeDocumentVersionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeDocumentVersions request.
func (r *DescribeDocumentVersionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
