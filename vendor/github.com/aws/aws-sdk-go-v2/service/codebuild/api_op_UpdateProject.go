// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codebuild

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/codebuild-2016-10-06/UpdateProjectInput
type UpdateProjectInput struct {
	_ struct{} `type:"structure"`

	// Information to be changed about the build output artifacts for the build
	// project.
	Artifacts *ProjectArtifacts `locationName:"artifacts" type:"structure"`

	// Set this to true to generate a publicly accessible URL for your project's
	// build badge.
	BadgeEnabled *bool `locationName:"badgeEnabled" type:"boolean"`

	// Stores recently used information so that it can be quickly accessed at a
	// later time.
	Cache *ProjectCache `locationName:"cache" type:"structure"`

	// A new or replacement description of the build project.
	Description *string `locationName:"description" type:"string"`

	// The AWS Key Management Service (AWS KMS) customer master key (CMK) to be
	// used for encrypting the build output artifacts.
	//
	// You can use a cross-account KMS key to encrypt the build output artifacts
	// if your service role has permission to that key.
	//
	// You can specify either the Amazon Resource Name (ARN) of the CMK or, if available,
	// the CMK's alias (using the format alias/alias-name ).
	EncryptionKey *string `locationName:"encryptionKey" min:"1" type:"string"`

	// Information to be changed about the build environment for the build project.
	Environment *ProjectEnvironment `locationName:"environment" type:"structure"`

	// Information about logs for the build project. A project can create logs in
	// Amazon CloudWatch Logs, logs in an S3 bucket, or both.
	LogsConfig *LogsConfig `locationName:"logsConfig" type:"structure"`

	// The name of the build project.
	//
	// You cannot change a build project's name.
	//
	// Name is a required field
	Name *string `locationName:"name" min:"1" type:"string" required:"true"`

	// The number of minutes a build is allowed to be queued before it times out.
	QueuedTimeoutInMinutes *int64 `locationName:"queuedTimeoutInMinutes" min:"5" type:"integer"`

	// An array of ProjectSource objects.
	SecondaryArtifacts []ProjectArtifacts `locationName:"secondaryArtifacts" type:"list"`

	// An array of ProjectSourceVersion objects. If secondarySourceVersions is specified
	// at the build level, then they take over these secondarySourceVersions (at
	// the project level).
	SecondarySourceVersions []ProjectSourceVersion `locationName:"secondarySourceVersions" type:"list"`

	// An array of ProjectSource objects.
	SecondarySources []ProjectSource `locationName:"secondarySources" type:"list"`

	// The replacement ARN of the AWS Identity and Access Management (IAM) role
	// that enables AWS CodeBuild to interact with dependent AWS services on behalf
	// of the AWS account.
	ServiceRole *string `locationName:"serviceRole" min:"1" type:"string"`

	// Information to be changed about the build input source code for the build
	// project.
	Source *ProjectSource `locationName:"source" type:"structure"`

	// A version of the build input to be built for this project. If not specified,
	// the latest version is used. If specified, it must be one of:
	//
	//    * For AWS CodeCommit: the commit ID to use.
	//
	//    * For GitHub: the commit ID, pull request ID, branch name, or tag name
	//    that corresponds to the version of the source code you want to build.
	//    If a pull request ID is specified, it must use the format pr/pull-request-ID
	//    (for example pr/25). If a branch name is specified, the branch's HEAD
	//    commit ID is used. If not specified, the default branch's HEAD commit
	//    ID is used.
	//
	//    * For Bitbucket: the commit ID, branch name, or tag name that corresponds
	//    to the version of the source code you want to build. If a branch name
	//    is specified, the branch's HEAD commit ID is used. If not specified, the
	//    default branch's HEAD commit ID is used.
	//
	//    * For Amazon Simple Storage Service (Amazon S3): the version ID of the
	//    object that represents the build input ZIP file to use.
	//
	// If sourceVersion is specified at the build level, then that version takes
	// precedence over this sourceVersion (at the project level).
	//
	// For more information, see Source Version Sample with CodeBuild (https://docs.aws.amazon.com/codebuild/latest/userguide/sample-source-version.html)
	// in the AWS CodeBuild User Guide.
	SourceVersion *string `locationName:"sourceVersion" type:"string"`

	// The replacement set of tags for this build project.
	//
	// These tags are available for use by AWS services that support AWS CodeBuild
	// build project tags.
	Tags []Tag `locationName:"tags" type:"list"`

	// The replacement value in minutes, from 5 to 480 (8 hours), for AWS CodeBuild
	// to wait before timing out any related build that did not get marked as completed.
	TimeoutInMinutes *int64 `locationName:"timeoutInMinutes" min:"5" type:"integer"`

	// VpcConfig enables AWS CodeBuild to access resources in an Amazon VPC.
	VpcConfig *VpcConfig `locationName:"vpcConfig" type:"structure"`
}

