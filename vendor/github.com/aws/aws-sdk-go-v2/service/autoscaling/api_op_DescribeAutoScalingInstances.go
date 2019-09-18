// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package autoscaling

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/autoscaling-2011-01-01/DescribeAutoScalingInstancesType
type DescribeAutoScalingInstancesInput struct {
	_ struct{} `type:"structure"`

	// The IDs of the instances. You can specify up to MaxRecords IDs. If you omit
	// this parameter, all Auto Scaling instances are described. If you specify
	// an ID that does not exist, it is ignored with no error.
	InstanceIds []string `type:"list"`

	// The maximum number of items to return with this call. The default value is
	// 50 and the maximum value is 50.
	MaxRecords *int64 `type:"integer"`

	// The token for the next set of items to return. (You received this token from
	// a previous call.)
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeAutoScalingInstancesInput) String() string {
	return awsutil.Prettify(s)
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/autoscaling-2011-01-01/AutoScalingInstancesType
type DescribeAutoScalingInstancesOutput struct {
	_ struct{} `type:"structure"`

	// The instances.
	AutoScalingInstances []AutoScalingInstanceDetails `type:"list"`

	// A string that indicates that the response contains more items than can be
	// returned in a single response. To receive additional items, specify this
	// string for the NextToken value when requesting the next set of items. This
	// value is null when there are no more items to return.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeAutoScalingInstancesOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeAutoScalingInstances = "DescribeAutoScalingInstances"

// DescribeAutoScalingInstancesRequest returns a request value for making API operation for
// Auto Scaling.
//
// Describes one or more Auto Scaling instances.
//
//    // Example sending a request using DescribeAutoScalingInstancesRequest.
//    req := client.DescribeAutoScalingInstancesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/autoscaling-2011-01-01/DescribeAutoScalingInstances
func (c *Client) DescribeAutoScalingInstancesRequest(input *DescribeAutoScalingInstancesInput) DescribeAutoScalingInstancesRequest {
	op := &aws.Operation{
		Name:       opDescribeAutoScalingInstances,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxRecords",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeAutoScalingInstancesInput{}
	}

	req := c.newRequest(op, input, &DescribeAutoScalingInstancesOutput{})
	return DescribeAutoScalingInstancesRequest{Request: req, Input: input, Copy: c.DescribeAutoScalingInstancesRequest}
}

// DescribeAutoScalingInstancesRequest is the request type for the
// DescribeAutoScalingInstances API operation.
type DescribeAutoScalingInstancesRequest struct {
	*aws.Request
	Input *DescribeAutoScalingInstancesInput
	Copy  func(*DescribeAutoScalingInstancesInput) DescribeAutoScalingInstancesRequest
}

// Send marshals and sends the DescribeAutoScalingInstances API request.
func (r DescribeAutoScalingInstancesRequest) Send(ctx context.Context) (*DescribeAutoScalingInstancesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeAutoScalingInstancesResponse{
		DescribeAutoScalingInstancesOutput: r.Request.Data.(*DescribeAutoScalingInstancesOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeAutoScalingInstancesRequestPaginator returns a paginator for DescribeAutoScalingInstances.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeAutoScalingInstancesRequest(input)
//   p := autoscaling.NewDescribeAutoScalingInstancesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeAutoScalingInstancesPaginator(req DescribeAutoScalingInstancesRequest) DescribeAutoScalingInstancesPaginator {
	return DescribeAutoScalingInstancesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeAutoScalingInstancesInput
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

// DescribeAutoScalingInstancesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeAutoScalingInstancesPaginator struct {
	aws.Pager
}

func (p *DescribeAutoScalingInstancesPaginator) CurrentPage() *DescribeAutoScalingInstancesOutput {
	return p.Pager.CurrentPage().(*DescribeAutoScalingInstancesOutput)
}

// DescribeAutoScalingInstancesResponse is the response type for the
// DescribeAutoScalingInstances API operation.
type DescribeAutoScalingInstancesResponse struct {
	*DescribeAutoScalingInstancesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeAutoScalingInstances request.
func (r *DescribeAutoScalingInstancesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
