// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticbeanstalk

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// A list of validation messages for a specified configuration template.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticbeanstalk-2010-12-01/ValidateConfigurationSettingsMessage
type ValidateConfigurationSettingsInput struct {
	_ struct{} `type:"structure"`

	// The name of the application that the configuration template or environment
	// belongs to.
	//
	// ApplicationName is a required field
	ApplicationName *string `min:"1" type:"string" required:"true"`

	// The name of the environment to validate the settings against.
	//
	// Condition: You cannot specify both this and a configuration template name.
	EnvironmentName *string `min:"4" type:"string"`

	// A list of the options and desired values to evaluate.
	//
	// OptionSettings is a required field
	OptionSettings []ConfigurationOptionSetting `type:"list" required:"true"`

	// The name of the configuration template to validate the settings against.
	//
	// Condition: You cannot specify both this and an environment name.
	TemplateName *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ValidateConfigurationSettingsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ValidateConfigurationSettingsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ValidateConfigurationSettingsInput"}

	if s.ApplicationName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationName"))
	}
	if s.ApplicationName != nil && len(*s.ApplicationName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ApplicationName", 1))
	}
	if s.EnvironmentName != nil && len(*s.EnvironmentName) < 4 {
		invalidParams.Add(aws.NewErrParamMinLen("EnvironmentName", 4))
	}

	if s.OptionSettings == nil {
		invalidParams.Add(aws.NewErrParamRequired("OptionSettings"))
	}
	if s.TemplateName != nil && len(*s.TemplateName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TemplateName", 1))
	}
	if s.OptionSettings != nil {
		for i, v := range s.OptionSettings {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "OptionSettings", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Provides a list of validation messages.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticbeanstalk-2010-12-01/ConfigurationSettingsValidationMessages
type ValidateConfigurationSettingsOutput struct {
	_ struct{} `type:"structure"`

	// A list of ValidationMessage.
	Messages []ValidationMessage `type:"list"`
}

// String returns the string representation
func (s ValidateConfigurationSettingsOutput) String() string {
	return awsutil.Prettify(s)
}

const opValidateConfigurationSettings = "ValidateConfigurationSettings"

// ValidateConfigurationSettingsRequest returns a request value for making API operation for
// AWS Elastic Beanstalk.
//
// Takes a set of configuration settings and either a configuration template
// or environment, and determines whether those values are valid.
//
// This action returns a list of messages indicating any errors or warnings
// associated with the selection of option values.
//
//    // Example sending a request using ValidateConfigurationSettingsRequest.
//    req := client.ValidateConfigurationSettingsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticbeanstalk-2010-12-01/ValidateConfigurationSettings
func (c *Client) ValidateConfigurationSettingsRequest(input *ValidateConfigurationSettingsInput) ValidateConfigurationSettingsRequest {
	op := &aws.Operation{
		Name:       opValidateConfigurationSettings,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ValidateConfigurationSettingsInput{}
	}

	req := c.newRequest(op, input, &ValidateConfigurationSettingsOutput{})
	return ValidateConfigurationSettingsRequest{Request: req, Input: input, Copy: c.ValidateConfigurationSettingsRequest}
}

// ValidateConfigurationSettingsRequest is the request type for the
// ValidateConfigurationSettings API operation.
type ValidateConfigurationSettingsRequest struct {
	*aws.Request
	Input *ValidateConfigurationSettingsInput
	Copy  func(*ValidateConfigurationSettingsInput) ValidateConfigurationSettingsRequest
}

// Send marshals and sends the ValidateConfigurationSettings API request.
func (r ValidateConfigurationSettingsRequest) Send(ctx context.Context) (*ValidateConfigurationSettingsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ValidateConfigurationSettingsResponse{
		ValidateConfigurationSettingsOutput: r.Request.Data.(*ValidateConfigurationSettingsOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ValidateConfigurationSettingsResponse is the response type for the
// ValidateConfigurationSettings API operation.
type ValidateConfigurationSettingsResponse struct {
	*ValidateConfigurationSettingsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ValidateConfigurationSettings request.
func (r *ValidateConfigurationSettingsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
