// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package devicefarm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the request to stop the remote access session.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/devicefarm-2015-06-23/StopRemoteAccessSessionRequest
type StopRemoteAccessSessionInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the remote access session you wish to stop.
	//
	// Arn is a required field
	Arn *string `locationName:"arn" min:"32" type:"string" required:"true"`
}

// String returns the string representation
func (s StopRemoteAccessSessionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StopRemoteAccessSessionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StopRemoteAccessSessionInput"}

	if s.Arn == nil {
		invalidParams.Add(aws.NewErrParamRequired("Arn"))
	}
	if s.Arn != nil && len(*s.Arn) < 32 {
		invalidParams.Add(aws.NewErrParamMinLen("Arn", 32))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the response from the server that describes the remote access
// session when AWS Device Farm stops the session.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/devicefarm-2015-06-23/StopRemoteAccessSessionResult
type StopRemoteAccessSessionOutput struct {
	_ struct{} `type:"structure"`

	// A container representing the metadata from the service about the remote access
	// session you are stopping.
	RemoteAccessSession *RemoteAccessSession `locationName:"remoteAccessSession" type:"structure"`
}

// String returns the string representation
func (s StopRemoteAccessSessionOutput) String() string {
	return awsutil.Prettify(s)
}

const opStopRemoteAccessSession = "StopRemoteAccessSession"

// StopRemoteAccessSessionRequest returns a request value for making API operation for
// AWS Device Farm.
//
// Ends a specified remote access session.
//
//    // Example sending a request using StopRemoteAccessSessionRequest.
//    req := client.StopRemoteAccessSessionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/devicefarm-2015-06-23/StopRemoteAccessSession
func (c *Client) StopRemoteAccessSessionRequest(input *StopRemoteAccessSessionInput) StopRemoteAccessSessionRequest {
	op := &aws.Operation{
		Name:       opStopRemoteAccessSession,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StopRemoteAccessSessionInput{}
	}

	req := c.newRequest(op, input, &StopRemoteAccessSessionOutput{})
	return StopRemoteAccessSessionRequest{Request: req, Input: input, Copy: c.StopRemoteAccessSessionRequest}
}

// StopRemoteAccessSessionRequest is the request type for the
// StopRemoteAccessSession API operation.
type StopRemoteAccessSessionRequest struct {
	*aws.Request
	Input *StopRemoteAccessSessionInput
	Copy  func(*StopRemoteAccessSessionInput) StopRemoteAccessSessionRequest
}

// Send marshals and sends the StopRemoteAccessSession API request.
func (r StopRemoteAccessSessionRequest) Send(ctx context.Context) (*StopRemoteAccessSessionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StopRemoteAccessSessionResponse{
		StopRemoteAccessSessionOutput: r.Request.Data.(*StopRemoteAccessSessionOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StopRemoteAccessSessionResponse is the response type for the
// StopRemoteAccessSession API operation.
type StopRemoteAccessSessionResponse struct {
	*StopRemoteAccessSessionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StopRemoteAccessSession request.
func (r *StopRemoteAccessSessionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
