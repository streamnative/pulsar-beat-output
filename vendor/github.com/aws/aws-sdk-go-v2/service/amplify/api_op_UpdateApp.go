// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package amplify

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Request structure for update App request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/amplify-2017-07-25/UpdateAppRequest
type UpdateAppInput struct {
	_ struct{} `type:"structure"`

	// Unique Id for an Amplify App.
	//
	// AppId is a required field
	AppId *string `location:"uri" locationName:"appId" min:"1" type:"string" required:"true"`

	// Automated branch creation config for the Amplify App.
	AutoBranchCreationConfig *AutoBranchCreationConfig `locationName:"autoBranchCreationConfig" type:"structure"`

	// Automated branch creation glob patterns for the Amplify App.
	AutoBranchCreationPatterns []string `locationName:"autoBranchCreationPatterns" type:"list"`

	// Basic Authorization credentials for an Amplify App.
	BasicAuthCredentials *string `locationName:"basicAuthCredentials" type:"string"`

	// BuildSpec for an Amplify App.
	BuildSpec *string `locationName:"buildSpec" min:"1" type:"string"`

	// Custom redirect / rewrite rules for an Amplify App.
	CustomRules []CustomRule `locationName:"customRules" type:"list"`

	// Description for an Amplify App.
	Description *string `locationName:"description" type:"string"`

	// Enables automated branch creation for the Amplify App.
	EnableAutoBranchCreation *bool `locationName:"enableAutoBranchCreation" type:"boolean"`

	// Enables Basic Authorization for an Amplify App.
	EnableBasicAuth *bool `locationName:"enableBasicAuth" type:"boolean"`

	// Enables branch auto-building for an Amplify App.
	EnableBranchAutoBuild *bool `locationName:"enableBranchAutoBuild" type:"boolean"`

	// Environment Variables for an Amplify App.
	EnvironmentVariables map[string]string `locationName:"environmentVariables" type:"map"`

	// IAM service role for an Amplify App.
	IamServiceRoleArn *string `locationName:"iamServiceRoleArn" min:"1" type:"string"`

	// Name for an Amplify App.
	Name *string `locationName:"name" min:"1" type:"string"`

	// Platform for an Amplify App.
	Platform Platform `locationName:"platform" type:"string" enum:"true"`
}

// String returns the string representation
func (s UpdateAppInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateAppInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateAppInput"}

	if s.AppId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AppId"))
	}
	if s.AppId != nil && len(*s.AppId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AppId", 1))
	}
	if s.BuildSpec != nil && len(*s.BuildSpec) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("BuildSpec", 1))
	}
	if s.IamServiceRoleArn != nil && len(*s.IamServiceRoleArn) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("IamServiceRoleArn", 1))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}
	if s.AutoBranchCreationConfig != nil {
		if err := s.AutoBranchCreationConfig.Validate(); err != nil {
			invalidParams.AddNested("AutoBranchCreationConfig", err.(aws.ErrInvalidParams))
		}
	}
	if s.CustomRules != nil {
		for i, v := range s.CustomRules {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "CustomRules", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateAppInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AutoBranchCreationConfig != nil {
		v := s.AutoBranchCreationConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "autoBranchCreationConfig", v, metadata)
	}
	if s.AutoBranchCreationPatterns != nil {
		v := s.AutoBranchCreationPatterns

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "autoBranchCreationPatterns", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.BasicAuthCredentials != nil {
		v := *s.BasicAuthCredentials

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "basicAuthCredentials", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.BuildSpec != nil {
		v := *s.BuildSpec

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "buildSpec", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CustomRules != nil {
		v := s.CustomRules

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "customRules", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.EnableAutoBranchCreation != nil {
		v := *s.EnableAutoBranchCreation

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "enableAutoBranchCreation", protocol.BoolValue(v), metadata)
	}
	if s.EnableBasicAuth != nil {
		v := *s.EnableBasicAuth

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "enableBasicAuth", protocol.BoolValue(v), metadata)
	}
	if s.EnableBranchAutoBuild != nil {
		v := *s.EnableBranchAutoBuild

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "enableBranchAutoBuild", protocol.BoolValue(v), metadata)
	}
	if s.EnvironmentVariables != nil {
		v := s.EnvironmentVariables

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "environmentVariables", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.IamServiceRoleArn != nil {
		v := *s.IamServiceRoleArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "iamServiceRoleArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Platform) > 0 {
		v := s.Platform

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "platform", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.AppId != nil {
		v := *s.AppId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "appId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Result structure for an Amplify App update request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/amplify-2017-07-25/UpdateAppResult
type UpdateAppOutput struct {
	_ struct{} `type:"structure"`

	// App structure for the updated App.
	//
	// App is a required field
	App *App `locationName:"app" type:"structure" required:"true"`
}

// String returns the string representation
func (s UpdateAppOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateAppOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.App != nil {
		v := s.App

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "app", v, metadata)
	}
	return nil
}

const opUpdateApp = "UpdateApp"

// UpdateAppRequest returns a request value for making API operation for
// AWS Amplify.
//
// Updates an existing Amplify App.
//
//    // Example sending a request using UpdateAppRequest.
//    req := client.UpdateAppRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/amplify-2017-07-25/UpdateApp
func (c *Client) UpdateAppRequest(input *UpdateAppInput) UpdateAppRequest {
	op := &aws.Operation{
		Name:       opUpdateApp,
		HTTPMethod: "POST",
		HTTPPath:   "/apps/{appId}",
	}

	if input == nil {
		input = &UpdateAppInput{}
	}

	req := c.newRequest(op, input, &UpdateAppOutput{})
	return UpdateAppRequest{Request: req, Input: input, Copy: c.UpdateAppRequest}
}

// UpdateAppRequest is the request type for the
// UpdateApp API operation.
type UpdateAppRequest struct {
	*aws.Request
	Input *UpdateAppInput
	Copy  func(*UpdateAppInput) UpdateAppRequest
}

// Send marshals and sends the UpdateApp API request.
func (r UpdateAppRequest) Send(ctx context.Context) (*UpdateAppResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateAppResponse{
		UpdateAppOutput: r.Request.Data.(*UpdateAppOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateAppResponse is the response type for the
// UpdateApp API operation.
type UpdateAppResponse struct {
	*UpdateAppOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateApp request.
func (r *UpdateAppResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
