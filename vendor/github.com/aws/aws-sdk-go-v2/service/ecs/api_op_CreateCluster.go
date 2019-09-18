// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ecs-2014-11-13/CreateClusterRequest
type CreateClusterInput struct {
	_ struct{} `type:"structure"`

	// The name of your cluster. If you do not specify a name for your cluster,
	// you create a cluster named default. Up to 255 letters (uppercase and lowercase),
	// numbers, and hyphens are allowed.
	ClusterName *string `locationName:"clusterName" type:"string"`

	// The setting to use when creating a cluster. This parameter is used to enable
	// CloudWatch Container Insights for a cluster. If this value is specified,
	// it will override the containerInsights value set with PutAccountSetting or
	// PutAccountSettingDefault.
	Settings []ClusterSetting `locationName:"settings" type:"list"`

	// The metadata that you apply to the cluster to help you categorize and organize
	// them. Each tag consists of a key and an optional value, both of which you
	// define.
	//
	// The following basic restrictions apply to tags:
	//
	//    * Maximum number of tags per resource - 50
	//
	//    * For each resource, each tag key must be unique, and each tag key can
	//    have only one value.
	//
	//    * Maximum key length - 128 Unicode characters in UTF-8
	//
	//    * Maximum value length - 256 Unicode characters in UTF-8
	//
	//    * If your tagging schema is used across multiple services and resources,
	//    remember that other services may have restrictions on allowed characters.
	//    Generally allowed characters are: letters, numbers, and spaces representable
	//    in UTF-8, and the following characters: + - = . _ : / @.
	//
	//    * Tag keys and values are case-sensitive.
	//
	//    * Do not use aws:, AWS:, or any upper or lowercase combination of such
	//    as a prefix for either keys or values as it is reserved for AWS use. You
	//    cannot edit or delete tag keys or values with this prefix. Tags with this
	//    prefix do not count against your tags per resource limit.
	Tags []Tag `locationName:"tags" type:"list"`
}

// String returns the string representation
func (s CreateClusterInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateClusterInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateClusterInput"}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ecs-2014-11-13/CreateClusterResponse
type CreateClusterOutput struct {
	_ struct{} `type:"structure"`

	// The full description of your new cluster.
	Cluster *Cluster `locationName:"cluster" type:"structure"`
}

// String returns the string representation
func (s CreateClusterOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateCluster = "CreateCluster"

// CreateClusterRequest returns a request value for making API operation for
// Amazon EC2 Container Service.
//
// Creates a new Amazon ECS cluster. By default, your account receives a default
// cluster when you launch your first container instance. However, you can create
// your own cluster with a unique name with the CreateCluster action.
//
// When you call the CreateCluster API operation, Amazon ECS attempts to create
// the service-linked role for your account so that required resources in other
// AWS services can be managed on your behalf. However, if the IAM user that
// makes the call does not have permissions to create the service-linked role,
// it is not created. For more information, see Using Service-Linked Roles for
// Amazon ECS (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using-service-linked-roles.html)
// in the Amazon Elastic Container Service Developer Guide.
//
//    // Example sending a request using CreateClusterRequest.
//    req := client.CreateClusterRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ecs-2014-11-13/CreateCluster
func (c *Client) CreateClusterRequest(input *CreateClusterInput) CreateClusterRequest {
	op := &aws.Operation{
		Name:       opCreateCluster,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateClusterInput{}
	}

	req := c.newRequest(op, input, &CreateClusterOutput{})
	return CreateClusterRequest{Request: req, Input: input, Copy: c.CreateClusterRequest}
}

// CreateClusterRequest is the request type for the
// CreateCluster API operation.
type CreateClusterRequest struct {
	*aws.Request
	Input *CreateClusterInput
	Copy  func(*CreateClusterInput) CreateClusterRequest
}

// Send marshals and sends the CreateCluster API request.
func (r CreateClusterRequest) Send(ctx context.Context) (*CreateClusterResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateClusterResponse{
		CreateClusterOutput: r.Request.Data.(*CreateClusterOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateClusterResponse is the response type for the
// CreateCluster API operation.
type CreateClusterResponse struct {
	*CreateClusterOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateCluster request.
func (r *CreateClusterResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
