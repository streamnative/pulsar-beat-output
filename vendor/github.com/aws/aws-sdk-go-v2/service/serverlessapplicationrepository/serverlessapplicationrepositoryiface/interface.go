// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package serverlessapplicationrepositoryiface provides an interface to enable mocking the AWSServerlessApplicationRepository service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package serverlessapplicationrepositoryiface

import (
	"github.com/aws/aws-sdk-go-v2/service/serverlessapplicationrepository"
)

// ClientAPI provides an interface to enable mocking the
// serverlessapplicationrepository.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWSServerlessApplicationRepository.
//    func myFunc(svc serverlessapplicationrepositoryiface.ClientAPI) bool {
//        // Make svc.CreateApplication request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := serverlessapplicationrepository.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        serverlessapplicationrepositoryiface.ClientPI
//    }
//    func (m *mockClientClient) CreateApplication(input *serverlessapplicationrepository.CreateApplicationInput) (*serverlessapplicationrepository.CreateApplicationOutput, error) {
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
	CreateApplicationRequest(*serverlessapplicationrepository.CreateApplicationInput) serverlessapplicationrepository.CreateApplicationRequest

	CreateApplicationVersionRequest(*serverlessapplicationrepository.CreateApplicationVersionInput) serverlessapplicationrepository.CreateApplicationVersionRequest

	CreateCloudFormationChangeSetRequest(*serverlessapplicationrepository.CreateCloudFormationChangeSetInput) serverlessapplicationrepository.CreateCloudFormationChangeSetRequest

	CreateCloudFormationTemplateRequest(*serverlessapplicationrepository.CreateCloudFormationTemplateInput) serverlessapplicationrepository.CreateCloudFormationTemplateRequest

	DeleteApplicationRequest(*serverlessapplicationrepository.DeleteApplicationInput) serverlessapplicationrepository.DeleteApplicationRequest

	GetApplicationRequest(*serverlessapplicationrepository.GetApplicationInput) serverlessapplicationrepository.GetApplicationRequest

	GetApplicationPolicyRequest(*serverlessapplicationrepository.GetApplicationPolicyInput) serverlessapplicationrepository.GetApplicationPolicyRequest

	GetCloudFormationTemplateRequest(*serverlessapplicationrepository.GetCloudFormationTemplateInput) serverlessapplicationrepository.GetCloudFormationTemplateRequest

	ListApplicationDependenciesRequest(*serverlessapplicationrepository.ListApplicationDependenciesInput) serverlessapplicationrepository.ListApplicationDependenciesRequest

	ListApplicationVersionsRequest(*serverlessapplicationrepository.ListApplicationVersionsInput) serverlessapplicationrepository.ListApplicationVersionsRequest

	ListApplicationsRequest(*serverlessapplicationrepository.ListApplicationsInput) serverlessapplicationrepository.ListApplicationsRequest

	PutApplicationPolicyRequest(*serverlessapplicationrepository.PutApplicationPolicyInput) serverlessapplicationrepository.PutApplicationPolicyRequest

	UpdateApplicationRequest(*serverlessapplicationrepository.UpdateApplicationInput) serverlessapplicationrepository.UpdateApplicationRequest
}

var _ ClientAPI = (*serverlessapplicationrepository.Client)(nil)
