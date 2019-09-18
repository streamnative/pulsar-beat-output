// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package cloudsearchiface provides an interface to enable mocking the Amazon CloudSearch service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package cloudsearchiface

import (
	"github.com/aws/aws-sdk-go-v2/service/cloudsearch"
)

// ClientAPI provides an interface to enable mocking the
// cloudsearch.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon CloudSearch.
//    func myFunc(svc cloudsearchiface.ClientAPI) bool {
//        // Make svc.BuildSuggesters request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := cloudsearch.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        cloudsearchiface.ClientPI
//    }
//    func (m *mockClientClient) BuildSuggesters(input *cloudsearch.BuildSuggestersInput) (*cloudsearch.BuildSuggestersOutput, error) {
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
	BuildSuggestersRequest(*cloudsearch.BuildSuggestersInput) cloudsearch.BuildSuggestersRequest

	CreateDomainRequest(*cloudsearch.CreateDomainInput) cloudsearch.CreateDomainRequest

	DefineAnalysisSchemeRequest(*cloudsearch.DefineAnalysisSchemeInput) cloudsearch.DefineAnalysisSchemeRequest

	DefineExpressionRequest(*cloudsearch.DefineExpressionInput) cloudsearch.DefineExpressionRequest

	DefineIndexFieldRequest(*cloudsearch.DefineIndexFieldInput) cloudsearch.DefineIndexFieldRequest

	DefineSuggesterRequest(*cloudsearch.DefineSuggesterInput) cloudsearch.DefineSuggesterRequest

	DeleteAnalysisSchemeRequest(*cloudsearch.DeleteAnalysisSchemeInput) cloudsearch.DeleteAnalysisSchemeRequest

	DeleteDomainRequest(*cloudsearch.DeleteDomainInput) cloudsearch.DeleteDomainRequest

	DeleteExpressionRequest(*cloudsearch.DeleteExpressionInput) cloudsearch.DeleteExpressionRequest

	DeleteIndexFieldRequest(*cloudsearch.DeleteIndexFieldInput) cloudsearch.DeleteIndexFieldRequest

	DeleteSuggesterRequest(*cloudsearch.DeleteSuggesterInput) cloudsearch.DeleteSuggesterRequest

	DescribeAnalysisSchemesRequest(*cloudsearch.DescribeAnalysisSchemesInput) cloudsearch.DescribeAnalysisSchemesRequest

	DescribeAvailabilityOptionsRequest(*cloudsearch.DescribeAvailabilityOptionsInput) cloudsearch.DescribeAvailabilityOptionsRequest

	DescribeDomainsRequest(*cloudsearch.DescribeDomainsInput) cloudsearch.DescribeDomainsRequest

	DescribeExpressionsRequest(*cloudsearch.DescribeExpressionsInput) cloudsearch.DescribeExpressionsRequest

	DescribeIndexFieldsRequest(*cloudsearch.DescribeIndexFieldsInput) cloudsearch.DescribeIndexFieldsRequest

	DescribeScalingParametersRequest(*cloudsearch.DescribeScalingParametersInput) cloudsearch.DescribeScalingParametersRequest

	DescribeServiceAccessPoliciesRequest(*cloudsearch.DescribeServiceAccessPoliciesInput) cloudsearch.DescribeServiceAccessPoliciesRequest

	DescribeSuggestersRequest(*cloudsearch.DescribeSuggestersInput) cloudsearch.DescribeSuggestersRequest

	IndexDocumentsRequest(*cloudsearch.IndexDocumentsInput) cloudsearch.IndexDocumentsRequest

	ListDomainNamesRequest(*cloudsearch.ListDomainNamesInput) cloudsearch.ListDomainNamesRequest

	UpdateAvailabilityOptionsRequest(*cloudsearch.UpdateAvailabilityOptionsInput) cloudsearch.UpdateAvailabilityOptionsRequest

	UpdateScalingParametersRequest(*cloudsearch.UpdateScalingParametersInput) cloudsearch.UpdateScalingParametersRequest

	UpdateServiceAccessPoliciesRequest(*cloudsearch.UpdateServiceAccessPoliciesInput) cloudsearch.UpdateServiceAccessPoliciesRequest
}

var _ ClientAPI = (*cloudsearch.Client)(nil)
