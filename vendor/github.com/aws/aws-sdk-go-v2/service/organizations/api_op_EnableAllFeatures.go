// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package organizations

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/organizations-2016-11-28/EnableAllFeaturesRequest
type EnableAllFeaturesInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s EnableAllFeaturesInput) String() string {
	return awsutil.Prettify(s)
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/organizations-2016-11-28/EnableAllFeaturesResponse
type EnableAllFeaturesOutput struct {
	_ struct{} `type:"structure"`

	// A structure that contains details about the handshake created to support
	// this request to enable all features in the organization.
	Handshake *Handshake `type:"structure"`
}

// String returns the string representation
func (s EnableAllFeaturesOutput) String() string {
	return awsutil.Prettify(s)
}

const opEnableAllFeatures = "EnableAllFeatures"

// EnableAllFeaturesRequest returns a request value for making API operation for
// AWS Organizations.
//
// Enables all features in an organization. This enables the use of organization
// policies that can restrict the services and actions that can be called in
// each account. Until you enable all features, you have access only to consolidated
// billing, and you can't use any of the advanced account administration features
// that AWS Organizations supports. For more information, see Enabling All Features
// in Your Organization (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_org_support-all-features.html)
// in the AWS Organizations User Guide.
//
// This operation is required only for organizations that were created explicitly
// with only the consolidated billing features enabled. Calling this operation
// sends a handshake to every invited account in the organization. The feature
// set change can be finalized and the additional features enabled only after
// all administrators in the invited accounts approve the change by accepting
// the handshake.
//
// After you enable all features, you can separately enable or disable individual
// policy types in a root using EnablePolicyType and DisablePolicyType. To see
// the status of policy types in a root, use ListRoots.
//
// After all invited member accounts accept the handshake, you finalize the
// feature set change by accepting the handshake that contains "Action": "ENABLE_ALL_FEATURES".
// This completes the change.
//
// After you enable all features in your organization, the master account in
// the organization can apply policies on all member accounts. These policies
// can restrict what users and even administrators in those accounts can do.
// The master account can apply policies that prevent accounts from leaving
// the organization. Ensure that your account administrators are aware of this.
//
// This operation can be called only from the organization's master account.
//
//    // Example sending a request using EnableAllFeaturesRequest.
//    req := client.EnableAllFeaturesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/organizations-2016-11-28/EnableAllFeatures
func (c *Client) EnableAllFeaturesRequest(input *EnableAllFeaturesInput) EnableAllFeaturesRequest {
	op := &aws.Operation{
		Name:       opEnableAllFeatures,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &EnableAllFeaturesInput{}
	}

	req := c.newRequest(op, input, &EnableAllFeaturesOutput{})
	return EnableAllFeaturesRequest{Request: req, Input: input, Copy: c.EnableAllFeaturesRequest}
}

// EnableAllFeaturesRequest is the request type for the
// EnableAllFeatures API operation.
type EnableAllFeaturesRequest struct {
	*aws.Request
	Input *EnableAllFeaturesInput
	Copy  func(*EnableAllFeaturesInput) EnableAllFeaturesRequest
}

// Send marshals and sends the EnableAllFeatures API request.
func (r EnableAllFeaturesRequest) Send(ctx context.Context) (*EnableAllFeaturesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &EnableAllFeaturesResponse{
		EnableAllFeaturesOutput: r.Request.Data.(*EnableAllFeaturesOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// EnableAllFeaturesResponse is the response type for the
// EnableAllFeatures API operation.
type EnableAllFeaturesResponse struct {
	*EnableAllFeaturesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// EnableAllFeatures request.
func (r *EnableAllFeaturesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
