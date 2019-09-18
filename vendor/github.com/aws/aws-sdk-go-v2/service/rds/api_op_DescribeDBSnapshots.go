// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/DescribeDBSnapshotsMessage
type DescribeDBSnapshotsInput struct {
	_ struct{} `type:"structure"`

	// The ID of the DB instance to retrieve the list of DB snapshots for. This
	// parameter can't be used in conjunction with DBSnapshotIdentifier. This parameter
	// is not case-sensitive.
	//
	// Constraints:
	//
	//    * If supplied, must match the identifier of an existing DBInstance.
	DBInstanceIdentifier *string `type:"string"`

	// A specific DB snapshot identifier to describe. This parameter can't be used
	// in conjunction with DBInstanceIdentifier. This value is stored as a lowercase
	// string.
	//
	// Constraints:
	//
	//    * If supplied, must match the identifier of an existing DBSnapshot.
	//
	//    * If this identifier is for an automated snapshot, the SnapshotType parameter
	//    must also be specified.
	DBSnapshotIdentifier *string `type:"string"`

	// A specific DB resource ID to describe.
	DbiResourceId *string `type:"string"`

	// A filter that specifies one or more DB snapshots to describe.
	//
	// Supported filters:
	//
	//    * db-instance-id - Accepts DB instance identifiers and DB instance Amazon
	//    Resource Names (ARNs).
	//
	//    * db-snapshot-id - Accepts DB snapshot identifiers.
	//
	//    * dbi-resource-id - Accepts identifiers of source DB instances.
	//
	//    * snapshot-type - Accepts types of DB snapshots.
	//
	//    * engine - Accepts names of database engines.
	Filters []Filter `locationNameList:"Filter" type:"list"`

	// A value that indicates whether to include manual DB cluster snapshots that
	// are public and can be copied or restored by any AWS account. By default,
	// the public snapshots are not included.
	//
	// You can share a manual DB snapshot as public by using the ModifyDBSnapshotAttribute
	// API.
	IncludePublic *bool `type:"boolean"`

	// A value that indicates whether to include shared manual DB cluster snapshots
	// from other AWS accounts that this AWS account has been given permission to
	// copy or restore. By default, these snapshots are not included.
	//
	// You can give an AWS account permission to restore a manual DB snapshot from
	// another AWS account by using the ModifyDBSnapshotAttribute API action.
	IncludeShared *bool `type:"boolean"`

	// An optional pagination token provided by a previous DescribeDBSnapshots request.
	// If this parameter is specified, the response includes only records beyond
	// the marker, up to the value specified by MaxRecords.
	Marker *string `type:"string"`

	// The maximum number of records to include in the response. If more records
	// exist than the specified MaxRecords value, a pagination token called a marker
	// is included in the response so that the remaining results can be retrieved.
	//
	// Default: 100
	//
	// Constraints: Minimum 20, maximum 100.
	MaxRecords *int64 `type:"integer"`

	// The type of snapshots to be returned. You can specify one of the following
	// values:
	//
	//    * automated - Return all DB snapshots that have been automatically taken
	//    by Amazon RDS for my AWS account.
	//
	//    * manual - Return all DB snapshots that have been taken by my AWS account.
	//
	//    * shared - Return all manual DB snapshots that have been shared to my
	//    AWS account.
	//
	//    * public - Return all DB snapshots that have been marked as public.
	//
	//    * awsbackup - Return the DB snapshots managed by the AWS Backup service.
	//    For information about AWS Backup, see the AWS Backup Developer Guide.
	//    (https://docs.aws.amazon.com/aws-backup/latest/devguide/whatisbackup.html)
	//    The awsbackup type does not apply to Aurora.
	//
	// If you don't specify a SnapshotType value, then both automated and manual
	// snapshots are returned. Shared and public DB snapshots are not included in
	// the returned results by default. You can include shared snapshots with these
	// results by enabling the IncludeShared parameter. You can include public snapshots
	// with these results by enabling the IncludePublic parameter.
	//
	// The IncludeShared and IncludePublic parameters don't apply for SnapshotType
	// values of manual or automated. The IncludePublic parameter doesn't apply
	// when SnapshotType is set to shared. The IncludeShared parameter doesn't apply
	// when SnapshotType is set to public.
	SnapshotType *string `type:"string"`
}

// String returns the string representation
func (s DescribeDBSnapshotsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeDBSnapshotsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeDBSnapshotsInput"}
	if s.Filters != nil {
		for i, v := range s.Filters {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Filters", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the result of a successful invocation of the DescribeDBSnapshots
// action.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/DBSnapshotMessage
type DescribeDBSnapshotsOutput struct {
	_ struct{} `type:"structure"`

	// A list of DBSnapshot instances.
	DBSnapshots []DBSnapshot `locationNameList:"DBSnapshot" type:"list"`

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to
	// the value specified by MaxRecords.
	Marker *string `type:"string"`
}

// String returns the string representation
func (s DescribeDBSnapshotsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeDBSnapshots = "DescribeDBSnapshots"

// DescribeDBSnapshotsRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Returns information about DB snapshots. This API action supports pagination.
//
//    // Example sending a request using DescribeDBSnapshotsRequest.
//    req := client.DescribeDBSnapshotsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/DescribeDBSnapshots
func (c *Client) DescribeDBSnapshotsRequest(input *DescribeDBSnapshotsInput) DescribeDBSnapshotsRequest {
	op := &aws.Operation{
		Name:       opDescribeDBSnapshots,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"Marker"},
			LimitToken:      "MaxRecords",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeDBSnapshotsInput{}
	}

	req := c.newRequest(op, input, &DescribeDBSnapshotsOutput{})
	return DescribeDBSnapshotsRequest{Request: req, Input: input, Copy: c.DescribeDBSnapshotsRequest}
}

// DescribeDBSnapshotsRequest is the request type for the
// DescribeDBSnapshots API operation.
type DescribeDBSnapshotsRequest struct {
	*aws.Request
	Input *DescribeDBSnapshotsInput
	Copy  func(*DescribeDBSnapshotsInput) DescribeDBSnapshotsRequest
}

// Send marshals and sends the DescribeDBSnapshots API request.
func (r DescribeDBSnapshotsRequest) Send(ctx context.Context) (*DescribeDBSnapshotsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeDBSnapshotsResponse{
		DescribeDBSnapshotsOutput: r.Request.Data.(*DescribeDBSnapshotsOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeDBSnapshotsRequestPaginator returns a paginator for DescribeDBSnapshots.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeDBSnapshotsRequest(input)
//   p := rds.NewDescribeDBSnapshotsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeDBSnapshotsPaginator(req DescribeDBSnapshotsRequest) DescribeDBSnapshotsPaginator {
	return DescribeDBSnapshotsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeDBSnapshotsInput
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

// DescribeDBSnapshotsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeDBSnapshotsPaginator struct {
	aws.Pager
}

func (p *DescribeDBSnapshotsPaginator) CurrentPage() *DescribeDBSnapshotsOutput {
	return p.Pager.CurrentPage().(*DescribeDBSnapshotsOutput)
}

// DescribeDBSnapshotsResponse is the response type for the
// DescribeDBSnapshots API operation.
type DescribeDBSnapshotsResponse struct {
	*DescribeDBSnapshotsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeDBSnapshots request.
func (r *DescribeDBSnapshotsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
