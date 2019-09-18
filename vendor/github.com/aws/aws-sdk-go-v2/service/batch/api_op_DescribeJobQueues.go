// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package batch

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/batch-2016-08-10/DescribeJobQueuesRequest
type DescribeJobQueuesInput struct {
	_ struct{} `type:"structure"`

	// A list of up to 100 queue names or full queue Amazon Resource Name (ARN)
	// entries.
	JobQueues []string `locationName:"jobQueues" type:"list"`

	// The maximum number of results returned by DescribeJobQueues in paginated
	// output. When this parameter is used, DescribeJobQueues only returns maxResults
	// results in a single page along with a nextToken response element. The remaining
	// results of the initial request can be seen by sending another DescribeJobQueues
	// request with the returned nextToken value. This value can be between 1 and
	// 100. If this parameter is not used, then DescribeJobQueues returns up to
	// 100 results and a nextToken value if applicable.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// The nextToken value returned from a previous paginated DescribeJobQueues
	// request where maxResults was used and the results exceeded the value of that
	// parameter. Pagination continues from the end of the previous results that
	// returned the nextToken value. This value is null when there are no more results
	// to return.
	//
	// This token should be treated as an opaque identifier that is only used to
	// retrieve the next items in a list and not for other programmatic purposes.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s DescribeJobQueuesInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeJobQueuesInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.JobQueues != nil {
		v := s.JobQueues

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "jobQueues", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/batch-2016-08-10/DescribeJobQueuesResponse
type DescribeJobQueuesOutput struct {
	_ struct{} `type:"structure"`

	// The list of job queues.
	JobQueues []JobQueueDetail `locationName:"jobQueues" type:"list"`

	// The nextToken value to include in a future DescribeJobQueues request. When
	// the results of a DescribeJobQueues request exceed maxResults, this value
	// can be used to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s DescribeJobQueuesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeJobQueuesOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.JobQueues != nil {
		v := s.JobQueues

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "jobQueues", metadata)
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

const opDescribeJobQueues = "DescribeJobQueues"

// DescribeJobQueuesRequest returns a request value for making API operation for
// AWS Batch.
//
// Describes one or more of your job queues.
//
//    // Example sending a request using DescribeJobQueuesRequest.
//    req := client.DescribeJobQueuesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/batch-2016-08-10/DescribeJobQueues
func (c *Client) DescribeJobQueuesRequest(input *DescribeJobQueuesInput) DescribeJobQueuesRequest {
	op := &aws.Operation{
		Name:       opDescribeJobQueues,
		HTTPMethod: "POST",
		HTTPPath:   "/v1/describejobqueues",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeJobQueuesInput{}
	}

	req := c.newRequest(op, input, &DescribeJobQueuesOutput{})
	return DescribeJobQueuesRequest{Request: req, Input: input, Copy: c.DescribeJobQueuesRequest}
}

// DescribeJobQueuesRequest is the request type for the
// DescribeJobQueues API operation.
type DescribeJobQueuesRequest struct {
	*aws.Request
	Input *DescribeJobQueuesInput
	Copy  func(*DescribeJobQueuesInput) DescribeJobQueuesRequest
}

// Send marshals and sends the DescribeJobQueues API request.
func (r DescribeJobQueuesRequest) Send(ctx context.Context) (*DescribeJobQueuesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeJobQueuesResponse{
		DescribeJobQueuesOutput: r.Request.Data.(*DescribeJobQueuesOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeJobQueuesRequestPaginator returns a paginator for DescribeJobQueues.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeJobQueuesRequest(input)
//   p := batch.NewDescribeJobQueuesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeJobQueuesPaginator(req DescribeJobQueuesRequest) DescribeJobQueuesPaginator {
	return DescribeJobQueuesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeJobQueuesInput
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

// DescribeJobQueuesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeJobQueuesPaginator struct {
	aws.Pager
}

func (p *DescribeJobQueuesPaginator) CurrentPage() *DescribeJobQueuesOutput {
	return p.Pager.CurrentPage().(*DescribeJobQueuesOutput)
}

// DescribeJobQueuesResponse is the response type for the
// DescribeJobQueues API operation.
type DescribeJobQueuesResponse struct {
	*DescribeJobQueuesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeJobQueues request.
func (r *DescribeJobQueuesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
