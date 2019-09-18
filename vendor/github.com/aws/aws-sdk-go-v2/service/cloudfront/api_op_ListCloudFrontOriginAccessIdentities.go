// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudfront

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// The request to list origin access identities.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudfront-2019-03-26/ListCloudFrontOriginAccessIdentitiesRequest
type ListCloudFrontOriginAccessIdentitiesInput struct {
	_ struct{} `type:"structure"`

	// Use this when paginating results to indicate where to begin in your list
	// of origin access identities. The results include identities in the list that
	// occur after the marker. To get the next page of results, set the Marker to
	// the value of the NextMarker from the current page's response (which is also
	// the ID of the last identity on that page).
	Marker *string `location:"querystring" locationName:"Marker" type:"string"`

	// The maximum number of origin access identities you want in the response body.
	MaxItems *int64 `location:"querystring" locationName:"MaxItems" type:"integer"`
}

// String returns the string representation
func (s ListCloudFrontOriginAccessIdentitiesInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListCloudFrontOriginAccessIdentitiesInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Marker != nil {
		v := *s.Marker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "Marker", protocol.StringValue(v), metadata)
	}
	if s.MaxItems != nil {
		v := *s.MaxItems

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "MaxItems", protocol.Int64Value(v), metadata)
	}
	return nil
}

// The returned result of the corresponding request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudfront-2019-03-26/ListCloudFrontOriginAccessIdentitiesResult
type ListCloudFrontOriginAccessIdentitiesOutput struct {
	_ struct{} `type:"structure" payload:"CloudFrontOriginAccessIdentityList"`

	// The CloudFrontOriginAccessIdentityList type.
	CloudFrontOriginAccessIdentityList *CloudFrontOriginAccessIdentityList `type:"structure"`
}

// String returns the string representation
func (s ListCloudFrontOriginAccessIdentitiesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListCloudFrontOriginAccessIdentitiesOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.CloudFrontOriginAccessIdentityList != nil {
		v := s.CloudFrontOriginAccessIdentityList

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "CloudFrontOriginAccessIdentityList", v, metadata)
	}
	return nil
}

const opListCloudFrontOriginAccessIdentities = "ListCloudFrontOriginAccessIdentities2019_03_26"

// ListCloudFrontOriginAccessIdentitiesRequest returns a request value for making API operation for
// Amazon CloudFront.
//
// Lists origin access identities.
//
//    // Example sending a request using ListCloudFrontOriginAccessIdentitiesRequest.
//    req := client.ListCloudFrontOriginAccessIdentitiesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudfront-2019-03-26/ListCloudFrontOriginAccessIdentities
func (c *Client) ListCloudFrontOriginAccessIdentitiesRequest(input *ListCloudFrontOriginAccessIdentitiesInput) ListCloudFrontOriginAccessIdentitiesRequest {
	op := &aws.Operation{
		Name:       opListCloudFrontOriginAccessIdentities,
		HTTPMethod: "GET",
		HTTPPath:   "/2019-03-26/origin-access-identity/cloudfront",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"CloudFrontOriginAccessIdentityList.NextMarker"},
			LimitToken:      "MaxItems",
			TruncationToken: "CloudFrontOriginAccessIdentityList.IsTruncated",
		},
	}

	if input == nil {
		input = &ListCloudFrontOriginAccessIdentitiesInput{}
	}

	req := c.newRequest(op, input, &ListCloudFrontOriginAccessIdentitiesOutput{})
	return ListCloudFrontOriginAccessIdentitiesRequest{Request: req, Input: input, Copy: c.ListCloudFrontOriginAccessIdentitiesRequest}
}

// ListCloudFrontOriginAccessIdentitiesRequest is the request type for the
// ListCloudFrontOriginAccessIdentities API operation.
type ListCloudFrontOriginAccessIdentitiesRequest struct {
	*aws.Request
	Input *ListCloudFrontOriginAccessIdentitiesInput
	Copy  func(*ListCloudFrontOriginAccessIdentitiesInput) ListCloudFrontOriginAccessIdentitiesRequest
}

// Send marshals and sends the ListCloudFrontOriginAccessIdentities API request.
func (r ListCloudFrontOriginAccessIdentitiesRequest) Send(ctx context.Context) (*ListCloudFrontOriginAccessIdentitiesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListCloudFrontOriginAccessIdentitiesResponse{
		ListCloudFrontOriginAccessIdentitiesOutput: r.Request.Data.(*ListCloudFrontOriginAccessIdentitiesOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListCloudFrontOriginAccessIdentitiesRequestPaginator returns a paginator for ListCloudFrontOriginAccessIdentities.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListCloudFrontOriginAccessIdentitiesRequest(input)
//   p := cloudfront.NewListCloudFrontOriginAccessIdentitiesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListCloudFrontOriginAccessIdentitiesPaginator(req ListCloudFrontOriginAccessIdentitiesRequest) ListCloudFrontOriginAccessIdentitiesPaginator {
	return ListCloudFrontOriginAccessIdentitiesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListCloudFrontOriginAccessIdentitiesInput
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

// ListCloudFrontOriginAccessIdentitiesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListCloudFrontOriginAccessIdentitiesPaginator struct {
	aws.Pager
}

func (p *ListCloudFrontOriginAccessIdentitiesPaginator) CurrentPage() *ListCloudFrontOriginAccessIdentitiesOutput {
	return p.Pager.CurrentPage().(*ListCloudFrontOriginAccessIdentitiesOutput)
}

// ListCloudFrontOriginAccessIdentitiesResponse is the response type for the
// ListCloudFrontOriginAccessIdentities API operation.
type ListCloudFrontOriginAccessIdentitiesResponse struct {
	*ListCloudFrontOriginAccessIdentitiesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListCloudFrontOriginAccessIdentities request.
func (r *ListCloudFrontOriginAccessIdentitiesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
