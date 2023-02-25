package util

import (
	"mime/multipart"
	"strings"
)

func IsImage(header *multipart.FileHeader) bool {
	// Get the content type of the file
	contentType := header.Header.Get("Content-Type")

	// Check if the content type is an image MIME type
	return strings.HasPrefix(contentType, "image/")
}
