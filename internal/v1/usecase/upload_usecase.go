package usecase

import (
	"context"
	"elotus/internal/v1/repository"
	"elotus/internal/v1/repository/model"
	"mime/multipart"
)

type UploadUseCase interface {
	UploadFile(ctx context.Context, file multipart.File) (*model.UploadedFile, error)
}

type UploadUseCaseImpl struct {
	fileRepo repository.UploadFileRepository
}

func NewUploadUseCase(fileRepo repository.UploadFileRepository) UploadUseCase {
	return &UploadUseCaseImpl{fileRepo: fileRepo}
}

func (u *UploadUseCaseImpl) UploadFile(ctx context.Context, file multipart.File) (*model.UploadedFile, error) {
	panic("implement me")
}
