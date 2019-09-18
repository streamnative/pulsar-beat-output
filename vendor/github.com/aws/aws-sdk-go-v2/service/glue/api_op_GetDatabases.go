// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/GetDatabasesRequest
type GetDatabasesInput struct {
	_ struct{} `type:"structure"`

	// The ID of the Data Catalog from which to retrieve Databases. If none is provided,
	// the AWS account ID is used by default.
	CatalogId *string `min:"1" type:"string"`

	// The maximum number of databases to return in one response.
	MaxResults *int64 `min:"1" type:"integer"`

	// A continuation token, if this is a continuation call.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s GetDatabasesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetDatabasesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetDatabasesInput"}
	if s.CatalogId != nil && len(*s.CatalogId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CatalogId", 1))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/GetDatabasesResponse
type GetDatabasesOutput struct {
	_ struct{} `type:"structure"`

	// A list of Database objects from the specified catalog.
	//
	// DatabaseList is a required field
	DatabaseList []Database `type:"list" required:"true"`

	// A continuation token for paginating the returned list of tokens, returned
	// if the current segment of the list is not the last.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s GetDatabasesOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetDatabases = "GetDatabases"

// GetDatabasesRequest returns a request value for making API operation for
// AWS Glue.
//
// Retrieves all databases defined in a given Data Catalog.
//
//    // Example sending a request using GetDatabasesRequest.
//    req := client.GetDatabasesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/GetDatabases
func (c *Client) GetDatabasesRequest(input *GetDatabasesInput) GetDatabasesRequest {
	op := &aws.Operation{
		Name:       opGetDatabases,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &GetDatabasesInput{}
	}

	req := c.newRequest(op, input, &GetDatabasesOutput{})
	return GetDatabasesRequest{Request: req, Input: input, Copy: c.GetDatabasesRequest}
}

// GetDatabasesRequest is the request type for the
// GetDatabases API operation.
type GetDatabasesRequest struct {
	*aws.Request
	Input *GetDatabasesInput
	Copy  func(*GetDatabasesInput) GetDatabasesRequest
}

// Send marshals and sends the GetDatabases API request.
func (r GetDatabasesRequest) Send(ctx context.Context) (*GetDatabasesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetDatabasesResponse{
		GetDatabasesOutput: r.Request.Data.(*GetDatabasesOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetDatabasesRequestPaginator returns a paginator for GetDatabases.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetDatabasesRequest(input)
//   p := glue.NewGetDatabasesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetDatabasesPaginator(req GetDatabasesRequest) GetDatabasesPaginator {
	return GetDatabasesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *GetDatabasesInput
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

// GetDatabasesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetDatabasesPaginator struct {
	aws.Pager
}

func (p *GetDatabasesPaginator) CurrentPage() *GetDatabasesOutput {
	return p.Pager.CurrentPage().(*GetDatabasesOutput)
}

// GetDatabasesResponse is the response type for the
// GetDatabases API operation.
type GetDatabasesResponse struct {
	*GetDatabasesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetDatabases request.
func (r *GetDatabasesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
