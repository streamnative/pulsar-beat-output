// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticbeanstalk

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticbeanstalk-2010-12-01/CreateStorageLocationInput
type CreateStorageLocationInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CreateStorageLocationInput) String() string {
	return awsutil.Prettify(s)
}

// Results of a CreateStorageLocationResult call.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticbeanstalk-2010-12-01/CreateStorageLocationResultMessage
type CreateStorageLocationOutput struct {
	_ struct{} `type:"structure"`

	// The name of the Amazon S3 bucket created.
	S3Bucket *string `type:"string"`
}

// String returns the string representation
func (s CreateStorageLocationOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateStorageLocation = "CreateStorageLocation"

// CreateStorageLocationRequest returns a request value for making API operation for
// AWS Elastic Beanstalk.
//
// Creates a bucket in Amazon S3 to store application versions, logs, and other
// files used by Elastic Beanstalk environments. The Elastic Beanstalk console
// and EB CLI call this API the first time you create an environment in a region.
// If the storage location already exists, CreateStorageLocation still returns
// the bucket name but does not create a new bucket.
//
//    // Example sending a request using CreateStorageLocationRequest.
//    req := client.CreateStorageLocationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticbeanstalk-2010-12-01/CreateStorageLocation
func (c *Client) CreateStorageLocationRequest(input *CreateStorageLocationInput) CreateStorageLocationRequest {
	op := &aws.Operation{
		Name:       opCreateStorageLocation,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateStorageLocationInput{}
	}

	req := c.newRequest(op, input, &CreateStorageLocationOutput{})
	return CreateStorageLocationRequest{Request: req, Input: input, Copy: c.CreateStorageLocationRequest}
}

// CreateStorageLocationRequest is the request type for the
// CreateStorageLocation API operation.
type CreateStorageLocationRequest struct {
	*aws.Request
	Input *CreateStorageLocationInput
	Copy  func(*CreateStorageLocationInput) CreateStorageLocationRequest
}

// Send marshals and sends the CreateStorageLocation API request.
func (r CreateStorageLocationRequest) Send(ctx context.Context) (*CreateStorageLocationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateStorageLocationResponse{
		CreateStorageLocationOutput: r.Request.Data.(*CreateStorageLocationOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateStorageLocationResponse is the response type for the
// CreateStorageLocation API operation.
type CreateStorageLocationResponse struct {
	*CreateStorageLocationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateStorageLocation request.
func (r *CreateStorageLocationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
