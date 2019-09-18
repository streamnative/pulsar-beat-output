// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package acmpca

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/acm-pca-2017-08-22/ImportCertificateAuthorityCertificateRequest
type ImportCertificateAuthorityCertificateInput struct {
	_ struct{} `type:"structure"`

	// The PEM-encoded certificate for a private CA. This may be a self-signed certificate
	// in the case of a root CA, or it may be signed by another CA that you control.
	//
	// Certificate is automatically base64 encoded/decoded by the SDK.
	//
	// Certificate is a required field
	Certificate []byte `min:"1" type:"blob" required:"true"`

	// The Amazon Resource Name (ARN) that was returned when you called CreateCertificateAuthority.
	// This must be of the form:
	//
	// arn:aws:acm-pca:region:account:certificate-authority/12345678-1234-1234-1234-123456789012
	//
	// CertificateAuthorityArn is a required field
	CertificateAuthorityArn *string `min:"5" type:"string" required:"true"`

	// A PEM-encoded file that contains all of your certificates, other than the
	// certificate you're importing, chaining up to your root CA. Your ACM Private
	// CA-hosted or on-premises root certificate is the last in the chain, and each
	// certificate in the chain signs the one preceding.
	//
	// This parameter must be supplied when you import a subordinate CA. When you
	// import a root CA, there is no chain.
	//
	// CertificateChain is automatically base64 encoded/decoded by the SDK.
	CertificateChain []byte `type:"blob"`
}

// String returns the string representation
func (s ImportCertificateAuthorityCertificateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ImportCertificateAuthorityCertificateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ImportCertificateAuthorityCertificateInput"}

	if s.Certificate == nil {
		invalidParams.Add(aws.NewErrParamRequired("Certificate"))
	}
	if s.Certificate != nil && len(s.Certificate) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Certificate", 1))
	}

	if s.CertificateAuthorityArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("CertificateAuthorityArn"))
	}
	if s.CertificateAuthorityArn != nil && len(*s.CertificateAuthorityArn) < 5 {
		invalidParams.Add(aws.NewErrParamMinLen("CertificateAuthorityArn", 5))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/acm-pca-2017-08-22/ImportCertificateAuthorityCertificateOutput
type ImportCertificateAuthorityCertificateOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s ImportCertificateAuthorityCertificateOutput) String() string {
	return awsutil.Prettify(s)
}

const opImportCertificateAuthorityCertificate = "ImportCertificateAuthorityCertificate"

// ImportCertificateAuthorityCertificateRequest returns a request value for making API operation for
// AWS Certificate Manager Private Certificate Authority.
//
// Imports a signed private CA certificate into ACM Private CA. This action
// is used when you are using a chain of trust whose root is located outside
// ACM Private CA. Before you can call this action, the following preparations
// must in place:
//
// In ACM Private CA, call the CreateCertificateAuthority action to create the
// private CA that that you plan to back with the imported certificate.
//
// Call the GetCertificateAuthorityCsr action to generate a certificate signing
// request (CSR).
//
// Sign the CSR using a root or intermediate CA hosted either by an on-premises
// PKI hierarchy or a commercial CA..
//
// Create a certificate chain and copy the signed certificate and the certificate
// chain to your working directory.
//
// The following requirements apply when you import a CA certificate.
//
//    * You cannot import a non-self-signed certificate for use as a root CA.
//
//    * You cannot import a self-signed certificate for use as a subordinate
//    CA.
//
//    * Your certificate chain must not include the private CA certificate that
//    you are importing.
//
//    * Your ACM Private CA-hosted or on-premises CA certificate must be the
//    last certificate in your chain. The subordinate certificate, if any, that
//    your root CA signed must be next to last. The subordinate certificate
//    signed by the preceding subordinate CA must come next, and so on until
//    your chain is built.
//
//    * The chain must be PEM-encoded.
//
//    // Example sending a request using ImportCertificateAuthorityCertificateRequest.
//    req := client.ImportCertificateAuthorityCertificateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/acm-pca-2017-08-22/ImportCertificateAuthorityCertificate
func (c *Client) ImportCertificateAuthorityCertificateRequest(input *ImportCertificateAuthorityCertificateInput) ImportCertificateAuthorityCertificateRequest {
	op := &aws.Operation{
		Name:       opImportCertificateAuthorityCertificate,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ImportCertificateAuthorityCertificateInput{}
	}

	req := c.newRequest(op, input, &ImportCertificateAuthorityCertificateOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return ImportCertificateAuthorityCertificateRequest{Request: req, Input: input, Copy: c.ImportCertificateAuthorityCertificateRequest}
}

// ImportCertificateAuthorityCertificateRequest is the request type for the
// ImportCertificateAuthorityCertificate API operation.
type ImportCertificateAuthorityCertificateRequest struct {
	*aws.Request
	Input *ImportCertificateAuthorityCertificateInput
	Copy  func(*ImportCertificateAuthorityCertificateInput) ImportCertificateAuthorityCertificateRequest
}

// Send marshals and sends the ImportCertificateAuthorityCertificate API request.
func (r ImportCertificateAuthorityCertificateRequest) Send(ctx context.Context) (*ImportCertificateAuthorityCertificateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ImportCertificateAuthorityCertificateResponse{
		ImportCertificateAuthorityCertificateOutput: r.Request.Data.(*ImportCertificateAuthorityCertificateOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ImportCertificateAuthorityCertificateResponse is the response type for the
// ImportCertificateAuthorityCertificate API operation.
type ImportCertificateAuthorityCertificateResponse struct {
	*ImportCertificateAuthorityCertificateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ImportCertificateAuthorityCertificate request.
func (r *ImportCertificateAuthorityCertificateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
