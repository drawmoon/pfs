package controller

import (
	"net/http"
	"strconv"

	"github.com/PowerReport/pfs/src/app/service"
	httputil "github.com/PowerReport/pfs/src/util/http"
	"github.com/gin-gonic/gin"
)

type DirectoryController struct {
	directoryService service.IDirectoryService
}

func NewDirectoryController() *DirectoryController {
	return &DirectoryController{}
}

// @Summary		获取目录详细信息
// @Tags		directories
// @Accept		json
// @Produce		json
// @Param		id		path		int		true		"目录id"
// @Success		200		{object}		directory.GetInfoResponse
// @Failure		400		{object}		http.ProblemDetails
// @Failure		404		{object}		http.ProblemDetails
// @Failure		500		{object}		http.ProblemDetails
// @Router		/api/fs/directories/{id}		[get]
func (c *DirectoryController) GetInfo(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		httputil.HandleException("目录id错误", ctx)
		return
	}

	ctx.JSON(http.StatusOK, c.directoryService.GetInfo(id))
}

// @Summary		获取指定目录下一层级的目录列表
// @Tags		directories
// @Accept		json
// @Produce		json
// @Param		id		path		int		true		"目录id"
// @Success		200		{object}		directory.GetDirectoriesResponse
// @Failure		400		{object}		http.ProblemDetails
// @Failure		404		{object}		http.ProblemDetails
// @Failure		500		{object}		http.ProblemDetails
// @Router		/api/fs/directories/{id}/directories		[get]
func (c *DirectoryController) GetDirectories(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		httputil.HandleException("目录id错误", ctx)
		return
	}

	search := ctx.Query("search")
	page, _ := strconv.ParseInt(ctx.DefaultQuery("page", "1"), 10, 64)
	pageSize, _ := strconv.ParseInt(ctx.DefaultQuery("pageSize", "10"), 10, 64)

	ctx.JSON(http.StatusOK, c.directoryService.GetDirectories(id, search, page, pageSize))
}

// @Summary		获取指定目录下一层级的文件列表
// @Tags		directories
// @Accept		json
// @Produce		json
// @Param		id		path		int		true		"目录id"
// @Success		200		{object}		directory.GetFilesResponse
// @Failure		400		{object}		http.ProblemDetails
// @Failure		404		{object}		http.ProblemDetails
// @Failure		500		{object}		http.ProblemDetails
// @Router		/api/fs/directories/{id}/files		[get]
func (c *DirectoryController) GetFiles(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		httputil.HandleException("目录id错误", ctx)
		return
	}

	search := ctx.Query("search")
	page, _ := strconv.ParseInt(ctx.DefaultQuery("page", "1"), 10, 64)
	pageSize, _ := strconv.ParseInt(ctx.DefaultQuery("pageSize", "10"), 10, 64)

	ctx.JSON(http.StatusOK, c.directoryService.GetFiles(id, search, page, pageSize))
}
