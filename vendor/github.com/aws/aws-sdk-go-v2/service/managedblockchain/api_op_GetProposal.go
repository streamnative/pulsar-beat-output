// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package managedblockchain

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/managedblockchain-2018-09-24/GetProposalInput
type GetProposalInput struct {
	_ struct{} `type:"structure"`

	// The unique identifier of the network for which the proposal is made.
	//
	// NetworkId is a required field
	NetworkId *string `location:"uri" locationName:"networkId" min:"1" type:"string" required:"true"`

	// The unique identifier of the proposal.
	//
	// ProposalId is a required field
	ProposalId *string `location:"uri" locationName:"proposalId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetProposalInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetProposalInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetProposalInput"}

	if s.NetworkId == nil {
		invalidParams.Add(aws.NewErrParamRequired("NetworkId"))
	}
	if s.NetworkId != nil && len(*s.NetworkId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NetworkId", 1))
	}

	if s.ProposalId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProposalId"))
	}
	if s.ProposalId != nil && len(*s.ProposalId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ProposalId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetProposalInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.NetworkId != nil {
		v := *s.NetworkId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "networkId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ProposalId != nil {
		v := *s.ProposalId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "proposalId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/managedblockchain-2018-09-24/GetProposalOutput
type GetProposalOutput struct {
	_ struct{} `type:"structure"`

	// Information about a proposal.
	Proposal *Proposal `type:"structure"`
}

// String returns the string representation
func (s GetProposalOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetProposalOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Proposal != nil {
		v := s.Proposal

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Proposal", v, metadata)
	}
	return nil
}

const opGetProposal = "GetProposal"

// GetProposalRequest returns a request value for making API operation for
// Amazon Managed Blockchain.
//
// Returns detailed information about a proposal.
//
//    // Example sending a request using GetProposalRequest.
//    req := client.GetProposalRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/managedblockchain-2018-09-24/GetProposal
func (c *Client) GetProposalRequest(input *GetProposalInput) GetProposalRequest {
	op := &aws.Operation{
		Name:       opGetProposal,
		HTTPMethod: "GET",
		HTTPPath:   "/networks/{networkId}/proposals/{proposalId}",
	}

	if input == nil {
		input = &GetProposalInput{}
	}

	req := c.newRequest(op, input, &GetProposalOutput{})
	return GetProposalRequest{Request: req, Input: input, Copy: c.GetProposalRequest}
}

// GetProposalRequest is the request type for the
// GetProposal API operation.
type GetProposalRequest struct {
	*aws.Request
	Input *GetProposalInput
	Copy  func(*GetProposalInput) GetProposalRequest
}

// Send marshals and sends the GetProposal API request.
func (r GetProposalRequest) Send(ctx context.Context) (*GetProposalResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetProposalResponse{
		GetProposalOutput: r.Request.Data.(*GetProposalOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetProposalResponse is the response type for the
// GetProposal API operation.
type GetProposalResponse struct {
	*GetProposalOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetProposal request.
func (r *GetProposalResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
