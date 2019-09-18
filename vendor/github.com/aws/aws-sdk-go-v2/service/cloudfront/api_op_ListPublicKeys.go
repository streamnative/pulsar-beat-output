// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudfront

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudfront-2019-03-26/ListPublicKeysRequest
type ListPublicKeysInput struct {
	_ struct{} `type:"structure"`

	// Use this when paginating results to indicate where to begin in your list
	// of public keys. The results include public keys in the list that occur after
	// the marker. To get the next page of results, set the Marker to the value
	// of the NextMarker from the current page's response (which is also the ID
	// of the last public key on that page).
	Marker *string `location:"querystring" locationName:"Marker" type:"string"`

	// The maximum number of public keys you want in the response body.
	MaxItems *int64 `location:"querystring" locationName:"MaxItems" type:"integer"`
}

// String returns the string representation
func (s ListPublicKeysInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListPublicKeysInput) MarshalFields(e protocol.FieldEncoder) error {

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

// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudfront-2019-03-26/ListPublicKeysResult
type ListPublicKeysOutput struct {
	_ struct{} `type:"structure" payload:"PublicKeyList"`

	// Returns a list of all public keys that have been added to CloudFront for
	// this account.
	PublicKeyList *PublicKeyList `type:"structure"`
}

// String returns the string representation
func (s ListPublicKeysOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListPublicKeysOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.PublicKeyList != nil {
		v := s.PublicKeyList

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "PublicKeyList", v, metadata)
	}
	return nil
}

const opListPublicKeys = "ListPublicKeys2019_03_26"

// ListPublicKeysRequest returns a request value for making API operation for
// Amazon CloudFront.
//
// List all public keys that have been added to CloudFront for this account.
//
//    // Example sending a request using ListPublicKeysRequest.
//    req := client.ListPublicKeysRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudfront-2019-03-26/ListPublicKeys
func (c *Client) ListPublicKeysRequest(input *ListPublicKeysInput) ListPublicKeysRequest {
	op := &aws.Operation{
		Name:       opListPublicKeys,
		HTTPMethod: "GET",
		HTTPPath:   "/2019-03-26/public-key",
	}

	if input == nil {
		input = &ListPublicKeysInput{}
	}

	req := c.newRequest(op, input, &ListPublicKeysOutput{})
	return ListPublicKeysRequest{Request: req, Input: input, Copy: c.ListPublicKeysRequest}
}

// ListPublicKeysRequest is the request type for the
// ListPublicKeys API operation.
type ListPublicKeysRequest struct {
	*aws.Request
	Input *ListPublicKeysInput
	Copy  func(*ListPublicKeysInput) ListPublicKeysRequest
}

// Send marshals and sends the ListPublicKeys API request.
func (r ListPublicKeysRequest) Send(ctx context.Context) (*ListPublicKeysResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListPublicKeysResponse{
		ListPublicKeysOutput: r.Request.Data.(*ListPublicKeysOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListPublicKeysResponse is the response type for the
// ListPublicKeys API operation.
type ListPublicKeysResponse struct {
	*ListPublicKeysOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListPublicKeys request.
func (r *ListPublicKeysResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
