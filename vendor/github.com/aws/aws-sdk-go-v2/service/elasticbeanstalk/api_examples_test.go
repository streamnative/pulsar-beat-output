// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticbeanstalk_test

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/awserr"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
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

// To abort a deployment
//
// The following code aborts a running application version deployment for an environment
// named my-env:
func ExampleClient_AbortEnvironmentUpdateRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := elasticbeanstalk.New(cfg)
	input := &elasticbeanstalk.AbortEnvironmentUpdateInput{
		EnvironmentName: aws.String("my-env"),
	}

	req := svc.AbortEnvironmentUpdateRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case elasticbeanstalk.ErrCodeInsufficientPrivilegesException:
				fmt.Println(elasticbeanstalk.ErrCodeInsufficientPrivilegesException, aerr.Error())
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

// To check the availability of a CNAME
//
// The following operation checks the availability of the subdomain my-cname:
func ExampleClient_CheckDNSAvailabilityRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := elasticbeanstalk.New(cfg)
	input := &elasticbeanstalk.CheckDNSAvailabilityInput{
		CNAMEPrefix: aws.String("my-cname"),
	}

	req := svc.CheckDNSAvailabilityRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
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

// To create a new application
//
// The following operation creates a new application named my-app:
func ExampleClient_CreateApplicationRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := elasticbeanstalk.New(cfg)
	input := &elasticbeanstalk.CreateApplicationInput{
		ApplicationName: aws.String("my-app"),
		Description:     aws.String("my application"),
	}

	req := svc.CreateApplicationRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case elasticbeanstalk.ErrCodeTooManyApplicationsException:
				fmt.Println(elasticbeanstalk.ErrCodeTooManyApplicationsException, aerr.Error())
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

// To create a new application
//
// The following operation creates a new version (v1) of an application named my-app:
func ExampleClient_CreateApplicationVersionRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := elasticbeanstalk.New(cfg)
	input := &elasticbeanstalk.CreateApplicationVersionInput{
		ApplicationName:       aws.String("my-app"),
		AutoCreateApplication: aws.Bool(true),
		Description:           aws.String("my-app-v1"),
		Process:               aws.Bool(true),
		SourceBundle: &elasticbeanstalk.S3Location{
			S3Bucket: aws.String("my-bucket"),
			S3Key:    aws.String("sample.war"),
		},
		VersionLabel: aws.String("v1"),
	}

	req := svc.CreateApplicationVersionRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case elasticbeanstalk.ErrCodeTooManyApplicationsException:
				fmt.Println(elasticbeanstalk.ErrCodeTooManyApplicationsException, aerr.Error())
			case elasticbeanstalk.ErrCodeTooManyApplicationVersionsException:
				fmt.Println(elasticbeanstalk.ErrCodeTooManyApplicationVersionsException, aerr.Error())
			case elasticbeanstalk.ErrCodeInsufficientPrivilegesException:
				fmt.Println(elasticbeanstalk.ErrCodeInsufficientPrivilegesException, aerr.Error())
			case elasticbeanstalk.ErrCodeS3LocationNotInServiceRegionException:
				fmt.Println(elasticbeanstalk.ErrCodeS3LocationNotInServiceRegionException, aerr.Error())
			case elasticbeanstalk.ErrCodeCodeBuildNotInServiceRegionException:
				fmt.Println(elasticbeanstalk.ErrCodeCodeBuildNotInServiceRegionException, aerr.Error())
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

// To create a configuration template
//
// The following operation creates a configuration template named my-app-v1 from the
// settings applied to an environment with the id e-rpqsewtp2j:
func ExampleClient_CreateConfigurationTemplateRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := elasticbeanstalk.New(cfg)
	input := &elasticbeanstalk.CreateConfigurationTemplateInput{
		ApplicationName: aws.String("my-app"),
		EnvironmentId:   aws.String("e-rpqsewtp2j"),
		TemplateName:    aws.String("my-app-v1"),
	}

	req := svc.CreateConfigurationTemplateRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case elasticbeanstalk.ErrCodeInsufficientPrivilegesException:
				fmt.Println(elasticbeanstalk.ErrCodeInsufficientPrivilegesException, aerr.Error())
			case elasticbeanstalk.ErrCodeTooManyBucketsException:
				fmt.Println(elasticbeanstalk.ErrCodeTooManyBucketsException, aerr.Error())
			case elasticbeanstalk.ErrCodeTooManyConfigurationTemplatesException:
				fmt.Println(elasticbeanstalk.ErrCodeTooManyConfigurationTemplatesException, aerr.Error())
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

