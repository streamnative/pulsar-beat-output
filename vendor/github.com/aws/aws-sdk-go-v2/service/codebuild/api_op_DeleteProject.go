// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codebuild

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/codebuild-2016-10-06/DeleteProjectInput
type DeleteProjectInput struct {
	_ struct{} `type:"structure"`

	// The name of the build project.
	//
	// Name is a required field
	Name *string `locationName:"name" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteProjectInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteProjectInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteProjectInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/codebuild-2016-10-06/DeleteProjectOutput
type DeleteProjectOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteProjectOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteProject = "DeleteProject"

// DeleteProjectRequest returns a request value for making API operation for
// AWS CodeBuild.
//
// Deletes a build project.
//
//    // Example sending a request using DeleteProjectRequest.
//    req := client.DeleteProjectRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codebuild-2016-10-06/DeleteProject
func (c *Client) DeleteProjectRequest(input *DeleteProjectInput) DeleteProjectRequest {
	op := &aws.Operation{
		Name:       opDeleteProject,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteProjectInput{}
	}

	req := c.newRequest(op, input, &DeleteProjectOutput{})
	return DeleteProjectRequest{Request: req, Input: input, Copy: c.DeleteProjectRequest}
}

// DeleteProjectRequest is the request type for the
// DeleteProject API operation.
type DeleteProjectRequest struct {
	*aws.Request
	Input *DeleteProjectInput
	Copy  func(*DeleteProjectInput) DeleteProjectRequest
}

// Send marshals and sends the DeleteProject API request.
func (r DeleteProjectRequest) Send(ctx context.Context) (*DeleteProjectResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteProjectResponse{
		DeleteProjectOutput: r.Request.Data.(*DeleteProjectOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteProjectResponse is the response type for the
// DeleteProject API operation.
type DeleteProjectResponse struct {
	*DeleteProjectOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteProject request.
func (r *DeleteProjectResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
