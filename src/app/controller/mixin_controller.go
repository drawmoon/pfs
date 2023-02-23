package controller

import (
	"net/http"
	"strconv"

	"github.com/PowerReport/pfs/src/app/service"
	"github.com/PowerReport/pfs/src/app/usecases/mixin"
	"github.com/PowerReport/pfs/src/util/identity"
	"github.com/PowerReport/pfs/src/util/woosh"
	"github.com/gin-gonic/gin"
)

type MixinController struct {
	mixinService service.IMixinService
}

func NewMixinController(mixinService service.IMixinService) *MixinController {
	return &MixinController{mixinService: mixinService}
}

// @Summary		获取文件或目录的详细信息
// @Tags		fs
// @Accept		json
// @Produce		json
// @Param		id		path		int		true		"操作id"
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

	info, err := c.mixinService.GetInfo(id)
	if err != nil {
		woosh.HandleException(err.Error(), ctx)
	}

	ctx.JSON(http.StatusOK, info)
}

// @Summary		获取指定目录下一层级的目录列表
// @Tags		fs
// @Accept		json
// @Produce		json
// @Param		id			path		int		true		"操作id"
// @Param		search		query		string	false		"搜索"
// @Param		page		query		int		false		"页码"
// @Param		pageSize	query		int		false		"页的大小"
// @Success		200			{object}	mixin.GetDirectoriesCaseRes
// @Failure		400			{object}	woosh.ProblemDetails
// @Failure		404			{object}	woosh.ProblemDetails
// @Failure		500			{object}	woosh.ProblemDetails
// @Router		/api/fs/{id}/directories		[get]
func (c *MixinController) GetDirectories(ctx *gin.Context) {
	id, err := identity.ParseOpId(ctx.Param("id"))
	if err != nil || !id.IsDirectory {
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
// @Param		id			path		int		true		"操作id"
// @Param		search		query		string	false		"搜索"
// @Param		page		query		int		false		"页码"
// @Param		pageSize	query		int		false		"页的大小"
// @Success		200			{object}	mixin.GetDirectoriesCaseRes
// @Failure		400			{object}	woosh.ProblemDetails
// @Failure		404			{object}	woosh.ProblemDetails
// @Failure		500			{object}	woosh.ProblemDetails
// @Router		/api/fs/{id}/files		[get]
func (c *MixinController) GetFiles(ctx *gin.Context) {
	id, err := identity.ParseOpId(ctx.Param("id"))
	if err != nil || !id.IsDirectory {
		woosh.HandleException("操作id错误", ctx)
		return
	}

	search := ctx.Query("search")
	page, _ := strconv.ParseInt(ctx.DefaultQuery("page", "1"), 10, 64)
	pageSize, _ := strconv.ParseInt(ctx.DefaultQuery("pageSize", "10"), 10, 64)

	ctx.JSON(http.StatusOK, c.mixinService.GetFiles(id, search, page, pageSize))
}

// @Summary		创建文件或目录
// @Tags		fs
// @Accept		json
// @Produce		json
// @Success		200			{object}	mixin.CreateCaseRes
// @Failure		400			{object}	woosh.ProblemDetails
// @Failure		404			{object}	woosh.ProblemDetails
// @Failure		500			{object}	woosh.ProblemDetails
// @Router		/api/fs		[post]
func (c *MixinController) Create(ctx *gin.Context) {
	req := mixin.CreateCaseReq{}

	ctx.JSON(http.StatusOK, c.mixinService.Create(req))
}

// @Summary		重命名文件或目录
// @Tags		fs
// @Accept		json
// @Produce		json
// @Success		200			{object}	mixin.RenameCaseRes
// @Failure		400			{object}	woosh.ProblemDetails
// @Failure		404			{object}	woosh.ProblemDetails
// @Failure		500			{object}	woosh.ProblemDetails
// @Router		/api/fs/rename		[put]
func (c *MixinController) Rename(ctx *gin.Context) {
	req := mixin.RenameCaseReq{}

	ctx.JSON(http.StatusOK, c.mixinService.Rename(req))
}

// @Summary		移动文件或目录
// @Tags		fs
// @Accept		json
// @Produce		json
// @Success		200			{object}	mixin.MoveCaseRes
// @Failure		400			{object}	woosh.ProblemDetails
// @Failure		404			{object}	woosh.ProblemDetails
// @Failure		500			{object}	woosh.ProblemDetails
// @Router		/api/fs/move		[put]
func (c *MixinController) Move(ctx *gin.Context) {
	req := mixin.MoveCaseReq{}

	ctx.JSON(http.StatusOK, c.mixinService.Move(req))
}

// @Summary		移除文件或目录
// @Tags		fs
// @Accept		json
// @Produce		json
// @Success		200			{object}	mixin.DeleteCaseRes
// @Failure		400			{object}	woosh.ProblemDetails
// @Failure		404			{object}	woosh.ProblemDetails
// @Failure		500			{object}	woosh.ProblemDetails
// @Router		/api/fs		[delete]
func (c *MixinController) Delete(ctx *gin.Context) {
	req := mixin.DeleteCaseReq{}

	ctx.JSON(http.StatusOK, c.mixinService.Delete(req))
}
