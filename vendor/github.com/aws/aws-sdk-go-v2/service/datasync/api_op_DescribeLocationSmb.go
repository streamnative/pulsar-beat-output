// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package datasync

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// DescribeLocationSmbRequest
// Please also see https://docs.aws.amazon.com/goto/WebAPI/datasync-2018-11-09/DescribeLocationSmbRequest
type DescribeLocationSmbInput struct {
	_ struct{} `type:"structure"`

	// The Amazon resource Name (ARN) of the SMB location to describe.
	//
	// LocationArn is a required field
	LocationArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeLocationSmbInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeLocationSmbInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeLocationSmbInput"}

	if s.LocationArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("LocationArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// DescribeLocationSmbResponse
// Please also see https://docs.aws.amazon.com/goto/WebAPI/datasync-2018-11-09/DescribeLocationSmbResponse
type DescribeLocationSmbOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the source SMB file system location that
	// is created.
	AgentArns []string `min:"1" type:"list"`

	// The time that the SMB location was created.
	CreationTime *time.Time `type:"timestamp"`

	// The name of the domain that the SMB server belongs to.
	Domain *string `type:"string"`

	// The Amazon resource Name (ARN) of the SMB location that was described.
	LocationArn *string `type:"string"`

	// The URL of the source SBM location that was described.
	LocationUri *string `type:"string"`

	// The mount options that are available for DataSync to use to access an SMB
	// location.
	MountOptions *SmbMountOptions `type:"structure"`

	// The user who is logged on the SMB server.
	User *string `type:"string"`
}

// String returns the string representation
func (s DescribeLocationSmbOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeLocationSmb = "DescribeLocationSmb"

// DescribeLocationSmbRequest returns a request value for making API operation for
// AWS DataSync.
//
// Returns metadata, such as the path and user information about a SMB location.
//
//    // Example sending a request using DescribeLocationSmbRequest.
//    req := client.DescribeLocationSmbRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/datasync-2018-11-09/DescribeLocationSmb
func (c *Client) DescribeLocationSmbRequest(input *DescribeLocationSmbInput) DescribeLocationSmbRequest {
	op := &aws.Operation{
		Name:       opDescribeLocationSmb,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeLocationSmbInput{}
	}

	req := c.newRequest(op, input, &DescribeLocationSmbOutput{})
	return DescribeLocationSmbRequest{Request: req, Input: input, Copy: c.DescribeLocationSmbRequest}
}

// DescribeLocationSmbRequest is the request type for the
// DescribeLocationSmb API operation.
type DescribeLocationSmbRequest struct {
	*aws.Request
	Input *DescribeLocationSmbInput
	Copy  func(*DescribeLocationSmbInput) DescribeLocationSmbRequest
}

// Send marshals and sends the DescribeLocationSmb API request.
func (r DescribeLocationSmbRequest) Send(ctx context.Context) (*DescribeLocationSmbResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeLocationSmbResponse{
		DescribeLocationSmbOutput: r.Request.Data.(*DescribeLocationSmbOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeLocationSmbResponse is the response type for the
// DescribeLocationSmb API operation.
type DescribeLocationSmbResponse struct {
	*DescribeLocationSmbOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeLocationSmb request.
func (r *DescribeLocationSmbResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
