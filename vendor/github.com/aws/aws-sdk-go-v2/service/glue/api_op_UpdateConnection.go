// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/UpdateConnectionRequest
type UpdateConnectionInput struct {
	_ struct{} `type:"structure"`

	// The ID of the Data Catalog in which the connection resides. If none is provided,
	// the AWS account ID is used by default.
	CatalogId *string `min:"1" type:"string"`

	// A ConnectionInput object that redefines the connection in question.
	//
	// ConnectionInput is a required field
	ConnectionInput *ConnectionInput `type:"structure" required:"true"`

	// The name of the connection definition to update.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateConnectionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateConnectionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateConnectionInput"}
	if s.CatalogId != nil && len(*s.CatalogId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CatalogId", 1))
	}

	if s.ConnectionInput == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConnectionInput"))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}
	if s.ConnectionInput != nil {
		if err := s.ConnectionInput.Validate(); err != nil {
			invalidParams.AddNested("ConnectionInput", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/UpdateConnectionResponse
type UpdateConnectionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateConnectionOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateConnection = "UpdateConnection"

// UpdateConnectionRequest returns a request value for making API operation for
// AWS Glue.
//
// Updates a connection definition in the Data Catalog.
//
//    // Example sending a request using UpdateConnectionRequest.
//    req := client.UpdateConnectionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/UpdateConnection
func (c *Client) UpdateConnectionRequest(input *UpdateConnectionInput) UpdateConnectionRequest {
	op := &aws.Operation{
		Name:       opUpdateConnection,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateConnectionInput{}
	}

	req := c.newRequest(op, input, &UpdateConnectionOutput{})
	return UpdateConnectionRequest{Request: req, Input: input, Copy: c.UpdateConnectionRequest}
}

// UpdateConnectionRequest is the request type for the
// UpdateConnection API operation.
type UpdateConnectionRequest struct {
	*aws.Request
	Input *UpdateConnectionInput
	Copy  func(*UpdateConnectionInput) UpdateConnectionRequest
}

// Send marshals and sends the UpdateConnection API request.
func (r UpdateConnectionRequest) Send(ctx context.Context) (*UpdateConnectionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateConnectionResponse{
		UpdateConnectionOutput: r.Request.Data.(*UpdateConnectionOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateConnectionResponse is the response type for the
// UpdateConnection API operation.
type UpdateConnectionResponse struct {
	*UpdateConnectionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateConnection request.
func (r *UpdateConnectionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
