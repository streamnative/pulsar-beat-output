// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workmail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/workmail-2017-10-01/ListGroupMembersRequest
type ListGroupMembersInput struct {
	_ struct{} `type:"structure"`

	// The identifier for the group to which the members (users or groups) are associated.
	//
	// GroupId is a required field
	GroupId *string `min:"12" type:"string" required:"true"`

	// The maximum number of results to return in a single call.
	MaxResults *int64 `min:"1" type:"integer"`

	// The token to use to retrieve the next page of results. The first call does
	// not contain any tokens.
	NextToken *string `min:"1" type:"string"`

	// The identifier for the organization under which the group exists.
	//
	// OrganizationId is a required field
	OrganizationId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ListGroupMembersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListGroupMembersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListGroupMembersInput"}

	if s.GroupId == nil {
		invalidParams.Add(aws.NewErrParamRequired("GroupId"))
	}
	if s.GroupId != nil && len(*s.GroupId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("GroupId", 12))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if s.OrganizationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("OrganizationId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/workmail-2017-10-01/ListGroupMembersResponse
type ListGroupMembersOutput struct {
	_ struct{} `type:"structure"`

	// The members associated to the group.
	Members []Member `type:"list"`

	// The token to use to retrieve the next page of results. The first call does
	// not contain any tokens.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListGroupMembersOutput) String() string {
	return awsutil.Prettify(s)
}

const opListGroupMembers = "ListGroupMembers"

// ListGroupMembersRequest returns a request value for making API operation for
// Amazon WorkMail.
//
// Returns an overview of the members of a group. Users and groups can be members
// of a group.
//
//    // Example sending a request using ListGroupMembersRequest.
//    req := client.ListGroupMembersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workmail-2017-10-01/ListGroupMembers
func (c *Client) ListGroupMembersRequest(input *ListGroupMembersInput) ListGroupMembersRequest {
	op := &aws.Operation{
		Name:       opListGroupMembers,
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
		input = &ListGroupMembersInput{}
	}

	req := c.newRequest(op, input, &ListGroupMembersOutput{})
	return ListGroupMembersRequest{Request: req, Input: input, Copy: c.ListGroupMembersRequest}
}

// ListGroupMembersRequest is the request type for the
// ListGroupMembers API operation.
type ListGroupMembersRequest struct {
	*aws.Request
	Input *ListGroupMembersInput
	Copy  func(*ListGroupMembersInput) ListGroupMembersRequest
}

// Send marshals and sends the ListGroupMembers API request.
func (r ListGroupMembersRequest) Send(ctx context.Context) (*ListGroupMembersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListGroupMembersResponse{
		ListGroupMembersOutput: r.Request.Data.(*ListGroupMembersOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListGroupMembersRequestPaginator returns a paginator for ListGroupMembers.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListGroupMembersRequest(input)
//   p := workmail.NewListGroupMembersRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListGroupMembersPaginator(req ListGroupMembersRequest) ListGroupMembersPaginator {
	return ListGroupMembersPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListGroupMembersInput
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

// ListGroupMembersPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListGroupMembersPaginator struct {
	aws.Pager
}

func (p *ListGroupMembersPaginator) CurrentPage() *ListGroupMembersOutput {
	return p.Pager.CurrentPage().(*ListGroupMembersOutput)
}

// ListGroupMembersResponse is the response type for the
// ListGroupMembers API operation.
type ListGroupMembersResponse struct {
	*ListGroupMembersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListGroupMembers request.
func (r *ListGroupMembersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
