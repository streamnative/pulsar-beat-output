// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package databasemigrationserviceiface provides an interface to enable mocking the AWS Database Migration Service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package databasemigrationserviceiface

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice"
)

// ClientAPI provides an interface to enable mocking the
// databasemigrationservice.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Database Migration Service.
//    func myFunc(svc databasemigrationserviceiface.ClientAPI) bool {
//        // Make svc.AddTagsToResource request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := databasemigrationservice.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        databasemigrationserviceiface.ClientPI
//    }
//    func (m *mockClientClient) AddTagsToResource(input *databasemigrationservice.AddTagsToResourceInput) (*databasemigrationservice.AddTagsToResourceOutput, error) {
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
	AddTagsToResourceRequest(*databasemigrationservice.AddTagsToResourceInput) databasemigrationservice.AddTagsToResourceRequest

	ApplyPendingMaintenanceActionRequest(*databasemigrationservice.ApplyPendingMaintenanceActionInput) databasemigrationservice.ApplyPendingMaintenanceActionRequest

	CreateEndpointRequest(*databasemigrationservice.CreateEndpointInput) databasemigrationservice.CreateEndpointRequest

	CreateEventSubscriptionRequest(*databasemigrationservice.CreateEventSubscriptionInput) databasemigrationservice.CreateEventSubscriptionRequest

	CreateReplicationInstanceRequest(*databasemigrationservice.CreateReplicationInstanceInput) databasemigrationservice.CreateReplicationInstanceRequest

	CreateReplicationSubnetGroupRequest(*databasemigrationservice.CreateReplicationSubnetGroupInput) databasemigrationservice.CreateReplicationSubnetGroupRequest

	CreateReplicationTaskRequest(*databasemigrationservice.CreateReplicationTaskInput) databasemigrationservice.CreateReplicationTaskRequest

	DeleteCertificateRequest(*databasemigrationservice.DeleteCertificateInput) databasemigrationservice.DeleteCertificateRequest

	DeleteEndpointRequest(*databasemigrationservice.DeleteEndpointInput) databasemigrationservice.DeleteEndpointRequest

	DeleteEventSubscriptionRequest(*databasemigrationservice.DeleteEventSubscriptionInput) databasemigrationservice.DeleteEventSubscriptionRequest

	DeleteReplicationInstanceRequest(*databasemigrationservice.DeleteReplicationInstanceInput) databasemigrationservice.DeleteReplicationInstanceRequest

	DeleteReplicationSubnetGroupRequest(*databasemigrationservice.DeleteReplicationSubnetGroupInput) databasemigrationservice.DeleteReplicationSubnetGroupRequest

	DeleteReplicationTaskRequest(*databasemigrationservice.DeleteReplicationTaskInput) databasemigrationservice.DeleteReplicationTaskRequest

	DescribeAccountAttributesRequest(*databasemigrationservice.DescribeAccountAttributesInput) databasemigrationservice.DescribeAccountAttributesRequest

	DescribeCertificatesRequest(*databasemigrationservice.DescribeCertificatesInput) databasemigrationservice.DescribeCertificatesRequest

	DescribeConnectionsRequest(*databasemigrationservice.DescribeConnectionsInput) databasemigrationservice.DescribeConnectionsRequest

	DescribeEndpointTypesRequest(*databasemigrationservice.DescribeEndpointTypesInput) databasemigrationservice.DescribeEndpointTypesRequest

	DescribeEndpointsRequest(*databasemigrationservice.DescribeEndpointsInput) databasemigrationservice.DescribeEndpointsRequest

	DescribeEventCategoriesRequest(*databasemigrationservice.DescribeEventCategoriesInput) databasemigrationservice.DescribeEventCategoriesRequest

	DescribeEventSubscriptionsRequest(*databasemigrationservice.DescribeEventSubscriptionsInput) databasemigrationservice.DescribeEventSubscriptionsRequest

	DescribeEventsRequest(*databasemigrationservice.DescribeEventsInput) databasemigrationservice.DescribeEventsRequest

	DescribeOrderableReplicationInstancesRequest(*databasemigrationservice.DescribeOrderableReplicationInstancesInput) databasemigrationservice.DescribeOrderableReplicationInstancesRequest

	DescribePendingMaintenanceActionsRequest(*databasemigrationservice.DescribePendingMaintenanceActionsInput) databasemigrationservice.DescribePendingMaintenanceActionsRequest

	DescribeRefreshSchemasStatusRequest(*databasemigrationservice.DescribeRefreshSchemasStatusInput) databasemigrationservice.DescribeRefreshSchemasStatusRequest

	DescribeReplicationInstanceTaskLogsRequest(*databasemigrationservice.DescribeReplicationInstanceTaskLogsInput) databasemigrationservice.DescribeReplicationInstanceTaskLogsRequest

	DescribeReplicationInstancesRequest(*databasemigrationservice.DescribeReplicationInstancesInput) databasemigrationservice.DescribeReplicationInstancesRequest

	DescribeReplicationSubnetGroupsRequest(*databasemigrationservice.DescribeReplicationSubnetGroupsInput) databasemigrationservice.DescribeReplicationSubnetGroupsRequest

	DescribeReplicationTaskAssessmentResultsRequest(*databasemigrationservice.DescribeReplicationTaskAssessmentResultsInput) databasemigrationservice.DescribeReplicationTaskAssessmentResultsRequest

	DescribeReplicationTasksRequest(*databasemigrationservice.DescribeReplicationTasksInput) databasemigrationservice.DescribeReplicationTasksRequest

	DescribeSchemasRequest(*databasemigrationservice.DescribeSchemasInput) databasemigrationservice.DescribeSchemasRequest

	DescribeTableStatisticsRequest(*databasemigrationservice.DescribeTableStatisticsInput) databasemigrationservice.DescribeTableStatisticsRequest

	ImportCertificateRequest(*databasemigrationservice.ImportCertificateInput) databasemigrationservice.ImportCertificateRequest

	ListTagsForResourceRequest(*databasemigrationservice.ListTagsForResourceInput) databasemigrationservice.ListTagsForResourceRequest

	ModifyEndpointRequest(*databasemigrationservice.ModifyEndpointInput) databasemigrationservice.ModifyEndpointRequest

	ModifyEventSubscriptionRequest(*databasemigrationservice.ModifyEventSubscriptionInput) databasemigrationservice.ModifyEventSubscriptionRequest

	ModifyReplicationInstanceRequest(*databasemigrationservice.ModifyReplicationInstanceInput) databasemigrationservice.ModifyReplicationInstanceRequest

	ModifyReplicationSubnetGroupRequest(*databasemigrationservice.ModifyReplicationSubnetGroupInput) databasemigrationservice.ModifyReplicationSubnetGroupRequest

	ModifyReplicationTaskRequest(*databasemigrationservice.ModifyReplicationTaskInput) databasemigrationservice.ModifyReplicationTaskRequest

	RebootReplicationInstanceRequest(*databasemigrationservice.RebootReplicationInstanceInput) databasemigrationservice.RebootReplicationInstanceRequest

	RefreshSchemasRequest(*databasemigrationservice.RefreshSchemasInput) databasemigrationservice.RefreshSchemasRequest

	ReloadTablesRequest(*databasemigrationservice.ReloadTablesInput) databasemigrationservice.ReloadTablesRequest

	RemoveTagsFromResourceRequest(*databasemigrationservice.RemoveTagsFromResourceInput) databasemigrationservice.RemoveTagsFromResourceRequest

	StartReplicationTaskRequest(*databasemigrationservice.StartReplicationTaskInput) databasemigrationservice.StartReplicationTaskRequest

	StartReplicationTaskAssessmentRequest(*databasemigrationservice.StartReplicationTaskAssessmentInput) databasemigrationservice.StartReplicationTaskAssessmentRequest

	StopReplicationTaskRequest(*databasemigrationservice.StopReplicationTaskInput) databasemigrationservice.StopReplicationTaskRequest

	TestConnectionRequest(*databasemigrationservice.TestConnectionInput) databasemigrationservice.TestConnectionRequest

	WaitUntilEndpointDeleted(context.Context, *databasemigrationservice.DescribeEndpointsInput, ...aws.WaiterOption) error

	WaitUntilReplicationInstanceAvailable(context.Context, *databasemigrationservice.DescribeReplicationInstancesInput, ...aws.WaiterOption) error

	WaitUntilReplicationInstanceDeleted(context.Context, *databasemigrationservice.DescribeReplicationInstancesInput, ...aws.WaiterOption) error

	WaitUntilReplicationTaskDeleted(context.Context, *databasemigrationservice.DescribeReplicationTasksInput, ...aws.WaiterOption) error

	WaitUntilReplicationTaskReady(context.Context, *databasemigrationservice.DescribeReplicationTasksInput, ...aws.WaiterOption) error

	WaitUntilReplicationTaskRunning(context.Context, *databasemigrationservice.DescribeReplicationTasksInput, ...aws.WaiterOption) error

	WaitUntilReplicationTaskStopped(context.Context, *databasemigrationservice.DescribeReplicationTasksInput, ...aws.WaiterOption) error

	WaitUntilTestConnectionSucceeds(context.Context, *databasemigrationservice.DescribeConnectionsInput, ...aws.WaiterOption) error
}

var _ ClientAPI = (*databasemigrationservice.Client)(nil)
