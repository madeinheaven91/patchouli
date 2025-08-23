package service

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var minioClient *minio.Client

func InitMinio() error {
	endpoint := os.Getenv("MINIO_ENDPOINT")
	accessKeyID := os.Getenv("MINIO_ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("MINIO_SECRET_ACCESS_KEY")
	// FIXME: Do i need ssl?
	useSSL := false

	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return err
	}
	minioClient = client

	err = makeBucket("books")

	return err
}

func makeBucket(name string) error {
	ctx := context.Background()
	// FIXME: region??
	err := minioClient.MakeBucket(ctx, name, minio.MakeBucketOptions{Region: "us-east-1"})
	if err != nil {
		exists, errBucketExists := minioClient.BucketExists(ctx, name)
		if errBucketExists == nil && exists {
			log.Printf("Bucket %s already exists\n", name)
		} else {
			return err
		}
	} else {
		log.Printf("Successfully created bucket %s\n", name)
	}
	return nil
}

func GenerateFileName(data []byte, header string) (string, string, error) {
	detected := http.DetectContentType(data)
	mimetype := detected

	if strings.HasPrefix(header, "text") && !strings.HasPrefix(detected, "text") {
		return "", detected, fmt.Errorf("wrong mimetype in header: %s, detected %s", header, detected)
	}
	if strings.HasPrefix(header, "text") {
		mimetype = header
	}

	var suffix string
	switch mimetype {
	case "application/pdf":
		suffix = "pdf"
	case "application/msword":
		suffix = "doc"
	case "application/vnd.openxmlformats-officedocument.wordprocessingml.document":
		suffix = "docx"
	case "application/vnd.oasis.opendocument.text":
		suffix = "odt"
	case "application/epub+zip":
		suffix = "epub"
	case "text/plain", "text/plain; charset=utf-8":
		suffix = "txt"
	case "text/html", "text/html; charset=utf-8":
		suffix = "html"
	case "text/markdown", "text/markdown; charset=utf-8":
		suffix = "md"
	default:
		return "", detected, fmt.Errorf("unsupported mimetype: %s", detected)
	}

	filename := "req_" + uuid.New().String() + "." + suffix
	return filename, mimetype, nil
}

func UploadBook(r *http.Request) (string, error) {
	// Read up to first 512 bytes to pass to http.DetectContentType in generateFileName
	data := make([]byte, 512)
	n, err := r.Body.Read(data)
	if err != nil && err != io.EOF {
		return "", err
	}
	filename, contentType, err := GenerateFileName(data[:n], r.Header.Get("Content-Type"))
	if err != nil {
		return "", fmt.Errorf("invalid mimetype: %v", err)
	}

	reader := io.MultiReader(bytes.NewReader(data[:n]), r.Body)

	log.Printf("Uploading new request %s: %s\n", contentType, filename)

	info, err := minioClient.PutObject(context.Background(),
		"books",
		filename,
		reader,
		-1,
		minio.PutObjectOptions{ContentType: contentType})
	return info.Key, err
}

func FetchBook(name string, r *http.Request) (*minio.Object, error) {
	return minioClient.GetObject(r.Context(), "books", name, minio.GetObjectOptions{})
}

func FetchBookStat(name string, r *http.Request) (minio.ObjectInfo, error) {
	return minioClient.StatObject(r.Context(), "books", name, minio.StatObjectOptions{})
}

func DeleteBook(name string, r *http.Request) error {
	return minioClient.RemoveObject(r.Context(), "books", name, minio.RemoveObjectOptions{})
}

func RenameBook(dst string, src string, r *http.Request) (minio.UploadInfo, error) {
	info, err := minioClient.CopyObject(r.Context(),
		minio.CopyDestOptions{Bucket: "books", Object: dst},
		minio.CopySrcOptions{Bucket: "books", Object: src},
	)
	if err != nil {
		return info, err
	}

	err = DeleteBook(src, r)
	return info, err
}
