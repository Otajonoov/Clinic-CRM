package v1

import (
	"fmt"
	"net/http"
	"path/filepath"

	"gitlab.com/clinic-crm/api-gateway/api/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Upload photo
// @Summary 		Upload media
// @Description 	Through this api frontent can upload photo and get the link to the media.
// @Tags 			Media
// @Security        BearerAuth
// @Accept 			multipart/form-data
// @Produce         json
// @Param 			filter  query models.FileUploadRequest false "Filter"
// @Param         	file                  formData file true "File"
// @Success         200					  {object} 	models.MediaResponse
// @Failure         default               {object}  models.DefaultResponse
// @Router          /v1/media/photo/ [post]
func (h *handlerV1) UploadMedia(ctx *gin.Context) {
	file := &models.File{}
	err := ctx.ShouldBind(&file)
	if HandleBadRequestErrWithMessage(ctx, &h.log, err, "ctx.ShouldBind(&file)") {
		return
	}

	fileSize := file.File.Size
	if fileSize > (int64(h.cfg.MaxImageSize) << 20) {
		ctx.JSON(http.StatusBadRequest, models.DefaultResponse{
			ErrorCode:    ErrorCodeImageSizeExceed,
			ErrorMessage: fmt.Sprintf("Image size should be less than %d mb", h.cfg.MaxImageSize),
		})
		return
	}

	ext := filepath.Ext(file.File.Filename)
	if ext != ".jpg" && ext != ".png" && ext != ".pdf" {
		ctx.JSON(http.StatusBadRequest, models.DefaultResponse{
			ErrorCode:    ErrorCodeImageExtensionNotAllowed,
			ErrorMessage: "Only .pdf, .jpg and .png images are allowed",
		})
		return
	}

	file.File.Filename = uuid.New().String() + filepath.Ext(file.File.Filename)

	err = ctx.SaveUploadedFile(file.File, "./media/"+file.File.Filename)
	if HandleInternalWithMessage(ctx, &h.log, err, "UploadMedia: c.SaveUploadedFile") {
		return
	}

	ctx.JSON(http.StatusOK, models.MediaResponse{
		ErrorCode:    ErrorSuccessCode,
		ErrorMessage: "",
		Body: models.UploadPhotoRes{
			URL: h.cfg.BaseUrl + "media/" + file.File.Filename,
		},
	})
}