// To create a new environment for an application
//
// The following operation creates a new environment for version v1 of a java application
// named my-app:
func ExampleClient_CreateEnvironmentRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := elasticbeanstalk.New(cfg)
	input := &elasticbeanstalk.CreateEnvironmentInput{
		ApplicationName:   aws.String("my-app"),
		CNAMEPrefix:       aws.String("my-app"),
		EnvironmentName:   aws.String("my-env"),
		SolutionStackName: aws.String("64bit Amazon Linux 2015.03 v2.0.0 running Tomcat 8 Java 8"),
		VersionLabel:      aws.String("v1"),
	}

	req := svc.CreateEnvironmentRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case elasticbeanstalk.ErrCodeTooManyEnvironmentsException:
				fmt.Println(elasticbeanstalk.ErrCodeTooManyEnvironmentsException, aerr.Error())
			case elasticbeanstalk.ErrCodeInsufficientPrivilegesException:
				fmt.Println(elasticbeanstalk.ErrCodeInsufficientPrivilegesException, aerr.Error())
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

// To create a new environment for an application
//
// The following operation creates a new environment for version v1 of a java application
// named my-app:
func ExampleClient_CreateStorageLocationRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := elasticbeanstalk.New(cfg)
	input := &elasticbeanstalk.CreateStorageLocationInput{}

	req := svc.CreateStorageLocationRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case elasticbeanstalk.ErrCodeTooManyBucketsException:
				fmt.Println(elasticbeanstalk.ErrCodeTooManyBucketsException, aerr.Error())
			case elasticbeanstalk.ErrCodeS3SubscriptionRequiredException:
				fmt.Println(elasticbeanstalk.ErrCodeS3SubscriptionRequiredException, aerr.Error())
			case elasticbeanstalk.ErrCodeInsufficientPrivilegesException:
				fmt.Println(elasticbeanstalk.ErrCodeInsufficientPrivilegesException, aerr.Error())
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

// To delete an application
//
// The following operation deletes an application named my-app:
func ExampleClient_DeleteApplicationRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := elasticbeanstalk.New(cfg)
	input := &elasticbeanstalk.DeleteApplicationInput{
		ApplicationName: aws.String("my-app"),
	}

	req := svc.DeleteApplicationRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case elasticbeanstalk.ErrCodeOperationInProgressException:
				fmt.Println(elasticbeanstalk.ErrCodeOperationInProgressException, aerr.Error())
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

// To delete an application version
//
// The following operation deletes an application version named 22a0-stage-150819_182129
// for an application named my-app:
func ExampleClient_DeleteApplicationVersionRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := elasticbeanstalk.New(cfg)
	input := &elasticbeanstalk.DeleteApplicationVersionInput{
		ApplicationName:    aws.String("my-app"),
		DeleteSourceBundle: aws.Bool(true),
		VersionLabel:       aws.String("22a0-stage-150819_182129"),
	}

	req := svc.DeleteApplicationVersionRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case elasticbeanstalk.ErrCodeSourceBundleDeletionException:
				fmt.Println(elasticbeanstalk.ErrCodeSourceBundleDeletionException, aerr.Error())
			case elasticbeanstalk.ErrCodeInsufficientPrivilegesException:
				fmt.Println(elasticbeanstalk.ErrCodeInsufficientPrivilegesException, aerr.Error())
			case elasticbeanstalk.ErrCodeOperationInProgressException:
				fmt.Println(elasticbeanstalk.ErrCodeOperationInProgressException, aerr.Error())
			case elasticbeanstalk.ErrCodeS3LocationNotInServiceRegionException:
				fmt.Println(elasticbeanstalk.ErrCodeS3LocationNotInServiceRegionException, aerr.Error())
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

