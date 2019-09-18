// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicecatalog

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/ListPrincipalsForPortfolioInput
type ListPrincipalsForPortfolioInput struct {
	_ struct{} `type:"structure"`

	// The language code.
	//
	//    * en - English (default)
	//
	//    * jp - Japanese
	//
	//    * zh - Chinese
	AcceptLanguage *string `type:"string"`

	// The maximum number of items to return with this call.
	PageSize *int64 `type:"integer"`

	// The page token for the next set of results. To retrieve the first set of
	// results, use null.
	PageToken *string `type:"string"`

	// The portfolio identifier.
	//
	// PortfolioId is a required field
	PortfolioId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s ListPrincipalsForPortfolioInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListPrincipalsForPortfolioInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListPrincipalsForPortfolioInput"}

	if s.PortfolioId == nil {
		invalidParams.Add(aws.NewErrParamRequired("PortfolioId"))
	}
	if s.PortfolioId != nil && len(*s.PortfolioId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PortfolioId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/ListPrincipalsForPortfolioOutput
type ListPrincipalsForPortfolioOutput struct {
	_ struct{} `type:"structure"`

	// The page token to use to retrieve the next set of results. If there are no
	// additional results, this value is null.
	NextPageToken *string `type:"string"`

	// The IAM principals (users or roles) associated with the portfolio.
	Principals []Principal `type:"list"`
}

// String returns the string representation
func (s ListPrincipalsForPortfolioOutput) String() string {
	return awsutil.Prettify(s)
}

const opListPrincipalsForPortfolio = "ListPrincipalsForPortfolio"

// ListPrincipalsForPortfolioRequest returns a request value for making API operation for
// AWS Service Catalog.
//
// Lists all principal ARNs associated with the specified portfolio.
//
//    // Example sending a request using ListPrincipalsForPortfolioRequest.
//    req := client.ListPrincipalsForPortfolioRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/ListPrincipalsForPortfolio
func (c *Client) ListPrincipalsForPortfolioRequest(input *ListPrincipalsForPortfolioInput) ListPrincipalsForPortfolioRequest {
	op := &aws.Operation{
		Name:       opListPrincipalsForPortfolio,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"PageToken"},
			OutputTokens:    []string{"NextPageToken"},
			LimitToken:      "PageSize",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListPrincipalsForPortfolioInput{}
	}

	req := c.newRequest(op, input, &ListPrincipalsForPortfolioOutput{})
	return ListPrincipalsForPortfolioRequest{Request: req, Input: input, Copy: c.ListPrincipalsForPortfolioRequest}
}

// ListPrincipalsForPortfolioRequest is the request type for the
// ListPrincipalsForPortfolio API operation.
type ListPrincipalsForPortfolioRequest struct {
	*aws.Request
	Input *ListPrincipalsForPortfolioInput
	Copy  func(*ListPrincipalsForPortfolioInput) ListPrincipalsForPortfolioRequest
}

// Send marshals and sends the ListPrincipalsForPortfolio API request.
func (r ListPrincipalsForPortfolioRequest) Send(ctx context.Context) (*ListPrincipalsForPortfolioResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListPrincipalsForPortfolioResponse{
		ListPrincipalsForPortfolioOutput: r.Request.Data.(*ListPrincipalsForPortfolioOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListPrincipalsForPortfolioRequestPaginator returns a paginator for ListPrincipalsForPortfolio.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListPrincipalsForPortfolioRequest(input)
//   p := servicecatalog.NewListPrincipalsForPortfolioRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListPrincipalsForPortfolioPaginator(req ListPrincipalsForPortfolioRequest) ListPrincipalsForPortfolioPaginator {
	return ListPrincipalsForPortfolioPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListPrincipalsForPortfolioInput
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

// ListPrincipalsForPortfolioPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListPrincipalsForPortfolioPaginator struct {
	aws.Pager
}

func (p *ListPrincipalsForPortfolioPaginator) CurrentPage() *ListPrincipalsForPortfolioOutput {
	return p.Pager.CurrentPage().(*ListPrincipalsForPortfolioOutput)
}

// ListPrincipalsForPortfolioResponse is the response type for the
// ListPrincipalsForPortfolio API operation.
type ListPrincipalsForPortfolioResponse struct {
	*ListPrincipalsForPortfolioOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListPrincipalsForPortfolio request.
func (r *ListPrincipalsForPortfolioResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
