// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package neptune

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/neptune-2014-10-31/DescribeValidDBInstanceModificationsMessage
type DescribeValidDBInstanceModificationsInput struct {
	_ struct{} `type:"structure"`

	// The customer identifier or the ARN of your DB instance.
	//
	// DBInstanceIdentifier is a required field
	DBInstanceIdentifier *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeValidDBInstanceModificationsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeValidDBInstanceModificationsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeValidDBInstanceModificationsInput"}

	if s.DBInstanceIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBInstanceIdentifier"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/neptune-2014-10-31/DescribeValidDBInstanceModificationsResult
type DescribeValidDBInstanceModificationsOutput struct {
	_ struct{} `type:"structure"`

	// Information about valid modifications that you can make to your DB instance.
	// Contains the result of a successful call to the DescribeValidDBInstanceModifications
	// action. You can use this information when you call ModifyDBInstance.
	ValidDBInstanceModificationsMessage *ValidDBInstanceModificationsMessage `type:"structure"`
}

// String returns the string representation
func (s DescribeValidDBInstanceModificationsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeValidDBInstanceModifications = "DescribeValidDBInstanceModifications"

// DescribeValidDBInstanceModificationsRequest returns a request value for making API operation for
// Amazon Neptune.
//
// You can call DescribeValidDBInstanceModifications to learn what modifications
// you can make to your DB instance. You can use this information when you call
// ModifyDBInstance.
//
//    // Example sending a request using DescribeValidDBInstanceModificationsRequest.
//    req := client.DescribeValidDBInstanceModificationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/neptune-2014-10-31/DescribeValidDBInstanceModifications
func (c *Client) DescribeValidDBInstanceModificationsRequest(input *DescribeValidDBInstanceModificationsInput) DescribeValidDBInstanceModificationsRequest {
	op := &aws.Operation{
		Name:       opDescribeValidDBInstanceModifications,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeValidDBInstanceModificationsInput{}
	}

	req := c.newRequest(op, input, &DescribeValidDBInstanceModificationsOutput{})
	return DescribeValidDBInstanceModificationsRequest{Request: req, Input: input, Copy: c.DescribeValidDBInstanceModificationsRequest}
}

// DescribeValidDBInstanceModificationsRequest is the request type for the
// DescribeValidDBInstanceModifications API operation.
type DescribeValidDBInstanceModificationsRequest struct {
	*aws.Request
	Input *DescribeValidDBInstanceModificationsInput
	Copy  func(*DescribeValidDBInstanceModificationsInput) DescribeValidDBInstanceModificationsRequest
}

// Send marshals and sends the DescribeValidDBInstanceModifications API request.
func (r DescribeValidDBInstanceModificationsRequest) Send(ctx context.Context) (*DescribeValidDBInstanceModificationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeValidDBInstanceModificationsResponse{
		DescribeValidDBInstanceModificationsOutput: r.Request.Data.(*DescribeValidDBInstanceModificationsOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeValidDBInstanceModificationsResponse is the response type for the
// DescribeValidDBInstanceModifications API operation.
type DescribeValidDBInstanceModificationsResponse struct {
	*DescribeValidDBInstanceModificationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeValidDBInstanceModifications request.
func (r *DescribeValidDBInstanceModificationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
