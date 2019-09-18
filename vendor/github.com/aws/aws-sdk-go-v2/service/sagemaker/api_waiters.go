// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
)

// WaitUntilEndpointDeleted uses the SageMaker API operation
// DescribeEndpoint to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Client) WaitUntilEndpointDeleted(ctx context.Context, input *DescribeEndpointInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilEndpointDeleted",
		MaxAttempts: 60,
		Delay:       aws.ConstantWaiterDelay(30 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:    aws.SuccessWaiterState,
				Matcher:  aws.ErrorWaiterMatch,
				Expected: "ValidationException",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathWaiterMatch, Argument: "EndpointStatus",
				Expected: "Failed",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeEndpointInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeEndpointRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}

// WaitUntilEndpointInService uses the SageMaker API operation
// DescribeEndpoint to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Client) WaitUntilEndpointInService(ctx context.Context, input *DescribeEndpointInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilEndpointInService",
		MaxAttempts: 120,
		Delay:       aws.ConstantWaiterDelay(30 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathWaiterMatch, Argument: "EndpointStatus",
				Expected: "InService",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathWaiterMatch, Argument: "EndpointStatus",
				Expected: "Failed",
			},
			{
				State:    aws.FailureWaiterState,
				Matcher:  aws.ErrorWaiterMatch,
				Expected: "ValidationException",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeEndpointInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeEndpointRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}

// WaitUntilNotebookInstanceDeleted uses the SageMaker API operation
// DescribeNotebookInstance to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Client) WaitUntilNotebookInstanceDeleted(ctx context.Context, input *DescribeNotebookInstanceInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilNotebookInstanceDeleted",
		MaxAttempts: 60,
		Delay:       aws.ConstantWaiterDelay(30 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:    aws.SuccessWaiterState,
				Matcher:  aws.ErrorWaiterMatch,
				Expected: "ValidationException",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathWaiterMatch, Argument: "NotebookInstanceStatus",
				Expected: "Failed",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeNotebookInstanceInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeNotebookInstanceRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}

// WaitUntilNotebookInstanceInService uses the SageMaker API operation
// DescribeNotebookInstance to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Client) WaitUntilNotebookInstanceInService(ctx context.Context, input *DescribeNotebookInstanceInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilNotebookInstanceInService",
		MaxAttempts: 60,
		Delay:       aws.ConstantWaiterDelay(30 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathWaiterMatch, Argument: "NotebookInstanceStatus",
				Expected: "InService",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathWaiterMatch, Argument: "NotebookInstanceStatus",
				Expected: "Failed",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeNotebookInstanceInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeNotebookInstanceRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}

// WaitUntilNotebookInstanceStopped uses the SageMaker API operation
// DescribeNotebookInstance to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Client) WaitUntilNotebookInstanceStopped(ctx context.Context, input *DescribeNotebookInstanceInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilNotebookInstanceStopped",
		MaxAttempts: 60,
		Delay:       aws.ConstantWaiterDelay(30 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathWaiterMatch, Argument: "NotebookInstanceStatus",
				Expected: "Stopped",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathWaiterMatch, Argument: "NotebookInstanceStatus",
				Expected: "Failed",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeNotebookInstanceInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeNotebookInstanceRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}

// WaitUntilTrainingJobCompletedOrStopped uses the SageMaker API operation
// DescribeTrainingJob to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Client) WaitUntilTrainingJobCompletedOrStopped(ctx context.Context, input *DescribeTrainingJobInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilTrainingJobCompletedOrStopped",
		MaxAttempts: 180,
		Delay:       aws.ConstantWaiterDelay(120 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathWaiterMatch, Argument: "TrainingJobStatus",
				Expected: "Completed",
			},
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathWaiterMatch, Argument: "TrainingJobStatus",
				Expected: "Stopped",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathWaiterMatch, Argument: "TrainingJobStatus",
				Expected: "Failed",
			},
			{
				State:    aws.FailureWaiterState,
				Matcher:  aws.ErrorWaiterMatch,
				Expected: "ValidationException",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeTrainingJobInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeTrainingJobRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}

// WaitUntilTransformJobCompletedOrStopped uses the SageMaker API operation
// DescribeTransformJob to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Client) WaitUntilTransformJobCompletedOrStopped(ctx context.Context, input *DescribeTransformJobInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilTransformJobCompletedOrStopped",
		MaxAttempts: 60,
		Delay:       aws.ConstantWaiterDelay(60 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathWaiterMatch, Argument: "TransformJobStatus",
				Expected: "Completed",
			},
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathWaiterMatch, Argument: "TransformJobStatus",
				Expected: "Stopped",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathWaiterMatch, Argument: "TransformJobStatus",
				Expected: "Failed",
			},
			{
				State:    aws.FailureWaiterState,
				Matcher:  aws.ErrorWaiterMatch,
				Expected: "ValidationException",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeTransformJobInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeTransformJobRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}
