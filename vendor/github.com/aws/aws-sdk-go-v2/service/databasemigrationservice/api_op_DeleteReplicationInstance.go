// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package databasemigrationservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/dms-2016-01-01/DeleteReplicationInstanceMessage
type DeleteReplicationInstanceInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the replication instance to be deleted.
	//
	// ReplicationInstanceArn is a required field
	ReplicationInstanceArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteReplicationInstanceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteReplicationInstanceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteReplicationInstanceInput"}

	if s.ReplicationInstanceArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ReplicationInstanceArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/dms-2016-01-01/DeleteReplicationInstanceResponse
type DeleteReplicationInstanceOutput struct {
	_ struct{} `type:"structure"`

	// The replication instance that was deleted.
	ReplicationInstance *ReplicationInstance `type:"structure"`
}

// String returns the string representation
func (s DeleteReplicationInstanceOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteReplicationInstance = "DeleteReplicationInstance"

// DeleteReplicationInstanceRequest returns a request value for making API operation for
// AWS Database Migration Service.
//
// Deletes the specified replication instance.
//
// You must delete any migration tasks that are associated with the replication
// instance before you can delete it.
//
//    // Example sending a request using DeleteReplicationInstanceRequest.
//    req := client.DeleteReplicationInstanceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dms-2016-01-01/DeleteReplicationInstance
func (c *Client) DeleteReplicationInstanceRequest(input *DeleteReplicationInstanceInput) DeleteReplicationInstanceRequest {
	op := &aws.Operation{
		Name:       opDeleteReplicationInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteReplicationInstanceInput{}
	}

	req := c.newRequest(op, input, &DeleteReplicationInstanceOutput{})
	return DeleteReplicationInstanceRequest{Request: req, Input: input, Copy: c.DeleteReplicationInstanceRequest}
}

// DeleteReplicationInstanceRequest is the request type for the
// DeleteReplicationInstance API operation.
type DeleteReplicationInstanceRequest struct {
	*aws.Request
	Input *DeleteReplicationInstanceInput
	Copy  func(*DeleteReplicationInstanceInput) DeleteReplicationInstanceRequest
}

// Send marshals and sends the DeleteReplicationInstance API request.
func (r DeleteReplicationInstanceRequest) Send(ctx context.Context) (*DeleteReplicationInstanceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteReplicationInstanceResponse{
		DeleteReplicationInstanceOutput: r.Request.Data.(*DeleteReplicationInstanceOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteReplicationInstanceResponse is the response type for the
// DeleteReplicationInstance API operation.
type DeleteReplicationInstanceResponse struct {
	*DeleteReplicationInstanceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteReplicationInstance request.
func (r *DeleteReplicationInstanceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
