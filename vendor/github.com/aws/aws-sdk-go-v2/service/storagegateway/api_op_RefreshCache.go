// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package storagegateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// RefreshCacheInput
// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/RefreshCacheInput
type RefreshCacheInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the file share you want to refresh.
	//
	// FileShareARN is a required field
	FileShareARN *string `min:"50" type:"string" required:"true"`

	// A comma-separated list of the paths of folders to refresh in the cache. The
	// default is ["/"]. The default refreshes objects and folders at the root of
	// the Amazon S3 bucket. If Recursive is set to "true", the entire S3 bucket
	// that the file share has access to is refreshed.
	FolderList []string `min:"1" type:"list"`

	// A value that specifies whether to recursively refresh folders in the cache.
	// The refresh includes folders that were in the cache the last time the gateway
	// listed the folder's contents. If this value set to "true", each folder that
	// is listed in FolderList is recursively updated. Otherwise, subfolders listed
	// in FolderList are not refreshed. Only objects that are in folders listed
	// directly under FolderList are found and used for the update. The default
	// is "true".
	Recursive *bool `type:"boolean"`
}

// String returns the string representation
func (s RefreshCacheInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RefreshCacheInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RefreshCacheInput"}

	if s.FileShareARN == nil {
		invalidParams.Add(aws.NewErrParamRequired("FileShareARN"))
	}
	if s.FileShareARN != nil && len(*s.FileShareARN) < 50 {
		invalidParams.Add(aws.NewErrParamMinLen("FileShareARN", 50))
	}
	if s.FolderList != nil && len(s.FolderList) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("FolderList", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// RefreshCacheOutput
// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/RefreshCacheOutput
type RefreshCacheOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the file share.
	FileShareARN *string `min:"50" type:"string"`

	// The randomly generated ID of the notification that was sent. This ID is in
	// UUID format.
	NotificationId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s RefreshCacheOutput) String() string {
	return awsutil.Prettify(s)
}

const opRefreshCache = "RefreshCache"

// RefreshCacheRequest returns a request value for making API operation for
// AWS Storage Gateway.
//
// Refreshes the cache for the specified file share. This operation finds objects
// in the Amazon S3 bucket that were added, removed or replaced since the gateway
// last listed the bucket's contents and cached the results. This operation
// is only supported in the file gateway type. You can subscribe to be notified
// through an Amazon CloudWatch event when your RefreshCache operation completes.
// For more information, see Getting Notified About File Operations (https://docs.aws.amazon.com/storagegateway/latest/userguide/monitoring-file-gateway.html#get-notification).
//
// When this API is called, it only initiates the refresh operation. When the
// API call completes and returns a success code, it doesn't necessarily mean
// that the file refresh has completed. You should use the refresh-complete
// notification to determine that the operation has completed before you check
// for new files on the gateway file share. You can subscribe to be notified
// through an CloudWatch event when your RefreshCache operation completes.
//
//    // Example sending a request using RefreshCacheRequest.
//    req := client.RefreshCacheRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/RefreshCache
func (c *Client) RefreshCacheRequest(input *RefreshCacheInput) RefreshCacheRequest {
	op := &aws.Operation{
		Name:       opRefreshCache,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RefreshCacheInput{}
	}

	req := c.newRequest(op, input, &RefreshCacheOutput{})
	return RefreshCacheRequest{Request: req, Input: input, Copy: c.RefreshCacheRequest}
}

// RefreshCacheRequest is the request type for the
// RefreshCache API operation.
type RefreshCacheRequest struct {
	*aws.Request
	Input *RefreshCacheInput
	Copy  func(*RefreshCacheInput) RefreshCacheRequest
}

// Send marshals and sends the RefreshCache API request.
func (r RefreshCacheRequest) Send(ctx context.Context) (*RefreshCacheResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RefreshCacheResponse{
		RefreshCacheOutput: r.Request.Data.(*RefreshCacheOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RefreshCacheResponse is the response type for the
// RefreshCache API operation.
type RefreshCacheResponse struct {
	*RefreshCacheOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RefreshCache request.
func (r *RefreshCacheResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
