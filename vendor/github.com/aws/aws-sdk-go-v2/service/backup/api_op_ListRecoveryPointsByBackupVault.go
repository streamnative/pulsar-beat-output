// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package backup

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/backup-2018-11-15/ListRecoveryPointsByBackupVaultInput
type ListRecoveryPointsByBackupVaultInput struct {
	_ struct{} `type:"structure"`

	// The name of a logical container where backups are stored. Backup vaults are
	// identified by names that are unique to the account used to create them and
	// the AWS Region where they are created. They consist of lowercase letters,
	// numbers, and hyphens.
	//
	// BackupVaultName is a required field
	BackupVaultName *string `location:"uri" locationName:"backupVaultName" type:"string" required:"true"`

	// Returns only recovery points that match the specified backup plan ID.
	ByBackupPlanId *string `location:"querystring" locationName:"backupPlanId" type:"string"`

	// Returns only recovery points that were created after the specified timestamp.
	ByCreatedAfter *time.Time `location:"querystring" locationName:"createdAfter" type:"timestamp"`

	// Returns only recovery points that were created before the specified timestamp.
	ByCreatedBefore *time.Time `location:"querystring" locationName:"createdBefore" type:"timestamp"`

	// Returns only recovery points that match the specified resource Amazon Resource
	// Name (ARN).
	ByResourceArn *string `location:"querystring" locationName:"resourceArn" type:"string"`

	// Returns only recovery points that match the specified resource type.
	ByResourceType *string `location:"querystring" locationName:"resourceType" type:"string"`

	// The maximum number of items to be returned.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// The next item following a partial list of returned items. For example, if
	// a request is made to return maxResults number of items, NextToken allows
	// you to return more items in your list starting at the location pointed to
	// by the next token.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListRecoveryPointsByBackupVaultInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListRecoveryPointsByBackupVaultInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListRecoveryPointsByBackupVaultInput"}

	if s.BackupVaultName == nil {
		invalidParams.Add(aws.NewErrParamRequired("BackupVaultName"))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListRecoveryPointsByBackupVaultInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.BackupVaultName != nil {
		v := *s.BackupVaultName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "backupVaultName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ByBackupPlanId != nil {
		v := *s.ByBackupPlanId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "backupPlanId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ByCreatedAfter != nil {
		v := *s.ByCreatedAfter

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "createdAfter",
			protocol.TimeValue{V: v, Format: protocol.ISO8601TimeFormatName, QuotedFormatTime: false}, metadata)
	}
	if s.ByCreatedBefore != nil {
		v := *s.ByCreatedBefore

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "createdBefore",
			protocol.TimeValue{V: v, Format: protocol.ISO8601TimeFormatName, QuotedFormatTime: false}, metadata)
	}
	if s.ByResourceArn != nil {
		v := *s.ByResourceArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "resourceArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ByResourceType != nil {
		v := *s.ByResourceType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "resourceType", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/backup-2018-11-15/ListRecoveryPointsByBackupVaultOutput
type ListRecoveryPointsByBackupVaultOutput struct {
	_ struct{} `type:"structure"`

	// The next item following a partial list of returned items. For example, if
	// a request is made to return maxResults number of items, NextToken allows
	// you to return more items in your list starting at the location pointed to
	// by the next token.
	NextToken *string `type:"string"`

	// An array of objects that contain detailed information about recovery points
	// saved in a backup vault.
	RecoveryPoints []RecoveryPointByBackupVault `type:"list"`
}

// String returns the string representation
func (s ListRecoveryPointsByBackupVaultOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListRecoveryPointsByBackupVaultOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RecoveryPoints != nil {
		v := s.RecoveryPoints

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "RecoveryPoints", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opListRecoveryPointsByBackupVault = "ListRecoveryPointsByBackupVault"

// ListRecoveryPointsByBackupVaultRequest returns a request value for making API operation for
// AWS Backup.
//
// Returns detailed information about the recovery points stored in a backup
// vault.
//
//    // Example sending a request using ListRecoveryPointsByBackupVaultRequest.
//    req := client.ListRecoveryPointsByBackupVaultRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/backup-2018-11-15/ListRecoveryPointsByBackupVault
func (c *Client) ListRecoveryPointsByBackupVaultRequest(input *ListRecoveryPointsByBackupVaultInput) ListRecoveryPointsByBackupVaultRequest {
	op := &aws.Operation{
		Name:       opListRecoveryPointsByBackupVault,
		HTTPMethod: "GET",
		HTTPPath:   "/backup-vaults/{backupVaultName}/recovery-points/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListRecoveryPointsByBackupVaultInput{}
	}

	req := c.newRequest(op, input, &ListRecoveryPointsByBackupVaultOutput{})
	return ListRecoveryPointsByBackupVaultRequest{Request: req, Input: input, Copy: c.ListRecoveryPointsByBackupVaultRequest}
}

// ListRecoveryPointsByBackupVaultRequest is the request type for the
// ListRecoveryPointsByBackupVault API operation.
type ListRecoveryPointsByBackupVaultRequest struct {
	*aws.Request
	Input *ListRecoveryPointsByBackupVaultInput
	Copy  func(*ListRecoveryPointsByBackupVaultInput) ListRecoveryPointsByBackupVaultRequest
}

// Send marshals and sends the ListRecoveryPointsByBackupVault API request.
func (r ListRecoveryPointsByBackupVaultRequest) Send(ctx context.Context) (*ListRecoveryPointsByBackupVaultResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListRecoveryPointsByBackupVaultResponse{
		ListRecoveryPointsByBackupVaultOutput: r.Request.Data.(*ListRecoveryPointsByBackupVaultOutput),
		response:                              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListRecoveryPointsByBackupVaultRequestPaginator returns a paginator for ListRecoveryPointsByBackupVault.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListRecoveryPointsByBackupVaultRequest(input)
//   p := backup.NewListRecoveryPointsByBackupVaultRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListRecoveryPointsByBackupVaultPaginator(req ListRecoveryPointsByBackupVaultRequest) ListRecoveryPointsByBackupVaultPaginator {
	return ListRecoveryPointsByBackupVaultPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListRecoveryPointsByBackupVaultInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// ListRecoveryPointsByBackupVaultPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListRecoveryPointsByBackupVaultPaginator struct {
	aws.Pager
}

func (p *ListRecoveryPointsByBackupVaultPaginator) CurrentPage() *ListRecoveryPointsByBackupVaultOutput {
	return p.Pager.CurrentPage().(*ListRecoveryPointsByBackupVaultOutput)
}

// ListRecoveryPointsByBackupVaultResponse is the response type for the
// ListRecoveryPointsByBackupVault API operation.
type ListRecoveryPointsByBackupVaultResponse struct {
	*ListRecoveryPointsByBackupVaultOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListRecoveryPointsByBackupVault request.
func (r *ListRecoveryPointsByBackupVaultResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
