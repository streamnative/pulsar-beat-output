// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package directconnect

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/directconnect-2012-10-25/DescribeInterconnectLoaRequest
type DescribeInterconnectLoaInput struct {
	_ struct{} `type:"structure"`

	// The ID of the interconnect.
	//
	// InterconnectId is a required field
	InterconnectId *string `locationName:"interconnectId" type:"string" required:"true"`

	// The standard media type for the LOA-CFA document. The only supported value
	// is application/pdf.
	LoaContentType LoaContentType `locationName:"loaContentType" type:"string" enum:"true"`

	// The name of the service provider who establishes connectivity on your behalf.
	// If you supply this parameter, the LOA-CFA lists the provider name alongside
	// your company name as the requester of the cross connect.
	ProviderName *string `locationName:"providerName" type:"string"`
}

// String returns the string representation
func (s DescribeInterconnectLoaInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeInterconnectLoaInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeInterconnectLoaInput"}

	if s.InterconnectId == nil {
		invalidParams.Add(aws.NewErrParamRequired("InterconnectId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/directconnect-2012-10-25/DescribeInterconnectLoaResponse
type DescribeInterconnectLoaOutput struct {
	_ struct{} `type:"structure"`

	// The Letter of Authorization - Connecting Facility Assignment (LOA-CFA).
	Loa *Loa `locationName:"loa" type:"structure"`
}

// String returns the string representation
func (s DescribeInterconnectLoaOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeInterconnectLoa = "DescribeInterconnectLoa"

// DescribeInterconnectLoaRequest returns a request value for making API operation for
// AWS Direct Connect.
//
// Deprecated. Use DescribeLoa instead.
//
// Gets the LOA-CFA for the specified interconnect.
//
// The Letter of Authorization - Connecting Facility Assignment (LOA-CFA) is
// a document that is used when establishing your cross connect to AWS at the
// colocation facility. For more information, see Requesting Cross Connects
// at AWS Direct Connect Locations (https://docs.aws.amazon.com/directconnect/latest/UserGuide/Colocation.html)
// in the AWS Direct Connect User Guide.
//
//    // Example sending a request using DescribeInterconnectLoaRequest.
//    req := client.DescribeInterconnectLoaRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/directconnect-2012-10-25/DescribeInterconnectLoa
func (c *Client) DescribeInterconnectLoaRequest(input *DescribeInterconnectLoaInput) DescribeInterconnectLoaRequest {
	if c.Client.Config.Logger != nil {
		c.Client.Config.Logger.Log("This operation, DescribeInterconnectLoa, has been deprecated")
	}
	op := &aws.Operation{
		Name:       opDescribeInterconnectLoa,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeInterconnectLoaInput{}
	}

	req := c.newRequest(op, input, &DescribeInterconnectLoaOutput{})
	return DescribeInterconnectLoaRequest{Request: req, Input: input, Copy: c.DescribeInterconnectLoaRequest}
}

// DescribeInterconnectLoaRequest is the request type for the
// DescribeInterconnectLoa API operation.
type DescribeInterconnectLoaRequest struct {
	*aws.Request
	Input *DescribeInterconnectLoaInput
	Copy  func(*DescribeInterconnectLoaInput) DescribeInterconnectLoaRequest
}

// Send marshals and sends the DescribeInterconnectLoa API request.
func (r DescribeInterconnectLoaRequest) Send(ctx context.Context) (*DescribeInterconnectLoaResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeInterconnectLoaResponse{
		DescribeInterconnectLoaOutput: r.Request.Data.(*DescribeInterconnectLoaOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeInterconnectLoaResponse is the response type for the
// DescribeInterconnectLoa API operation.
type DescribeInterconnectLoaResponse struct {
	*DescribeInterconnectLoaOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeInterconnectLoa request.
func (r *DescribeInterconnectLoaResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
