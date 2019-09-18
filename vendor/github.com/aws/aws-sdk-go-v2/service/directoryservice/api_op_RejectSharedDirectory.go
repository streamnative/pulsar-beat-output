// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package directoryservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ds-2015-04-16/RejectSharedDirectoryRequest
type RejectSharedDirectoryInput struct {
	_ struct{} `type:"structure"`

	// Identifier of the shared directory in the directory consumer account. This
	// identifier is different for each directory owner account.
	//
	// SharedDirectoryId is a required field
	SharedDirectoryId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s RejectSharedDirectoryInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RejectSharedDirectoryInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RejectSharedDirectoryInput"}

	if s.SharedDirectoryId == nil {
		invalidParams.Add(aws.NewErrParamRequired("SharedDirectoryId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ds-2015-04-16/RejectSharedDirectoryResult
type RejectSharedDirectoryOutput struct {
	_ struct{} `type:"structure"`

	// Identifier of the shared directory in the directory consumer account.
	SharedDirectoryId *string `type:"string"`
}

// String returns the string representation
func (s RejectSharedDirectoryOutput) String() string {
	return awsutil.Prettify(s)
}

const opRejectSharedDirectory = "RejectSharedDirectory"

// RejectSharedDirectoryRequest returns a request value for making API operation for
// AWS Directory Service.
//
// Rejects a directory sharing request that was sent from the directory owner
// account.
//
//    // Example sending a request using RejectSharedDirectoryRequest.
//    req := client.RejectSharedDirectoryRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ds-2015-04-16/RejectSharedDirectory
func (c *Client) RejectSharedDirectoryRequest(input *RejectSharedDirectoryInput) RejectSharedDirectoryRequest {
	op := &aws.Operation{
		Name:       opRejectSharedDirectory,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RejectSharedDirectoryInput{}
	}

	req := c.newRequest(op, input, &RejectSharedDirectoryOutput{})
	return RejectSharedDirectoryRequest{Request: req, Input: input, Copy: c.RejectSharedDirectoryRequest}
}

// RejectSharedDirectoryRequest is the request type for the
// RejectSharedDirectory API operation.
type RejectSharedDirectoryRequest struct {
	*aws.Request
	Input *RejectSharedDirectoryInput
	Copy  func(*RejectSharedDirectoryInput) RejectSharedDirectoryRequest
}

// Send marshals and sends the RejectSharedDirectory API request.
func (r RejectSharedDirectoryRequest) Send(ctx context.Context) (*RejectSharedDirectoryResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RejectSharedDirectoryResponse{
		RejectSharedDirectoryOutput: r.Request.Data.(*RejectSharedDirectoryOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RejectSharedDirectoryResponse is the response type for the
// RejectSharedDirectory API operation.
type RejectSharedDirectoryResponse struct {
	*RejectSharedDirectoryOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RejectSharedDirectory request.
func (r *RejectSharedDirectoryResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
