// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package directconnect

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/directconnect-2012-10-25/CreateInterconnectRequest
type CreateInterconnectInput struct {
	_ struct{} `type:"structure"`

	// The port bandwidth, in Gbps. The possible values are 1 and 10.
	//
	// Bandwidth is a required field
	Bandwidth *string `locationName:"bandwidth" type:"string" required:"true"`

	// The name of the interconnect.
	//
	// InterconnectName is a required field
	InterconnectName *string `locationName:"interconnectName" type:"string" required:"true"`

	// The ID of the LAG.
	LagId *string `locationName:"lagId" type:"string"`

	// The location of the interconnect.
	//
	// Location is a required field
	Location *string `locationName:"location" type:"string" required:"true"`

	// The tags to assign to the interconnect,
	Tags []Tag `locationName:"tags" min:"1" type:"list"`
}

// String returns the string representation
func (s CreateInterconnectInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateInterconnectInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateInterconnectInput"}

	if s.Bandwidth == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bandwidth"))
	}

	if s.InterconnectName == nil {
		invalidParams.Add(aws.NewErrParamRequired("InterconnectName"))
	}

	if s.Location == nil {
		invalidParams.Add(aws.NewErrParamRequired("Location"))
	}
	if s.Tags != nil && len(s.Tags) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Tags", 1))
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Information about an interconnect.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/directconnect-2012-10-25/Interconnect
type CreateInterconnectOutput struct {
	_ struct{} `type:"structure"`

	// The Direct Connect endpoint on which the physical connection terminates.
	AwsDevice *string `locationName:"awsDevice" deprecated:"true" type:"string"`

	// The Direct Connect endpoint on which the physical connection terminates.
	AwsDeviceV2 *string `locationName:"awsDeviceV2" type:"string"`

	// The bandwidth of the connection.
	Bandwidth *string `locationName:"bandwidth" type:"string"`

	// Indicates whether the interconnect supports a secondary BGP in the same address
	// family (IPv4/IPv6).
	HasLogicalRedundancy HasLogicalRedundancy `locationName:"hasLogicalRedundancy" type:"string" enum:"true"`

	// The ID of the interconnect.
	InterconnectId *string `locationName:"interconnectId" type:"string"`

	// The name of the interconnect.
	InterconnectName *string `locationName:"interconnectName" type:"string"`

	// The state of the interconnect. The following are the possible values:
	//
	//    * requested: The initial state of an interconnect. The interconnect stays
	//    in the requested state until the Letter of Authorization (LOA) is sent
	//    to the customer.
	//
	//    * pending: The interconnect is approved, and is being initialized.
	//
	//    * available: The network link is up, and the interconnect is ready for
	//    use.
	//
	//    * down: The network link is down.
	//
	//    * deleting: The interconnect is being deleted.
	//
	//    * deleted: The interconnect is deleted.
	//
	//    * unknown: The state of the interconnect is not available.
	InterconnectState InterconnectState `locationName:"interconnectState" type:"string" enum:"true"`

	// Indicates whether jumbo frames (9001 MTU) are supported.
	JumboFrameCapable *bool `locationName:"jumboFrameCapable" type:"boolean"`

	// The ID of the LAG.
	LagId *string `locationName:"lagId" type:"string"`

	// The time of the most recent call to DescribeLoa for this connection.
	LoaIssueTime *time.Time `locationName:"loaIssueTime" type:"timestamp"`

	// The location of the connection.
	Location *string `locationName:"location" type:"string"`

	// The AWS Region where the connection is located.
	Region *string `locationName:"region" type:"string"`

	// Any tags assigned to the interconnect.
	Tags []Tag `locationName:"tags" min:"1" type:"list"`
}

// String returns the string representation
func (s CreateInterconnectOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateInterconnect = "CreateInterconnect"

// CreateInterconnectRequest returns a request value for making API operation for
// AWS Direct Connect.
//
// Creates an interconnect between an AWS Direct Connect Partner's network and
// a specific AWS Direct Connect location.
//
// An interconnect is a connection that is capable of hosting other connections.
// The AWS Direct Connect partner can use an interconnect to provide AWS Direct
// Connect hosted connections to customers through their own network services.
// Like a standard connection, an interconnect links the partner's network to
// an AWS Direct Connect location over a standard Ethernet fiber-optic cable.
// One end is connected to the partner's router, the other to an AWS Direct
// Connect router.
//
// You can automatically add the new interconnect to a link aggregation group
// (LAG) by specifying a LAG ID in the request. This ensures that the new interconnect
// is allocated on the same AWS Direct Connect endpoint that hosts the specified
// LAG. If there are no available ports on the endpoint, the request fails and
// no interconnect is created.
//
// For each end customer, the AWS Direct Connect Partner provisions a connection
// on their interconnect by calling AllocateHostedConnection. The end customer
// can then connect to AWS resources by creating a virtual interface on their
// connection, using the VLAN assigned to them by the AWS Direct Connect Partner.
//
// Intended for use by AWS Direct Connect Partners only.
//
//    // Example sending a request using CreateInterconnectRequest.
//    req := client.CreateInterconnectRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/directconnect-2012-10-25/CreateInterconnect
func (c *Client) CreateInterconnectRequest(input *CreateInterconnectInput) CreateInterconnectRequest {
	op := &aws.Operation{
		Name:       opCreateInterconnect,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateInterconnectInput{}
	}

	req := c.newRequest(op, input, &CreateInterconnectOutput{})
	return CreateInterconnectRequest{Request: req, Input: input, Copy: c.CreateInterconnectRequest}
}

// CreateInterconnectRequest is the request type for the
// CreateInterconnect API operation.
type CreateInterconnectRequest struct {
	*aws.Request
	Input *CreateInterconnectInput
	Copy  func(*CreateInterconnectInput) CreateInterconnectRequest
}

// Send marshals and sends the CreateInterconnect API request.
func (r CreateInterconnectRequest) Send(ctx context.Context) (*CreateInterconnectResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateInterconnectResponse{
		CreateInterconnectOutput: r.Request.Data.(*CreateInterconnectOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateInterconnectResponse is the response type for the
// CreateInterconnect API operation.
type CreateInterconnectResponse struct {
	*CreateInterconnectOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateInterconnect request.
func (r *CreateInterconnectResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
