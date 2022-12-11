package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type FileController struct {
}

// @Summary 获取文件详细信息
// @Tags files
// @Accept json
// @Produce json
// @Param id path int true "文件id"
// @Success 200 {object} object
// @Failure 400 {object} object
// @Failure 404 {object} object
// @Failure 500 {object} object
// @Router /api/fs/file/{id} [get]
func (file *FileController) GetInfo(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Method not implemented.")
}

func NewFileController() *FileController {
	return &FileController{}
}
