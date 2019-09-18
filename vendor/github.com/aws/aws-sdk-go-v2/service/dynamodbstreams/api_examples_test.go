// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package dynamodbstreams_test

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/awserr"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/dynamodbstreams"
)

var _ time.Duration
var _ strings.Reader
var _ aws.Config

func parseTime(layout, value string) *time.Time {
	t, err := time.Parse(layout, value)
	if err != nil {
		panic(err)
	}
	return &t
}

// To describe a stream with a given stream ARN
//
// The following example describes a stream with a given stream ARN.
func ExampleClient_DescribeStreamRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := dynamodbstreams.New(cfg)
	input := &dynamodbstreams.DescribeStreamInput{
		StreamArn: aws.String("arn:aws:dynamodb:us-west-2:111122223333:table/Forum/stream/2015-05-20T20:51:10.252"),
	}

	req := svc.DescribeStreamRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodbstreams.ErrCodeResourceNotFoundException:
				fmt.Println(dynamodbstreams.ErrCodeResourceNotFoundException, aerr.Error())
			case dynamodbstreams.ErrCodeInternalServerError:
				fmt.Println(dynamodbstreams.ErrCodeInternalServerError, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To retrieve all the stream records from a shard
//
// The following example retrieves all the stream records from a shard.
func ExampleClient_GetRecordsRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := dynamodbstreams.New(cfg)
	input := &dynamodbstreams.GetRecordsInput{
		ShardIterator: aws.String("arn:aws:dynamodb:us-west-2:111122223333:table/Forum/stream/2015-05-20T20:51:10.252|1|AAAAAAAAAAEvJp6D+zaQ...  <remaining characters omitted> ..."),
	}

	req := svc.GetRecordsRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodbstreams.ErrCodeResourceNotFoundException:
				fmt.Println(dynamodbstreams.ErrCodeResourceNotFoundException, aerr.Error())
			case dynamodbstreams.ErrCodeLimitExceededException:
				fmt.Println(dynamodbstreams.ErrCodeLimitExceededException, aerr.Error())
			case dynamodbstreams.ErrCodeInternalServerError:
				fmt.Println(dynamodbstreams.ErrCodeInternalServerError, aerr.Error())
			case dynamodbstreams.ErrCodeExpiredIteratorException:
				fmt.Println(dynamodbstreams.ErrCodeExpiredIteratorException, aerr.Error())
			case dynamodbstreams.ErrCodeTrimmedDataAccessException:
				fmt.Println(dynamodbstreams.ErrCodeTrimmedDataAccessException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To obtain a shard iterator for the provided stream ARN and shard ID
//
// The following example returns a shard iterator for the provided stream ARN and shard
// ID.
func ExampleClient_GetShardIteratorRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := dynamodbstreams.New(cfg)
	input := &dynamodbstreams.GetShardIteratorInput{
		ShardId:           aws.String("00000001414576573621-f55eea83"),
		ShardIteratorType: dynamodbstreams.ShardIteratorTypeTrimHorizon,
		StreamArn:         aws.String("arn:aws:dynamodb:us-west-2:111122223333:table/Forum/stream/2015-05-20T20:51:10.252"),
	}

	req := svc.GetShardIteratorRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodbstreams.ErrCodeResourceNotFoundException:
				fmt.Println(dynamodbstreams.ErrCodeResourceNotFoundException, aerr.Error())
			case dynamodbstreams.ErrCodeInternalServerError:
				fmt.Println(dynamodbstreams.ErrCodeInternalServerError, aerr.Error())
			case dynamodbstreams.ErrCodeTrimmedDataAccessException:
				fmt.Println(dynamodbstreams.ErrCodeTrimmedDataAccessException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To list all of the stream ARNs
//
// The following example lists all of the stream ARNs.
func ExampleClient_ListStreamsRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := dynamodbstreams.New(cfg)
	input := &dynamodbstreams.ListStreamsInput{}

	req := svc.ListStreamsRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodbstreams.ErrCodeResourceNotFoundException:
				fmt.Println(dynamodbstreams.ErrCodeResourceNotFoundException, aerr.Error())
			case dynamodbstreams.ErrCodeInternalServerError:
				fmt.Println(dynamodbstreams.ErrCodeInternalServerError, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}
