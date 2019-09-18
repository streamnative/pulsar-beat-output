// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/DeleteWorkteamRequest
type DeleteWorkteamInput struct {
	_ struct{} `type:"structure"`

	// The name of the work team to delete.
	//
	// WorkteamName is a required field
	WorkteamName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteWorkteamInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteWorkteamInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteWorkteamInput"}

	if s.WorkteamName == nil {
		invalidParams.Add(aws.NewErrParamRequired("WorkteamName"))
	}
	if s.WorkteamName != nil && len(*s.WorkteamName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("WorkteamName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/DeleteWorkteamResponse
type DeleteWorkteamOutput struct {
	_ struct{} `type:"structure"`

	// Returns true if the work team was successfully deleted; otherwise, returns
	// false.
	//
	// Success is a required field
	Success *bool `type:"boolean" required:"true"`
}

// String returns the string representation
func (s DeleteWorkteamOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteWorkteam = "DeleteWorkteam"

// DeleteWorkteamRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// Deletes an existing work team. This operation can't be undone.
//
//    // Example sending a request using DeleteWorkteamRequest.
//    req := client.DeleteWorkteamRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/DeleteWorkteam
func (c *Client) DeleteWorkteamRequest(input *DeleteWorkteamInput) DeleteWorkteamRequest {
	op := &aws.Operation{
		Name:       opDeleteWorkteam,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteWorkteamInput{}
	}

	req := c.newRequest(op, input, &DeleteWorkteamOutput{})
	return DeleteWorkteamRequest{Request: req, Input: input, Copy: c.DeleteWorkteamRequest}
}

// DeleteWorkteamRequest is the request type for the
// DeleteWorkteam API operation.
type DeleteWorkteamRequest struct {
	*aws.Request
	Input *DeleteWorkteamInput
	Copy  func(*DeleteWorkteamInput) DeleteWorkteamRequest
}

// Send marshals and sends the DeleteWorkteam API request.
func (r DeleteWorkteamRequest) Send(ctx context.Context) (*DeleteWorkteamResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteWorkteamResponse{
		DeleteWorkteamOutput: r.Request.Data.(*DeleteWorkteamOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteWorkteamResponse is the response type for the
// DeleteWorkteam API operation.
type DeleteWorkteamResponse struct {
	*DeleteWorkteamOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteWorkteam request.
func (r *DeleteWorkteamResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
