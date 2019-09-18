// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codestar

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/codestar-2017-04-19/DeleteUserProfileRequest
type DeleteUserProfileInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the user to delete from AWS CodeStar.
	//
	// UserArn is a required field
	UserArn *string `locationName:"userArn" min:"32" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteUserProfileInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteUserProfileInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteUserProfileInput"}

	if s.UserArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserArn"))
	}
	if s.UserArn != nil && len(*s.UserArn) < 32 {
		invalidParams.Add(aws.NewErrParamMinLen("UserArn", 32))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/codestar-2017-04-19/DeleteUserProfileResult
type DeleteUserProfileOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the user deleted from AWS CodeStar.
	//
	// UserArn is a required field
	UserArn *string `locationName:"userArn" min:"32" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteUserProfileOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteUserProfile = "DeleteUserProfile"

// DeleteUserProfileRequest returns a request value for making API operation for
// AWS CodeStar.
//
// Deletes a user profile in AWS CodeStar, including all personal preference
// data associated with that profile, such as display name and email address.
// It does not delete the history of that user, for example the history of commits
// made by that user.
//
//    // Example sending a request using DeleteUserProfileRequest.
//    req := client.DeleteUserProfileRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codestar-2017-04-19/DeleteUserProfile
func (c *Client) DeleteUserProfileRequest(input *DeleteUserProfileInput) DeleteUserProfileRequest {
	op := &aws.Operation{
		Name:       opDeleteUserProfile,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteUserProfileInput{}
	}

	req := c.newRequest(op, input, &DeleteUserProfileOutput{})
	return DeleteUserProfileRequest{Request: req, Input: input, Copy: c.DeleteUserProfileRequest}
}

// DeleteUserProfileRequest is the request type for the
// DeleteUserProfile API operation.
type DeleteUserProfileRequest struct {
	*aws.Request
	Input *DeleteUserProfileInput
	Copy  func(*DeleteUserProfileInput) DeleteUserProfileRequest
}

// Send marshals and sends the DeleteUserProfile API request.
func (r DeleteUserProfileRequest) Send(ctx context.Context) (*DeleteUserProfileResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteUserProfileResponse{
		DeleteUserProfileOutput: r.Request.Data.(*DeleteUserProfileOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteUserProfileResponse is the response type for the
// DeleteUserProfile API operation.
type DeleteUserProfileResponse struct {
	*DeleteUserProfileOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteUserProfile request.
func (r *DeleteUserProfileResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
