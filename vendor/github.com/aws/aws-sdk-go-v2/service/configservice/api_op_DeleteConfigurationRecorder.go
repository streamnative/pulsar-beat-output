// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package configservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

// The request object for the DeleteConfigurationRecorder action.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/config-2014-11-12/DeleteConfigurationRecorderRequest
type DeleteConfigurationRecorderInput struct {
	_ struct{} `type:"structure"`

	// The name of the configuration recorder to be deleted. You can retrieve the
	// name of your configuration recorder by using the DescribeConfigurationRecorders
	// action.
	//
	// ConfigurationRecorderName is a required field
	ConfigurationRecorderName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteConfigurationRecorderInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteConfigurationRecorderInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteConfigurationRecorderInput"}

	if s.ConfigurationRecorderName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConfigurationRecorderName"))
	}
	if s.ConfigurationRecorderName != nil && len(*s.ConfigurationRecorderName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ConfigurationRecorderName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/config-2014-11-12/DeleteConfigurationRecorderOutput
type DeleteConfigurationRecorderOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteConfigurationRecorderOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteConfigurationRecorder = "DeleteConfigurationRecorder"

// DeleteConfigurationRecorderRequest returns a request value for making API operation for
// AWS Config.
//
// Deletes the configuration recorder.
//
// After the configuration recorder is deleted, AWS Config will not record resource
// configuration changes until you create a new configuration recorder.
//
// This action does not delete the configuration information that was previously
// recorded. You will be able to access the previously recorded information
// by using the GetResourceConfigHistory action, but you will not be able to
// access this information in the AWS Config console until you create a new
// configuration recorder.
//
//    // Example sending a request using DeleteConfigurationRecorderRequest.
//    req := client.DeleteConfigurationRecorderRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/config-2014-11-12/DeleteConfigurationRecorder
func (c *Client) DeleteConfigurationRecorderRequest(input *DeleteConfigurationRecorderInput) DeleteConfigurationRecorderRequest {
	op := &aws.Operation{
		Name:       opDeleteConfigurationRecorder,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteConfigurationRecorderInput{}
	}

	req := c.newRequest(op, input, &DeleteConfigurationRecorderOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteConfigurationRecorderRequest{Request: req, Input: input, Copy: c.DeleteConfigurationRecorderRequest}
}

// DeleteConfigurationRecorderRequest is the request type for the
// DeleteConfigurationRecorder API operation.
type DeleteConfigurationRecorderRequest struct {
	*aws.Request
	Input *DeleteConfigurationRecorderInput
	Copy  func(*DeleteConfigurationRecorderInput) DeleteConfigurationRecorderRequest
}

// Send marshals and sends the DeleteConfigurationRecorder API request.
func (r DeleteConfigurationRecorderRequest) Send(ctx context.Context) (*DeleteConfigurationRecorderResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteConfigurationRecorderResponse{
		DeleteConfigurationRecorderOutput: r.Request.Data.(*DeleteConfigurationRecorderOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteConfigurationRecorderResponse is the response type for the
// DeleteConfigurationRecorder API operation.
type DeleteConfigurationRecorderResponse struct {
	*DeleteConfigurationRecorderOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteConfigurationRecorder request.
func (r *DeleteConfigurationRecorderResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
