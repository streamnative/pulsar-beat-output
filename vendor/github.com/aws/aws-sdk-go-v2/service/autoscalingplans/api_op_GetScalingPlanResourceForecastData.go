// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package autoscalingplans

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/autoscaling-plans-2018-01-06/GetScalingPlanResourceForecastDataRequest
type GetScalingPlanResourceForecastDataInput struct {
	_ struct{} `type:"structure"`

	// The exclusive end time of the time range for the forecast data to get. The
	// maximum time duration between the start and end time is seven days.
	//
	// Although this parameter can accept a date and time that is more than two
	// days in the future, the availability of forecast data has limits. AWS Auto
	// Scaling only issues forecasts for periods of two days in advance.
	//
	// EndTime is a required field
	EndTime *time.Time `type:"timestamp" required:"true"`

	// The type of forecast data to get.
	//
	//    * LoadForecast: The load metric forecast.
	//
	//    * CapacityForecast: The capacity forecast.
	//
	//    * ScheduledActionMinCapacity: The minimum capacity for each scheduled
	//    scaling action. This data is calculated as the larger of two values: the
	//    capacity forecast or the minimum capacity in the scaling instruction.
	//
	//    * ScheduledActionMaxCapacity: The maximum capacity for each scheduled
	//    scaling action. The calculation used is determined by the predictive scaling
	//    maximum capacity behavior setting in the scaling instruction.
	//
	// ForecastDataType is a required field
	ForecastDataType ForecastDataType `type:"string" required:"true" enum:"true"`

	// The ID of the resource. This string consists of the resource type and unique
	// identifier.
	//
	//    * Auto Scaling group - The resource type is autoScalingGroup and the unique
	//    identifier is the name of the Auto Scaling group. Example: autoScalingGroup/my-asg.
	//
	//    * ECS service - The resource type is service and the unique identifier
	//    is the cluster name and service name. Example: service/default/sample-webapp.
	//
	//    * Spot Fleet request - The resource type is spot-fleet-request and the
	//    unique identifier is the Spot Fleet request ID. Example: spot-fleet-request/sfr-73fbd2ce-aa30-494c-8788-1cee4EXAMPLE.
	//
	//    * DynamoDB table - The resource type is table and the unique identifier
	//    is the resource ID. Example: table/my-table.
	//
	//    * DynamoDB global secondary index - The resource type is index and the
	//    unique identifier is the resource ID. Example: table/my-table/index/my-table-index.
	//
	//    * Aurora DB cluster - The resource type is cluster and the unique identifier
	//    is the cluster name. Example: cluster:my-db-cluster.
	//
	// ResourceId is a required field
	ResourceId *string `type:"string" required:"true"`

	// The scalable dimension for the resource.
	//
	// ScalableDimension is a required field
	ScalableDimension ScalableDimension `type:"string" required:"true" enum:"true"`

	// The name of the scaling plan.
	//
	// ScalingPlanName is a required field
	ScalingPlanName *string `min:"1" type:"string" required:"true"`

	// The version number of the scaling plan.
	//
	// ScalingPlanVersion is a required field
	ScalingPlanVersion *int64 `type:"long" required:"true"`

	// The namespace of the AWS service.
	//
	// ServiceNamespace is a required field
	ServiceNamespace ServiceNamespace `type:"string" required:"true" enum:"true"`

	// The inclusive start time of the time range for the forecast data to get.
	// The date and time can be at most 56 days before the current date and time.
	//
	// StartTime is a required field
	StartTime *time.Time `type:"timestamp" required:"true"`
}

// String returns the string representation
func (s GetScalingPlanResourceForecastDataInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetScalingPlanResourceForecastDataInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetScalingPlanResourceForecastDataInput"}

	if s.EndTime == nil {
		invalidParams.Add(aws.NewErrParamRequired("EndTime"))
	}
	if len(s.ForecastDataType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("ForecastDataType"))
	}

	if s.ResourceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceId"))
	}
	if len(s.ScalableDimension) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("ScalableDimension"))
	}

	if s.ScalingPlanName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ScalingPlanName"))
	}
	if s.ScalingPlanName != nil && len(*s.ScalingPlanName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ScalingPlanName", 1))
	}

	if s.ScalingPlanVersion == nil {
		invalidParams.Add(aws.NewErrParamRequired("ScalingPlanVersion"))
	}
	if len(s.ServiceNamespace) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("ServiceNamespace"))
	}

	if s.StartTime == nil {
		invalidParams.Add(aws.NewErrParamRequired("StartTime"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/autoscaling-plans-2018-01-06/GetScalingPlanResourceForecastDataResponse
type GetScalingPlanResourceForecastDataOutput struct {
	_ struct{} `type:"structure"`

	// The data points to return.
	//
	// Datapoints is a required field
	Datapoints []Datapoint `type:"list" required:"true"`
}

// String returns the string representation
func (s GetScalingPlanResourceForecastDataOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetScalingPlanResourceForecastData = "GetScalingPlanResourceForecastData"

// GetScalingPlanResourceForecastDataRequest returns a request value for making API operation for
// AWS Auto Scaling Plans.
//
// Retrieves the forecast data for a scalable resource.
//
// Capacity forecasts are represented as predicted values, or data points, that
// are calculated using historical data points from a specified CloudWatch load
// metric. Data points are available for up to 56 days.
//
//    // Example sending a request using GetScalingPlanResourceForecastDataRequest.
//    req := client.GetScalingPlanResourceForecastDataRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/autoscaling-plans-2018-01-06/GetScalingPlanResourceForecastData
func (c *Client) GetScalingPlanResourceForecastDataRequest(input *GetScalingPlanResourceForecastDataInput) GetScalingPlanResourceForecastDataRequest {
	op := &aws.Operation{
		Name:       opGetScalingPlanResourceForecastData,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetScalingPlanResourceForecastDataInput{}
	}

	req := c.newRequest(op, input, &GetScalingPlanResourceForecastDataOutput{})
	return GetScalingPlanResourceForecastDataRequest{Request: req, Input: input, Copy: c.GetScalingPlanResourceForecastDataRequest}
}

// GetScalingPlanResourceForecastDataRequest is the request type for the
// GetScalingPlanResourceForecastData API operation.
type GetScalingPlanResourceForecastDataRequest struct {
	*aws.Request
	Input *GetScalingPlanResourceForecastDataInput
	Copy  func(*GetScalingPlanResourceForecastDataInput) GetScalingPlanResourceForecastDataRequest
}

// Send marshals and sends the GetScalingPlanResourceForecastData API request.
func (r GetScalingPlanResourceForecastDataRequest) Send(ctx context.Context) (*GetScalingPlanResourceForecastDataResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetScalingPlanResourceForecastDataResponse{
		GetScalingPlanResourceForecastDataOutput: r.Request.Data.(*GetScalingPlanResourceForecastDataOutput),
		response:                                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetScalingPlanResourceForecastDataResponse is the response type for the
// GetScalingPlanResourceForecastData API operation.
type GetScalingPlanResourceForecastDataResponse struct {
	*GetScalingPlanResourceForecastDataOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetScalingPlanResourceForecastData request.
func (r *GetScalingPlanResourceForecastDataResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
