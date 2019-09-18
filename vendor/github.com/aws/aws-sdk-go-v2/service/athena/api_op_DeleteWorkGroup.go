// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package athena

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/athena-2017-05-18/DeleteWorkGroupInput
type DeleteWorkGroupInput struct {
	_ struct{} `type:"structure"`

	// The option to delete the workgroup and its contents even if the workgroup
	// contains any named queries.
	RecursiveDeleteOption *bool `type:"boolean"`

	// The unique name of the workgroup to delete.
	//
	// WorkGroup is a required field
	WorkGroup *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteWorkGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteWorkGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteWorkGroupInput"}

	if s.WorkGroup == nil {
		invalidParams.Add(aws.NewErrParamRequired("WorkGroup"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/athena-2017-05-18/DeleteWorkGroupOutput
type DeleteWorkGroupOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteWorkGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteWorkGroup = "DeleteWorkGroup"

// DeleteWorkGroupRequest returns a request value for making API operation for
// Amazon Athena.
//
// Deletes the workgroup with the specified name. The primary workgroup cannot
// be deleted.
//
//    // Example sending a request using DeleteWorkGroupRequest.
//    req := client.DeleteWorkGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/athena-2017-05-18/DeleteWorkGroup
func (c *Client) DeleteWorkGroupRequest(input *DeleteWorkGroupInput) DeleteWorkGroupRequest {
	op := &aws.Operation{
		Name:       opDeleteWorkGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteWorkGroupInput{}
	}

	req := c.newRequest(op, input, &DeleteWorkGroupOutput{})
	return DeleteWorkGroupRequest{Request: req, Input: input, Copy: c.DeleteWorkGroupRequest}
}

// DeleteWorkGroupRequest is the request type for the
// DeleteWorkGroup API operation.
type DeleteWorkGroupRequest struct {
	*aws.Request
	Input *DeleteWorkGroupInput
	Copy  func(*DeleteWorkGroupInput) DeleteWorkGroupRequest
}

// Send marshals and sends the DeleteWorkGroup API request.
func (r DeleteWorkGroupRequest) Send(ctx context.Context) (*DeleteWorkGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteWorkGroupResponse{
		DeleteWorkGroupOutput: r.Request.Data.(*DeleteWorkGroupOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteWorkGroupResponse is the response type for the
// DeleteWorkGroup API operation.
type DeleteWorkGroupResponse struct {
	*DeleteWorkGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteWorkGroup request.
func (r *DeleteWorkGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
