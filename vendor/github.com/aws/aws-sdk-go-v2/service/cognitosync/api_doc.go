// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package cognitosync provides the client and types for making API
// requests to Amazon Cognito Sync.
//
// Amazon Cognito Sync provides an AWS service and client library that enable
// cross-device syncing of application-related user data. High-level client
// libraries are available for both iOS and Android. You can use these libraries
// to persist data locally so that it's available even if the device is offline.
// Developer credentials don't need to be stored on the mobile device to access
// the service. You can use Amazon Cognito to obtain a normalized user ID and
// credentials. User data is persisted in a dataset that can store up to 1 MB
// of key-value pairs, and you can have up to 20 datasets per user identity.
//
// With Amazon Cognito Sync, the data stored for each identity is accessible
// only to credentials assigned to that identity. In order to use the Cognito
// Sync service, you need to make API calls using credentials retrieved with
// Amazon Cognito Identity service (http://docs.aws.amazon.com/cognitoidentity/latest/APIReference/Welcome.html).
//
// If you want to use Cognito Sync in an Android or iOS application, you will
// probably want to make API calls via the AWS Mobile SDK. To learn more, see
// the Developer Guide for Android (http://docs.aws.amazon.com/mobile/sdkforandroid/developerguide/cognito-sync.html)
// and the Developer Guide for iOS (http://docs.aws.amazon.com/mobile/sdkforios/developerguide/cognito-sync.html).
//
// See https://docs.aws.amazon.com/goto/WebAPI/cognito-sync-2014-06-30 for more information on this service.
//
// See cognitosync package documentation for more information.
// https://docs.aws.amazon.com/sdk-for-go/api/service/cognitosync/
//
// Using the Client
//
// To use Amazon Cognito Sync with the SDK use the New function to create
// a new service client. With that client you can make API requests to the service.
// These clients are safe to use concurrently.
//
// See the SDK's documentation for more information on how to use the SDK.
// https://docs.aws.amazon.com/sdk-for-go/api/
//
// See aws.Config documentation for more information on configuring SDK clients.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
//
// See the Amazon Cognito Sync client for more information on
// creating client for this service.
// https://docs.aws.amazon.com/sdk-for-go/api/service/cognitosync/#New
package cognitosync
