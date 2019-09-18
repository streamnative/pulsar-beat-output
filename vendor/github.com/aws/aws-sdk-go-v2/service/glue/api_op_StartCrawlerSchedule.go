// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/StartCrawlerScheduleRequest
type StartCrawlerScheduleInput struct {
	_ struct{} `type:"structure"`

	// Name of the crawler to schedule.
	//
	// CrawlerName is a required field
	CrawlerName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s StartCrawlerScheduleInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartCrawlerScheduleInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartCrawlerScheduleInput"}

	if s.CrawlerName == nil {
		invalidParams.Add(aws.NewErrParamRequired("CrawlerName"))
	}
	if s.CrawlerName != nil && len(*s.CrawlerName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CrawlerName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/StartCrawlerScheduleResponse
type StartCrawlerScheduleOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s StartCrawlerScheduleOutput) String() string {
	return awsutil.Prettify(s)
}

const opStartCrawlerSchedule = "StartCrawlerSchedule"

// StartCrawlerScheduleRequest returns a request value for making API operation for
// AWS Glue.
//
// Changes the schedule state of the specified crawler to SCHEDULED, unless
// the crawler is already running or the schedule state is already SCHEDULED.
//
//    // Example sending a request using StartCrawlerScheduleRequest.
//    req := client.StartCrawlerScheduleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/StartCrawlerSchedule
func (c *Client) StartCrawlerScheduleRequest(input *StartCrawlerScheduleInput) StartCrawlerScheduleRequest {
	op := &aws.Operation{
		Name:       opStartCrawlerSchedule,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StartCrawlerScheduleInput{}
	}

	req := c.newRequest(op, input, &StartCrawlerScheduleOutput{})
	return StartCrawlerScheduleRequest{Request: req, Input: input, Copy: c.StartCrawlerScheduleRequest}
}

// StartCrawlerScheduleRequest is the request type for the
// StartCrawlerSchedule API operation.
type StartCrawlerScheduleRequest struct {
	*aws.Request
	Input *StartCrawlerScheduleInput
	Copy  func(*StartCrawlerScheduleInput) StartCrawlerScheduleRequest
}

// Send marshals and sends the StartCrawlerSchedule API request.
func (r StartCrawlerScheduleRequest) Send(ctx context.Context) (*StartCrawlerScheduleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartCrawlerScheduleResponse{
		StartCrawlerScheduleOutput: r.Request.Data.(*StartCrawlerScheduleOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartCrawlerScheduleResponse is the response type for the
// StartCrawlerSchedule API operation.
type StartCrawlerScheduleResponse struct {
	*StartCrawlerScheduleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartCrawlerSchedule request.
func (r *StartCrawlerScheduleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
