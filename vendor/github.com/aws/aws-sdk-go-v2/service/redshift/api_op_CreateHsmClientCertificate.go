// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package redshift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/CreateHsmClientCertificateMessage
type CreateHsmClientCertificateInput struct {
	_ struct{} `type:"structure"`

	// The identifier to be assigned to the new HSM client certificate that the
	// cluster will use to connect to the HSM to use the database encryption keys.
	//
	// HsmClientCertificateIdentifier is a required field
	HsmClientCertificateIdentifier *string `type:"string" required:"true"`

	// A list of tag instances.
	Tags []Tag `locationNameList:"Tag" type:"list"`
}

// String returns the string representation
func (s CreateHsmClientCertificateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateHsmClientCertificateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateHsmClientCertificateInput"}

	if s.HsmClientCertificateIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("HsmClientCertificateIdentifier"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/CreateHsmClientCertificateResult
type CreateHsmClientCertificateOutput struct {
	_ struct{} `type:"structure"`

	// Returns information about an HSM client certificate. The certificate is stored
	// in a secure Hardware Storage Module (HSM), and used by the Amazon Redshift
	// cluster to encrypt data files.
	HsmClientCertificate *HsmClientCertificate `type:"structure"`
}

// String returns the string representation
func (s CreateHsmClientCertificateOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateHsmClientCertificate = "CreateHsmClientCertificate"

// CreateHsmClientCertificateRequest returns a request value for making API operation for
// Amazon Redshift.
//
// Creates an HSM client certificate that an Amazon Redshift cluster will use
// to connect to the client's HSM in order to store and retrieve the keys used
// to encrypt the cluster databases.
//
// The command returns a public key, which you must store in the HSM. In addition
// to creating the HSM certificate, you must create an Amazon Redshift HSM configuration
// that provides a cluster the information needed to store and use encryption
// keys in the HSM. For more information, go to Hardware Security Modules (https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-HSM.html)
// in the Amazon Redshift Cluster Management Guide.
//
//    // Example sending a request using CreateHsmClientCertificateRequest.
//    req := client.CreateHsmClientCertificateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/CreateHsmClientCertificate
func (c *Client) CreateHsmClientCertificateRequest(input *CreateHsmClientCertificateInput) CreateHsmClientCertificateRequest {
	op := &aws.Operation{
		Name:       opCreateHsmClientCertificate,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateHsmClientCertificateInput{}
	}

	req := c.newRequest(op, input, &CreateHsmClientCertificateOutput{})
	return CreateHsmClientCertificateRequest{Request: req, Input: input, Copy: c.CreateHsmClientCertificateRequest}
}

// CreateHsmClientCertificateRequest is the request type for the
// CreateHsmClientCertificate API operation.
type CreateHsmClientCertificateRequest struct {
	*aws.Request
	Input *CreateHsmClientCertificateInput
	Copy  func(*CreateHsmClientCertificateInput) CreateHsmClientCertificateRequest
}

// Send marshals and sends the CreateHsmClientCertificate API request.
func (r CreateHsmClientCertificateRequest) Send(ctx context.Context) (*CreateHsmClientCertificateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateHsmClientCertificateResponse{
		CreateHsmClientCertificateOutput: r.Request.Data.(*CreateHsmClientCertificateOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateHsmClientCertificateResponse is the response type for the
// CreateHsmClientCertificate API operation.
type CreateHsmClientCertificateResponse struct {
	*CreateHsmClientCertificateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateHsmClientCertificate request.
func (r *CreateHsmClientCertificateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
