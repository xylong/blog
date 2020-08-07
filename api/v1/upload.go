package v1

import (
	"blog/internal"
	"blog/pkg/file"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/http"
)

var Upload = &UploadController{}

type UploadController struct {
}

// @Summary 上传
// @Description 上传
// @Tags 上传
// @Produce  json
// @Param body body dto.RegisterInput true "body"
// @Success 200 {json} json "{"code":200,"data":null,"msg":"ok"}"
// @Router /api/v1/register [post]
func (u UploadController) Upload(c *gin.Context) {
	var (
		//file  multipart.File
		image *multipart.FileHeader
		err   error
	)

	if image, err = c.FormFile("image"); err != nil {
		internal.PanicCode(internal.UploadInvalidFile)
	}
	if file.CheckImageExt(file.GetExt(image.Filename)) {
		internal.PanicCode(internal.UploadFormat)
	}
	fileName := file.GenerateImageName(image.Filename)
	if err = c.SaveUploadedFile(image, fileName); err != nil {
		internal.PanicCode(internal.UploadSave)
	}

	//if file, image, err = c.Request.FormFile("image"); err != nil {
	//	internal.PanicCode(internal.UploadInvalidFile)
	//} else if image == nil {
	//	internal.PanicCode(internal.InvalidParams)
	//}

	//imageName := file2.GetImageName(image.Filename)
	//if file2.CheckImageExt(imageName) {
	//	internal.PanicCode(internal.UploadFormat)
	//}
	//if file2.CheckImageSize(file) {
	//	internal.PanicCode(internal.UploadSize)
	//}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "ok",
		"data": fileName,
	})

	//fullPath := file2.GetImageFullPath()
	//savePath := file2.GetImagePath()
	//src := fullPath + imageName

}