// To delete a configuration template
//
// The following operation deletes a configuration template named my-template for an
// application named my-app:
func ExampleClient_DeleteConfigurationTemplateRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := elasticbeanstalk.New(cfg)
	input := &elasticbeanstalk.DeleteConfigurationTemplateInput{
		ApplicationName: aws.String("my-app"),
		TemplateName:    aws.String("my-template"),
	}

	req := svc.DeleteConfigurationTemplateRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case elasticbeanstalk.ErrCodeOperationInProgressException:
				fmt.Println(elasticbeanstalk.ErrCodeOperationInProgressException, aerr.Error())
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

// To delete a draft configuration
//
// The following operation deletes a draft configuration for an environment named my-env:
func ExampleClient_DeleteEnvironmentConfigurationRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := elasticbeanstalk.New(cfg)
	input := &elasticbeanstalk.DeleteEnvironmentConfigurationInput{
		ApplicationName: aws.String("my-app"),
		EnvironmentName: aws.String("my-env"),
	}

	req := svc.DeleteEnvironmentConfigurationRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
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

// To view information about an application version
//
// The following operation retrieves information about an application version labeled
// v2:
func ExampleClient_DescribeApplicationVersionsRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := elasticbeanstalk.New(cfg)
	input := &elasticbeanstalk.DescribeApplicationVersionsInput{
		ApplicationName: aws.String("my-app"),
		VersionLabels: []string{
			"v2",
		},
	}

	req := svc.DescribeApplicationVersionsRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
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

// To view a list of applications
//
// The following operation retrieves information about applications in the current region:
func ExampleClient_DescribeApplicationsRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := elasticbeanstalk.New(cfg)
	input := &elasticbeanstalk.DescribeApplicationsInput{}

	req := svc.DescribeApplicationsRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
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

// To view configuration options for an environment
//
// The following operation retrieves descriptions of all available configuration options
// for an environment named my-env:
func ExampleClient_DescribeConfigurationOptionsRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := elasticbeanstalk.New(cfg)
	input := &elasticbeanstalk.DescribeConfigurationOptionsInput{
		ApplicationName: aws.String("my-app"),
		EnvironmentName: aws.String("my-env"),
	}

	req := svc.DescribeConfigurationOptionsRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case elasticbeanstalk.ErrCodeTooManyBucketsException:
				fmt.Println(elasticbeanstalk.ErrCodeTooManyBucketsException, aerr.Error())
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

// To view configurations settings for an environment
//
// The following operation retrieves configuration settings for an environment named
// my-env:
func ExampleClient_DescribeConfigurationSettingsRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := elasticbeanstalk.New(cfg)
	input := &elasticbeanstalk.DescribeConfigurationSettingsInput{
		ApplicationName: aws.String("my-app"),
		EnvironmentName: aws.String("my-env"),
	}

	req := svc.DescribeConfigurationSettingsRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case elasticbeanstalk.ErrCodeTooManyBucketsException:
				fmt.Println(elasticbeanstalk.ErrCodeTooManyBucketsException, aerr.Error())
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

// To view environment health
//
// The following operation retrieves overall health information for an environment named
// my-env:
func ExampleClient_DescribeEnvironmentHealthRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := elasticbeanstalk.New(cfg)
	input := &elasticbeanstalk.DescribeEnvironmentHealthInput{
		AttributeNames: []elasticbeanstalk.EnvironmentHealthAttribute{
			elasticbeanstalk.EnvironmentHealthAttribute("All"),
		},
		EnvironmentName: aws.String("my-env"),
	}

	req := svc.DescribeEnvironmentHealthRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case elasticbeanstalk.ErrCodeInvalidRequestException:
				fmt.Println(elasticbeanstalk.ErrCodeInvalidRequestException, aerr.Error())
			case elasticbeanstalk.ErrCodeElasticBeanstalkServiceException:
				fmt.Println(elasticbeanstalk.ErrCodeElasticBeanstalkServiceException, aerr.Error())
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

