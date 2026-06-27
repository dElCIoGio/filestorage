package contracts

import (
	"context"
	"io"

	"github.com/dElCIoGio/filestorage/internal/modules/signing/domain"
)

type Principal struct {
	AccountID   string
	AccessKeyID string
}

type CreateMultipartUploadCommand struct {
	Principal   Principal
	Bucket      string
	Key         string
	ContentType domain.ContentType
	Metadata    map[string]string
}

type CreateMultipartUploadResult struct {
	UploadID string
	Bucket   string
	Key      string
}

type UploadPartCommand struct {
	Principal  Principal
	Bucket     string
	Key        string
	UploadID   string
	PartNumber int
	Body       io.Reader
	SizeBytes  int64
}

type UploadPartResult struct {
	ETag       string
	PartNumber int
	SizeBytes  int64
}

type UploadedPartData struct {
	PartNumber int
	ETag       string
}

type CompleteMultipartUploadCommand struct {
	Principal Principal
	Bucket    string
	Key       string
	UploadID  string
	Parts     []UploadedPartData
}

type CompleteMultipartUploadResult struct {
	Bucket    string
	Key       string
	ETag      string
	SizeBytes int64
}

type AbortMultipartUploadCommand struct {
	Principal Principal
	Bucket    string
	Key       string
	UploadID  string
}

type API interface {
	CreateMultipartUpload(ctx context.Context, data CreateMultipartUploadCommand) (CreateMultipartUploadResult, error)
	UploadPart(ctx context.Context, data UploadPartCommand) (UploadPartResult, error)
	CompleteMultipartUpload(ctx context.Context, data CompleteMultipartUploadCommand) (CompleteMultipartUploadResult, error)
	AbortMultipartUpload(ctx context.Context, data AbortMultipartUploadCommand) error
}
