// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/GetBotRequest
type GetBotInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Chime account ID.
	//
	// AccountId is a required field
	AccountId *string `location:"uri" locationName:"accountId" type:"string" required:"true"`

	// The bot ID.
	//
	// BotId is a required field
	BotId *string `location:"uri" locationName:"botId" type:"string" required:"true"`
}

// String returns the string representation
func (s GetBotInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetBotInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetBotInput"}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}

	if s.BotId == nil {
		invalidParams.Add(aws.NewErrParamRequired("BotId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetBotInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AccountId != nil {
		v := *s.AccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "accountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.BotId != nil {
		v := *s.BotId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "botId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/GetBotResponse
type GetBotOutput struct {
	_ struct{} `type:"structure"`

	// The chat bot details.
	Bot *Bot `type:"structure"`
}

// String returns the string representation
func (s GetBotOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetBotOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Bot != nil {
		v := s.Bot

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Bot", v, metadata)
	}
	return nil
}

const opGetBot = "GetBot"

// GetBotRequest returns a request value for making API operation for
// Amazon Chime.
//
// Retrieves details for the specified bot, such as bot email address, bot type,
// status, and display name.
//
//    // Example sending a request using GetBotRequest.
//    req := client.GetBotRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/GetBot
func (c *Client) GetBotRequest(input *GetBotInput) GetBotRequest {
	op := &aws.Operation{
		Name:       opGetBot,
		HTTPMethod: "GET",
		HTTPPath:   "/accounts/{accountId}/bots/{botId}",
	}

	if input == nil {
		input = &GetBotInput{}
	}

	req := c.newRequest(op, input, &GetBotOutput{})
	return GetBotRequest{Request: req, Input: input, Copy: c.GetBotRequest}
}

// GetBotRequest is the request type for the
// GetBot API operation.
type GetBotRequest struct {
	*aws.Request
	Input *GetBotInput
	Copy  func(*GetBotInput) GetBotRequest
}

// Send marshals and sends the GetBot API request.
func (r GetBotRequest) Send(ctx context.Context) (*GetBotResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetBotResponse{
		GetBotOutput: r.Request.Data.(*GetBotOutput),
		response:     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetBotResponse is the response type for the
// GetBot API operation.
type GetBotResponse struct {
	*GetBotOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetBot request.
func (r *GetBotResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
