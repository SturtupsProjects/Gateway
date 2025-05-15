package minio

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
	"mime/multipart"
	"net/http"
)

var MinioClient *minio.Client
var Endpoint = "minio.smartadmin.uz" // Уберите https://

func InitMiniOClient() error {
	accessKeyID := "minioadmin"
	secretAccessKey := "minioadmin"

	minioClient, err := minio.New(Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: true, // Установите true для использования HTTPS
	})

	if err != nil {
		log.Println(err)
		return err
	}

	MinioClient = minioClient

	return nil
}

func UploadMedia(fileHeader *multipart.FileHeader) (string, error) {
	// Open the uploaded file
	file, err := fileHeader.Open()
	if err != nil {
		log.Println("1-", err)
		return "", err
	}
	defer file.Close()

	// Detect the content type dynamically
	contentType := fileHeader.Header.Get("Content-Type")
	if contentType == "" {
		buffer := make([]byte, 512) // Read the first 512 bytes to detect the MIME type
		_, err = file.Read(buffer)
		if err != nil {
			log.Println("2-", err)
			return "", err
		}
		contentType = http.DetectContentType(buffer)
		// Reset the file pointer to the beginning
		if _, err = file.Seek(0, 0); err != nil {
			log.Println("3-", err)
			return "", err
		}
	}

	fileHeader.Filename += uuid.NewString()

	// Upload the file to MinIO
	_, err = MinioClient.PutObject(context.Background(), "media", fileHeader.Filename, file, fileHeader.Size, minio.PutObjectOptions{
		ContentType: contentType,
	})

	if err != nil {
		log.Println("4-", err)
		return "", err
	}

	// Generate the file URL
	imageUrl := fmt.Sprintf("https://%s/%s/%s", Endpoint, "media", fileHeader.Filename)

	return imageUrl, nil
}