// To view information about the AWS resources in your environment
//
// The following operation retrieves information about resources in an environment named
// my-env:
func ExampleClient_DescribeEnvironmentResourcesRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := elasticbeanstalk.New(cfg)
	input := &elasticbeanstalk.DescribeEnvironmentResourcesInput{
		EnvironmentName: aws.String("my-env"),
	}

	req := svc.DescribeEnvironmentResourcesRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case elasticbeanstalk.ErrCodeInsufficientPrivilegesException:
				fmt.Println(elasticbeanstalk.ErrCodeInsufficientPrivilegesException, aerr.Error())
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

// To view information about an environment
//
// The following operation retrieves information about an environment named my-env:
func ExampleClient_DescribeEnvironmentsRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := elasticbeanstalk.New(cfg)
	input := &elasticbeanstalk.DescribeEnvironmentsInput{
		EnvironmentNames: []string{
			"my-env",
		},
	}

	req := svc.DescribeEnvironmentsRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
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

// To view events for an environment
//
// The following operation retrieves events for an environment named my-env:
func ExampleClient_DescribeEventsRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := elasticbeanstalk.New(cfg)
	input := &elasticbeanstalk.DescribeEventsInput{
		EnvironmentName: aws.String("my-env"),
	}

	req := svc.DescribeEventsRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
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

// To view environment health
//
// The following operation retrieves health information for instances in an environment
// named my-env:
func ExampleClient_DescribeInstancesHealthRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := elasticbeanstalk.New(cfg)
	input := &elasticbeanstalk.DescribeInstancesHealthInput{
		AttributeNames: []elasticbeanstalk.InstancesHealthAttribute{
			elasticbeanstalk.InstancesHealthAttribute("All"),
		},
		EnvironmentName: aws.String("my-env"),
	}

	req := svc.DescribeInstancesHealthRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case elasticbeanstalk.ErrCodeInvalidRequestException:
				fmt.Println(elasticbeanstalk.ErrCodeInvalidRequestException, aerr.Error())
			case elasticbeanstalk.ErrCodeElasticBeanstalkServiceException:
				fmt.Println(elasticbeanstalk.ErrCodeElasticBeanstalkServiceException, aerr.Error())
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

// To view solution stacks
//
// The following operation lists solution stacks for all currently available platform
// configurations and any that you have used in the past:
func ExampleClient_ListAvailableSolutionStacksRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := elasticbeanstalk.New(cfg)
	input := &elasticbeanstalk.ListAvailableSolutionStacksInput{}

	req := svc.ListAvailableSolutionStacksRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
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

// To rebuild an environment
//
// The following operation terminates and recreates the resources in an environment
// named my-env:
func ExampleClient_RebuildEnvironmentRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := elasticbeanstalk.New(cfg)
	input := &elasticbeanstalk.RebuildEnvironmentInput{
		EnvironmentName: aws.String("my-env"),
	}

	req := svc.RebuildEnvironmentRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case elasticbeanstalk.ErrCodeInsufficientPrivilegesException:
				fmt.Println(elasticbeanstalk.ErrCodeInsufficientPrivilegesException, aerr.Error())
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

// To request tailed logs
//
// The following operation requests logs from an environment named my-env:
func ExampleClient_RequestEnvironmentInfoRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := elasticbeanstalk.New(cfg)
	input := &elasticbeanstalk.RequestEnvironmentInfoInput{
		EnvironmentName: aws.String("my-env"),
		InfoType:        elasticbeanstalk.EnvironmentInfoTypeTail,
	}

	req := svc.RequestEnvironmentInfoRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
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

// To restart application servers
//
// The following operation restarts application servers on all instances in an environment
// named my-env:
func ExampleClient_RestartAppServerRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := elasticbeanstalk.New(cfg)
	input := &elasticbeanstalk.RestartAppServerInput{
		EnvironmentName: aws.String("my-env"),
	}

	req := svc.RestartAppServerRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
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

// To retrieve tailed logs
//
// The following operation retrieves a link to logs from an environment named my-env:
func ExampleClient_RetrieveEnvironmentInfoRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := elasticbeanstalk.New(cfg)
	input := &elasticbeanstalk.RetrieveEnvironmentInfoInput{
		EnvironmentName: aws.String("my-env"),
		InfoType:        elasticbeanstalk.EnvironmentInfoTypeTail,
	}

	req := svc.RetrieveEnvironmentInfoRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
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

