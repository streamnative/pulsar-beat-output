// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/DeleteKnownHostKeysRequest
type DeleteKnownHostKeysInput struct {
	_ struct{} `type:"structure"`

	// The name of the instance for which you want to reset the host key or certificate.
	//
	// InstanceName is a required field
	InstanceName *string `locationName:"instanceName" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteKnownHostKeysInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteKnownHostKeysInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteKnownHostKeysInput"}

	if s.InstanceName == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/DeleteKnownHostKeysResult
type DeleteKnownHostKeysOutput struct {
	_ struct{} `type:"structure"`

	// A list of objects describing the API operation.
	Operations []Operation `locationName:"operations" type:"list"`
}

// String returns the string representation
func (s DeleteKnownHostKeysOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteKnownHostKeys = "DeleteKnownHostKeys"

// DeleteKnownHostKeysRequest returns a request value for making API operation for
// Amazon Lightsail.
//
// Deletes the known host key or certificate used by the Amazon Lightsail browser-based
// SSH or RDP clients to authenticate an instance. This operation enables the
// Lightsail browser-based SSH or RDP clients to connect to the instance after
// a host key mismatch.
//
// Perform this operation only if you were expecting the host key or certificate
// mismatch or if you are familiar with the new host key or certificate on the
// instance. For more information, see Troubleshooting connection issues when
// using the Amazon Lightsail browser-based SSH or RDP client (https://lightsail.aws.amazon.com/ls/docs/en/articles/amazon-lightsail-troubleshooting-browser-based-ssh-rdp-client-connection).
//
//    // Example sending a request using DeleteKnownHostKeysRequest.
//    req := client.DeleteKnownHostKeysRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/DeleteKnownHostKeys
func (c *Client) DeleteKnownHostKeysRequest(input *DeleteKnownHostKeysInput) DeleteKnownHostKeysRequest {
	op := &aws.Operation{
		Name:       opDeleteKnownHostKeys,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteKnownHostKeysInput{}
	}

	req := c.newRequest(op, input, &DeleteKnownHostKeysOutput{})
	return DeleteKnownHostKeysRequest{Request: req, Input: input, Copy: c.DeleteKnownHostKeysRequest}
}

// DeleteKnownHostKeysRequest is the request type for the
// DeleteKnownHostKeys API operation.
type DeleteKnownHostKeysRequest struct {
	*aws.Request
	Input *DeleteKnownHostKeysInput
	Copy  func(*DeleteKnownHostKeysInput) DeleteKnownHostKeysRequest
}

// Send marshals and sends the DeleteKnownHostKeys API request.
func (r DeleteKnownHostKeysRequest) Send(ctx context.Context) (*DeleteKnownHostKeysResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteKnownHostKeysResponse{
		DeleteKnownHostKeysOutput: r.Request.Data.(*DeleteKnownHostKeysOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteKnownHostKeysResponse is the response type for the
// DeleteKnownHostKeys API operation.
type DeleteKnownHostKeysResponse struct {
	*DeleteKnownHostKeysOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteKnownHostKeys request.
func (r *DeleteKnownHostKeysResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
