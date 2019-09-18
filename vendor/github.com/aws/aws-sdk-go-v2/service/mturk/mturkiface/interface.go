// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package mturkiface provides an interface to enable mocking the Amazon Mechanical Turk service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package mturkiface

import (
	"github.com/aws/aws-sdk-go-v2/service/mturk"
)

// ClientAPI provides an interface to enable mocking the
// mturk.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon MTurk.
//    func myFunc(svc mturkiface.ClientAPI) bool {
//        // Make svc.AcceptQualificationRequest request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := mturk.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        mturkiface.ClientPI
//    }
//    func (m *mockClientClient) AcceptQualificationRequest(input *mturk.AcceptQualificationRequestInput) (*mturk.AcceptQualificationRequestOutput, error) {
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
	AcceptQualificationRequestRequest(*mturk.AcceptQualificationRequestInput) mturk.AcceptQualificationRequestRequest

	ApproveAssignmentRequest(*mturk.ApproveAssignmentInput) mturk.ApproveAssignmentRequest

	AssociateQualificationWithWorkerRequest(*mturk.AssociateQualificationWithWorkerInput) mturk.AssociateQualificationWithWorkerRequest

	CreateAdditionalAssignmentsForHITRequest(*mturk.CreateAdditionalAssignmentsForHITInput) mturk.CreateAdditionalAssignmentsForHITRequest

	CreateHITRequest(*mturk.CreateHITInput) mturk.CreateHITRequest

	CreateHITTypeRequest(*mturk.CreateHITTypeInput) mturk.CreateHITTypeRequest

	CreateHITWithHITTypeRequest(*mturk.CreateHITWithHITTypeInput) mturk.CreateHITWithHITTypeRequest

	CreateQualificationTypeRequest(*mturk.CreateQualificationTypeInput) mturk.CreateQualificationTypeRequest

	CreateWorkerBlockRequest(*mturk.CreateWorkerBlockInput) mturk.CreateWorkerBlockRequest

	DeleteHITRequest(*mturk.DeleteHITInput) mturk.DeleteHITRequest

	DeleteQualificationTypeRequest(*mturk.DeleteQualificationTypeInput) mturk.DeleteQualificationTypeRequest

	DeleteWorkerBlockRequest(*mturk.DeleteWorkerBlockInput) mturk.DeleteWorkerBlockRequest

	DisassociateQualificationFromWorkerRequest(*mturk.DisassociateQualificationFromWorkerInput) mturk.DisassociateQualificationFromWorkerRequest

	GetAccountBalanceRequest(*mturk.GetAccountBalanceInput) mturk.GetAccountBalanceRequest

	GetAssignmentRequest(*mturk.GetAssignmentInput) mturk.GetAssignmentRequest

	GetFileUploadURLRequest(*mturk.GetFileUploadURLInput) mturk.GetFileUploadURLRequest

	GetHITRequest(*mturk.GetHITInput) mturk.GetHITRequest

	GetQualificationScoreRequest(*mturk.GetQualificationScoreInput) mturk.GetQualificationScoreRequest

	GetQualificationTypeRequest(*mturk.GetQualificationTypeInput) mturk.GetQualificationTypeRequest

	ListAssignmentsForHITRequest(*mturk.ListAssignmentsForHITInput) mturk.ListAssignmentsForHITRequest

	ListBonusPaymentsRequest(*mturk.ListBonusPaymentsInput) mturk.ListBonusPaymentsRequest

	ListHITsRequest(*mturk.ListHITsInput) mturk.ListHITsRequest

	ListHITsForQualificationTypeRequest(*mturk.ListHITsForQualificationTypeInput) mturk.ListHITsForQualificationTypeRequest

	ListQualificationRequestsRequest(*mturk.ListQualificationRequestsInput) mturk.ListQualificationRequestsRequest

	ListQualificationTypesRequest(*mturk.ListQualificationTypesInput) mturk.ListQualificationTypesRequest

	ListReviewPolicyResultsForHITRequest(*mturk.ListReviewPolicyResultsForHITInput) mturk.ListReviewPolicyResultsForHITRequest

	ListReviewableHITsRequest(*mturk.ListReviewableHITsInput) mturk.ListReviewableHITsRequest

	ListWorkerBlocksRequest(*mturk.ListWorkerBlocksInput) mturk.ListWorkerBlocksRequest

	ListWorkersWithQualificationTypeRequest(*mturk.ListWorkersWithQualificationTypeInput) mturk.ListWorkersWithQualificationTypeRequest

	NotifyWorkersRequest(*mturk.NotifyWorkersInput) mturk.NotifyWorkersRequest

	RejectAssignmentRequest(*mturk.RejectAssignmentInput) mturk.RejectAssignmentRequest

	RejectQualificationRequestRequest(*mturk.RejectQualificationRequestInput) mturk.RejectQualificationRequestRequest

	SendBonusRequest(*mturk.SendBonusInput) mturk.SendBonusRequest

	SendTestEventNotificationRequest(*mturk.SendTestEventNotificationInput) mturk.SendTestEventNotificationRequest

	UpdateExpirationForHITRequest(*mturk.UpdateExpirationForHITInput) mturk.UpdateExpirationForHITRequest

	UpdateHITReviewStatusRequest(*mturk.UpdateHITReviewStatusInput) mturk.UpdateHITReviewStatusRequest

	UpdateHITTypeOfHITRequest(*mturk.UpdateHITTypeOfHITInput) mturk.UpdateHITTypeOfHITRequest

	UpdateNotificationSettingsRequest(*mturk.UpdateNotificationSettingsInput) mturk.UpdateNotificationSettingsRequest

	UpdateQualificationTypeRequest(*mturk.UpdateQualificationTypeInput) mturk.UpdateQualificationTypeRequest
}

var _ ClientAPI = (*mturk.Client)(nil)
