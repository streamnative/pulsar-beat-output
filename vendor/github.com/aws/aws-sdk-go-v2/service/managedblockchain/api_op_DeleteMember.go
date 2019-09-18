// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package managedblockchain

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/managedblockchain-2018-09-24/DeleteMemberInput
type DeleteMemberInput struct {
	_ struct{} `type:"structure"`

	// The unique identifier of the member to remove.
	//
	// MemberId is a required field
	MemberId *string `location:"uri" locationName:"memberId" min:"1" type:"string" required:"true"`

	// The unique identifier of the network from which the member is removed.
	//
	// NetworkId is a required field
	NetworkId *string `location:"uri" locationName:"networkId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteMemberInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteMemberInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteMemberInput"}

	if s.MemberId == nil {
		invalidParams.Add(aws.NewErrParamRequired("MemberId"))
	}
	if s.MemberId != nil && len(*s.MemberId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("MemberId", 1))
	}

	if s.NetworkId == nil {
		invalidParams.Add(aws.NewErrParamRequired("NetworkId"))
	}
	if s.NetworkId != nil && len(*s.NetworkId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NetworkId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteMemberInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.MemberId != nil {
		v := *s.MemberId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "memberId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.NetworkId != nil {
		v := *s.NetworkId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "networkId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/managedblockchain-2018-09-24/DeleteMemberOutput
type DeleteMemberOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteMemberOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteMemberOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteMember = "DeleteMember"

// DeleteMemberRequest returns a request value for making API operation for
// Amazon Managed Blockchain.
//
// Deletes a member. Deleting a member removes the member and all associated
// resources from the network. DeleteMember can only be called for a specified
// MemberId if the principal performing the action is associated with the AWS
// account that owns the member. In all other cases, the DeleteMember action
// is carried out as the result of an approved proposal to remove a member.
// If MemberId is the last member in a network specified by the last AWS account,
// the network is deleted also.
//
//    // Example sending a request using DeleteMemberRequest.
//    req := client.DeleteMemberRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/managedblockchain-2018-09-24/DeleteMember
func (c *Client) DeleteMemberRequest(input *DeleteMemberInput) DeleteMemberRequest {
	op := &aws.Operation{
		Name:       opDeleteMember,
		HTTPMethod: "DELETE",
		HTTPPath:   "/networks/{networkId}/members/{memberId}",
	}

	if input == nil {
		input = &DeleteMemberInput{}
	}

	req := c.newRequest(op, input, &DeleteMemberOutput{})
	return DeleteMemberRequest{Request: req, Input: input, Copy: c.DeleteMemberRequest}
}

// DeleteMemberRequest is the request type for the
// DeleteMember API operation.
type DeleteMemberRequest struct {
	*aws.Request
	Input *DeleteMemberInput
	Copy  func(*DeleteMemberInput) DeleteMemberRequest
}

// Send marshals and sends the DeleteMember API request.
func (r DeleteMemberRequest) Send(ctx context.Context) (*DeleteMemberResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteMemberResponse{
		DeleteMemberOutput: r.Request.Data.(*DeleteMemberOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteMemberResponse is the response type for the
// DeleteMember API operation.
type DeleteMemberResponse struct {
	*DeleteMemberOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteMember request.
func (r *DeleteMemberResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
