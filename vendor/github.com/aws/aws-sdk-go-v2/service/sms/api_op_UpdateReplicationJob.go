// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sms

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/sms-2016-10-24/UpdateReplicationJobRequest
type UpdateReplicationJobInput struct {
	_ struct{} `type:"structure"`

	// The description of the replication job.
	Description *string `locationName:"description" type:"string"`

	// When true, the replication job produces encrypted AMIs . See also KmsKeyId
	// below.
	Encrypted *bool `locationName:"encrypted" type:"boolean"`

	// The time between consecutive replication runs, in hours.
	Frequency *int64 `locationName:"frequency" type:"integer"`

	// KMS key ID for replication jobs that produce encrypted AMIs. Can be any of
	// the following:
	//
	//    * KMS key ID
	//
	//    * KMS key alias
	//
	//    * ARN referring to KMS key ID
	//
	//    * ARN referring to KMS key alias
	//
	// If encrypted is true but a KMS key id is not specified, the customer's default
	// KMS key for EBS is used.
	KmsKeyId *string `locationName:"kmsKeyId" type:"string"`

	// The license type to be used for the AMI created by a successful replication
	// run.
	LicenseType LicenseType `locationName:"licenseType" type:"string" enum:"true"`

	// The start time of the next replication run.
	NextReplicationRunStartTime *time.Time `locationName:"nextReplicationRunStartTime" type:"timestamp"`

	// The maximum number of SMS-created AMIs to retain. The oldest will be deleted
	// once the maximum number is reached and a new AMI is created.
	NumberOfRecentAmisToKeep *int64 `locationName:"numberOfRecentAmisToKeep" type:"integer"`

	// The identifier of the replication job.
	//
	// ReplicationJobId is a required field
	ReplicationJobId *string `locationName:"replicationJobId" type:"string" required:"true"`

	// The name of the IAM role to be used by AWS SMS.
	RoleName *string `locationName:"roleName" type:"string"`
}

// String returns the string representation
func (s UpdateReplicationJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateReplicationJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateReplicationJobInput"}

	if s.ReplicationJobId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ReplicationJobId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/sms-2016-10-24/UpdateReplicationJobResponse
type UpdateReplicationJobOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateReplicationJobOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateReplicationJob = "UpdateReplicationJob"

// UpdateReplicationJobRequest returns a request value for making API operation for
// AWS Server Migration Service.
//
// Updates the specified settings for the specified replication job.
//
//    // Example sending a request using UpdateReplicationJobRequest.
//    req := client.UpdateReplicationJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sms-2016-10-24/UpdateReplicationJob
func (c *Client) UpdateReplicationJobRequest(input *UpdateReplicationJobInput) UpdateReplicationJobRequest {
	op := &aws.Operation{
		Name:       opUpdateReplicationJob,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateReplicationJobInput{}
	}

	req := c.newRequest(op, input, &UpdateReplicationJobOutput{})
	return UpdateReplicationJobRequest{Request: req, Input: input, Copy: c.UpdateReplicationJobRequest}
}

// UpdateReplicationJobRequest is the request type for the
// UpdateReplicationJob API operation.
type UpdateReplicationJobRequest struct {
	*aws.Request
	Input *UpdateReplicationJobInput
	Copy  func(*UpdateReplicationJobInput) UpdateReplicationJobRequest
}

// Send marshals and sends the UpdateReplicationJob API request.
func (r UpdateReplicationJobRequest) Send(ctx context.Context) (*UpdateReplicationJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateReplicationJobResponse{
		UpdateReplicationJobOutput: r.Request.Data.(*UpdateReplicationJobOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateReplicationJobResponse is the response type for the
// UpdateReplicationJob API operation.
type UpdateReplicationJobResponse struct {
	*UpdateReplicationJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateReplicationJob request.
func (r *UpdateReplicationJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
