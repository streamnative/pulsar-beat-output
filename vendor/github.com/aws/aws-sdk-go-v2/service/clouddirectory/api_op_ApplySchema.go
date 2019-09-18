// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package clouddirectory

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/ApplySchemaRequest
type ApplySchemaInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) that is associated with the Directory into
	// which the schema is copied. For more information, see arns.
	//
	// DirectoryArn is a required field
	DirectoryArn *string `location:"header" locationName:"x-amz-data-partition" type:"string" required:"true"`

	// Published schema Amazon Resource Name (ARN) that needs to be copied. For
	// more information, see arns.
	//
	// PublishedSchemaArn is a required field
	PublishedSchemaArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ApplySchemaInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ApplySchemaInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ApplySchemaInput"}

	if s.DirectoryArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("DirectoryArn"))
	}

	if s.PublishedSchemaArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("PublishedSchemaArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ApplySchemaInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.PublishedSchemaArn != nil {
		v := *s.PublishedSchemaArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "PublishedSchemaArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DirectoryArn != nil {
		v := *s.DirectoryArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-data-partition", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/ApplySchemaResponse
type ApplySchemaOutput struct {
	_ struct{} `type:"structure"`

	// The applied schema ARN that is associated with the copied schema in the Directory.
	// You can use this ARN to describe the schema information applied on this directory.
	// For more information, see arns.
	AppliedSchemaArn *string `type:"string"`

	// The ARN that is associated with the Directory. For more information, see
	// arns.
	DirectoryArn *string `type:"string"`
}

// String returns the string representation
func (s ApplySchemaOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ApplySchemaOutput) MarshalFields(e protocol.FieldEncoder) error {
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
	return nil
}

const opApplySchema = "ApplySchema"

// ApplySchemaRequest returns a request value for making API operation for
// Amazon CloudDirectory.
//
// Copies the input published schema, at the specified version, into the Directory
// with the same name and version as that of the published schema.
//
//    // Example sending a request using ApplySchemaRequest.
//    req := client.ApplySchemaRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/ApplySchema
func (c *Client) ApplySchemaRequest(input *ApplySchemaInput) ApplySchemaRequest {
	op := &aws.Operation{
		Name:       opApplySchema,
		HTTPMethod: "PUT",
		HTTPPath:   "/amazonclouddirectory/2017-01-11/schema/apply",
	}

	if input == nil {
		input = &ApplySchemaInput{}
	}

	req := c.newRequest(op, input, &ApplySchemaOutput{})
	return ApplySchemaRequest{Request: req, Input: input, Copy: c.ApplySchemaRequest}
}

// ApplySchemaRequest is the request type for the
// ApplySchema API operation.
type ApplySchemaRequest struct {
	*aws.Request
	Input *ApplySchemaInput
	Copy  func(*ApplySchemaInput) ApplySchemaRequest
}

// Send marshals and sends the ApplySchema API request.
func (r ApplySchemaRequest) Send(ctx context.Context) (*ApplySchemaResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ApplySchemaResponse{
		ApplySchemaOutput: r.Request.Data.(*ApplySchemaOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ApplySchemaResponse is the response type for the
// ApplySchema API operation.
type ApplySchemaResponse struct {
	*ApplySchemaOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ApplySchema request.
func (r *ApplySchemaResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
