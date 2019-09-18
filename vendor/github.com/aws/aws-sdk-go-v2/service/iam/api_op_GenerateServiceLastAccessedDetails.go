// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/GenerateServiceLastAccessedDetailsRequest
type GenerateServiceLastAccessedDetailsInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the IAM resource (user, group, role, or managed policy) used to
	// generate information about when the resource was last used in an attempt
	// to access an AWS service.
	//
	// Arn is a required field
	Arn *string `min:"20" type:"string" required:"true"`
}

// String returns the string representation
func (s GenerateServiceLastAccessedDetailsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GenerateServiceLastAccessedDetailsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GenerateServiceLastAccessedDetailsInput"}

	if s.Arn == nil {
		invalidParams.Add(aws.NewErrParamRequired("Arn"))
	}
	if s.Arn != nil && len(*s.Arn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("Arn", 20))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/GenerateServiceLastAccessedDetailsResponse
type GenerateServiceLastAccessedDetailsOutput struct {
	_ struct{} `type:"structure"`

	// The job ID that you can use in the GetServiceLastAccessedDetails or GetServiceLastAccessedDetailsWithEntities
	// operations.
	JobId *string `min:"36" type:"string"`
}

// String returns the string representation
func (s GenerateServiceLastAccessedDetailsOutput) String() string {
	return awsutil.Prettify(s)
}

const opGenerateServiceLastAccessedDetails = "GenerateServiceLastAccessedDetails"

// GenerateServiceLastAccessedDetailsRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Generates a report that includes details about when an IAM resource (user,
// group, role, or policy) was last used in an attempt to access AWS services.
// Recent activity usually appears within four hours. IAM reports activity for
// the last 365 days, or less if your Region began supporting this feature within
// the last year. For more information, see Regions Where Data Is Tracked (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_access-advisor.html#access-advisor_tracking-period).
//
// The service last accessed data includes all attempts to access an AWS API,
// not just the successful ones. This includes all attempts that were made using
// the AWS Management Console, the AWS API through any of the SDKs, or any of
// the command line tools. An unexpected entry in the service last accessed
// data does not mean that your account has been compromised, because the request
// might have been denied. Refer to your CloudTrail logs as the authoritative
// source for information about all API calls and whether they were successful
// or denied access. For more information, see Logging IAM Events with CloudTrail
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/cloudtrail-integration.html)
// in the IAM User Guide.
//
// The GenerateServiceLastAccessedDetails operation returns a JobId. Use this
// parameter in the following operations to retrieve the following details from
// your report:
//
//    * GetServiceLastAccessedDetails – Use this operation for users, groups,
//    roles, or policies to list every AWS service that the resource could access
//    using permissions policies. For each service, the response includes information
//    about the most recent access attempt.
//
//    * GetServiceLastAccessedDetailsWithEntities – Use this operation for
//    groups and policies to list information about the associated entities
//    (users or roles) that attempted to access a specific AWS service.
//
// To check the status of the GenerateServiceLastAccessedDetails request, use
// the JobId parameter in the same operations and test the JobStatus response
// parameter.
//
// For additional information about the permissions policies that allow an identity
// (user, group, or role) to access specific services, use the ListPoliciesGrantingServiceAccess
// operation.
//
// Service last accessed data does not use other policy types when determining
// whether a resource could access a service. These other policy types include
// resource-based policies, access control lists, AWS Organizations policies,
// IAM permissions boundaries, and AWS STS assume role policies. It only applies
// permissions policy logic. For more about the evaluation of policy types,
// see Evaluating Policies (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_evaluation-logic.html#policy-eval-basics)
// in the IAM User Guide.
//
// For more information about service last accessed data, see Reducing Policy
// Scope by Viewing User Activity (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_access-advisor.html)
// in the IAM User Guide.
//
//    // Example sending a request using GenerateServiceLastAccessedDetailsRequest.
//    req := client.GenerateServiceLastAccessedDetailsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/GenerateServiceLastAccessedDetails
func (c *Client) GenerateServiceLastAccessedDetailsRequest(input *GenerateServiceLastAccessedDetailsInput) GenerateServiceLastAccessedDetailsRequest {
	op := &aws.Operation{
		Name:       opGenerateServiceLastAccessedDetails,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GenerateServiceLastAccessedDetailsInput{}
	}

	req := c.newRequest(op, input, &GenerateServiceLastAccessedDetailsOutput{})
	return GenerateServiceLastAccessedDetailsRequest{Request: req, Input: input, Copy: c.GenerateServiceLastAccessedDetailsRequest}
}

// GenerateServiceLastAccessedDetailsRequest is the request type for the
// GenerateServiceLastAccessedDetails API operation.
type GenerateServiceLastAccessedDetailsRequest struct {
	*aws.Request
	Input *GenerateServiceLastAccessedDetailsInput
	Copy  func(*GenerateServiceLastAccessedDetailsInput) GenerateServiceLastAccessedDetailsRequest
}

// Send marshals and sends the GenerateServiceLastAccessedDetails API request.
func (r GenerateServiceLastAccessedDetailsRequest) Send(ctx context.Context) (*GenerateServiceLastAccessedDetailsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GenerateServiceLastAccessedDetailsResponse{
		GenerateServiceLastAccessedDetailsOutput: r.Request.Data.(*GenerateServiceLastAccessedDetailsOutput),
		response:                                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GenerateServiceLastAccessedDetailsResponse is the response type for the
// GenerateServiceLastAccessedDetails API operation.
type GenerateServiceLastAccessedDetailsResponse struct {
	*GenerateServiceLastAccessedDetailsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GenerateServiceLastAccessedDetails request.
func (r *GenerateServiceLastAccessedDetailsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
