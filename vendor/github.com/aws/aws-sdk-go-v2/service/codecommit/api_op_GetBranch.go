// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codecommit

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a get branch operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codecommit-2015-04-13/GetBranchInput
type GetBranchInput struct {
	_ struct{} `type:"structure"`

	// The name of the branch for which you want to retrieve information.
	BranchName *string `locationName:"branchName" min:"1" type:"string"`

	// The name of the repository that contains the branch for which you want to
	// retrieve information.
	RepositoryName *string `locationName:"repositoryName" min:"1" type:"string"`
}

// String returns the string representation
func (s GetBranchInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetBranchInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetBranchInput"}
	if s.BranchName != nil && len(*s.BranchName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("BranchName", 1))
	}
	if s.RepositoryName != nil && len(*s.RepositoryName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RepositoryName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of a get branch operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codecommit-2015-04-13/GetBranchOutput
type GetBranchOutput struct {
	_ struct{} `type:"structure"`

	// The name of the branch.
	Branch *BranchInfo `locationName:"branch" type:"structure"`
}

// String returns the string representation
func (s GetBranchOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetBranch = "GetBranch"

// GetBranchRequest returns a request value for making API operation for
// AWS CodeCommit.
//
// Returns information about a repository branch, including its name and the
// last commit ID.
//
//    // Example sending a request using GetBranchRequest.
//    req := client.GetBranchRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codecommit-2015-04-13/GetBranch
func (c *Client) GetBranchRequest(input *GetBranchInput) GetBranchRequest {
	op := &aws.Operation{
		Name:       opGetBranch,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetBranchInput{}
	}

	req := c.newRequest(op, input, &GetBranchOutput{})
	return GetBranchRequest{Request: req, Input: input, Copy: c.GetBranchRequest}
}

// GetBranchRequest is the request type for the
// GetBranch API operation.
type GetBranchRequest struct {
	*aws.Request
	Input *GetBranchInput
	Copy  func(*GetBranchInput) GetBranchRequest
}

// Send marshals and sends the GetBranch API request.
func (r GetBranchRequest) Send(ctx context.Context) (*GetBranchResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetBranchResponse{
		GetBranchOutput: r.Request.Data.(*GetBranchOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetBranchResponse is the response type for the
// GetBranch API operation.
type GetBranchResponse struct {
	*GetBranchOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetBranch request.
func (r *GetBranchResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
