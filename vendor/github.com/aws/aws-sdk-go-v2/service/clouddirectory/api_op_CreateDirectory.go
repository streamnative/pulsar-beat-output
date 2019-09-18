// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package clouddirectory

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/CreateDirectoryRequest
type CreateDirectoryInput struct {
	_ struct{} `type:"structure"`

	// The name of the Directory. Should be unique per account, per region.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`

	// The Amazon Resource Name (ARN) of the published schema that will be copied
	// into the data Directory. For more information, see arns.
	//
	// SchemaArn is a required field
	SchemaArn *string `location:"header" locationName:"x-amz-data-partition" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateDirectoryInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDirectoryInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateDirectoryInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if s.SchemaArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("SchemaArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateDirectoryInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SchemaArn != nil {
		v := *s.SchemaArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-data-partition", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/CreateDirectoryResponse
type CreateDirectoryOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the published schema in the Directory. Once a published schema
	// is copied into the directory, it has its own ARN, which is referred to applied
	// schema ARN. For more information, see arns.
	//
	// AppliedSchemaArn is a required field
	AppliedSchemaArn *string `type:"string" required:"true"`

	// The ARN that is associated with the Directory. For more information, see
	// arns.
	//
	// DirectoryArn is a required field
	DirectoryArn *string `type:"string" required:"true"`

	// The name of the Directory.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`

	// The root object node of the created directory.
	//
	// ObjectIdentifier is a required field
	ObjectIdentifier *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateDirectoryOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateDirectoryOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.AppliedSchemaArn != nil {
		v := *s.AppliedSchemaArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "AppliedSchemaArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DirectoryArn != nil {
		v := *s.DirectoryArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DirectoryArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ObjectIdentifier != nil {
		v := *s.ObjectIdentifier

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ObjectIdentifier", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opCreateDirectory = "CreateDirectory"

// CreateDirectoryRequest returns a request value for making API operation for
// Amazon CloudDirectory.
//
// Creates a Directory by copying the published schema into the directory. A
// directory cannot be created without a schema.
//
// You can also quickly create a directory using a managed schema, called the
// QuickStartSchema. For more information, see Managed Schema (https://docs.aws.amazon.com/clouddirectory/latest/developerguide/schemas_managed.html)
// in the Amazon Cloud Directory Developer Guide.
//
//    // Example sending a request using CreateDirectoryRequest.
//    req := client.CreateDirectoryRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/CreateDirectory
func (c *Client) CreateDirectoryRequest(input *CreateDirectoryInput) CreateDirectoryRequest {
	op := &aws.Operation{
		Name:       opCreateDirectory,
		HTTPMethod: "PUT",
		HTTPPath:   "/amazonclouddirectory/2017-01-11/directory/create",
	}

	if input == nil {
		input = &CreateDirectoryInput{}
	}

	req := c.newRequest(op, input, &CreateDirectoryOutput{})
	return CreateDirectoryRequest{Request: req, Input: input, Copy: c.CreateDirectoryRequest}
}

// CreateDirectoryRequest is the request type for the
// CreateDirectory API operation.
type CreateDirectoryRequest struct {
	*aws.Request
	Input *CreateDirectoryInput
	Copy  func(*CreateDirectoryInput) CreateDirectoryRequest
}

// Send marshals and sends the CreateDirectory API request.
func (r CreateDirectoryRequest) Send(ctx context.Context) (*CreateDirectoryResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateDirectoryResponse{
		CreateDirectoryOutput: r.Request.Data.(*CreateDirectoryOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateDirectoryResponse is the response type for the
// CreateDirectory API operation.
type CreateDirectoryResponse struct {
	*CreateDirectoryOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateDirectory request.
func (r *CreateDirectoryResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
