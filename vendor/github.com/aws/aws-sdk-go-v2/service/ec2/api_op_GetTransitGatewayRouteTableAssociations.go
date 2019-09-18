// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/GetTransitGatewayRouteTableAssociationsRequest
type GetTransitGatewayRouteTableAssociationsInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// One or more filters. The possible values are:
	//
	//    * resource-id - The ID of the resource.
	//
	//    * resource-type - The resource type (vpc | vpn).
	//
	//    * transit-gateway-attachment-id - The ID of the attachment.
	Filters []Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int64 `min:"5" type:"integer"`

	// The token for the next page of results.
	NextToken *string `type:"string"`

	// The ID of the transit gateway route table.
	//
	// TransitGatewayRouteTableId is a required field
	TransitGatewayRouteTableId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s GetTransitGatewayRouteTableAssociationsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetTransitGatewayRouteTableAssociationsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetTransitGatewayRouteTableAssociationsInput"}
	if s.MaxResults != nil && *s.MaxResults < 5 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 5))
	}

	if s.TransitGatewayRouteTableId == nil {
		invalidParams.Add(aws.NewErrParamRequired("TransitGatewayRouteTableId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/GetTransitGatewayRouteTableAssociationsResult
type GetTransitGatewayRouteTableAssociationsOutput struct {
	_ struct{} `type:"structure"`

	// Information about the associations.
	Associations []TransitGatewayRouteTableAssociation `locationName:"associations" locationNameList:"item" type:"list"`

	// The token to use to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s GetTransitGatewayRouteTableAssociationsOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetTransitGatewayRouteTableAssociations = "GetTransitGatewayRouteTableAssociations"

// GetTransitGatewayRouteTableAssociationsRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Gets information about the associations for the specified transit gateway
// route table.
//
//    // Example sending a request using GetTransitGatewayRouteTableAssociationsRequest.
//    req := client.GetTransitGatewayRouteTableAssociationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/GetTransitGatewayRouteTableAssociations
func (c *Client) GetTransitGatewayRouteTableAssociationsRequest(input *GetTransitGatewayRouteTableAssociationsInput) GetTransitGatewayRouteTableAssociationsRequest {
	op := &aws.Operation{
		Name:       opGetTransitGatewayRouteTableAssociations,
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
		input = &GetTransitGatewayRouteTableAssociationsInput{}
	}

	req := c.newRequest(op, input, &GetTransitGatewayRouteTableAssociationsOutput{})
	return GetTransitGatewayRouteTableAssociationsRequest{Request: req, Input: input, Copy: c.GetTransitGatewayRouteTableAssociationsRequest}
}

// GetTransitGatewayRouteTableAssociationsRequest is the request type for the
// GetTransitGatewayRouteTableAssociations API operation.
type GetTransitGatewayRouteTableAssociationsRequest struct {
	*aws.Request
	Input *GetTransitGatewayRouteTableAssociationsInput
	Copy  func(*GetTransitGatewayRouteTableAssociationsInput) GetTransitGatewayRouteTableAssociationsRequest
}

// Send marshals and sends the GetTransitGatewayRouteTableAssociations API request.
func (r GetTransitGatewayRouteTableAssociationsRequest) Send(ctx context.Context) (*GetTransitGatewayRouteTableAssociationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetTransitGatewayRouteTableAssociationsResponse{
		GetTransitGatewayRouteTableAssociationsOutput: r.Request.Data.(*GetTransitGatewayRouteTableAssociationsOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetTransitGatewayRouteTableAssociationsRequestPaginator returns a paginator for GetTransitGatewayRouteTableAssociations.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetTransitGatewayRouteTableAssociationsRequest(input)
//   p := ec2.NewGetTransitGatewayRouteTableAssociationsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetTransitGatewayRouteTableAssociationsPaginator(req GetTransitGatewayRouteTableAssociationsRequest) GetTransitGatewayRouteTableAssociationsPaginator {
	return GetTransitGatewayRouteTableAssociationsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *GetTransitGatewayRouteTableAssociationsInput
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

// GetTransitGatewayRouteTableAssociationsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetTransitGatewayRouteTableAssociationsPaginator struct {
	aws.Pager
}

func (p *GetTransitGatewayRouteTableAssociationsPaginator) CurrentPage() *GetTransitGatewayRouteTableAssociationsOutput {
	return p.Pager.CurrentPage().(*GetTransitGatewayRouteTableAssociationsOutput)
}

// GetTransitGatewayRouteTableAssociationsResponse is the response type for the
// GetTransitGatewayRouteTableAssociations API operation.
type GetTransitGatewayRouteTableAssociationsResponse struct {
	*GetTransitGatewayRouteTableAssociationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetTransitGatewayRouteTableAssociations request.
func (r *GetTransitGatewayRouteTableAssociationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
