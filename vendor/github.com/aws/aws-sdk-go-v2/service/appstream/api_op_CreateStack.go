// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appstream

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/appstream-2016-12-01/CreateStackRequest
type CreateStackInput struct {
	_ struct{} `type:"structure"`

	// The list of virtual private cloud (VPC) interface endpoint objects. Users
	// of the stack can connect to AppStream 2.0 only through the specified endpoints.
	AccessEndpoints []AccessEndpoint `min:"1" type:"list"`

	// The persistent application settings for users of a stack. When these settings
	// are enabled, changes that users make to applications and Windows settings
	// are automatically saved after each session and applied to the next session.
	ApplicationSettings *ApplicationSettings `type:"structure"`

	// The description to display.
	Description *string `type:"string"`

	// The stack name to display.
	DisplayName *string `type:"string"`

	// The URL that users are redirected to after they click the Send Feedback link.
	// If no URL is specified, no Send Feedback link is displayed.
	FeedbackURL *string `type:"string"`

	// The name of the stack.
	//
	// Name is a required field
	Name *string `type:"string" required:"true"`

	// The URL that users are redirected to after their streaming session ends.
	RedirectURL *string `type:"string"`

	// The storage connectors to enable.
	StorageConnectors []StorageConnector `type:"list"`

	// The tags to associate with the stack. A tag is a key-value pair, and the
	// value is optional. For example, Environment=Test. If you do not specify a
	// value, Environment=.
	//
	// If you do not specify a value, the value is set to an empty string.
	//
	// Generally allowed characters are: letters, numbers, and spaces representable
	// in UTF-8, and the following special characters:
	//
	// _ . : / = + \ - @
	//
	// For more information about tags, see Tagging Your Resources (https://docs.aws.amazon.com/appstream2/latest/developerguide/tagging-basic.html)
	// in the Amazon AppStream 2.0 Administration Guide.
	Tags map[string]string `min:"1" type:"map"`

	// The actions that are enabled or disabled for users during their streaming
	// sessions. By default, these actions are enabled.
	UserSettings []UserSetting `min:"1" type:"list"`
}

// String returns the string representation
func (s CreateStackInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateStackInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateStackInput"}
	if s.AccessEndpoints != nil && len(s.AccessEndpoints) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AccessEndpoints", 1))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Tags != nil && len(s.Tags) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Tags", 1))
	}
	if s.UserSettings != nil && len(s.UserSettings) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserSettings", 1))
	}
	if s.AccessEndpoints != nil {
		for i, v := range s.AccessEndpoints {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "AccessEndpoints", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.ApplicationSettings != nil {
		if err := s.ApplicationSettings.Validate(); err != nil {
			invalidParams.AddNested("ApplicationSettings", err.(aws.ErrInvalidParams))
		}
	}
	if s.StorageConnectors != nil {
		for i, v := range s.StorageConnectors {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "StorageConnectors", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.UserSettings != nil {
		for i, v := range s.UserSettings {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "UserSettings", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/appstream-2016-12-01/CreateStackResult
type CreateStackOutput struct {
	_ struct{} `type:"structure"`

	// Information about the stack.
	Stack *Stack `type:"structure"`
}

// String returns the string representation
func (s CreateStackOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateStack = "CreateStack"

// CreateStackRequest returns a request value for making API operation for
// Amazon AppStream.
//
// Creates a stack to start streaming applications to users. A stack consists
// of an associated fleet, user access policies, and storage configurations.
//
//    // Example sending a request using CreateStackRequest.
//    req := client.CreateStackRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appstream-2016-12-01/CreateStack
func (c *Client) CreateStackRequest(input *CreateStackInput) CreateStackRequest {
	op := &aws.Operation{
		Name:       opCreateStack,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateStackInput{}
	}

	req := c.newRequest(op, input, &CreateStackOutput{})
	return CreateStackRequest{Request: req, Input: input, Copy: c.CreateStackRequest}
}

// CreateStackRequest is the request type for the
// CreateStack API operation.
type CreateStackRequest struct {
	*aws.Request
	Input *CreateStackInput
	Copy  func(*CreateStackInput) CreateStackRequest
}

// Send marshals and sends the CreateStack API request.
func (r CreateStackRequest) Send(ctx context.Context) (*CreateStackResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateStackResponse{
		CreateStackOutput: r.Request.Data.(*CreateStackOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateStackResponse is the response type for the
// CreateStack API operation.
type CreateStackResponse struct {
	*CreateStackOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateStack request.
func (r *CreateStackResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
