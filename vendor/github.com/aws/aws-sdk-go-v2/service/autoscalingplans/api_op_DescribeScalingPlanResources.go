// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package autoscalingplans

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/autoscaling-plans-2018-01-06/DescribeScalingPlanResourcesRequest
type DescribeScalingPlanResourcesInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of scalable resources to return. The value must be between
	// 1 and 50. The default value is 50.
	MaxResults *int64 `type:"integer"`

	// The token for the next set of results.
	NextToken *string `type:"string"`

	// The name of the scaling plan.
	//
	// ScalingPlanName is a required field
	ScalingPlanName *string `min:"1" type:"string" required:"true"`

	// The version number of the scaling plan.
	//
	// ScalingPlanVersion is a required field
	ScalingPlanVersion *int64 `type:"long" required:"true"`
}

// String returns the string representation
func (s DescribeScalingPlanResourcesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeScalingPlanResourcesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeScalingPlanResourcesInput"}

	if s.ScalingPlanName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ScalingPlanName"))
	}
	if s.ScalingPlanName != nil && len(*s.ScalingPlanName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ScalingPlanName", 1))
	}

	if s.ScalingPlanVersion == nil {
		invalidParams.Add(aws.NewErrParamRequired("ScalingPlanVersion"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/autoscaling-plans-2018-01-06/DescribeScalingPlanResourcesResponse
type DescribeScalingPlanResourcesOutput struct {
	_ struct{} `type:"structure"`

	// The token required to get the next set of results. This value is null if
	// there are no more results to return.
	NextToken *string `type:"string"`

	// Information about the scalable resources.
	ScalingPlanResources []ScalingPlanResource `type:"list"`
}

// String returns the string representation
func (s DescribeScalingPlanResourcesOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeScalingPlanResources = "DescribeScalingPlanResources"

// DescribeScalingPlanResourcesRequest returns a request value for making API operation for
// AWS Auto Scaling Plans.
//
// Describes the scalable resources in the specified scaling plan.
//
//    // Example sending a request using DescribeScalingPlanResourcesRequest.
//    req := client.DescribeScalingPlanResourcesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/autoscaling-plans-2018-01-06/DescribeScalingPlanResources
func (c *Client) DescribeScalingPlanResourcesRequest(input *DescribeScalingPlanResourcesInput) DescribeScalingPlanResourcesRequest {
	op := &aws.Operation{
		Name:       opDescribeScalingPlanResources,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeScalingPlanResourcesInput{}
	}

	req := c.newRequest(op, input, &DescribeScalingPlanResourcesOutput{})
	return DescribeScalingPlanResourcesRequest{Request: req, Input: input, Copy: c.DescribeScalingPlanResourcesRequest}
}

// DescribeScalingPlanResourcesRequest is the request type for the
// DescribeScalingPlanResources API operation.
type DescribeScalingPlanResourcesRequest struct {
	*aws.Request
	Input *DescribeScalingPlanResourcesInput
	Copy  func(*DescribeScalingPlanResourcesInput) DescribeScalingPlanResourcesRequest
}

// Send marshals and sends the DescribeScalingPlanResources API request.
func (r DescribeScalingPlanResourcesRequest) Send(ctx context.Context) (*DescribeScalingPlanResourcesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeScalingPlanResourcesResponse{
		DescribeScalingPlanResourcesOutput: r.Request.Data.(*DescribeScalingPlanResourcesOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeScalingPlanResourcesResponse is the response type for the
// DescribeScalingPlanResources API operation.
type DescribeScalingPlanResourcesResponse struct {
	*DescribeScalingPlanResourcesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeScalingPlanResources request.
func (r *DescribeScalingPlanResourcesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
