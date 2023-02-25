package controller

import (
	_const "elotus/const"
	"elotus/internal/v1/entity"
	"elotus/internal/v1/usecase"

	"github.com/gofiber/fiber/v2"
)

type UploadController struct {
	uploadUseCase usecase.UploadUseCase
}

func NewUploadController(uploadUseCase usecase.UploadUseCase) *UploadController {
	return &UploadController{uploadUseCase: uploadUseCase}
}

// UploadFile godoc
// @Summary upload an image
// @Tags Uploader
// @ID upload-an-image
// @Accept multipart/form-data
// @Produce json
// @Param data formData file true "File to upload"
// @Success 200 {object} model.UploadedFile
// @Failure 400 {object} entity.ResponseError
// @Router /api/v1/upload [post]
// @Security BearerAuth
func (u *UploadController) UploadFile(ctx *fiber.Ctx) error {

	// Parse form data
	form, err := ctx.MultipartForm()
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(entity.ResponseError{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	// Get file from form data
	fileHeaders := form.File["data"]
	if len(fileHeaders) == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(entity.ResponseError{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	res, err := u.uploadUseCase.UploadFile(ctx.Context(), fileHeaders[0])
	if err != nil {
		errorCode := fiber.StatusInternalServerError
		if code, ok := _const.ErrorCode[err.Error()]; ok {
			errorCode = code
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(entity.ResponseError{
			Code:    errorCode,
			Message: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(res)
}
