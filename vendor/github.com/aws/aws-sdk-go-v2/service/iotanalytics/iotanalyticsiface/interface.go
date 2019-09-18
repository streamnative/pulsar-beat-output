// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package iotanalyticsiface provides an interface to enable mocking the AWS IoT Analytics service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package iotanalyticsiface

import (
	"github.com/aws/aws-sdk-go-v2/service/iotanalytics"
)

// ClientAPI provides an interface to enable mocking the
// iotanalytics.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS IoT Analytics.
//    func myFunc(svc iotanalyticsiface.ClientAPI) bool {
//        // Make svc.BatchPutMessage request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := iotanalytics.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        iotanalyticsiface.ClientPI
//    }
//    func (m *mockClientClient) BatchPutMessage(input *iotanalytics.BatchPutMessageInput) (*iotanalytics.BatchPutMessageOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockClientClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type ClientAPI interface {
	BatchPutMessageRequest(*iotanalytics.BatchPutMessageInput) iotanalytics.BatchPutMessageRequest

	CancelPipelineReprocessingRequest(*iotanalytics.CancelPipelineReprocessingInput) iotanalytics.CancelPipelineReprocessingRequest

	CreateChannelRequest(*iotanalytics.CreateChannelInput) iotanalytics.CreateChannelRequest

	CreateDatasetRequest(*iotanalytics.CreateDatasetInput) iotanalytics.CreateDatasetRequest

	CreateDatasetContentRequest(*iotanalytics.CreateDatasetContentInput) iotanalytics.CreateDatasetContentRequest

	CreateDatastoreRequest(*iotanalytics.CreateDatastoreInput) iotanalytics.CreateDatastoreRequest

	CreatePipelineRequest(*iotanalytics.CreatePipelineInput) iotanalytics.CreatePipelineRequest

	DeleteChannelRequest(*iotanalytics.DeleteChannelInput) iotanalytics.DeleteChannelRequest

	DeleteDatasetRequest(*iotanalytics.DeleteDatasetInput) iotanalytics.DeleteDatasetRequest

	DeleteDatasetContentRequest(*iotanalytics.DeleteDatasetContentInput) iotanalytics.DeleteDatasetContentRequest

	DeleteDatastoreRequest(*iotanalytics.DeleteDatastoreInput) iotanalytics.DeleteDatastoreRequest

	DeletePipelineRequest(*iotanalytics.DeletePipelineInput) iotanalytics.DeletePipelineRequest

	DescribeChannelRequest(*iotanalytics.DescribeChannelInput) iotanalytics.DescribeChannelRequest

	DescribeDatasetRequest(*iotanalytics.DescribeDatasetInput) iotanalytics.DescribeDatasetRequest

	DescribeDatastoreRequest(*iotanalytics.DescribeDatastoreInput) iotanalytics.DescribeDatastoreRequest

	DescribeLoggingOptionsRequest(*iotanalytics.DescribeLoggingOptionsInput) iotanalytics.DescribeLoggingOptionsRequest

	DescribePipelineRequest(*iotanalytics.DescribePipelineInput) iotanalytics.DescribePipelineRequest

	GetDatasetContentRequest(*iotanalytics.GetDatasetContentInput) iotanalytics.GetDatasetContentRequest

	ListChannelsRequest(*iotanalytics.ListChannelsInput) iotanalytics.ListChannelsRequest

	ListDatasetContentsRequest(*iotanalytics.ListDatasetContentsInput) iotanalytics.ListDatasetContentsRequest

	ListDatasetsRequest(*iotanalytics.ListDatasetsInput) iotanalytics.ListDatasetsRequest

	ListDatastoresRequest(*iotanalytics.ListDatastoresInput) iotanalytics.ListDatastoresRequest

	ListPipelinesRequest(*iotanalytics.ListPipelinesInput) iotanalytics.ListPipelinesRequest

	ListTagsForResourceRequest(*iotanalytics.ListTagsForResourceInput) iotanalytics.ListTagsForResourceRequest

	PutLoggingOptionsRequest(*iotanalytics.PutLoggingOptionsInput) iotanalytics.PutLoggingOptionsRequest

	RunPipelineActivityRequest(*iotanalytics.RunPipelineActivityInput) iotanalytics.RunPipelineActivityRequest

	SampleChannelDataRequest(*iotanalytics.SampleChannelDataInput) iotanalytics.SampleChannelDataRequest

	StartPipelineReprocessingRequest(*iotanalytics.StartPipelineReprocessingInput) iotanalytics.StartPipelineReprocessingRequest

	TagResourceRequest(*iotanalytics.TagResourceInput) iotanalytics.TagResourceRequest

	UntagResourceRequest(*iotanalytics.UntagResourceInput) iotanalytics.UntagResourceRequest

	UpdateChannelRequest(*iotanalytics.UpdateChannelInput) iotanalytics.UpdateChannelRequest

	UpdateDatasetRequest(*iotanalytics.UpdateDatasetInput) iotanalytics.UpdateDatasetRequest

	UpdateDatastoreRequest(*iotanalytics.UpdateDatastoreInput) iotanalytics.UpdateDatastoreRequest

	UpdatePipelineRequest(*iotanalytics.UpdatePipelineInput) iotanalytics.UpdatePipelineRequest
}

var _ ClientAPI = (*iotanalytics.Client)(nil)
