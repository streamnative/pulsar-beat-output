// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package applicationinsights

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/application-insights-2018-11-25/UpdateComponentConfigurationRequest
type UpdateComponentConfigurationInput struct {
	_ struct{} `type:"structure"`

	// The configuration settings of the component. The value is the escaped JSON
	// of the configuration. For more information about the JSON format, see Working
	// with JSON (https://docs.aws.amazon.com/sdk-for-javascript/v2/developer-guide/working-with-json.html).
	// You can send a request to DescribeComponentConfigurationRecommendation to
	// see the recommended configuration for a component.
	ComponentConfiguration *string `type:"string"`

	// The name of the component.
	//
	// ComponentName is a required field
	ComponentName *string `type:"string" required:"true"`

	// Indicates whether the application component is monitored.
	Monitor *bool `type:"boolean"`

	// The name of the resource group.
	//
	// ResourceGroupName is a required field
	ResourceGroupName *string `type:"string" required:"true"`

	// The tier of the application component. Supported tiers include DOT_NET_WORKER,
	// DOT_NET_WEB, SQL_SERVER, and DEFAULT.
	Tier *string `type:"string"`
}

// String returns the string representation
func (s UpdateComponentConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateComponentConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateComponentConfigurationInput"}

	if s.ComponentName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ComponentName"))
	}

	if s.ResourceGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceGroupName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/application-insights-2018-11-25/UpdateComponentConfigurationResponse
type UpdateComponentConfigurationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateComponentConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateComponentConfiguration = "UpdateComponentConfiguration"

// UpdateComponentConfigurationRequest returns a request value for making API operation for
// Amazon CloudWatch Application Insights.
//
// Updates the monitoring configurations for the component. The configuration
// input parameter is an escaped JSON of the configuration and should match
// the schema of what is returned by DescribeComponentConfigurationRecommendation.
//
//    // Example sending a request using UpdateComponentConfigurationRequest.
//    req := client.UpdateComponentConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/application-insights-2018-11-25/UpdateComponentConfiguration
func (c *Client) UpdateComponentConfigurationRequest(input *UpdateComponentConfigurationInput) UpdateComponentConfigurationRequest {
	op := &aws.Operation{
		Name:       opUpdateComponentConfiguration,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateComponentConfigurationInput{}
	}

	req := c.newRequest(op, input, &UpdateComponentConfigurationOutput{})
	return UpdateComponentConfigurationRequest{Request: req, Input: input, Copy: c.UpdateComponentConfigurationRequest}
}

// UpdateComponentConfigurationRequest is the request type for the
// UpdateComponentConfiguration API operation.
type UpdateComponentConfigurationRequest struct {
	*aws.Request
	Input *UpdateComponentConfigurationInput
	Copy  func(*UpdateComponentConfigurationInput) UpdateComponentConfigurationRequest
}

// Send marshals and sends the UpdateComponentConfiguration API request.
func (r UpdateComponentConfigurationRequest) Send(ctx context.Context) (*UpdateComponentConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateComponentConfigurationResponse{
		UpdateComponentConfigurationOutput: r.Request.Data.(*UpdateComponentConfigurationOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateComponentConfigurationResponse is the response type for the
// UpdateComponentConfiguration API operation.
type UpdateComponentConfigurationResponse struct {
	*UpdateComponentConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateComponentConfiguration request.
func (r *UpdateComponentConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
