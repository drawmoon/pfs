package controller

import (
	"net/http"
	"strconv"

	"github.com/PowerReport/pfs/src/app/service"
	"github.com/PowerReport/pfs/src/util/identity"
	"github.com/PowerReport/pfs/src/util/woosh"
	"github.com/gin-gonic/gin"
)

type MixinController struct {
	mixinService service.IMixinService
}

func NewMixinController() *MixinController {
	return &MixinController{}
}

// @Summary		获取文件或目录的详细信息
// @Tags		fs
// @Accept		json
// @Produce		json
// @Param		id		path		int		true		"文件id"
// @Success		200		{object}		mixin.GetInfoCaseRes
// @Failure		400		{object}		woosh.ProblemDetails
// @Failure		404		{object}		woosh.ProblemDetails
// @Failure		500		{object}		woosh.ProblemDetails
// @Router		/api/fs/{id}		[get]
func (c *MixinController) GetInfo(ctx *gin.Context) {
	id, err := identity.ParseOpId(ctx.Param("id"))
	if err != nil {
		woosh.HandleException("操作id错误", ctx)
		return
	}

	ctx.JSON(http.StatusOK, c.mixinService.GetInfo(id))
}

// @Summary		获取指定目录下一层级的目录列表
// @Tags		fs
// @Accept		json
// @Produce		json
// @Param		id		path		int		true		"目录id"
// @Success		200		{object}		mixin.GetDirectoriesCaseRes
// @Failure		400		{object}		woosh.ProblemDetails
// @Failure		404		{object}		woosh.ProblemDetails
// @Failure		500		{object}		woosh.ProblemDetails
// @Router		/api/fs/{id}/directories		[get]
func (c *MixinController) GetDirectories(ctx *gin.Context) {
	id, err := identity.ParseOpId(ctx.Param("id"))
	if err != nil {
		woosh.HandleException("操作id错误", ctx)
		return
	}

	search := ctx.Query("search")
	page, _ := strconv.ParseInt(ctx.DefaultQuery("page", "1"), 10, 64)
	pageSize, _ := strconv.ParseInt(ctx.DefaultQuery("pageSize", "10"), 10, 64)

	ctx.JSON(http.StatusOK, c.mixinService.GetDirectories(id, search, page, pageSize))
}

// @Summary		获取指定目录下一层级的文件列表
// @Tags		fs
// @Accept		json
// @Produce		json
// @Param		id		path		int		true		"目录id"
// @Success		200		{object}		mixin.GetFilesCaseRes
// @Failure		400		{object}		woosh.ProblemDetails
// @Failure		404		{object}		woosh.ProblemDetails
// @Failure		500		{object}		woosh.ProblemDetails
// @Router		/api/fs/{id}/files		[get]
func (c *MixinController) GetFiles(ctx *gin.Context) {
	id, err := identity.ParseOpId(ctx.Param("id"))
	if err != nil {
		woosh.HandleException("操作id错误", ctx)
		return
	}

	search := ctx.Query("search")
	page, _ := strconv.ParseInt(ctx.DefaultQuery("page", "1"), 10, 64)
	pageSize, _ := strconv.ParseInt(ctx.DefaultQuery("pageSize", "10"), 10, 64)

	ctx.JSON(http.StatusOK, c.mixinService.GetFiles(id, search, page, pageSize))
}
