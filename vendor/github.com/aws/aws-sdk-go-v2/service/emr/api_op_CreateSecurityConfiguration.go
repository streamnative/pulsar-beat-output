// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package emr

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticmapreduce-2009-03-31/CreateSecurityConfigurationInput
type CreateSecurityConfigurationInput struct {
	_ struct{} `type:"structure"`

	// The name of the security configuration.
	//
	// Name is a required field
	Name *string `type:"string" required:"true"`

	// The security configuration details in JSON format. For JSON parameters and
	// examples, see Use Security Configurations to Set Up Cluster Security (https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-security-configurations.html)
	// in the Amazon EMR Management Guide.
	//
	// SecurityConfiguration is a required field
	SecurityConfiguration *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateSecurityConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateSecurityConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateSecurityConfigurationInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if s.SecurityConfiguration == nil {
		invalidParams.Add(aws.NewErrParamRequired("SecurityConfiguration"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticmapreduce-2009-03-31/CreateSecurityConfigurationOutput
type CreateSecurityConfigurationOutput struct {
	_ struct{} `type:"structure"`

	// The date and time the security configuration was created.
	//
	// CreationDateTime is a required field
	CreationDateTime *time.Time `type:"timestamp" required:"true"`

	// The name of the security configuration.
	//
	// Name is a required field
	Name *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateSecurityConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateSecurityConfiguration = "CreateSecurityConfiguration"

// CreateSecurityConfigurationRequest returns a request value for making API operation for
// Amazon Elastic MapReduce.
//
// Creates a security configuration, which is stored in the service and can
// be specified when a cluster is created.
//
//    // Example sending a request using CreateSecurityConfigurationRequest.
//    req := client.CreateSecurityConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticmapreduce-2009-03-31/CreateSecurityConfiguration
func (c *Client) CreateSecurityConfigurationRequest(input *CreateSecurityConfigurationInput) CreateSecurityConfigurationRequest {
	op := &aws.Operation{
		Name:       opCreateSecurityConfiguration,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateSecurityConfigurationInput{}
	}

	req := c.newRequest(op, input, &CreateSecurityConfigurationOutput{})
	return CreateSecurityConfigurationRequest{Request: req, Input: input, Copy: c.CreateSecurityConfigurationRequest}
}

// CreateSecurityConfigurationRequest is the request type for the
// CreateSecurityConfiguration API operation.
type CreateSecurityConfigurationRequest struct {
	*aws.Request
	Input *CreateSecurityConfigurationInput
	Copy  func(*CreateSecurityConfigurationInput) CreateSecurityConfigurationRequest
}

// Send marshals and sends the CreateSecurityConfiguration API request.
func (r CreateSecurityConfigurationRequest) Send(ctx context.Context) (*CreateSecurityConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateSecurityConfigurationResponse{
		CreateSecurityConfigurationOutput: r.Request.Data.(*CreateSecurityConfigurationOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateSecurityConfigurationResponse is the response type for the
// CreateSecurityConfiguration API operation.
type CreateSecurityConfigurationResponse struct {
	*CreateSecurityConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateSecurityConfiguration request.
func (r *CreateSecurityConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
