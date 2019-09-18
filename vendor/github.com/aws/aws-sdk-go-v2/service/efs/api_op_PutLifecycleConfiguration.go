// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package efs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticfilesystem-2015-02-01/PutLifecycleConfigurationRequest
type PutLifecycleConfigurationInput struct {
	_ struct{} `type:"structure"`

	// The ID of the file system for which you are creating the LifecycleConfiguration
	// object (String).
	//
	// FileSystemId is a required field
	FileSystemId *string `location:"uri" locationName:"FileSystemId" type:"string" required:"true"`

	// An array of LifecyclePolicy objects that define the file system's LifecycleConfiguration
	// object. A LifecycleConfiguration object tells lifecycle management when to
	// transition files from the Standard storage class to the Infrequent Access
	// storage class.
	//
	// LifecyclePolicies is a required field
	LifecyclePolicies []LifecyclePolicy `type:"list" required:"true"`
}

// String returns the string representation
func (s PutLifecycleConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutLifecycleConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutLifecycleConfigurationInput"}

	if s.FileSystemId == nil {
		invalidParams.Add(aws.NewErrParamRequired("FileSystemId"))
	}

	if s.LifecyclePolicies == nil {
		invalidParams.Add(aws.NewErrParamRequired("LifecyclePolicies"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutLifecycleConfigurationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.LifecyclePolicies != nil {
		v := s.LifecyclePolicies

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "LifecyclePolicies", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.FileSystemId != nil {
		v := *s.FileSystemId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "FileSystemId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticfilesystem-2015-02-01/LifecycleConfigurationDescription
type PutLifecycleConfigurationOutput struct {
	_ struct{} `type:"structure"`

	// An array of lifecycle management policies. Currently, EFS supports a maximum
	// of one policy per file system.
	LifecyclePolicies []LifecyclePolicy `type:"list"`
}

// String returns the string representation
func (s PutLifecycleConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutLifecycleConfigurationOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.LifecyclePolicies != nil {
		v := s.LifecyclePolicies

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "LifecyclePolicies", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opPutLifecycleConfiguration = "PutLifecycleConfiguration"

// PutLifecycleConfigurationRequest returns a request value for making API operation for
// Amazon Elastic File System.
//
// Enables lifecycle management by creating a new LifecycleConfiguration object.
// A LifecycleConfiguration object defines when files in an Amazon EFS file
// system are automatically transitioned to the lower-cost EFS Infrequent Access
// (IA) storage class. A LifecycleConfiguration applies to all files in a file
// system.
//
// Each Amazon EFS file system supports one lifecycle configuration, which applies
// to all files in the file system. If a LifecycleConfiguration object already
// exists for the specified file system, a PutLifecycleConfiguration call modifies
// the existing configuration. A PutLifecycleConfiguration call with an empty
// LifecyclePolicies array in the request body deletes any existing LifecycleConfiguration
// and disables lifecycle management.
//
// In the request, specify the following:
//
//    * The ID for the file system for which you are enabling, disabling, or
//    modifying lifecycle management.
//
//    * A LifecyclePolicies array of LifecyclePolicy objects that define when
//    files are moved to the IA storage class. The array can contain only one
//    LifecyclePolicy item.
//
// This operation requires permissions for the elasticfilesystem:PutLifecycleConfiguration
// operation.
//
// To apply a LifecycleConfiguration object to an encrypted file system, you
// need the same AWS Key Management Service (AWS KMS) permissions as when you
// created the encrypted file system.
//
//    // Example sending a request using PutLifecycleConfigurationRequest.
//    req := client.PutLifecycleConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticfilesystem-2015-02-01/PutLifecycleConfiguration
func (c *Client) PutLifecycleConfigurationRequest(input *PutLifecycleConfigurationInput) PutLifecycleConfigurationRequest {
	op := &aws.Operation{
		Name:       opPutLifecycleConfiguration,
		HTTPMethod: "PUT",
		HTTPPath:   "/2015-02-01/file-systems/{FileSystemId}/lifecycle-configuration",
	}

	if input == nil {
		input = &PutLifecycleConfigurationInput{}
	}

	req := c.newRequest(op, input, &PutLifecycleConfigurationOutput{})
	return PutLifecycleConfigurationRequest{Request: req, Input: input, Copy: c.PutLifecycleConfigurationRequest}
}

// PutLifecycleConfigurationRequest is the request type for the
// PutLifecycleConfiguration API operation.
type PutLifecycleConfigurationRequest struct {
	*aws.Request
	Input *PutLifecycleConfigurationInput
	Copy  func(*PutLifecycleConfigurationInput) PutLifecycleConfigurationRequest
}

// Send marshals and sends the PutLifecycleConfiguration API request.
func (r PutLifecycleConfigurationRequest) Send(ctx context.Context) (*PutLifecycleConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutLifecycleConfigurationResponse{
		PutLifecycleConfigurationOutput: r.Request.Data.(*PutLifecycleConfigurationOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutLifecycleConfigurationResponse is the response type for the
// PutLifecycleConfiguration API operation.
type PutLifecycleConfigurationResponse struct {
	*PutLifecycleConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutLifecycleConfiguration request.
func (r *PutLifecycleConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
