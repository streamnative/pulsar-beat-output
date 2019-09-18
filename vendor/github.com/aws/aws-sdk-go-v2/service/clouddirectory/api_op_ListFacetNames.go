// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package clouddirectory

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/ListFacetNamesRequest
type ListFacetNamesInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results to retrieve.
	MaxResults *int64 `min:"1" type:"integer"`

	// The pagination token.
	NextToken *string `type:"string"`

	// The Amazon Resource Name (ARN) to retrieve facet names from.
	//
	// SchemaArn is a required field
	SchemaArn *string `location:"header" locationName:"x-amz-data-partition" type:"string" required:"true"`
}

// String returns the string representation
func (s ListFacetNamesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListFacetNamesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListFacetNamesInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if s.SchemaArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("SchemaArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListFacetNamesInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "MaxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SchemaArn != nil {
		v := *s.SchemaArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-data-partition", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/ListFacetNamesResponse
type ListFacetNamesOutput struct {
	_ struct{} `type:"structure"`

	// The names of facets that exist within the schema.
	FacetNames []string `type:"list"`

	// The pagination token.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListFacetNamesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListFacetNamesOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.FacetNames != nil {
		v := s.FacetNames

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "FacetNames", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListFacetNames = "ListFacetNames"

// ListFacetNamesRequest returns a request value for making API operation for
// Amazon CloudDirectory.
//
// Retrieves the names of facets that exist in a schema.
//
//    // Example sending a request using ListFacetNamesRequest.
//    req := client.ListFacetNamesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/ListFacetNames
func (c *Client) ListFacetNamesRequest(input *ListFacetNamesInput) ListFacetNamesRequest {
	op := &aws.Operation{
		Name:       opListFacetNames,
		HTTPMethod: "POST",
		HTTPPath:   "/amazonclouddirectory/2017-01-11/facet/list",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListFacetNamesInput{}
	}

	req := c.newRequest(op, input, &ListFacetNamesOutput{})
	return ListFacetNamesRequest{Request: req, Input: input, Copy: c.ListFacetNamesRequest}
}

// ListFacetNamesRequest is the request type for the
// ListFacetNames API operation.
type ListFacetNamesRequest struct {
	*aws.Request
	Input *ListFacetNamesInput
	Copy  func(*ListFacetNamesInput) ListFacetNamesRequest
}

// Send marshals and sends the ListFacetNames API request.
func (r ListFacetNamesRequest) Send(ctx context.Context) (*ListFacetNamesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListFacetNamesResponse{
		ListFacetNamesOutput: r.Request.Data.(*ListFacetNamesOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListFacetNamesRequestPaginator returns a paginator for ListFacetNames.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListFacetNamesRequest(input)
//   p := clouddirectory.NewListFacetNamesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListFacetNamesPaginator(req ListFacetNamesRequest) ListFacetNamesPaginator {
	return ListFacetNamesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListFacetNamesInput
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

// ListFacetNamesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListFacetNamesPaginator struct {
	aws.Pager
}

func (p *ListFacetNamesPaginator) CurrentPage() *ListFacetNamesOutput {
	return p.Pager.CurrentPage().(*ListFacetNamesOutput)
}

// ListFacetNamesResponse is the response type for the
// ListFacetNames API operation.
type ListFacetNamesResponse struct {
	*ListFacetNamesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListFacetNames request.
func (r *ListFacetNamesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
