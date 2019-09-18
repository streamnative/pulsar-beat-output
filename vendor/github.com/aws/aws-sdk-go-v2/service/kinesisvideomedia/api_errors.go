// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesisvideomedia

const (

	// ErrCodeConnectionLimitExceededException for service response error code
	// "ConnectionLimitExceededException".
	//
	// Kinesis Video Streams has throttled the request because you have exceeded
	// the limit of allowed client connections.
	ErrCodeConnectionLimitExceededException = "ConnectionLimitExceededException"

	// ErrCodeInvalidArgumentException for service response error code
	// "InvalidArgumentException".
	//
	// The value for this input parameter is invalid.
	ErrCodeInvalidArgumentException = "InvalidArgumentException"

	// ErrCodeInvalidEndpointException for service response error code
	// "InvalidEndpointException".
	//
	// Status Code: 400, Caller used wrong endpoint to write data to a stream. On
	// receiving such an exception, the user must call GetDataEndpoint with AccessMode
	// set to "READ" and use the endpoint Kinesis Video returns in the next GetMedia
	// call.
	ErrCodeInvalidEndpointException = "InvalidEndpointException"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	//
	// Kinesis Video Streams has throttled the request because you have exceeded
	// the limit of allowed client calls. Try making the call later.
	ErrCodeLimitExceededException = "LimitExceededException"

	// ErrCodeNotAuthorizedException for service response error code
	// "NotAuthorizedException".
	//
	// Status Code: 403, The caller is not authorized to perform an operation on
	// the given stream, or the token has expired.
	ErrCodeNotAuthorizedException = "NotAuthorizedException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// Status Code: 404, The stream with the given name does not exist.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"
)
