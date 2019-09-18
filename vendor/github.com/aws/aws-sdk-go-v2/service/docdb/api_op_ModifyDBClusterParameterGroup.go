// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package docdb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input to ModifyDBClusterParameterGroup.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/docdb-2014-10-31/ModifyDBClusterParameterGroupMessage
type ModifyDBClusterParameterGroupInput struct {
	_ struct{} `type:"structure"`

	// The name of the DB cluster parameter group to modify.
	//
	// DBClusterParameterGroupName is a required field
	DBClusterParameterGroupName *string `type:"string" required:"true"`

	// A list of parameters in the DB cluster parameter group to modify.
	//
	// Parameters is a required field
	Parameters []Parameter `locationNameList:"Parameter" type:"list" required:"true"`
}

// String returns the string representation
func (s ModifyDBClusterParameterGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyDBClusterParameterGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ModifyDBClusterParameterGroupInput"}

	if s.DBClusterParameterGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBClusterParameterGroupName"))
	}

	if s.Parameters == nil {
		invalidParams.Add(aws.NewErrParamRequired("Parameters"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the name of a DB cluster parameter group.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/docdb-2014-10-31/DBClusterParameterGroupNameMessage
type ModifyDBClusterParameterGroupOutput struct {
	_ struct{} `type:"structure"`

	// The name of a DB cluster parameter group.
	//
	// Constraints:
	//
	//    * Must be from 1 to 255 letters or numbers.
	//
	//    * The first character must be a letter.
	//
	//    * Cannot end with a hyphen or contain two consecutive hyphens.
	//
	// This value is stored as a lowercase string.
	DBClusterParameterGroupName *string `type:"string"`
}

// String returns the string representation
func (s ModifyDBClusterParameterGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opModifyDBClusterParameterGroup = "ModifyDBClusterParameterGroup"

// ModifyDBClusterParameterGroupRequest returns a request value for making API operation for
// Amazon DocumentDB with MongoDB compatibility.
//
// Modifies the parameters of a DB cluster parameter group. To modify more than
// one parameter, submit a list of the following: ParameterName, ParameterValue,
// and ApplyMethod. A maximum of 20 parameters can be modified in a single request.
//
// Changes to dynamic parameters are applied immediately. Changes to static
// parameters require a reboot or maintenance window before the change can take
// effect.
//
// After you create a DB cluster parameter group, you should wait at least 5
// minutes before creating your first DB cluster that uses that DB cluster parameter
// group as the default parameter group. This allows Amazon DocumentDB to fully
// complete the create action before the parameter group is used as the default
// for a new DB cluster. This step is especially important for parameters that
// are critical when creating the default database for a DB cluster, such as
// the character set for the default database defined by the character_set_database
// parameter.
//
//    // Example sending a request using ModifyDBClusterParameterGroupRequest.
//    req := client.ModifyDBClusterParameterGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/docdb-2014-10-31/ModifyDBClusterParameterGroup
func (c *Client) ModifyDBClusterParameterGroupRequest(input *ModifyDBClusterParameterGroupInput) ModifyDBClusterParameterGroupRequest {
	op := &aws.Operation{
		Name:       opModifyDBClusterParameterGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyDBClusterParameterGroupInput{}
	}

	req := c.newRequest(op, input, &ModifyDBClusterParameterGroupOutput{})
	return ModifyDBClusterParameterGroupRequest{Request: req, Input: input, Copy: c.ModifyDBClusterParameterGroupRequest}
}

// ModifyDBClusterParameterGroupRequest is the request type for the
// ModifyDBClusterParameterGroup API operation.
type ModifyDBClusterParameterGroupRequest struct {
	*aws.Request
	Input *ModifyDBClusterParameterGroupInput
	Copy  func(*ModifyDBClusterParameterGroupInput) ModifyDBClusterParameterGroupRequest
}

// Send marshals and sends the ModifyDBClusterParameterGroup API request.
func (r ModifyDBClusterParameterGroupRequest) Send(ctx context.Context) (*ModifyDBClusterParameterGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ModifyDBClusterParameterGroupResponse{
		ModifyDBClusterParameterGroupOutput: r.Request.Data.(*ModifyDBClusterParameterGroupOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ModifyDBClusterParameterGroupResponse is the response type for the
// ModifyDBClusterParameterGroup API operation.
type ModifyDBClusterParameterGroupResponse struct {
	*ModifyDBClusterParameterGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ModifyDBClusterParameterGroup request.
func (r *ModifyDBClusterParameterGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
