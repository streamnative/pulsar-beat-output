// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package comprehendmedical

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/comprehendmedical-2018-10-30/DetectEntitiesRequest
type DetectEntitiesInput struct {
	_ struct{} `type:"structure"`

	// A UTF-8 text string containing the clinical content being examined for entities.
	// Each string must contain fewer than 20,000 bytes of characters.
	//
	// Text is a required field
	Text *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DetectEntitiesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DetectEntitiesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DetectEntitiesInput"}

	if s.Text == nil {
		invalidParams.Add(aws.NewErrParamRequired("Text"))
	}
	if s.Text != nil && len(*s.Text) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Text", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/comprehendmedical-2018-10-30/DetectEntitiesResponse
type DetectEntitiesOutput struct {
	_ struct{} `type:"structure"`

	// The collection of medical entities extracted from the input text and their
	// associated information. For each entity, the response provides the entity
	// text, the entity category, where the entity text begins and ends, and the
	// level of confidence that Comprehend Medical has in the detection and analysis.
	// Attributes and traits of the entity are also returned.
	//
	// Entities is a required field
	Entities []Entity `type:"list" required:"true"`

	// If the result of the previous request to DetectEntities was truncated, include
	// the Paginationtoken to fetch the next page of entities.
	PaginationToken *string `min:"1" type:"string"`

	// Attributes extracted from the input text that we were unable to relate to
	// an entity.
	UnmappedAttributes []UnmappedAttribute `type:"list"`
}

// String returns the string representation
func (s DetectEntitiesOutput) String() string {
	return awsutil.Prettify(s)
}

const opDetectEntities = "DetectEntities"

// DetectEntitiesRequest returns a request value for making API operation for
// AWS Comprehend Medical.
//
// Inspects the clinical text for a variety of medical entities and returns
// specific information about them such as entity category, location, and confidence
// score on that information .
//
//    // Example sending a request using DetectEntitiesRequest.
//    req := client.DetectEntitiesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/comprehendmedical-2018-10-30/DetectEntities
func (c *Client) DetectEntitiesRequest(input *DetectEntitiesInput) DetectEntitiesRequest {
	op := &aws.Operation{
		Name:       opDetectEntities,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DetectEntitiesInput{}
	}

	req := c.newRequest(op, input, &DetectEntitiesOutput{})
	return DetectEntitiesRequest{Request: req, Input: input, Copy: c.DetectEntitiesRequest}
}

// DetectEntitiesRequest is the request type for the
// DetectEntities API operation.
type DetectEntitiesRequest struct {
	*aws.Request
	Input *DetectEntitiesInput
	Copy  func(*DetectEntitiesInput) DetectEntitiesRequest
}

// Send marshals and sends the DetectEntities API request.
func (r DetectEntitiesRequest) Send(ctx context.Context) (*DetectEntitiesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DetectEntitiesResponse{
		DetectEntitiesOutput: r.Request.Data.(*DetectEntitiesOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DetectEntitiesResponse is the response type for the
// DetectEntities API operation.
type DetectEntitiesResponse struct {
	*DetectEntitiesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DetectEntities request.
func (r *DetectEntitiesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
