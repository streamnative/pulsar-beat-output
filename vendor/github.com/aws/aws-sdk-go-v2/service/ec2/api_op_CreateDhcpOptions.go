// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CreateDhcpOptionsRequest
type CreateDhcpOptionsInput struct {
	_ struct{} `type:"structure"`

	// A DHCP configuration option.
	//
	// DhcpConfigurations is a required field
	DhcpConfigurations []NewDhcpConfiguration `locationName:"dhcpConfiguration" locationNameList:"item" type:"list" required:"true"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`
}

// String returns the string representation
func (s CreateDhcpOptionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDhcpOptionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateDhcpOptionsInput"}

	if s.DhcpConfigurations == nil {
		invalidParams.Add(aws.NewErrParamRequired("DhcpConfigurations"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CreateDhcpOptionsResult
type CreateDhcpOptionsOutput struct {
	_ struct{} `type:"structure"`

	// A set of DHCP options.
	DhcpOptions *DhcpOptions `locationName:"dhcpOptions" type:"structure"`
}

// String returns the string representation
func (s CreateDhcpOptionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateDhcpOptions = "CreateDhcpOptions"

// CreateDhcpOptionsRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Creates a set of DHCP options for your VPC. After creating the set, you must
// associate it with the VPC, causing all existing and new instances that you
// launch in the VPC to use this set of DHCP options. The following are the
// individual DHCP options you can specify. For more information about the options,
// see RFC 2132 (http://www.ietf.org/rfc/rfc2132.txt).
//
//    * domain-name-servers - The IP addresses of up to four domain name servers,
//    or AmazonProvidedDNS. The default DHCP option set specifies AmazonProvidedDNS.
//    If specifying more than one domain name server, specify the IP addresses
//    in a single parameter, separated by commas. To have your instance receive
//    a custom DNS hostname as specified in domain-name, you must set domain-name-servers
//    to a custom DNS server.
//
//    * domain-name - If you're using AmazonProvidedDNS in us-east-1, specify
//    ec2.internal. If you're using AmazonProvidedDNS in another Region, specify
//    region.compute.internal (for example, ap-northeast-1.compute.internal).
//    Otherwise, specify a domain name (for example, MyCompany.com). This value
//    is used to complete unqualified DNS hostnames. Important: Some Linux operating
//    systems accept multiple domain names separated by spaces. However, Windows
//    and other Linux operating systems treat the value as a single domain,
//    which results in unexpected behavior. If your DHCP options set is associated
//    with a VPC that has instances with multiple operating systems, specify
//    only one domain name.
//
//    * ntp-servers - The IP addresses of up to four Network Time Protocol (NTP)
//    servers.
//
//    * netbios-name-servers - The IP addresses of up to four NetBIOS name servers.
//
//    * netbios-node-type - The NetBIOS node type (1, 2, 4, or 8). We recommend
//    that you specify 2 (broadcast and multicast are not currently supported).
//    For more information about these node types, see RFC 2132 (http://www.ietf.org/rfc/rfc2132.txt).
//
// Your VPC automatically starts out with a set of DHCP options that includes
// only a DNS server that we provide (AmazonProvidedDNS). If you create a set
// of options, and if your VPC has an internet gateway, make sure to set the
// domain-name-servers option either to AmazonProvidedDNS or to a domain name
// server of your choice. For more information, see DHCP Options Sets (https://docs.aws.amazon.com/vpc/latest/userguide/VPC_DHCP_Options.html)
// in the Amazon Virtual Private Cloud User Guide.
//
//    // Example sending a request using CreateDhcpOptionsRequest.
//    req := client.CreateDhcpOptionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CreateDhcpOptions
func (c *Client) CreateDhcpOptionsRequest(input *CreateDhcpOptionsInput) CreateDhcpOptionsRequest {
	op := &aws.Operation{
		Name:       opCreateDhcpOptions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateDhcpOptionsInput{}
	}

	req := c.newRequest(op, input, &CreateDhcpOptionsOutput{})
	return CreateDhcpOptionsRequest{Request: req, Input: input, Copy: c.CreateDhcpOptionsRequest}
}

// CreateDhcpOptionsRequest is the request type for the
// CreateDhcpOptions API operation.
type CreateDhcpOptionsRequest struct {
	*aws.Request
	Input *CreateDhcpOptionsInput
	Copy  func(*CreateDhcpOptionsInput) CreateDhcpOptionsRequest
}

// Send marshals and sends the CreateDhcpOptions API request.
func (r CreateDhcpOptionsRequest) Send(ctx context.Context) (*CreateDhcpOptionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateDhcpOptionsResponse{
		CreateDhcpOptionsOutput: r.Request.Data.(*CreateDhcpOptionsOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateDhcpOptionsResponse is the response type for the
// CreateDhcpOptions API operation.
type CreateDhcpOptionsResponse struct {
	*CreateDhcpOptionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateDhcpOptions request.
func (r *CreateDhcpOptionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
