// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package redshift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/CreateClusterParameterGroupMessage
type CreateClusterParameterGroupInput struct {
	_ struct{} `type:"structure"`

	// A description of the parameter group.
	//
	// Description is a required field
	Description *string `type:"string" required:"true"`

	// The Amazon Redshift engine version to which the cluster parameter group applies.
	// The cluster engine version determines the set of parameters.
	//
	// To get a list of valid parameter group family names, you can call DescribeClusterParameterGroups.
	// By default, Amazon Redshift returns a list of all the parameter groups that
	// are owned by your AWS account, including the default parameter groups for
	// each Amazon Redshift engine version. The parameter group family names associated
	// with the default parameter groups provide you the valid values. For example,
	// a valid family name is "redshift-1.0".
	//
	// ParameterGroupFamily is a required field
	ParameterGroupFamily *string `type:"string" required:"true"`

	// The name of the cluster parameter group.
	//
	// Constraints:
	//
	//    * Must be 1 to 255 alphanumeric characters or hyphens
	//
	//    * First character must be a letter.
	//
	//    * Cannot end with a hyphen or contain two consecutive hyphens.
	//
	//    * Must be unique withing your AWS account.
	//
	// This value is stored as a lower-case string.
	//
	// ParameterGroupName is a required field
	ParameterGroupName *string `type:"string" required:"true"`

	// A list of tag instances.
	Tags []Tag `locationNameList:"Tag" type:"list"`
}

// String returns the string representation
func (s CreateClusterParameterGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateClusterParameterGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateClusterParameterGroupInput"}

	if s.Description == nil {
		invalidParams.Add(aws.NewErrParamRequired("Description"))
	}

	if s.ParameterGroupFamily == nil {
		invalidParams.Add(aws.NewErrParamRequired("ParameterGroupFamily"))
	}

	if s.ParameterGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ParameterGroupName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/CreateClusterParameterGroupResult
type CreateClusterParameterGroupOutput struct {
	_ struct{} `type:"structure"`

	// Describes a parameter group.
	ClusterParameterGroup *ClusterParameterGroup `type:"structure"`
}

// String returns the string representation
func (s CreateClusterParameterGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateClusterParameterGroup = "CreateClusterParameterGroup"

// CreateClusterParameterGroupRequest returns a request value for making API operation for
// Amazon Redshift.
//
// Creates an Amazon Redshift parameter group.
//
// Creating parameter groups is independent of creating clusters. You can associate
// a cluster with a parameter group when you create the cluster. You can also
// associate an existing cluster with a parameter group after the cluster is
// created by using ModifyCluster.
//
// Parameters in the parameter group define specific behavior that applies to
// the databases you create on the cluster. For more information about parameters
// and parameter groups, go to Amazon Redshift Parameter Groups (https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-parameter-groups.html)
// in the Amazon Redshift Cluster Management Guide.
//
//    // Example sending a request using CreateClusterParameterGroupRequest.
//    req := client.CreateClusterParameterGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/CreateClusterParameterGroup
func (c *Client) CreateClusterParameterGroupRequest(input *CreateClusterParameterGroupInput) CreateClusterParameterGroupRequest {
	op := &aws.Operation{
		Name:       opCreateClusterParameterGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateClusterParameterGroupInput{}
	}

	req := c.newRequest(op, input, &CreateClusterParameterGroupOutput{})
	return CreateClusterParameterGroupRequest{Request: req, Input: input, Copy: c.CreateClusterParameterGroupRequest}
}

// CreateClusterParameterGroupRequest is the request type for the
// CreateClusterParameterGroup API operation.
type CreateClusterParameterGroupRequest struct {
	*aws.Request
	Input *CreateClusterParameterGroupInput
	Copy  func(*CreateClusterParameterGroupInput) CreateClusterParameterGroupRequest
}

// Send marshals and sends the CreateClusterParameterGroup API request.
func (r CreateClusterParameterGroupRequest) Send(ctx context.Context) (*CreateClusterParameterGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateClusterParameterGroupResponse{
		CreateClusterParameterGroupOutput: r.Request.Data.(*CreateClusterParameterGroupOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateClusterParameterGroupResponse is the response type for the
// CreateClusterParameterGroup API operation.
type CreateClusterParameterGroupResponse struct {
	*CreateClusterParameterGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateClusterParameterGroup request.
func (r *CreateClusterParameterGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
