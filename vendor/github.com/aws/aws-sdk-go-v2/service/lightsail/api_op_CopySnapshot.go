// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/CopySnapshotRequest
type CopySnapshotInput struct {
	_ struct{} `type:"structure"`

	// The AWS Region where the source snapshot is located.
	//
	// SourceRegion is a required field
	SourceRegion RegionName `locationName:"sourceRegion" type:"string" required:"true" enum:"true"`

	// The name of the source instance or disk snapshot to be copied.
	//
	// SourceSnapshotName is a required field
	SourceSnapshotName *string `locationName:"sourceSnapshotName" type:"string" required:"true"`

	// The name of the new instance or disk snapshot to be created as a copy.
	//
	// TargetSnapshotName is a required field
	TargetSnapshotName *string `locationName:"targetSnapshotName" type:"string" required:"true"`
}

// String returns the string representation
func (s CopySnapshotInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CopySnapshotInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CopySnapshotInput"}
	if len(s.SourceRegion) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("SourceRegion"))
	}

	if s.SourceSnapshotName == nil {
		invalidParams.Add(aws.NewErrParamRequired("SourceSnapshotName"))
	}

	if s.TargetSnapshotName == nil {
		invalidParams.Add(aws.NewErrParamRequired("TargetSnapshotName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/CopySnapshotResult
type CopySnapshotOutput struct {
	_ struct{} `type:"structure"`

	// A list of objects describing the API operation.
	Operations []Operation `locationName:"operations" type:"list"`
}

// String returns the string representation
func (s CopySnapshotOutput) String() string {
	return awsutil.Prettify(s)
}

const opCopySnapshot = "CopySnapshot"

// CopySnapshotRequest returns a request value for making API operation for
// Amazon Lightsail.
//
// Copies an instance or disk snapshot from one AWS Region to another in Amazon
// Lightsail.
//
//    // Example sending a request using CopySnapshotRequest.
//    req := client.CopySnapshotRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/CopySnapshot
func (c *Client) CopySnapshotRequest(input *CopySnapshotInput) CopySnapshotRequest {
	op := &aws.Operation{
		Name:       opCopySnapshot,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CopySnapshotInput{}
	}

	req := c.newRequest(op, input, &CopySnapshotOutput{})
	return CopySnapshotRequest{Request: req, Input: input, Copy: c.CopySnapshotRequest}
}

// CopySnapshotRequest is the request type for the
// CopySnapshot API operation.
type CopySnapshotRequest struct {
	*aws.Request
	Input *CopySnapshotInput
	Copy  func(*CopySnapshotInput) CopySnapshotRequest
}

// Send marshals and sends the CopySnapshot API request.
func (r CopySnapshotRequest) Send(ctx context.Context) (*CopySnapshotResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CopySnapshotResponse{
		CopySnapshotOutput: r.Request.Data.(*CopySnapshotOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CopySnapshotResponse is the response type for the
// CopySnapshot API operation.
type CopySnapshotResponse struct {
	*CopySnapshotOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CopySnapshot request.
func (r *CopySnapshotResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
