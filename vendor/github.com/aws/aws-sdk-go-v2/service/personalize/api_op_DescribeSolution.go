// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package personalize

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/personalize-2018-05-22/DescribeSolutionRequest
type DescribeSolutionInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the solution to describe.
	//
	// SolutionArn is a required field
	SolutionArn *string `locationName:"solutionArn" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeSolutionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeSolutionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeSolutionInput"}

	if s.SolutionArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("SolutionArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/personalize-2018-05-22/DescribeSolutionResponse
type DescribeSolutionOutput struct {
	_ struct{} `type:"structure"`

	// An object that describes the solution.
	Solution *Solution `locationName:"solution" type:"structure"`
}

// String returns the string representation
func (s DescribeSolutionOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeSolution = "DescribeSolution"

// DescribeSolutionRequest returns a request value for making API operation for
// Amazon Personalize.
//
// Describes a solution. For more information on solutions, see CreateSolution.
//
//    // Example sending a request using DescribeSolutionRequest.
//    req := client.DescribeSolutionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/personalize-2018-05-22/DescribeSolution
func (c *Client) DescribeSolutionRequest(input *DescribeSolutionInput) DescribeSolutionRequest {
	op := &aws.Operation{
		Name:       opDescribeSolution,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeSolutionInput{}
	}

	req := c.newRequest(op, input, &DescribeSolutionOutput{})
	return DescribeSolutionRequest{Request: req, Input: input, Copy: c.DescribeSolutionRequest}
}

// DescribeSolutionRequest is the request type for the
// DescribeSolution API operation.
type DescribeSolutionRequest struct {
	*aws.Request
	Input *DescribeSolutionInput
	Copy  func(*DescribeSolutionInput) DescribeSolutionRequest
}

// Send marshals and sends the DescribeSolution API request.
func (r DescribeSolutionRequest) Send(ctx context.Context) (*DescribeSolutionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeSolutionResponse{
		DescribeSolutionOutput: r.Request.Data.(*DescribeSolutionOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeSolutionResponse is the response type for the
// DescribeSolution API operation.
type DescribeSolutionResponse struct {
	*DescribeSolutionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeSolution request.
func (r *DescribeSolutionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
