// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package simpledb

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
)

type BatchDeleteAttributesInput struct {
	_ struct{} `type:"structure"`

	// The name of the domain in which the attributes are being deleted.
	//
	// DomainName is a required field
	DomainName *string `type:"string" required:"true"`

	// A list of items on which to perform the operation.
	//
	// Items is a required field
	Items []DeletableItem `locationNameList:"Item" type:"list" flattened:"true" required:"true"`
}

// String returns the string representation
func (s BatchDeleteAttributesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BatchDeleteAttributesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "BatchDeleteAttributesInput"}

	if s.DomainName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainName"))
	}

	if s.Items == nil {
		invalidParams.Add(aws.NewErrParamRequired("Items"))
	}
	if s.Items != nil {
		for i, v := range s.Items {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Items", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type BatchDeleteAttributesOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s BatchDeleteAttributesOutput) String() string {
	return awsutil.Prettify(s)
}

const opBatchDeleteAttributes = "BatchDeleteAttributes"

// BatchDeleteAttributesRequest returns a request value for making API operation for
// Amazon SimpleDB.
//
// Performs multiple DeleteAttributes operations in a single call, which reduces
// round trips and latencies. This enables Amazon SimpleDB to optimize requests,
// which generally yields better throughput.
//
// If you specify BatchDeleteAttributes without attributes or values, all the
// attributes for the item are deleted.
//
// BatchDeleteAttributes is an idempotent operation; running it multiple times
// on the same item or attribute doesn't result in an error.
//
// The BatchDeleteAttributes operation succeeds or fails in its entirety. There
// are no partial deletes. You can execute multiple BatchDeleteAttributes operations
// and other operations in parallel. However, large numbers of concurrent BatchDeleteAttributes
// calls can result in Service Unavailable (503) responses.
//
// This operation is vulnerable to exceeding the maximum URL size when making
// a REST request using the HTTP GET method.
//
// This operation does not support conditions using Expected.X.Name, Expected.X.Value,
// or Expected.X.Exists.
//
// The following limitations are enforced for this operation:
//    * 1 MB request size
//
//    * 25 item limit per BatchDeleteAttributes operation
//
//    // Example sending a request using BatchDeleteAttributesRequest.
//    req := client.BatchDeleteAttributesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) BatchDeleteAttributesRequest(input *BatchDeleteAttributesInput) BatchDeleteAttributesRequest {
	op := &aws.Operation{
		Name:       opBatchDeleteAttributes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &BatchDeleteAttributesInput{}
	}

	req := c.newRequest(op, input, &BatchDeleteAttributesOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return BatchDeleteAttributesRequest{Request: req, Input: input, Copy: c.BatchDeleteAttributesRequest}
}

// BatchDeleteAttributesRequest is the request type for the
// BatchDeleteAttributes API operation.
type BatchDeleteAttributesRequest struct {
	*aws.Request
	Input *BatchDeleteAttributesInput
	Copy  func(*BatchDeleteAttributesInput) BatchDeleteAttributesRequest
}

// Send marshals and sends the BatchDeleteAttributes API request.
func (r BatchDeleteAttributesRequest) Send(ctx context.Context) (*BatchDeleteAttributesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &BatchDeleteAttributesResponse{
		BatchDeleteAttributesOutput: r.Request.Data.(*BatchDeleteAttributesOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// BatchDeleteAttributesResponse is the response type for the
// BatchDeleteAttributes API operation.
type BatchDeleteAttributesResponse struct {
	*BatchDeleteAttributesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// BatchDeleteAttributes request.
func (r *BatchDeleteAttributesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
