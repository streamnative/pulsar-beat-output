// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesis

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

// Represents the input for MergeShards.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesis-2013-12-02/MergeShardsInput
type MergeShardsInput struct {
	_ struct{} `type:"structure"`

	// The shard ID of the adjacent shard for the merge.
	//
	// AdjacentShardToMerge is a required field
	AdjacentShardToMerge *string `min:"1" type:"string" required:"true"`

	// The shard ID of the shard to combine with the adjacent shard for the merge.
	//
	// ShardToMerge is a required field
	ShardToMerge *string `min:"1" type:"string" required:"true"`

	// The name of the stream for the merge.
	//
	// StreamName is a required field
	StreamName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s MergeShardsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *MergeShardsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "MergeShardsInput"}

	if s.AdjacentShardToMerge == nil {
		invalidParams.Add(aws.NewErrParamRequired("AdjacentShardToMerge"))
	}
	if s.AdjacentShardToMerge != nil && len(*s.AdjacentShardToMerge) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AdjacentShardToMerge", 1))
	}

	if s.ShardToMerge == nil {
		invalidParams.Add(aws.NewErrParamRequired("ShardToMerge"))
	}
	if s.ShardToMerge != nil && len(*s.ShardToMerge) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ShardToMerge", 1))
	}

	if s.StreamName == nil {
		invalidParams.Add(aws.NewErrParamRequired("StreamName"))
	}
	if s.StreamName != nil && len(*s.StreamName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("StreamName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesis-2013-12-02/MergeShardsOutput
type MergeShardsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s MergeShardsOutput) String() string {
	return awsutil.Prettify(s)
}

const opMergeShards = "MergeShards"

// MergeShardsRequest returns a request value for making API operation for
// Amazon Kinesis.
//
// Merges two adjacent shards in a Kinesis data stream and combines them into
// a single shard to reduce the stream's capacity to ingest and transport data.
// Two shards are considered adjacent if the union of the hash key ranges for
// the two shards form a contiguous set with no gaps. For example, if you have
// two shards, one with a hash key range of 276...381 and the other with a hash
// key range of 382...454, then you could merge these two shards into a single
// shard that would have a hash key range of 276...454. After the merge, the
// single child shard receives data for all hash key values covered by the two
// parent shards.
//
// MergeShards is called when there is a need to reduce the overall capacity
// of a stream because of excess capacity that is not being used. You must specify
// the shard to be merged and the adjacent shard for a stream. For more information
// about merging shards, see Merge Two Shards (http://docs.aws.amazon.com/kinesis/latest/dev/kinesis-using-sdk-java-resharding-merge.html)
// in the Amazon Kinesis Data Streams Developer Guide.
//
// If the stream is in the ACTIVE state, you can call MergeShards. If a stream
// is in the CREATING, UPDATING, or DELETING state, MergeShards returns a ResourceInUseException.
// If the specified stream does not exist, MergeShards returns a ResourceNotFoundException.
//
// You can use DescribeStream to check the state of the stream, which is returned
// in StreamStatus.
//
// MergeShards is an asynchronous operation. Upon receiving a MergeShards request,
// Amazon Kinesis Data Streams immediately returns a response and sets the StreamStatus
// to UPDATING. After the operation is completed, Kinesis Data Streams sets
// the StreamStatus to ACTIVE. Read and write operations continue to work while
// the stream is in the UPDATING state.
//
// You use DescribeStream to determine the shard IDs that are specified in the
// MergeShards request.
//
// If you try to operate on too many streams in parallel using CreateStream,
// DeleteStream, MergeShards, or SplitShard, you receive a LimitExceededException.
//
// MergeShards has a limit of five transactions per second per account.
//
//    // Example sending a request using MergeShardsRequest.
//    req := client.MergeShardsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesis-2013-12-02/MergeShards
func (c *Client) MergeShardsRequest(input *MergeShardsInput) MergeShardsRequest {
	op := &aws.Operation{
		Name:       opMergeShards,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &MergeShardsInput{}
	}

	req := c.newRequest(op, input, &MergeShardsOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return MergeShardsRequest{Request: req, Input: input, Copy: c.MergeShardsRequest}
}

// MergeShardsRequest is the request type for the
// MergeShards API operation.
type MergeShardsRequest struct {
	*aws.Request
	Input *MergeShardsInput
	Copy  func(*MergeShardsInput) MergeShardsRequest
}

// Send marshals and sends the MergeShards API request.
func (r MergeShardsRequest) Send(ctx context.Context) (*MergeShardsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &MergeShardsResponse{
		MergeShardsOutput: r.Request.Data.(*MergeShardsOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// MergeShardsResponse is the response type for the
// MergeShards API operation.
type MergeShardsResponse struct {
	*MergeShardsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// MergeShards request.
func (r *MergeShardsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
