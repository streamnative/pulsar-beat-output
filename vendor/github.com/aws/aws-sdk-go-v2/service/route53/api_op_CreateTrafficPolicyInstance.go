// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// A complex type that contains information about the resource record sets that
// you want to create based on a specified traffic policy.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53-2013-04-01/CreateTrafficPolicyInstanceRequest
type CreateTrafficPolicyInstanceInput struct {
	_ struct{} `locationName:"CreateTrafficPolicyInstanceRequest" type:"structure" xmlURI:"https://route53.amazonaws.com/doc/2013-04-01/"`

	// The ID of the hosted zone that you want Amazon Route 53 to create resource
	// record sets in by using the configuration in a traffic policy.
	//
	// HostedZoneId is a required field
	HostedZoneId *string `type:"string" required:"true"`

	// The domain name (such as example.com) or subdomain name (such as www.example.com)
	// for which Amazon Route 53 responds to DNS queries by using the resource record
	// sets that Route 53 creates for this traffic policy instance.
	//
	// Name is a required field
	Name *string `type:"string" required:"true"`

	// (Optional) The TTL that you want Amazon Route 53 to assign to all of the
	// resource record sets that it creates in the specified hosted zone.
	//
	// TTL is a required field
	TTL *int64 `type:"long" required:"true"`

	// The ID of the traffic policy that you want to use to create resource record
	// sets in the specified hosted zone.
	//
	// TrafficPolicyId is a required field
	TrafficPolicyId *string `min:"1" type:"string" required:"true"`

	// The version of the traffic policy that you want to use to create resource
	// record sets in the specified hosted zone.
	//
	// TrafficPolicyVersion is a required field
	TrafficPolicyVersion *int64 `min:"1" type:"integer" required:"true"`
}

// String returns the string representation
func (s CreateTrafficPolicyInstanceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateTrafficPolicyInstanceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateTrafficPolicyInstanceInput"}

	if s.HostedZoneId == nil {
		invalidParams.Add(aws.NewErrParamRequired("HostedZoneId"))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if s.TTL == nil {
		invalidParams.Add(aws.NewErrParamRequired("TTL"))
	}

	if s.TrafficPolicyId == nil {
		invalidParams.Add(aws.NewErrParamRequired("TrafficPolicyId"))
	}
	if s.TrafficPolicyId != nil && len(*s.TrafficPolicyId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TrafficPolicyId", 1))
	}

	if s.TrafficPolicyVersion == nil {
		invalidParams.Add(aws.NewErrParamRequired("TrafficPolicyVersion"))
	}
	if s.TrafficPolicyVersion != nil && *s.TrafficPolicyVersion < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("TrafficPolicyVersion", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateTrafficPolicyInstanceInput) MarshalFields(e protocol.FieldEncoder) error {

	e.SetFields(protocol.BodyTarget, "CreateTrafficPolicyInstanceRequest", protocol.FieldMarshalerFunc(func(e protocol.FieldEncoder) error {
		if s.HostedZoneId != nil {
			v := *s.HostedZoneId

			metadata := protocol.Metadata{}
			e.SetValue(protocol.BodyTarget, "HostedZoneId", protocol.StringValue(v), metadata)
		}
		if s.Name != nil {
			v := *s.Name

			metadata := protocol.Metadata{}
			e.SetValue(protocol.BodyTarget, "Name", protocol.StringValue(v), metadata)
		}
		if s.TTL != nil {
			v := *s.TTL

			metadata := protocol.Metadata{}
			e.SetValue(protocol.BodyTarget, "TTL", protocol.Int64Value(v), metadata)
		}
		if s.TrafficPolicyId != nil {
			v := *s.TrafficPolicyId

			metadata := protocol.Metadata{}
			e.SetValue(protocol.BodyTarget, "TrafficPolicyId", protocol.StringValue(v), metadata)
		}
		if s.TrafficPolicyVersion != nil {
			v := *s.TrafficPolicyVersion

			metadata := protocol.Metadata{}
			e.SetValue(protocol.BodyTarget, "TrafficPolicyVersion", protocol.Int64Value(v), metadata)
		}
		return nil
	}), protocol.Metadata{XMLNamespaceURI: "https://route53.amazonaws.com/doc/2013-04-01/"})
	return nil
}

// A complex type that contains the response information for the CreateTrafficPolicyInstance
// request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53-2013-04-01/CreateTrafficPolicyInstanceResponse
type CreateTrafficPolicyInstanceOutput struct {
	_ struct{} `type:"structure"`

	// A unique URL that represents a new traffic policy instance.
	//
	// Location is a required field
	Location *string `location:"header" locationName:"Location" type:"string" required:"true"`

	// A complex type that contains settings for the new traffic policy instance.
	//
	// TrafficPolicyInstance is a required field
	TrafficPolicyInstance *TrafficPolicyInstance `type:"structure" required:"true"`
}

// String returns the string representation
func (s CreateTrafficPolicyInstanceOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateTrafficPolicyInstanceOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.TrafficPolicyInstance != nil {
		v := s.TrafficPolicyInstance

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "TrafficPolicyInstance", v, metadata)
	}
	if s.Location != nil {
		v := *s.Location

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Location", protocol.StringValue(v), metadata)
	}
	return nil
}

const opCreateTrafficPolicyInstance = "CreateTrafficPolicyInstance"

// CreateTrafficPolicyInstanceRequest returns a request value for making API operation for
// Amazon Route 53.
//
// Creates resource record sets in a specified hosted zone based on the settings
// in a specified traffic policy version. In addition, CreateTrafficPolicyInstance
// associates the resource record sets with a specified domain name (such as
// example.com) or subdomain name (such as www.example.com). Amazon Route 53
// responds to DNS queries for the domain or subdomain name by using the resource
// record sets that CreateTrafficPolicyInstance created.
//
//    // Example sending a request using CreateTrafficPolicyInstanceRequest.
//    req := client.CreateTrafficPolicyInstanceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53-2013-04-01/CreateTrafficPolicyInstance
func (c *Client) CreateTrafficPolicyInstanceRequest(input *CreateTrafficPolicyInstanceInput) CreateTrafficPolicyInstanceRequest {
	op := &aws.Operation{
		Name:       opCreateTrafficPolicyInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/2013-04-01/trafficpolicyinstance",
	}

	if input == nil {
		input = &CreateTrafficPolicyInstanceInput{}
	}

	req := c.newRequest(op, input, &CreateTrafficPolicyInstanceOutput{})
	return CreateTrafficPolicyInstanceRequest{Request: req, Input: input, Copy: c.CreateTrafficPolicyInstanceRequest}
}

// CreateTrafficPolicyInstanceRequest is the request type for the
// CreateTrafficPolicyInstance API operation.
type CreateTrafficPolicyInstanceRequest struct {
	*aws.Request
	Input *CreateTrafficPolicyInstanceInput
	Copy  func(*CreateTrafficPolicyInstanceInput) CreateTrafficPolicyInstanceRequest
}

// Send marshals and sends the CreateTrafficPolicyInstance API request.
func (r CreateTrafficPolicyInstanceRequest) Send(ctx context.Context) (*CreateTrafficPolicyInstanceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateTrafficPolicyInstanceResponse{
		CreateTrafficPolicyInstanceOutput: r.Request.Data.(*CreateTrafficPolicyInstanceOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateTrafficPolicyInstanceResponse is the response type for the
// CreateTrafficPolicyInstance API operation.
type CreateTrafficPolicyInstanceResponse struct {
	*CreateTrafficPolicyInstanceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateTrafficPolicyInstance request.
func (r *CreateTrafficPolicyInstanceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
