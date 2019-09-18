// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesis

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

// Represents the input for SplitShard.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesis-2013-12-02/SplitShardInput
type SplitShardInput struct {
	_ struct{} `type:"structure"`

	// A hash key value for the starting hash key of one of the child shards created
	// by the split. The hash key range for a given shard constitutes a set of ordered
	// contiguous positive integers. The value for NewStartingHashKey must be in
	// the range of hash keys being mapped into the shard. The NewStartingHashKey
	// hash key value and all higher hash key values in hash key range are distributed
	// to one of the child shards. All the lower hash key values in the range are
	// distributed to the other child shard.
	//
	// NewStartingHashKey is a required field
	NewStartingHashKey *string `type:"string" required:"true"`

	// The shard ID of the shard to split.
	//
	// ShardToSplit is a required field
	ShardToSplit *string `min:"1" type:"string" required:"true"`

	// The name of the stream for the shard split.
	//
	// StreamName is a required field
	StreamName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s SplitShardInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SplitShardInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SplitShardInput"}

	if s.NewStartingHashKey == nil {
		invalidParams.Add(aws.NewErrParamRequired("NewStartingHashKey"))
	}

	if s.ShardToSplit == nil {
		invalidParams.Add(aws.NewErrParamRequired("ShardToSplit"))
	}
	if s.ShardToSplit != nil && len(*s.ShardToSplit) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ShardToSplit", 1))
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

// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesis-2013-12-02/SplitShardOutput
type SplitShardOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s SplitShardOutput) String() string {
	return awsutil.Prettify(s)
}

const opSplitShard = "SplitShard"

// SplitShardRequest returns a request value for making API operation for
// Amazon Kinesis.
//
// Splits a shard into two new shards in the Kinesis data stream, to increase
// the stream's capacity to ingest and transport data. SplitShard is called
// when there is a need to increase the overall capacity of a stream because
// of an expected increase in the volume of data records being ingested.
//
// You can also use SplitShard when a shard appears to be approaching its maximum
// utilization; for example, the producers sending data into the specific shard
// are suddenly sending more than previously anticipated. You can also call
// SplitShard to increase stream capacity, so that more Kinesis Data Streams
// applications can simultaneously read data from the stream for real-time processing.
//
// You must specify the shard to be split and the new hash key, which is the
// position in the shard where the shard gets split in two. In many cases, the
// new hash key might be the average of the beginning and ending hash key, but
// it can be any hash key value in the range being mapped into the shard. For
// more information, see Split a Shard (http://docs.aws.amazon.com/kinesis/latest/dev/kinesis-using-sdk-java-resharding-split.html)
// in the Amazon Kinesis Data Streams Developer Guide.
//
// You can use DescribeStream to determine the shard ID and hash key values
// for the ShardToSplit and NewStartingHashKey parameters that are specified
// in the SplitShard request.
//
// SplitShard is an asynchronous operation. Upon receiving a SplitShard request,
// Kinesis Data Streams immediately returns a response and sets the stream status
// to UPDATING. After the operation is completed, Kinesis Data Streams sets
// the stream status to ACTIVE. Read and write operations continue to work while
// the stream is in the UPDATING state.
//
// You can use DescribeStream to check the status of the stream, which is returned
// in StreamStatus. If the stream is in the ACTIVE state, you can call SplitShard.
// If a stream is in CREATING or UPDATING or DELETING states, DescribeStream
// returns a ResourceInUseException.
//
// If the specified stream does not exist, DescribeStream returns a ResourceNotFoundException.
// If you try to create more shards than are authorized for your account, you
// receive a LimitExceededException.
//
// For the default shard limit for an AWS account, see Kinesis Data Streams
// Limits (http://docs.aws.amazon.com/kinesis/latest/dev/service-sizes-and-limits.html)
// in the Amazon Kinesis Data Streams Developer Guide. To increase this limit,
// contact AWS Support (http://docs.aws.amazon.com/general/latest/gr/aws_service_limits.html).
//
// If you try to operate on too many streams simultaneously using CreateStream,
// DeleteStream, MergeShards, and/or SplitShard, you receive a LimitExceededException.
//
// SplitShard has a limit of five transactions per second per account.
//
//    // Example sending a request using SplitShardRequest.
//    req := client.SplitShardRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesis-2013-12-02/SplitShard
func (c *Client) SplitShardRequest(input *SplitShardInput) SplitShardRequest {
	op := &aws.Operation{
		Name:       opSplitShard,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &SplitShardInput{}
	}

	req := c.newRequest(op, input, &SplitShardOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return SplitShardRequest{Request: req, Input: input, Copy: c.SplitShardRequest}
}

// SplitShardRequest is the request type for the
// SplitShard API operation.
type SplitShardRequest struct {
	*aws.Request
	Input *SplitShardInput
	Copy  func(*SplitShardInput) SplitShardRequest
}

// Send marshals and sends the SplitShard API request.
func (r SplitShardRequest) Send(ctx context.Context) (*SplitShardResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SplitShardResponse{
		SplitShardOutput: r.Request.Data.(*SplitShardOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// SplitShardResponse is the response type for the
// SplitShard API operation.
type SplitShardResponse struct {
	*SplitShardOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SplitShard request.
func (r *SplitShardResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
