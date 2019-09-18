// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicediscovery

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicediscovery-2017-03-14/GetNamespaceRequest
type GetNamespaceInput struct {
	_ struct{} `type:"structure"`

	// The ID of the namespace that you want to get information about.
	//
	// Id is a required field
	Id *string `type:"string" required:"true"`
}

// String returns the string representation
func (s GetNamespaceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetNamespaceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetNamespaceInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicediscovery-2017-03-14/GetNamespaceResponse
type GetNamespaceOutput struct {
	_ struct{} `type:"structure"`

	// A complex type that contains information about the specified namespace.
	Namespace *Namespace `type:"structure"`
}

// String returns the string representation
func (s GetNamespaceOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetNamespace = "GetNamespace"

// GetNamespaceRequest returns a request value for making API operation for
// AWS Cloud Map.
//
// Gets information about a namespace.
//
//    // Example sending a request using GetNamespaceRequest.
//    req := client.GetNamespaceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicediscovery-2017-03-14/GetNamespace
func (c *Client) GetNamespaceRequest(input *GetNamespaceInput) GetNamespaceRequest {
	op := &aws.Operation{
		Name:       opGetNamespace,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetNamespaceInput{}
	}

	req := c.newRequest(op, input, &GetNamespaceOutput{})
	return GetNamespaceRequest{Request: req, Input: input, Copy: c.GetNamespaceRequest}
}

// GetNamespaceRequest is the request type for the
// GetNamespace API operation.
type GetNamespaceRequest struct {
	*aws.Request
	Input *GetNamespaceInput
	Copy  func(*GetNamespaceInput) GetNamespaceRequest
}

// Send marshals and sends the GetNamespace API request.
func (r GetNamespaceRequest) Send(ctx context.Context) (*GetNamespaceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetNamespaceResponse{
		GetNamespaceOutput: r.Request.Data.(*GetNamespaceOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetNamespaceResponse is the response type for the
// GetNamespace API operation.
type GetNamespaceResponse struct {
	*GetNamespaceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetNamespace request.
func (r *GetNamespaceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
