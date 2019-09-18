// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package databasemigrationservice

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/dms-2016-01-01/DescribeConnectionsMessage
type DescribeConnectionsInput struct {
	_ struct{} `type:"structure"`

	// The filters applied to the connection.
	//
	// Valid filter names: endpoint-arn | replication-instance-arn
	Filters []Filter `type:"list"`

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to
	// the value specified by MaxRecords.
	Marker *string `type:"string"`

	// The maximum number of records to include in the response. If more records
	// exist than the specified MaxRecords value, a pagination token called a marker
	// is included in the response so that the remaining results can be retrieved.
	//
	// Default: 100
	//
	// Constraints: Minimum 20, maximum 100.
	MaxRecords *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribeConnectionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeConnectionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeConnectionsInput"}
	if s.Filters != nil {
		for i, v := range s.Filters {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Filters", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/dms-2016-01-01/DescribeConnectionsResponse
type DescribeConnectionsOutput struct {
	_ struct{} `type:"structure"`

	// A description of the connections.
	Connections []Connection `type:"list"`

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to
	// the value specified by MaxRecords.
	Marker *string `type:"string"`
}

// String returns the string representation
func (s DescribeConnectionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeConnections = "DescribeConnections"

// DescribeConnectionsRequest returns a request value for making API operation for
// AWS Database Migration Service.
//
// Describes the status of the connections that have been made between the replication
// instance and an endpoint. Connections are created when you test an endpoint.
//
//    // Example sending a request using DescribeConnectionsRequest.
//    req := client.DescribeConnectionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dms-2016-01-01/DescribeConnections
func (c *Client) DescribeConnectionsRequest(input *DescribeConnectionsInput) DescribeConnectionsRequest {
	op := &aws.Operation{
		Name:       opDescribeConnections,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"Marker"},
			LimitToken:      "MaxRecords",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeConnectionsInput{}
	}

	req := c.newRequest(op, input, &DescribeConnectionsOutput{})
	return DescribeConnectionsRequest{Request: req, Input: input, Copy: c.DescribeConnectionsRequest}
}

// DescribeConnectionsRequest is the request type for the
// DescribeConnections API operation.
type DescribeConnectionsRequest struct {
	*aws.Request
	Input *DescribeConnectionsInput
	Copy  func(*DescribeConnectionsInput) DescribeConnectionsRequest
}

// Send marshals and sends the DescribeConnections API request.
func (r DescribeConnectionsRequest) Send(ctx context.Context) (*DescribeConnectionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeConnectionsResponse{
		DescribeConnectionsOutput: r.Request.Data.(*DescribeConnectionsOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeConnectionsRequestPaginator returns a paginator for DescribeConnections.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeConnectionsRequest(input)
//   p := databasemigrationservice.NewDescribeConnectionsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeConnectionsPaginator(req DescribeConnectionsRequest) DescribeConnectionsPaginator {
	return DescribeConnectionsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeConnectionsInput
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

// DescribeConnectionsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeConnectionsPaginator struct {
	aws.Pager
}

func (p *DescribeConnectionsPaginator) CurrentPage() *DescribeConnectionsOutput {
	return p.Pager.CurrentPage().(*DescribeConnectionsOutput)
}

// DescribeConnectionsResponse is the response type for the
// DescribeConnections API operation.
type DescribeConnectionsResponse struct {
	*DescribeConnectionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeConnections request.
func (r *DescribeConnectionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
