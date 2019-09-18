// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Request to get a signing certificate from Cognito.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/GetSigningCertificateRequest
type GetSigningCertificateInput struct {
	_ struct{} `type:"structure"`

	// The user pool ID.
	//
	// UserPoolId is a required field
	UserPoolId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetSigningCertificateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetSigningCertificateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetSigningCertificateInput"}

	if s.UserPoolId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserPoolId"))
	}
	if s.UserPoolId != nil && len(*s.UserPoolId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserPoolId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Response from Cognito for a signing certificate request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/GetSigningCertificateResponse
type GetSigningCertificateOutput struct {
	_ struct{} `type:"structure"`

	// The signing certificate.
	Certificate *string `type:"string"`
}

// String returns the string representation
func (s GetSigningCertificateOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetSigningCertificate = "GetSigningCertificate"

// GetSigningCertificateRequest returns a request value for making API operation for
// Amazon Cognito Identity Provider.
//
// This method takes a user pool ID, and returns the signing certificate.
//
//    // Example sending a request using GetSigningCertificateRequest.
//    req := client.GetSigningCertificateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/GetSigningCertificate
func (c *Client) GetSigningCertificateRequest(input *GetSigningCertificateInput) GetSigningCertificateRequest {
	op := &aws.Operation{
		Name:       opGetSigningCertificate,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetSigningCertificateInput{}
	}

	req := c.newRequest(op, input, &GetSigningCertificateOutput{})
	return GetSigningCertificateRequest{Request: req, Input: input, Copy: c.GetSigningCertificateRequest}
}

// GetSigningCertificateRequest is the request type for the
// GetSigningCertificate API operation.
type GetSigningCertificateRequest struct {
	*aws.Request
	Input *GetSigningCertificateInput
	Copy  func(*GetSigningCertificateInput) GetSigningCertificateRequest
}

// Send marshals and sends the GetSigningCertificate API request.
func (r GetSigningCertificateRequest) Send(ctx context.Context) (*GetSigningCertificateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetSigningCertificateResponse{
		GetSigningCertificateOutput: r.Request.Data.(*GetSigningCertificateOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetSigningCertificateResponse is the response type for the
// GetSigningCertificate API operation.
type GetSigningCertificateResponse struct {
	*GetSigningCertificateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetSigningCertificate request.
func (r *GetSigningCertificateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