// To swap environment CNAMES
//
// The following operation swaps the assigned subdomains of two environments:
func ExampleClient_SwapEnvironmentCNAMEsRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := elasticbeanstalk.New(cfg)
	input := &elasticbeanstalk.SwapEnvironmentCNAMEsInput{
		DestinationEnvironmentName: aws.String("my-env-green"),
		SourceEnvironmentName:      aws.String("my-env-blue"),
	}

	req := svc.SwapEnvironmentCNAMEsRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
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

// To terminate an environment
//
// The following operation terminates an Elastic Beanstalk environment named my-env:
func ExampleClient_TerminateEnvironmentRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := elasticbeanstalk.New(cfg)
	input := &elasticbeanstalk.TerminateEnvironmentInput{
		EnvironmentName: aws.String("my-env"),
	}

	req := svc.TerminateEnvironmentRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case elasticbeanstalk.ErrCodeInsufficientPrivilegesException:
				fmt.Println(elasticbeanstalk.ErrCodeInsufficientPrivilegesException, aerr.Error())
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

// To change an application's description
//
// The following operation updates the description of an application named my-app:
func ExampleClient_UpdateApplicationRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := elasticbeanstalk.New(cfg)
	input := &elasticbeanstalk.UpdateApplicationInput{
		ApplicationName: aws.String("my-app"),
		Description:     aws.String("my Elastic Beanstalk application"),
	}

	req := svc.UpdateApplicationRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
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

// To change an application version's description
//
// The following operation updates the description of an application version named 22a0-stage-150819_185942:
func ExampleClient_UpdateApplicationVersionRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := elasticbeanstalk.New(cfg)
	input := &elasticbeanstalk.UpdateApplicationVersionInput{
		ApplicationName: aws.String("my-app"),
		Description:     aws.String("new description"),
		VersionLabel:    aws.String("22a0-stage-150819_185942"),
	}

	req := svc.UpdateApplicationVersionRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
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

// To update a configuration template
//
// The following operation removes the configured CloudWatch custom health metrics configuration
// ConfigDocument from a saved configuration template named my-template:
func ExampleClient_UpdateConfigurationTemplateRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := elasticbeanstalk.New(cfg)
	input := &elasticbeanstalk.UpdateConfigurationTemplateInput{
		ApplicationName: aws.String("my-app"),
		OptionsToRemove: []elasticbeanstalk.OptionSpecification{
			{
				Namespace:  aws.String("aws:elasticbeanstalk:healthreporting:system"),
				OptionName: aws.String("ConfigDocument"),
			},
		},
		TemplateName: aws.String("my-template"),
	}

	req := svc.UpdateConfigurationTemplateRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case elasticbeanstalk.ErrCodeInsufficientPrivilegesException:
				fmt.Println(elasticbeanstalk.ErrCodeInsufficientPrivilegesException, aerr.Error())
			case elasticbeanstalk.ErrCodeTooManyBucketsException:
				fmt.Println(elasticbeanstalk.ErrCodeTooManyBucketsException, aerr.Error())
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

// To update an environment to a new version
//
// The following operation updates an environment named "my-env" to version "v2" of
// the application to which it belongs:
func ExampleClient_UpdateEnvironmentRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := elasticbeanstalk.New(cfg)
	input := &elasticbeanstalk.UpdateEnvironmentInput{
		EnvironmentName: aws.String("my-env"),
		VersionLabel:    aws.String("v2"),
	}

	req := svc.UpdateEnvironmentRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case elasticbeanstalk.ErrCodeInsufficientPrivilegesException:
				fmt.Println(elasticbeanstalk.ErrCodeInsufficientPrivilegesException, aerr.Error())
			case elasticbeanstalk.ErrCodeTooManyBucketsException:
				fmt.Println(elasticbeanstalk.ErrCodeTooManyBucketsException, aerr.Error())
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

