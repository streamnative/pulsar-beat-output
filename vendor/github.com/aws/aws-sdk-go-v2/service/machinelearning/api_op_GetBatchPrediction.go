// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package machinelearning

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetBatchPredictionInput struct {
	_ struct{} `type:"structure"`

	// An ID assigned to the BatchPrediction at creation.
	//
	// BatchPredictionId is a required field
	BatchPredictionId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetBatchPredictionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetBatchPredictionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetBatchPredictionInput"}

	if s.BatchPredictionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("BatchPredictionId"))
	}
	if s.BatchPredictionId != nil && len(*s.BatchPredictionId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("BatchPredictionId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of a GetBatchPrediction operation and describes a BatchPrediction.
type GetBatchPredictionOutput struct {
	_ struct{} `type:"structure"`

	// The ID of the DataSource that was used to create the BatchPrediction.
	BatchPredictionDataSourceId *string `min:"1" type:"string"`

	// An ID assigned to the BatchPrediction at creation. This value should be identical
	// to the value of the BatchPredictionID in the request.
	BatchPredictionId *string `min:"1" type:"string"`

	// The approximate CPU time in milliseconds that Amazon Machine Learning spent
	// processing the BatchPrediction, normalized and scaled on computation resources.
	// ComputeTime is only available if the BatchPrediction is in the COMPLETED
	// state.
	ComputeTime *int64 `type:"long"`

	// The time when the BatchPrediction was created. The time is expressed in epoch
	// time.
	CreatedAt *time.Time `type:"timestamp"`

	// The AWS user account that invoked the BatchPrediction. The account type can
	// be either an AWS root account or an AWS Identity and Access Management (IAM)
	// user account.
	CreatedByIamUser *string `type:"string"`

	// The epoch time when Amazon Machine Learning marked the BatchPrediction as
	// COMPLETED or FAILED. FinishedAt is only available when the BatchPrediction
	// is in the COMPLETED or FAILED state.
	FinishedAt *time.Time `type:"timestamp"`

	// The location of the data file or directory in Amazon Simple Storage Service
	// (Amazon S3).
	InputDataLocationS3 *string `type:"string"`

	// The number of invalid records that Amazon Machine Learning saw while processing
	// the BatchPrediction.
	InvalidRecordCount *int64 `type:"long"`

	// The time of the most recent edit to BatchPrediction. The time is expressed
	// in epoch time.
	LastUpdatedAt *time.Time `type:"timestamp"`

	// A link to the file that contains logs of the CreateBatchPrediction operation.
	LogUri *string `type:"string"`

	// The ID of the MLModel that generated predictions for the BatchPrediction
	// request.
	MLModelId *string `min:"1" type:"string"`

	// A description of the most recent details about processing the batch prediction
	// request.
	Message *string `type:"string"`

	// A user-supplied name or description of the BatchPrediction.
	Name *string `type:"string"`

	// The location of an Amazon S3 bucket or directory to receive the operation
	// results.
	OutputUri *string `type:"string"`

	// The epoch time when Amazon Machine Learning marked the BatchPrediction as
	// INPROGRESS. StartedAt isn't available if the BatchPrediction is in the PENDING
	// state.
	StartedAt *time.Time `type:"timestamp"`

	// The status of the BatchPrediction, which can be one of the following values:
	//
	//    * PENDING - Amazon Machine Learning (Amazon ML) submitted a request to
	//    generate batch predictions.
	//
	//    * INPROGRESS - The batch predictions are in progress.
	//
	//    * FAILED - The request to perform a batch prediction did not run to completion.
	//    It is not usable.
	//
	//    * COMPLETED - The batch prediction process completed successfully.
	//
	//    * DELETED - The BatchPrediction is marked as deleted. It is not usable.
	Status EntityStatus `type:"string" enum:"true"`

	// The number of total records that Amazon Machine Learning saw while processing
	// the BatchPrediction.
	TotalRecordCount *int64 `type:"long"`
}

// String returns the string representation
func (s GetBatchPredictionOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetBatchPrediction = "GetBatchPrediction"

// GetBatchPredictionRequest returns a request value for making API operation for
// Amazon Machine Learning.
//
// Returns a BatchPrediction that includes detailed metadata, status, and data
// file information for a Batch Prediction request.
//
//    // Example sending a request using GetBatchPredictionRequest.
//    req := client.GetBatchPredictionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) GetBatchPredictionRequest(input *GetBatchPredictionInput) GetBatchPredictionRequest {
	op := &aws.Operation{
		Name:       opGetBatchPrediction,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetBatchPredictionInput{}
	}

	req := c.newRequest(op, input, &GetBatchPredictionOutput{})
	return GetBatchPredictionRequest{Request: req, Input: input, Copy: c.GetBatchPredictionRequest}
}

// GetBatchPredictionRequest is the request type for the
// GetBatchPrediction API operation.
type GetBatchPredictionRequest struct {
	*aws.Request
	Input *GetBatchPredictionInput
	Copy  func(*GetBatchPredictionInput) GetBatchPredictionRequest
}

// Send marshals and sends the GetBatchPrediction API request.
func (r GetBatchPredictionRequest) Send(ctx context.Context) (*GetBatchPredictionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetBatchPredictionResponse{
		GetBatchPredictionOutput: r.Request.Data.(*GetBatchPredictionOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetBatchPredictionResponse is the response type for the
// GetBatchPrediction API operation.
type GetBatchPredictionResponse struct {
	*GetBatchPredictionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetBatchPrediction request.
func (r *GetBatchPredictionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
