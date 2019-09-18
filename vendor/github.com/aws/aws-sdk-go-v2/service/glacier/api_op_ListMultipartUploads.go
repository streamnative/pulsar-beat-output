// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glacier

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Provides options for retrieving list of in-progress multipart uploads for
// an Amazon Glacier vault.
type ListMultipartUploadsInput struct {
	_ struct{} `type:"structure"`

	// The AccountId value is the AWS account ID of the account that owns the vault.
	// You can either specify an AWS account ID or optionally a single '-' (hyphen),
	// in which case Amazon S3 Glacier uses the AWS account ID associated with the
	// credentials used to sign the request. If you use an account ID, do not include
	// any hyphens ('-') in the ID.
	//
	// AccountId is a required field
	AccountId *string `location:"uri" locationName:"accountId" type:"string" required:"true"`

	// Specifies the maximum number of uploads returned in the response body. If
	// this value is not specified, the List Uploads operation returns up to 50
	// uploads.
	Limit *string `location:"querystring" locationName:"limit" type:"string"`

	// An opaque string used for pagination. This value specifies the upload at
	// which the listing of uploads should begin. Get the marker value from a previous
	// List Uploads response. You need only include the marker if you are continuing
	// the pagination of results started in a previous List Uploads request.
	Marker *string `location:"querystring" locationName:"marker" type:"string"`

	// The name of the vault.
	//
	// VaultName is a required field
	VaultName *string `location:"uri" locationName:"vaultName" type:"string" required:"true"`
}

// String returns the string representation
func (s ListMultipartUploadsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListMultipartUploadsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListMultipartUploadsInput"}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}

	if s.VaultName == nil {
		invalidParams.Add(aws.NewErrParamRequired("VaultName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListMultipartUploadsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AccountId != nil {
		v := *s.AccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "accountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VaultName != nil {
		v := *s.VaultName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "vaultName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Limit != nil {
		v := *s.Limit

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "limit", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Marker != nil {
		v := *s.Marker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "marker", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Contains the Amazon S3 Glacier response to your request.
type ListMultipartUploadsOutput struct {
	_ struct{} `type:"structure"`

	// An opaque string that represents where to continue pagination of the results.
	// You use the marker in a new List Multipart Uploads request to obtain more
	// uploads in the list. If there are no more uploads, this value is null.
	Marker *string `type:"string"`

	// A list of in-progress multipart uploads.
	UploadsList []UploadListElement `type:"list"`
}

// String returns the string representation
func (s ListMultipartUploadsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListMultipartUploadsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Marker != nil {
		v := *s.Marker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Marker", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.UploadsList != nil {
		v := s.UploadsList

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "UploadsList", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opListMultipartUploads = "ListMultipartUploads"

// ListMultipartUploadsRequest returns a request value for making API operation for
// Amazon Glacier.
//
// This operation lists in-progress multipart uploads for the specified vault.
// An in-progress multipart upload is a multipart upload that has been initiated
// by an InitiateMultipartUpload request, but has not yet been completed or
// aborted. The list returned in the List Multipart Upload response has no guaranteed
// order.
//
// The List Multipart Uploads operation supports pagination. By default, this
// operation returns up to 50 multipart uploads in the response. You should
// always check the response for a marker at which to continue the list; if
// there are no more items the marker is null. To return a list of multipart
// uploads that begins at a specific upload, set the marker request parameter
// to the value you obtained from a previous List Multipart Upload request.
// You can also limit the number of uploads returned in the response by specifying
// the limit parameter in the request.
//
// Note the difference between this operation and listing parts (ListParts).
// The List Multipart Uploads operation lists all multipart uploads for a vault
// and does not require a multipart upload ID. The List Parts operation requires
// a multipart upload ID since parts are associated with a single upload.
//
// An AWS account has full permission to perform all operations (actions). However,
// AWS Identity and Access Management (IAM) users don't have any permissions
// by default. You must grant them explicit permission to perform specific actions.
// For more information, see Access Control Using AWS Identity and Access Management
// (IAM) (https://docs.aws.amazon.com/amazonglacier/latest/dev/using-iam-with-amazon-glacier.html).
//
// For conceptual information and the underlying REST API, see Working with
// Archives in Amazon S3 Glacier (https://docs.aws.amazon.com/amazonglacier/latest/dev/working-with-archives.html)
// and List Multipart Uploads (https://docs.aws.amazon.com/amazonglacier/latest/dev/api-multipart-list-uploads.html)
// in the Amazon Glacier Developer Guide.
//
//    // Example sending a request using ListMultipartUploadsRequest.
//    req := client.ListMultipartUploadsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) ListMultipartUploadsRequest(input *ListMultipartUploadsInput) ListMultipartUploadsRequest {
	op := &aws.Operation{
		Name:       opListMultipartUploads,
		HTTPMethod: "GET",
		HTTPPath:   "/{accountId}/vaults/{vaultName}/multipart-uploads",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"marker"},
			OutputTokens:    []string{"Marker"},
			LimitToken:      "limit",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListMultipartUploadsInput{}
	}

	req := c.newRequest(op, input, &ListMultipartUploadsOutput{})
	return ListMultipartUploadsRequest{Request: req, Input: input, Copy: c.ListMultipartUploadsRequest}
}

// ListMultipartUploadsRequest is the request type for the
// ListMultipartUploads API operation.
type ListMultipartUploadsRequest struct {
	*aws.Request
	Input *ListMultipartUploadsInput
	Copy  func(*ListMultipartUploadsInput) ListMultipartUploadsRequest
}

// Send marshals and sends the ListMultipartUploads API request.
func (r ListMultipartUploadsRequest) Send(ctx context.Context) (*ListMultipartUploadsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListMultipartUploadsResponse{
		ListMultipartUploadsOutput: r.Request.Data.(*ListMultipartUploadsOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListMultipartUploadsRequestPaginator returns a paginator for ListMultipartUploads.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListMultipartUploadsRequest(input)
//   p := glacier.NewListMultipartUploadsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListMultipartUploadsPaginator(req ListMultipartUploadsRequest) ListMultipartUploadsPaginator {
	return ListMultipartUploadsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListMultipartUploadsInput
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

// ListMultipartUploadsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListMultipartUploadsPaginator struct {
	aws.Pager
}

func (p *ListMultipartUploadsPaginator) CurrentPage() *ListMultipartUploadsOutput {
	return p.Pager.CurrentPage().(*ListMultipartUploadsOutput)
}

// ListMultipartUploadsResponse is the response type for the
// ListMultipartUploads API operation.
type ListMultipartUploadsResponse struct {
	*ListMultipartUploadsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListMultipartUploads request.
func (r *ListMultipartUploadsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
