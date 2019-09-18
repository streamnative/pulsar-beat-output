// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restxml"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/PutBucketAclRequest
type PutBucketAclInput struct {
	_ struct{} `type:"structure" payload:"AccessControlPolicy"`

	// The canned ACL to apply to the bucket.
	ACL BucketCannedACL `location:"header" locationName:"x-amz-acl" type:"string" enum:"true"`

	// Contains the elements that set the ACL permissions for an object per grantee.
	AccessControlPolicy *AccessControlPolicy `locationName:"AccessControlPolicy" type:"structure" xmlURI:"http://s3.amazonaws.com/doc/2006-03-01/"`

	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`

	// Allows grantee the read, write, read ACP, and write ACP permissions on the
	// bucket.
	GrantFullControl *string `location:"header" locationName:"x-amz-grant-full-control" type:"string"`

	// Allows grantee to list the objects in the bucket.
	GrantRead *string `location:"header" locationName:"x-amz-grant-read" type:"string"`

	// Allows grantee to read the bucket ACL.
	GrantReadACP *string `location:"header" locationName:"x-amz-grant-read-acp" type:"string"`

	// Allows grantee to create, overwrite, and delete any object in the bucket.
	GrantWrite *string `location:"header" locationName:"x-amz-grant-write" type:"string"`

	// Allows grantee to write the ACL for the applicable bucket.
	GrantWriteACP *string `location:"header" locationName:"x-amz-grant-write-acp" type:"string"`
}

// String returns the string representation
func (s PutBucketAclInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutBucketAclInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutBucketAclInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}
	if s.AccessControlPolicy != nil {
		if err := s.AccessControlPolicy.Validate(); err != nil {
			invalidParams.AddNested("AccessControlPolicy", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *PutBucketAclInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutBucketAclInput) MarshalFields(e protocol.FieldEncoder) error {

	if len(s.ACL) > 0 {
		v := s.ACL

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-acl", v, metadata)
	}
	if s.GrantFullControl != nil {
		v := *s.GrantFullControl

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-grant-full-control", protocol.StringValue(v), metadata)
	}
	if s.GrantRead != nil {
		v := *s.GrantRead

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-grant-read", protocol.StringValue(v), metadata)
	}
	if s.GrantReadACP != nil {
		v := *s.GrantReadACP

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-grant-read-acp", protocol.StringValue(v), metadata)
	}
	if s.GrantWrite != nil {
		v := *s.GrantWrite

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-grant-write", protocol.StringValue(v), metadata)
	}
	if s.GrantWriteACP != nil {
		v := *s.GrantWriteACP

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-grant-write-acp", protocol.StringValue(v), metadata)
	}
	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Bucket", protocol.StringValue(v), metadata)
	}
	if s.AccessControlPolicy != nil {
		v := s.AccessControlPolicy

		metadata := protocol.Metadata{XMLNamespaceURI: "http://s3.amazonaws.com/doc/2006-03-01/"}
		e.SetFields(protocol.PayloadTarget, "AccessControlPolicy", v, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/PutBucketAclOutput
type PutBucketAclOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutBucketAclOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutBucketAclOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opPutBucketAcl = "PutBucketAcl"

// PutBucketAclRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Sets the permissions on a bucket using access control lists (ACL).
//
//    // Example sending a request using PutBucketAclRequest.
//    req := client.PutBucketAclRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/PutBucketAcl
func (c *Client) PutBucketAclRequest(input *PutBucketAclInput) PutBucketAclRequest {
	op := &aws.Operation{
		Name:       opPutBucketAcl,
		HTTPMethod: "PUT",
		HTTPPath:   "/{Bucket}?acl",
	}

	if input == nil {
		input = &PutBucketAclInput{}
	}

	req := c.newRequest(op, input, &PutBucketAclOutput{})
	req.Handlers.Unmarshal.Remove(restxml.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return PutBucketAclRequest{Request: req, Input: input, Copy: c.PutBucketAclRequest}
}

// PutBucketAclRequest is the request type for the
// PutBucketAcl API operation.
type PutBucketAclRequest struct {
	*aws.Request
	Input *PutBucketAclInput
	Copy  func(*PutBucketAclInput) PutBucketAclRequest
}

// Send marshals and sends the PutBucketAcl API request.
func (r PutBucketAclRequest) Send(ctx context.Context) (*PutBucketAclResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutBucketAclResponse{
		PutBucketAclOutput: r.Request.Data.(*PutBucketAclOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutBucketAclResponse is the response type for the
// PutBucketAcl API operation.
type PutBucketAclResponse struct {
	*PutBucketAclOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutBucketAcl request.
func (r *PutBucketAclResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
