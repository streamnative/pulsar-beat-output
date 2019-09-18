// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package quicksight

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/quicksight-2018-04-01/CreateGroupMembershipRequest
type CreateGroupMembershipInput struct {
	_ struct{} `type:"structure"`

	// The ID for the AWS account that the group is in. Currently, you use the ID
	// for the AWS account that contains your Amazon QuickSight account.
	//
	// AwsAccountId is a required field
	AwsAccountId *string `location:"uri" locationName:"AwsAccountId" min:"12" type:"string" required:"true"`

	// The name of the group that you want to add the user to.
	//
	// GroupName is a required field
	GroupName *string `location:"uri" locationName:"GroupName" min:"1" type:"string" required:"true"`

	// The name of the user that you want to add to the group membership.
	//
	// MemberName is a required field
	MemberName *string `location:"uri" locationName:"MemberName" min:"1" type:"string" required:"true"`

	// The namespace. Currently, you should set this to default.
	//
	// Namespace is a required field
	Namespace *string `location:"uri" locationName:"Namespace" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateGroupMembershipInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateGroupMembershipInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateGroupMembershipInput"}

	if s.AwsAccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AwsAccountId"))
	}
	if s.AwsAccountId != nil && len(*s.AwsAccountId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("AwsAccountId", 12))
	}

	if s.GroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("GroupName"))
	}
	if s.GroupName != nil && len(*s.GroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("GroupName", 1))
	}

	if s.MemberName == nil {
		invalidParams.Add(aws.NewErrParamRequired("MemberName"))
	}
	if s.MemberName != nil && len(*s.MemberName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("MemberName", 1))
	}

	if s.Namespace == nil {
		invalidParams.Add(aws.NewErrParamRequired("Namespace"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateGroupMembershipInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AwsAccountId != nil {
		v := *s.AwsAccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "AwsAccountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.GroupName != nil {
		v := *s.GroupName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "GroupName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MemberName != nil {
		v := *s.MemberName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "MemberName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Namespace != nil {
		v := *s.Namespace

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Namespace", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/quicksight-2018-04-01/CreateGroupMembershipResponse
type CreateGroupMembershipOutput struct {
	_ struct{} `type:"structure"`

	// The group member.
	GroupMember *GroupMember `type:"structure"`

	// The AWS request ID for this operation.
	RequestId *string `type:"string"`

	// The http status of the request.
	Status *int64 `location:"statusCode" type:"integer"`
}

// String returns the string representation
func (s CreateGroupMembershipOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateGroupMembershipOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.GroupMember != nil {
		v := s.GroupMember

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "GroupMember", v, metadata)
	}
	if s.RequestId != nil {
		v := *s.RequestId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "RequestId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	// ignoring invalid encode state, StatusCode. Status
	return nil
}

const opCreateGroupMembership = "CreateGroupMembership"

// CreateGroupMembershipRequest returns a request value for making API operation for
// Amazon QuickSight.
//
// Adds an Amazon QuickSight user to an Amazon QuickSight group.
//
// The permissions resource is arn:aws:quicksight:us-east-1:<aws-account-id>:group/default/<group-name> .
//
// The condition resource is the user name.
//
// The condition key is quicksight:UserName.
//
// The response is the group member object.
//
// CLI Sample:
//
// aws quicksight create-group-membership --aws-account-id=111122223333 --namespace=default
// --group-name=Sales --member-name=Pat
//
//    // Example sending a request using CreateGroupMembershipRequest.
//    req := client.CreateGroupMembershipRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/quicksight-2018-04-01/CreateGroupMembership
func (c *Client) CreateGroupMembershipRequest(input *CreateGroupMembershipInput) CreateGroupMembershipRequest {
	op := &aws.Operation{
		Name:       opCreateGroupMembership,
		HTTPMethod: "PUT",
		HTTPPath:   "/accounts/{AwsAccountId}/namespaces/{Namespace}/groups/{GroupName}/members/{MemberName}",
	}

	if input == nil {
		input = &CreateGroupMembershipInput{}
	}

	req := c.newRequest(op, input, &CreateGroupMembershipOutput{})
	return CreateGroupMembershipRequest{Request: req, Input: input, Copy: c.CreateGroupMembershipRequest}
}

// CreateGroupMembershipRequest is the request type for the
// CreateGroupMembership API operation.
type CreateGroupMembershipRequest struct {
	*aws.Request
	Input *CreateGroupMembershipInput
	Copy  func(*CreateGroupMembershipInput) CreateGroupMembershipRequest
}

// Send marshals and sends the CreateGroupMembership API request.
func (r CreateGroupMembershipRequest) Send(ctx context.Context) (*CreateGroupMembershipResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateGroupMembershipResponse{
		CreateGroupMembershipOutput: r.Request.Data.(*CreateGroupMembershipOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateGroupMembershipResponse is the response type for the
// CreateGroupMembership API operation.
type CreateGroupMembershipResponse struct {
	*CreateGroupMembershipOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateGroupMembership request.
func (r *CreateGroupMembershipResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
