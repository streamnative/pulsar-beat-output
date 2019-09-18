// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package signeriface provides an interface to enable mocking the AWS Signer service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package signeriface

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/signer"
)

// ClientAPI provides an interface to enable mocking the
// signer.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // signer.
//    func myFunc(svc signeriface.ClientAPI) bool {
//        // Make svc.CancelSigningProfile request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := signer.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        signeriface.ClientPI
//    }
//    func (m *mockClientClient) CancelSigningProfile(input *signer.CancelSigningProfileInput) (*signer.CancelSigningProfileOutput, error) {
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
	CancelSigningProfileRequest(*signer.CancelSigningProfileInput) signer.CancelSigningProfileRequest

	DescribeSigningJobRequest(*signer.DescribeSigningJobInput) signer.DescribeSigningJobRequest

	GetSigningPlatformRequest(*signer.GetSigningPlatformInput) signer.GetSigningPlatformRequest

	GetSigningProfileRequest(*signer.GetSigningProfileInput) signer.GetSigningProfileRequest

	ListSigningJobsRequest(*signer.ListSigningJobsInput) signer.ListSigningJobsRequest

	ListSigningPlatformsRequest(*signer.ListSigningPlatformsInput) signer.ListSigningPlatformsRequest

	ListSigningProfilesRequest(*signer.ListSigningProfilesInput) signer.ListSigningProfilesRequest

	PutSigningProfileRequest(*signer.PutSigningProfileInput) signer.PutSigningProfileRequest

	StartSigningJobRequest(*signer.StartSigningJobInput) signer.StartSigningJobRequest

	WaitUntilSuccessfulSigningJob(context.Context, *signer.DescribeSigningJobInput, ...aws.WaiterOption) error
}

var _ ClientAPI = (*signer.Client)(nil)
