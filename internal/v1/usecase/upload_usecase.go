package usecase

import (
	"context"
	"elotus/config"
	_const "elotus/const"
	"elotus/global"
	"elotus/internal/v1/repository"
	"elotus/internal/v1/repository/model"
	"elotus/package/util"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"time"

	"go.uber.org/zap"
)

type UploadUseCase interface {
	UploadFile(ctx context.Context, fileHeader *multipart.FileHeader) (*model.UploadedFile, error)
}

type UploadUseCaseImpl struct {
	fileRepo repository.UploadFileRepository
}

func NewUploadUseCase(fileRepo repository.UploadFileRepository) UploadUseCase {
	return &UploadUseCaseImpl{fileRepo: fileRepo}
}

func (u *UploadUseCaseImpl) UploadFile(ctx context.Context, fileHeader *multipart.FileHeader) (*model.UploadedFile, error) {
	// open file
	file, err := fileHeader.Open()
	if err != nil {
		global.Logger.Error("open file error", zap.Error(err))
		return nil, fmt.Errorf(_const.FailedToOpenFile)
	}
	defer file.Close()

	// check file type
	if !util.IsImage(fileHeader) {
		return nil, fmt.Errorf(_const.InvalidFileType)
	}

	// check file size
	if fileHeader.Size > 1024*1024*8 {
		return nil, fmt.Errorf(_const.FileTooLarge)
	}

	// Create temporary file in /tmp directory
	fileName := time.Now().Format(time.DateTime) + "." + fileHeader.Filename
	tempFile, err := ioutil.TempFile(config.StoragePath, fileName)
	if err != nil {
		global.Logger.Error("create temp file error", zap.Error(err))
		return nil, fmt.Errorf(_const.InternalServerErr)
	}
	defer tempFile.Close()

	// Write file to temporary file
	_, err = io.Copy(tempFile, file)
	if err != nil {
		global.Logger.Error("copy file error", zap.Error(err))
		return nil, fmt.Errorf(_const.InternalServerErr)
	}

	//write file mete data to db
	data, _ := json.Marshal(fileHeader)
	meta := &model.UploadedFile{
		Path:     fmt.Sprintf("%s/%s", config.StoragePath, fileName),
		MeteData: string(data),
	}
	err = u.fileRepo.Create(meta)
	if err != nil {
		global.Logger.Error("create file error", zap.Error(err))
		return nil, fmt.Errorf(_const.InternalServerErr)
	}

	return meta, nil
}
