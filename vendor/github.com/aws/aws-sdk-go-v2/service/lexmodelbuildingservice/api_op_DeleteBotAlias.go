// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lexmodelbuildingservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lex-models-2017-04-19/DeleteBotAliasRequest
type DeleteBotAliasInput struct {
	_ struct{} `type:"structure"`

	// The name of the bot that the alias points to.
	//
	// BotName is a required field
	BotName *string `location:"uri" locationName:"botName" min:"2" type:"string" required:"true"`

	// The name of the alias to delete. The name is case sensitive.
	//
	// Name is a required field
	Name *string `location:"uri" locationName:"name" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteBotAliasInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteBotAliasInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteBotAliasInput"}

	if s.BotName == nil {
		invalidParams.Add(aws.NewErrParamRequired("BotName"))
	}
	if s.BotName != nil && len(*s.BotName) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("BotName", 2))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteBotAliasInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.BotName != nil {
		v := *s.BotName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "botName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lex-models-2017-04-19/DeleteBotAliasOutput
type DeleteBotAliasOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteBotAliasOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteBotAliasOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteBotAlias = "DeleteBotAlias"

// DeleteBotAliasRequest returns a request value for making API operation for
// Amazon Lex Model Building Service.
//
// Deletes an alias for the specified bot.
//
// You can't delete an alias that is used in the association between a bot and
// a messaging channel. If an alias is used in a channel association, the DeleteBot
// operation returns a ResourceInUseException exception that includes a reference
// to the channel association that refers to the bot. You can remove the reference
// to the alias by deleting the channel association. If you get the same exception
// again, delete the referring association until the DeleteBotAlias operation
// is successful.
//
//    // Example sending a request using DeleteBotAliasRequest.
//    req := client.DeleteBotAliasRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lex-models-2017-04-19/DeleteBotAlias
func (c *Client) DeleteBotAliasRequest(input *DeleteBotAliasInput) DeleteBotAliasRequest {
	op := &aws.Operation{
		Name:       opDeleteBotAlias,
		HTTPMethod: "DELETE",
		HTTPPath:   "/bots/{botName}/aliases/{name}",
	}

	if input == nil {
		input = &DeleteBotAliasInput{}
	}

	req := c.newRequest(op, input, &DeleteBotAliasOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteBotAliasRequest{Request: req, Input: input, Copy: c.DeleteBotAliasRequest}
}

// DeleteBotAliasRequest is the request type for the
// DeleteBotAlias API operation.
type DeleteBotAliasRequest struct {
	*aws.Request
	Input *DeleteBotAliasInput
	Copy  func(*DeleteBotAliasInput) DeleteBotAliasRequest
}

// Send marshals and sends the DeleteBotAlias API request.
func (r DeleteBotAliasRequest) Send(ctx context.Context) (*DeleteBotAliasResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteBotAliasResponse{
		DeleteBotAliasOutput: r.Request.Data.(*DeleteBotAliasOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteBotAliasResponse is the response type for the
// DeleteBotAlias API operation.
type DeleteBotAliasResponse struct {
	*DeleteBotAliasOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteBotAlias request.
func (r *DeleteBotAliasResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
