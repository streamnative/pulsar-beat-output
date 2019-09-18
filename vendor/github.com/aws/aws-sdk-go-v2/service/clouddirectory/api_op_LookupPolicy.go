// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package clouddirectory

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/LookupPolicyRequest
type LookupPolicyInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) that is associated with the Directory. For
	// more information, see arns.
	//
	// DirectoryArn is a required field
	DirectoryArn *string `location:"header" locationName:"x-amz-data-partition" type:"string" required:"true"`

	// The maximum number of items to be retrieved in a single call. This is an
	// approximate number.
	MaxResults *int64 `min:"1" type:"integer"`

	// The token to request the next page of results.
	NextToken *string `type:"string"`

	// Reference that identifies the object whose policies will be looked up.
	//
	// ObjectReference is a required field
	ObjectReference *ObjectReference `type:"structure" required:"true"`
}

// String returns the string representation
func (s LookupPolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *LookupPolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "LookupPolicyInput"}

	if s.DirectoryArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("DirectoryArn"))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if s.ObjectReference == nil {
		invalidParams.Add(aws.NewErrParamRequired("ObjectReference"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s LookupPolicyInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "MaxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ObjectReference != nil {
		v := s.ObjectReference

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "ObjectReference", v, metadata)
	}
	if s.DirectoryArn != nil {
		v := *s.DirectoryArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-data-partition", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/LookupPolicyResponse
type LookupPolicyOutput struct {
	_ struct{} `type:"structure"`

	// The pagination token.
	NextToken *string `type:"string"`

	// Provides list of path to policies. Policies contain PolicyId, ObjectIdentifier,
	// and PolicyType. For more information, see Policies (https://docs.aws.amazon.com/clouddirectory/latest/developerguide/key_concepts_directory.html#key_concepts_policies).
	PolicyToPathList []PolicyToPath `type:"list"`
}

// String returns the string representation
func (s LookupPolicyOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s LookupPolicyOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PolicyToPathList != nil {
		v := s.PolicyToPathList

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "PolicyToPathList", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opLookupPolicy = "LookupPolicy"

// LookupPolicyRequest returns a request value for making API operation for
// Amazon CloudDirectory.
//
// Lists all policies from the root of the Directory to the object specified.
// If there are no policies present, an empty list is returned. If policies
// are present, and if some objects don't have the policies attached, it returns
// the ObjectIdentifier for such objects. If policies are present, it returns
// ObjectIdentifier, policyId, and policyType. Paths that don't lead to the
// root from the target object are ignored. For more information, see Policies
// (https://docs.aws.amazon.com/clouddirectory/latest/developerguide/key_concepts_directory.html#key_concepts_policies).
//
//    // Example sending a request using LookupPolicyRequest.
//    req := client.LookupPolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/LookupPolicy
func (c *Client) LookupPolicyRequest(input *LookupPolicyInput) LookupPolicyRequest {
	op := &aws.Operation{
		Name:       opLookupPolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/amazonclouddirectory/2017-01-11/policy/lookup",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &LookupPolicyInput{}
	}

	req := c.newRequest(op, input, &LookupPolicyOutput{})
	return LookupPolicyRequest{Request: req, Input: input, Copy: c.LookupPolicyRequest}
}

// LookupPolicyRequest is the request type for the
// LookupPolicy API operation.
type LookupPolicyRequest struct {
	*aws.Request
	Input *LookupPolicyInput
	Copy  func(*LookupPolicyInput) LookupPolicyRequest
}

// Send marshals and sends the LookupPolicy API request.
func (r LookupPolicyRequest) Send(ctx context.Context) (*LookupPolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &LookupPolicyResponse{
		LookupPolicyOutput: r.Request.Data.(*LookupPolicyOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewLookupPolicyRequestPaginator returns a paginator for LookupPolicy.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.LookupPolicyRequest(input)
//   p := clouddirectory.NewLookupPolicyRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewLookupPolicyPaginator(req LookupPolicyRequest) LookupPolicyPaginator {
	return LookupPolicyPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *LookupPolicyInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// LookupPolicyPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type LookupPolicyPaginator struct {
	aws.Pager
}

func (p *LookupPolicyPaginator) CurrentPage() *LookupPolicyOutput {
	return p.Pager.CurrentPage().(*LookupPolicyOutput)
}

// LookupPolicyResponse is the response type for the
// LookupPolicy API operation.
type LookupPolicyResponse struct {
	*LookupPolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// LookupPolicy request.
func (r *LookupPolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
