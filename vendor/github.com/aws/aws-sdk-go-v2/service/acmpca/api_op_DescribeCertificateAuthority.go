// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package acmpca

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/acm-pca-2017-08-22/DescribeCertificateAuthorityRequest
type DescribeCertificateAuthorityInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) that was returned when you called CreateCertificateAuthority.
	// This must be of the form:
	//
	// arn:aws:acm-pca:region:account:certificate-authority/12345678-1234-1234-1234-123456789012 .
	//
	// CertificateAuthorityArn is a required field
	CertificateAuthorityArn *string `min:"5" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeCertificateAuthorityInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeCertificateAuthorityInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeCertificateAuthorityInput"}

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

// Please also see https://docs.aws.amazon.com/goto/WebAPI/acm-pca-2017-08-22/DescribeCertificateAuthorityResponse
type DescribeCertificateAuthorityOutput struct {
	_ struct{} `type:"structure"`

	// A CertificateAuthority structure that contains information about your private
	// CA.
	CertificateAuthority *CertificateAuthority `type:"structure"`
}

// String returns the string representation
func (s DescribeCertificateAuthorityOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeCertificateAuthority = "DescribeCertificateAuthority"

// DescribeCertificateAuthorityRequest returns a request value for making API operation for
// AWS Certificate Manager Private Certificate Authority.
//
// Lists information about your private certificate authority (CA). You specify
// the private CA on input by its ARN (Amazon Resource Name). The output contains
// the status of your CA. This can be any of the following:
//
//    * CREATING - ACM Private CA is creating your private certificate authority.
//
//    * PENDING_CERTIFICATE - The certificate is pending. You must use your
//    ACM Private CA-hosted or on-premises root or subordinate CA to sign your
//    private CA CSR and then import it into PCA.
//
//    * ACTIVE - Your private CA is active.
//
//    * DISABLED - Your private CA has been disabled.
//
//    * EXPIRED - Your private CA certificate has expired.
//
//    * FAILED - Your private CA has failed. Your CA can fail because of problems
//    such a network outage or backend AWS failure or other errors. A failed
//    CA can never return to the pending state. You must create a new CA.
//
//    * DELETED - Your private CA is within the restoration period, after which
//    it is permanently deleted. The length of time remaining in the CA's restoration
//    period is also included in this action's output.
//
//    // Example sending a request using DescribeCertificateAuthorityRequest.
//    req := client.DescribeCertificateAuthorityRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/acm-pca-2017-08-22/DescribeCertificateAuthority
func (c *Client) DescribeCertificateAuthorityRequest(input *DescribeCertificateAuthorityInput) DescribeCertificateAuthorityRequest {
	op := &aws.Operation{
		Name:       opDescribeCertificateAuthority,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeCertificateAuthorityInput{}
	}

	req := c.newRequest(op, input, &DescribeCertificateAuthorityOutput{})
	return DescribeCertificateAuthorityRequest{Request: req, Input: input, Copy: c.DescribeCertificateAuthorityRequest}
}

// DescribeCertificateAuthorityRequest is the request type for the
// DescribeCertificateAuthority API operation.
type DescribeCertificateAuthorityRequest struct {
	*aws.Request
	Input *DescribeCertificateAuthorityInput
	Copy  func(*DescribeCertificateAuthorityInput) DescribeCertificateAuthorityRequest
}

// Send marshals and sends the DescribeCertificateAuthority API request.
func (r DescribeCertificateAuthorityRequest) Send(ctx context.Context) (*DescribeCertificateAuthorityResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeCertificateAuthorityResponse{
		DescribeCertificateAuthorityOutput: r.Request.Data.(*DescribeCertificateAuthorityOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeCertificateAuthorityResponse is the response type for the
// DescribeCertificateAuthority API operation.
type DescribeCertificateAuthorityResponse struct {
	*DescribeCertificateAuthorityOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeCertificateAuthority request.
func (r *DescribeCertificateAuthorityResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
