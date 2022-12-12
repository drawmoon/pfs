package controller

import (
	"net/http"
	"strconv"

	"github.com/PowerReport/pfs/src/app/service"
	httputil "github.com/PowerReport/pfs/src/util/http"
	"github.com/gin-gonic/gin"
)

type FileController struct {
	fileService service.IFileService
}

func NewFileController() *FileController {
	return &FileController{}
}

// @Summary		获取文件详细信息
// @Tags		files
// @Accept		json
// @Produce		json
// @Param		id		path		int		true		"文件id"
// @Success		200		{object}		file.GetInfoResponse
// @Failure		400		{object}		http.ProblemDetails
// @Failure		404		{object}		http.ProblemDetails
// @Failure		500		{object}		http.ProblemDetails
// @Router		/api/fs/files/{id}		[get]
func (c *FileController) GetInfo(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		httputil.HandleException("文件id错误", ctx)
		return
	}

	ctx.JSON(http.StatusOK, c.fileService.GetInfo(id))
}