// String returns the string representation
func (s UpdateProjectInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateProjectInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateProjectInput"}
	if s.EncryptionKey != nil && len(*s.EncryptionKey) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("EncryptionKey", 1))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}
	if s.QueuedTimeoutInMinutes != nil && *s.QueuedTimeoutInMinutes < 5 {
		invalidParams.Add(aws.NewErrParamMinValue("QueuedTimeoutInMinutes", 5))
	}
	if s.ServiceRole != nil && len(*s.ServiceRole) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ServiceRole", 1))
	}
	if s.TimeoutInMinutes != nil && *s.TimeoutInMinutes < 5 {
		invalidParams.Add(aws.NewErrParamMinValue("TimeoutInMinutes", 5))
	}
	if s.Artifacts != nil {
		if err := s.Artifacts.Validate(); err != nil {
			invalidParams.AddNested("Artifacts", err.(aws.ErrInvalidParams))
		}
	}
	if s.Cache != nil {
		if err := s.Cache.Validate(); err != nil {
			invalidParams.AddNested("Cache", err.(aws.ErrInvalidParams))
		}
	}
	if s.Environment != nil {
		if err := s.Environment.Validate(); err != nil {
			invalidParams.AddNested("Environment", err.(aws.ErrInvalidParams))
		}
	}
	if s.LogsConfig != nil {
		if err := s.LogsConfig.Validate(); err != nil {
			invalidParams.AddNested("LogsConfig", err.(aws.ErrInvalidParams))
		}
	}
	if s.SecondaryArtifacts != nil {
		for i, v := range s.SecondaryArtifacts {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "SecondaryArtifacts", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.SecondarySourceVersions != nil {
		for i, v := range s.SecondarySourceVersions {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "SecondarySourceVersions", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.SecondarySources != nil {
		for i, v := range s.SecondarySources {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "SecondarySources", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.Source != nil {
		if err := s.Source.Validate(); err != nil {
			invalidParams.AddNested("Source", err.(aws.ErrInvalidParams))
		}
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.VpcConfig != nil {
		if err := s.VpcConfig.Validate(); err != nil {
			invalidParams.AddNested("VpcConfig", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/codebuild-2016-10-06/UpdateProjectOutput
type UpdateProjectOutput struct {
	_ struct{} `type:"structure"`

	// Information about the build project that was changed.
	Project *Project `locationName:"project" type:"structure"`
}

// String returns the string representation
func (s UpdateProjectOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateProject = "UpdateProject"

// UpdateProjectRequest returns a request value for making API operation for
// AWS CodeBuild.
//
// Changes the settings of a build project.
//
//    // Example sending a request using UpdateProjectRequest.
//    req := client.UpdateProjectRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codebuild-2016-10-06/UpdateProject
func (c *Client) UpdateProjectRequest(input *UpdateProjectInput) UpdateProjectRequest {
	op := &aws.Operation{
		Name:       opUpdateProject,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateProjectInput{}
	}

	req := c.newRequest(op, input, &UpdateProjectOutput{})
	return UpdateProjectRequest{Request: req, Input: input, Copy: c.UpdateProjectRequest}
}

// UpdateProjectRequest is the request type for the
// UpdateProject API operation.
type UpdateProjectRequest struct {
	*aws.Request
	Input *UpdateProjectInput
	Copy  func(*UpdateProjectInput) UpdateProjectRequest
}

// Send marshals and sends the UpdateProject API request.
func (r UpdateProjectRequest) Send(ctx context.Context) (*UpdateProjectResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateProjectResponse{
		UpdateProjectOutput: r.Request.Data.(*UpdateProjectOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateProjectResponse is the response type for the
// UpdateProject API operation.
type UpdateProjectResponse struct {
	*UpdateProjectOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateProject request.
func (r *UpdateProjectResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
