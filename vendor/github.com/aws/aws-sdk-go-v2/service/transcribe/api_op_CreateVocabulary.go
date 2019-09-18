// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package transcribe

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/transcribe-2017-10-26/CreateVocabularyRequest
type CreateVocabularyInput struct {
	_ struct{} `type:"structure"`

	// The language code of the vocabulary entries.
	//
	// LanguageCode is a required field
	LanguageCode LanguageCode `type:"string" required:"true" enum:"true"`

	// An array of strings that contains the vocabulary entries.
	Phrases []string `type:"list"`

	// The S3 location of the text file that contains the definition of the custom
	// vocabulary. The URI must be in the same region as the API endpoint that you
	// are calling. The general form is
	//
	// https://s3-<aws-region>.amazonaws.com/<bucket-name>/<keyprefix>/<objectkey>
	//
	// For example:
	//
	// https://s3-us-east-1.amazonaws.com/examplebucket/vocab.txt
	//
	// For more information about S3 object names, see Object Keys (http://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html#object-keys)
	// in the Amazon S3 Developer Guide.
	//
	// For more information about custom vocabularies, see Custom Vocabularies (http://docs.aws.amazon.com/transcribe/latest/dg/how-it-works.html#how-vocabulary).
	VocabularyFileUri *string `min:"1" type:"string"`

	// The name of the vocabulary. The name must be unique within an AWS account.
	// The name is case-sensitive.
	//
	// VocabularyName is a required field
	VocabularyName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateVocabularyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateVocabularyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateVocabularyInput"}
	if len(s.LanguageCode) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("LanguageCode"))
	}
	if s.VocabularyFileUri != nil && len(*s.VocabularyFileUri) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("VocabularyFileUri", 1))
	}

	if s.VocabularyName == nil {
		invalidParams.Add(aws.NewErrParamRequired("VocabularyName"))
	}
	if s.VocabularyName != nil && len(*s.VocabularyName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("VocabularyName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/transcribe-2017-10-26/CreateVocabularyResponse
type CreateVocabularyOutput struct {
	_ struct{} `type:"structure"`

	// If the VocabularyState field is FAILED, this field contains information about
	// why the job failed.
	FailureReason *string `type:"string"`

	// The language code of the vocabulary entries.
	LanguageCode LanguageCode `type:"string" enum:"true"`

	// The date and time that the vocabulary was created.
	LastModifiedTime *time.Time `type:"timestamp"`

	// The name of the vocabulary.
	VocabularyName *string `min:"1" type:"string"`

	// The processing state of the vocabulary. When the VocabularyState field contains
	// READY the vocabulary is ready to be used in a StartTranscriptionJob request.
	VocabularyState VocabularyState `type:"string" enum:"true"`
}

// String returns the string representation
func (s CreateVocabularyOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateVocabulary = "CreateVocabulary"

// CreateVocabularyRequest returns a request value for making API operation for
// Amazon Transcribe Service.
//
// Creates a new custom vocabulary that you can use to change the way Amazon
// Transcribe handles transcription of an audio file.
//
//    // Example sending a request using CreateVocabularyRequest.
//    req := client.CreateVocabularyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/transcribe-2017-10-26/CreateVocabulary
func (c *Client) CreateVocabularyRequest(input *CreateVocabularyInput) CreateVocabularyRequest {
	op := &aws.Operation{
		Name:       opCreateVocabulary,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateVocabularyInput{}
	}

	req := c.newRequest(op, input, &CreateVocabularyOutput{})
	return CreateVocabularyRequest{Request: req, Input: input, Copy: c.CreateVocabularyRequest}
}

// CreateVocabularyRequest is the request type for the
// CreateVocabulary API operation.
type CreateVocabularyRequest struct {
	*aws.Request
	Input *CreateVocabularyInput
	Copy  func(*CreateVocabularyInput) CreateVocabularyRequest
}

// Send marshals and sends the CreateVocabulary API request.
func (r CreateVocabularyRequest) Send(ctx context.Context) (*CreateVocabularyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateVocabularyResponse{
		CreateVocabularyOutput: r.Request.Data.(*CreateVocabularyOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateVocabularyResponse is the response type for the
// CreateVocabulary API operation.
type CreateVocabularyResponse struct {
	*CreateVocabularyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateVocabulary request.
func (r *CreateVocabularyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
