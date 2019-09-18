// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the request to get the device.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/GetDeviceRequest
type GetDeviceInput struct {
	_ struct{} `type:"structure"`

	// The access token.
	AccessToken *string `type:"string"`

	// The device key.
	//
	// DeviceKey is a required field
	DeviceKey *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetDeviceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetDeviceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetDeviceInput"}

	if s.DeviceKey == nil {
		invalidParams.Add(aws.NewErrParamRequired("DeviceKey"))
	}
	if s.DeviceKey != nil && len(*s.DeviceKey) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DeviceKey", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Gets the device response.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/GetDeviceResponse
type GetDeviceOutput struct {
	_ struct{} `type:"structure"`

	// The device.
	//
	// Device is a required field
	Device *DeviceType `type:"structure" required:"true"`
}

// String returns the string representation
func (s GetDeviceOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetDevice = "GetDevice"

// GetDeviceRequest returns a request value for making API operation for
// Amazon Cognito Identity Provider.
//
// Gets the device.
//
//    // Example sending a request using GetDeviceRequest.
//    req := client.GetDeviceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/GetDevice
func (c *Client) GetDeviceRequest(input *GetDeviceInput) GetDeviceRequest {
	op := &aws.Operation{
		Name:       opGetDevice,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetDeviceInput{}
	}

	req := c.newRequest(op, input, &GetDeviceOutput{})
	return GetDeviceRequest{Request: req, Input: input, Copy: c.GetDeviceRequest}
}

// GetDeviceRequest is the request type for the
// GetDevice API operation.
type GetDeviceRequest struct {
	*aws.Request
	Input *GetDeviceInput
	Copy  func(*GetDeviceInput) GetDeviceRequest
}

// Send marshals and sends the GetDevice API request.
func (r GetDeviceRequest) Send(ctx context.Context) (*GetDeviceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetDeviceResponse{
		GetDeviceOutput: r.Request.Data.(*GetDeviceOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetDeviceResponse is the response type for the
// GetDevice API operation.
type GetDeviceResponse struct {
	*GetDeviceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetDevice request.
func (r *GetDeviceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
