// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package gamelift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input for a request action.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/DescribeMatchmakingInput
type DescribeMatchmakingInput struct {
	_ struct{} `type:"structure"`

	// Unique identifier for a matchmaking ticket. You can include up to 10 ID values.
	//
	// TicketIds is a required field
	TicketIds []string `type:"list" required:"true"`
}

// String returns the string representation
func (s DescribeMatchmakingInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeMatchmakingInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeMatchmakingInput"}

	if s.TicketIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("TicketIds"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the returned data in response to a request action.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/DescribeMatchmakingOutput
type DescribeMatchmakingOutput struct {
	_ struct{} `type:"structure"`

	// Collection of existing matchmaking ticket objects matching the request.
	TicketList []MatchmakingTicket `type:"list"`
}

// String returns the string representation
func (s DescribeMatchmakingOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeMatchmaking = "DescribeMatchmaking"

// DescribeMatchmakingRequest returns a request value for making API operation for
// Amazon GameLift.
//
// Retrieves one or more matchmaking tickets. Use this operation to retrieve
// ticket information, including status and--once a successful match is made--acquire
// connection information for the resulting new game session.
//
// You can use this operation to track the progress of matchmaking requests
// (through polling) as an alternative to using event notifications. See more
// details on tracking matchmaking requests through polling or notifications
// in StartMatchmaking.
//
// To request matchmaking tickets, provide a list of up to 10 ticket IDs. If
// the request is successful, a ticket object is returned for each requested
// ID that currently exists.
//
// Learn more
//
//  Add FlexMatch to a Game Client (https://docs.aws.amazon.com/gamelift/latest/developerguide/match-client.html)
//
//  Set Up FlexMatch Event Notification (https://docs.aws.amazon.com/gamelift/latest/developerguidematch-notification.html)
//
// Related operations
//
//    * StartMatchmaking
//
//    * DescribeMatchmaking
//
//    * StopMatchmaking
//
//    * AcceptMatch
//
//    * StartMatchBackfill
//
//    // Example sending a request using DescribeMatchmakingRequest.
//    req := client.DescribeMatchmakingRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/DescribeMatchmaking
func (c *Client) DescribeMatchmakingRequest(input *DescribeMatchmakingInput) DescribeMatchmakingRequest {
	op := &aws.Operation{
		Name:       opDescribeMatchmaking,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeMatchmakingInput{}
	}

	req := c.newRequest(op, input, &DescribeMatchmakingOutput{})
	return DescribeMatchmakingRequest{Request: req, Input: input, Copy: c.DescribeMatchmakingRequest}
}

// DescribeMatchmakingRequest is the request type for the
// DescribeMatchmaking API operation.
type DescribeMatchmakingRequest struct {
	*aws.Request
	Input *DescribeMatchmakingInput
	Copy  func(*DescribeMatchmakingInput) DescribeMatchmakingRequest
}

// Send marshals and sends the DescribeMatchmaking API request.
func (r DescribeMatchmakingRequest) Send(ctx context.Context) (*DescribeMatchmakingResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeMatchmakingResponse{
		DescribeMatchmakingOutput: r.Request.Data.(*DescribeMatchmakingOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeMatchmakingResponse is the response type for the
// DescribeMatchmaking API operation.
type DescribeMatchmakingResponse struct {
	*DescribeMatchmakingOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeMatchmaking request.
func (r *DescribeMatchmakingResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