// To configure option settings
//
// The following operation configures several options in the aws:elb:loadbalancer namespace:
func ExampleClient_UpdateEnvironmentRequest_shared01() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := elasticbeanstalk.New(cfg)
	input := &elasticbeanstalk.UpdateEnvironmentInput{
		EnvironmentName: aws.String("my-env"),
		OptionSettings: []elasticbeanstalk.ConfigurationOptionSetting{
			{
				Namespace:  aws.String("aws:elb:healthcheck"),
				OptionName: aws.String("Interval"),
				Value:      aws.String("15"),
			},
			{
				Namespace:  aws.String("aws:elb:healthcheck"),
				OptionName: aws.String("Timeout"),
				Value:      aws.String("8"),
			},
			{
				Namespace:  aws.String("aws:elb:healthcheck"),
				OptionName: aws.String("HealthyThreshold"),
				Value:      aws.String("2"),
			},
			{
				Namespace:  aws.String("aws:elb:healthcheck"),
				OptionName: aws.String("UnhealthyThreshold"),
				Value:      aws.String("3"),
			},
		},
	}

	req := svc.UpdateEnvironmentRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case elasticbeanstalk.ErrCodeInsufficientPrivilegesException:
				fmt.Println(elasticbeanstalk.ErrCodeInsufficientPrivilegesException, aerr.Error())
			case elasticbeanstalk.ErrCodeTooManyBucketsException:
				fmt.Println(elasticbeanstalk.ErrCodeTooManyBucketsException, aerr.Error())
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

// To validate configuration settings
//
// The following operation validates a CloudWatch custom metrics config document:
func ExampleClient_ValidateConfigurationSettingsRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := elasticbeanstalk.New(cfg)
	input := &elasticbeanstalk.ValidateConfigurationSettingsInput{
		ApplicationName: aws.String("my-app"),
		EnvironmentName: aws.String("my-env"),
		OptionSettings: []elasticbeanstalk.ConfigurationOptionSetting{
			{
				Namespace:  aws.String("aws:elasticbeanstalk:healthreporting:system"),
				OptionName: aws.String("ConfigDocument"),
				Value:      aws.String("{\"CloudWatchMetrics\": {\"Environment\": {\"ApplicationLatencyP99.9\": null,\"InstancesSevere\": 60,\"ApplicationLatencyP90\": 60,\"ApplicationLatencyP99\": null,\"ApplicationLatencyP95\": 60,\"InstancesUnknown\": 60,\"ApplicationLatencyP85\": 60,\"InstancesInfo\": null,\"ApplicationRequests2xx\": null,\"InstancesDegraded\": null,\"InstancesWarning\": 60,\"ApplicationLatencyP50\": 60,\"ApplicationRequestsTotal\": null,\"InstancesNoData\": null,\"InstancesPending\": 60,\"ApplicationLatencyP10\": null,\"ApplicationRequests5xx\": null,\"ApplicationLatencyP75\": null,\"InstancesOk\": 60,\"ApplicationRequests3xx\": null,\"ApplicationRequests4xx\": null},\"Instance\": {\"ApplicationLatencyP99.9\": null,\"ApplicationLatencyP90\": 60,\"ApplicationLatencyP99\": null,\"ApplicationLatencyP95\": null,\"ApplicationLatencyP85\": null,\"CPUUser\": 60,\"ApplicationRequests2xx\": null,\"CPUIdle\": null,\"ApplicationLatencyP50\": null,\"ApplicationRequestsTotal\": 60,\"RootFilesystemUtil\": null,\"LoadAverage1min\": null,\"CPUIrq\": null,\"CPUNice\": 60,\"CPUIowait\": 60,\"ApplicationLatencyP10\": null,\"LoadAverage5min\": null,\"ApplicationRequests5xx\": null,\"ApplicationLatencyP75\": 60,\"CPUSystem\": 60,\"ApplicationRequests3xx\": 60,\"ApplicationRequests4xx\": null,\"InstanceHealth\": null,\"CPUSoftirq\": 60}},\"Version\": 1}"),
			},
		},
	}

	req := svc.ValidateConfigurationSettingsRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case elasticbeanstalk.ErrCodeInsufficientPrivilegesException:
				fmt.Println(elasticbeanstalk.ErrCodeInsufficientPrivilegesException, aerr.Error())
			case elasticbeanstalk.ErrCodeTooManyBucketsException:
				fmt.Println(elasticbeanstalk.ErrCodeTooManyBucketsException, aerr.Error())
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
