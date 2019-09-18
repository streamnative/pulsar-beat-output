// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeDhcpOptionsRequest
type DescribeDhcpOptionsInput struct {
	_ struct{} `type:"structure"`

	// The IDs of one or more DHCP options sets.
	//
	// Default: Describes all your DHCP options sets.
	DhcpOptionsIds []string `locationName:"DhcpOptionsId" locationNameList:"DhcpOptionsId" type:"list"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// One or more filters.
	//
	//    * dhcp-options-id - The ID of a DHCP options set.
	//
	//    * key - The key for one of the options (for example, domain-name).
	//
	//    * value - The value for one of the options.
	//
	//    * owner-id - The ID of the AWS account that owns the DHCP options set.
	//
	//    * tag:<key> - The key/value combination of a tag assigned to the resource.
	//    Use the tag key in the filter name and the tag value as the filter value.
	//    For example, to find all resources that have a tag with the key Owner
	//    and the value TeamA, specify tag:Owner for the filter name and TeamA for
	//    the filter value.
	//
	//    * tag-key - The key of a tag assigned to the resource. Use this filter
	//    to find all resources assigned a tag with a specific key, regardless of
	//    the tag value.
	Filters []Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int64 `min:"5" type:"integer"`

	// The token for the next page of results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeDhcpOptionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeDhcpOptionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeDhcpOptionsInput"}
	if s.MaxResults != nil && *s.MaxResults < 5 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 5))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeDhcpOptionsResult
type DescribeDhcpOptionsOutput struct {
	_ struct{} `type:"structure"`

	// Information about one or more DHCP options sets.
	DhcpOptions []DhcpOptions `locationName:"dhcpOptionsSet" locationNameList:"item" type:"list"`

	// The token to use to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s DescribeDhcpOptionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeDhcpOptions = "DescribeDhcpOptions"

// DescribeDhcpOptionsRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes one or more of your DHCP options sets.
//
// For more information, see DHCP Options Sets (https://docs.aws.amazon.com/vpc/latest/userguide/VPC_DHCP_Options.html)
// in the Amazon Virtual Private Cloud User Guide.
//
//    // Example sending a request using DescribeDhcpOptionsRequest.
//    req := client.DescribeDhcpOptionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeDhcpOptions
func (c *Client) DescribeDhcpOptionsRequest(input *DescribeDhcpOptionsInput) DescribeDhcpOptionsRequest {
	op := &aws.Operation{
		Name:       opDescribeDhcpOptions,
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
		input = &DescribeDhcpOptionsInput{}
	}

	req := c.newRequest(op, input, &DescribeDhcpOptionsOutput{})
	return DescribeDhcpOptionsRequest{Request: req, Input: input, Copy: c.DescribeDhcpOptionsRequest}
}

// DescribeDhcpOptionsRequest is the request type for the
// DescribeDhcpOptions API operation.
type DescribeDhcpOptionsRequest struct {
	*aws.Request
	Input *DescribeDhcpOptionsInput
	Copy  func(*DescribeDhcpOptionsInput) DescribeDhcpOptionsRequest
}

// Send marshals and sends the DescribeDhcpOptions API request.
func (r DescribeDhcpOptionsRequest) Send(ctx context.Context) (*DescribeDhcpOptionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeDhcpOptionsResponse{
		DescribeDhcpOptionsOutput: r.Request.Data.(*DescribeDhcpOptionsOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeDhcpOptionsRequestPaginator returns a paginator for DescribeDhcpOptions.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeDhcpOptionsRequest(input)
//   p := ec2.NewDescribeDhcpOptionsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeDhcpOptionsPaginator(req DescribeDhcpOptionsRequest) DescribeDhcpOptionsPaginator {
	return DescribeDhcpOptionsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeDhcpOptionsInput
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

// DescribeDhcpOptionsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeDhcpOptionsPaginator struct {
	aws.Pager
}

func (p *DescribeDhcpOptionsPaginator) CurrentPage() *DescribeDhcpOptionsOutput {
	return p.Pager.CurrentPage().(*DescribeDhcpOptionsOutput)
}

// DescribeDhcpOptionsResponse is the response type for the
// DescribeDhcpOptions API operation.
type DescribeDhcpOptionsResponse struct {
	*DescribeDhcpOptionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeDhcpOptions request.
func (r *DescribeDhcpOptionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
