// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/CreateCrawlerRequest
type CreateCrawlerInput struct {
	_ struct{} `type:"structure"`

	// A list of custom classifiers that the user has registered. By default, all
	// built-in classifiers are included in a crawl, but these custom classifiers
	// always override the default classifiers for a given classification.
	Classifiers []string `type:"list"`

	// The crawler configuration information. This versioned JSON string allows
	// users to specify aspects of a crawler's behavior. For more information, see
	// Configuring a Crawler (http://docs.aws.amazon.com/glue/latest/dg/crawler-configuration.html).
	Configuration *string `type:"string"`

	// The name of the SecurityConfiguration structure to be used by this crawler.
	CrawlerSecurityConfiguration *string `type:"string"`

	// The AWS Glue database where results are written, such as: arn:aws:daylight:us-east-1::database/sometable/*.
	DatabaseName *string `type:"string"`

	// A description of the new crawler.
	Description *string `type:"string"`

	// Name of the new crawler.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`

	// The IAM role or Amazon Resource Name (ARN) of an IAM role used by the new
	// crawler to access customer resources.
	//
	// Role is a required field
	Role *string `type:"string" required:"true"`

	// A cron expression used to specify the schedule. For more information, see
	// Time-Based Schedules for Jobs and Crawlers (http://docs.aws.amazon.com/glue/latest/dg/monitor-data-warehouse-schedule.html).
	// For example, to run something every day at 12:15 UTC, specify cron(15 12
	// * * ? *).
	Schedule *string `type:"string"`

	// The policy for the crawler's update and deletion behavior.
	SchemaChangePolicy *SchemaChangePolicy `type:"structure"`

	// The table prefix used for catalog tables that are created.
	TablePrefix *string `type:"string"`

	// The tags to use with this crawler request. You can use tags to limit access
	// to the crawler. For more information, see AWS Tags in AWS Glue (http://docs.aws.amazon.com/glue/latest/dg/monitor-tags.html).
	Tags map[string]string `type:"map"`

	// A list of collection of targets to crawl.
	//
	// Targets is a required field
	Targets *CrawlerTargets `type:"structure" required:"true"`
}

// String returns the string representation
func (s CreateCrawlerInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateCrawlerInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateCrawlerInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if s.Role == nil {
		invalidParams.Add(aws.NewErrParamRequired("Role"))
	}

	if s.Targets == nil {
		invalidParams.Add(aws.NewErrParamRequired("Targets"))
	}
	if s.Targets != nil {
		if err := s.Targets.Validate(); err != nil {
			invalidParams.AddNested("Targets", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/CreateCrawlerResponse
type CreateCrawlerOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CreateCrawlerOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateCrawler = "CreateCrawler"

// CreateCrawlerRequest returns a request value for making API operation for
// AWS Glue.
//
// Creates a new crawler with specified targets, role, configuration, and optional
// schedule. At least one crawl target must be specified, in the s3Targets field,
// the jdbcTargets field, or the DynamoDBTargets field.
//
//    // Example sending a request using CreateCrawlerRequest.
//    req := client.CreateCrawlerRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/CreateCrawler
func (c *Client) CreateCrawlerRequest(input *CreateCrawlerInput) CreateCrawlerRequest {
	op := &aws.Operation{
		Name:       opCreateCrawler,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateCrawlerInput{}
	}

	req := c.newRequest(op, input, &CreateCrawlerOutput{})
	return CreateCrawlerRequest{Request: req, Input: input, Copy: c.CreateCrawlerRequest}
}

// CreateCrawlerRequest is the request type for the
// CreateCrawler API operation.
type CreateCrawlerRequest struct {
	*aws.Request
	Input *CreateCrawlerInput
	Copy  func(*CreateCrawlerInput) CreateCrawlerRequest
}

// Send marshals and sends the CreateCrawler API request.
func (r CreateCrawlerRequest) Send(ctx context.Context) (*CreateCrawlerResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateCrawlerResponse{
		CreateCrawlerOutput: r.Request.Data.(*CreateCrawlerOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateCrawlerResponse is the response type for the
// CreateCrawler API operation.
type CreateCrawlerResponse struct {
	*CreateCrawlerOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateCrawler request.
func (r *CreateCrawlerResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
