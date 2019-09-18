// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package batch

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/batch-2016-08-10/DescribeJobDefinitionsRequest
type DescribeJobDefinitionsInput struct {
	_ struct{} `type:"structure"`

	// The name of the job definition to describe.
	JobDefinitionName *string `locationName:"jobDefinitionName" type:"string"`

	// A list of up to 100 job definition names or full Amazon Resource Name (ARN)
	// entries.
	JobDefinitions []string `locationName:"jobDefinitions" type:"list"`

	// The maximum number of results returned by DescribeJobDefinitions in paginated
	// output. When this parameter is used, DescribeJobDefinitions only returns
	// maxResults results in a single page along with a nextToken response element.
	// The remaining results of the initial request can be seen by sending another
	// DescribeJobDefinitions request with the returned nextToken value. This value
	// can be between 1 and 100. If this parameter is not used, then DescribeJobDefinitions
	// returns up to 100 results and a nextToken value if applicable.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// The nextToken value returned from a previous paginated DescribeJobDefinitions
	// request where maxResults was used and the results exceeded the value of that
	// parameter. Pagination continues from the end of the previous results that
	// returned the nextToken value. This value is null when there are no more results
	// to return.
	//
	// This token should be treated as an opaque identifier that is only used to
	// retrieve the next items in a list and not for other programmatic purposes.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The status with which to filter job definitions.
	Status *string `locationName:"status" type:"string"`
}

// String returns the string representation
func (s DescribeJobDefinitionsInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeJobDefinitionsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.JobDefinitionName != nil {
		v := *s.JobDefinitionName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "jobDefinitionName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.JobDefinitions != nil {
		v := s.JobDefinitions

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "jobDefinitions", metadata)
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
	if s.Status != nil {
		v := *s.Status

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "status", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/batch-2016-08-10/DescribeJobDefinitionsResponse
type DescribeJobDefinitionsOutput struct {
	_ struct{} `type:"structure"`

	// The list of job definitions.
	JobDefinitions []JobDefinition `locationName:"jobDefinitions" type:"list"`

	// The nextToken value to include in a future DescribeJobDefinitions request.
	// When the results of a DescribeJobDefinitions request exceed maxResults, this
	// value can be used to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s DescribeJobDefinitionsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeJobDefinitionsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.JobDefinitions != nil {
		v := s.JobDefinitions

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "jobDefinitions", metadata)
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

const opDescribeJobDefinitions = "DescribeJobDefinitions"

// DescribeJobDefinitionsRequest returns a request value for making API operation for
// AWS Batch.
//
// Describes a list of job definitions. You can specify a status (such as ACTIVE)
// to only return job definitions that match that status.
//
//    // Example sending a request using DescribeJobDefinitionsRequest.
//    req := client.DescribeJobDefinitionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/batch-2016-08-10/DescribeJobDefinitions
func (c *Client) DescribeJobDefinitionsRequest(input *DescribeJobDefinitionsInput) DescribeJobDefinitionsRequest {
	op := &aws.Operation{
		Name:       opDescribeJobDefinitions,
		HTTPMethod: "POST",
		HTTPPath:   "/v1/describejobdefinitions",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeJobDefinitionsInput{}
	}

	req := c.newRequest(op, input, &DescribeJobDefinitionsOutput{})
	return DescribeJobDefinitionsRequest{Request: req, Input: input, Copy: c.DescribeJobDefinitionsRequest}
}

// DescribeJobDefinitionsRequest is the request type for the
// DescribeJobDefinitions API operation.
type DescribeJobDefinitionsRequest struct {
	*aws.Request
	Input *DescribeJobDefinitionsInput
	Copy  func(*DescribeJobDefinitionsInput) DescribeJobDefinitionsRequest
}

// Send marshals and sends the DescribeJobDefinitions API request.
func (r DescribeJobDefinitionsRequest) Send(ctx context.Context) (*DescribeJobDefinitionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeJobDefinitionsResponse{
		DescribeJobDefinitionsOutput: r.Request.Data.(*DescribeJobDefinitionsOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeJobDefinitionsRequestPaginator returns a paginator for DescribeJobDefinitions.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeJobDefinitionsRequest(input)
//   p := batch.NewDescribeJobDefinitionsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeJobDefinitionsPaginator(req DescribeJobDefinitionsRequest) DescribeJobDefinitionsPaginator {
	return DescribeJobDefinitionsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeJobDefinitionsInput
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

// DescribeJobDefinitionsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeJobDefinitionsPaginator struct {
	aws.Pager
}

func (p *DescribeJobDefinitionsPaginator) CurrentPage() *DescribeJobDefinitionsOutput {
	return p.Pager.CurrentPage().(*DescribeJobDefinitionsOutput)
}

// DescribeJobDefinitionsResponse is the response type for the
// DescribeJobDefinitions API operation.
type DescribeJobDefinitionsResponse struct {
	*DescribeJobDefinitionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeJobDefinitions request.
func (r *DescribeJobDefinitionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
