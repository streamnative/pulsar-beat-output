// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package devicefarm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the request sent to retrieve the account settings.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/devicefarm-2015-06-23/GetAccountSettingsRequest
type GetAccountSettingsInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s GetAccountSettingsInput) String() string {
	return awsutil.Prettify(s)
}

// Represents the account settings return values from the GetAccountSettings
// request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/devicefarm-2015-06-23/GetAccountSettingsResult
type GetAccountSettingsOutput struct {
	_ struct{} `type:"structure"`

	// The account settings.
	AccountSettings *AccountSettings `locationName:"accountSettings" type:"structure"`
}

// String returns the string representation
func (s GetAccountSettingsOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetAccountSettings = "GetAccountSettings"

// GetAccountSettingsRequest returns a request value for making API operation for
// AWS Device Farm.
//
// Returns the number of unmetered iOS and/or unmetered Android devices that
// have been purchased by the account.
//
//    // Example sending a request using GetAccountSettingsRequest.
//    req := client.GetAccountSettingsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/devicefarm-2015-06-23/GetAccountSettings
func (c *Client) GetAccountSettingsRequest(input *GetAccountSettingsInput) GetAccountSettingsRequest {
	op := &aws.Operation{
		Name:       opGetAccountSettings,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetAccountSettingsInput{}
	}

	req := c.newRequest(op, input, &GetAccountSettingsOutput{})
	return GetAccountSettingsRequest{Request: req, Input: input, Copy: c.GetAccountSettingsRequest}
}

// GetAccountSettingsRequest is the request type for the
// GetAccountSettings API operation.
type GetAccountSettingsRequest struct {
	*aws.Request
	Input *GetAccountSettingsInput
	Copy  func(*GetAccountSettingsInput) GetAccountSettingsRequest
}

// Send marshals and sends the GetAccountSettings API request.
func (r GetAccountSettingsRequest) Send(ctx context.Context) (*GetAccountSettingsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetAccountSettingsResponse{
		GetAccountSettingsOutput: r.Request.Data.(*GetAccountSettingsOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetAccountSettingsResponse is the response type for the
// GetAccountSettings API operation.
type GetAccountSettingsResponse struct {
	*GetAccountSettingsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetAccountSettings request.
func (r *GetAccountSettingsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
