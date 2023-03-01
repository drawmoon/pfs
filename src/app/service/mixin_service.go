package service

import (
	"github.com/PowerReport/pfs/src/app/usecases/mixin"
	ds "github.com/PowerReport/pfs/src/domain/directory/service"
	fs "github.com/PowerReport/pfs/src/domain/file/service"
	"github.com/PowerReport/pfs/src/util/identity"
	"github.com/jinzhu/copier"
)

// 混合服务（不区分目录或文件的管理服务）
type IMixinService interface {
	// 获取目录或文件的详细信息
	GetInfo(id identity.OpId) (mixin.GetInfoCaseRes, error)
	// 获取指定目录下的目录列表
	GetDirectories(id identity.OpId, search string, page int64, pageSize int64) mixin.GetDirectoriesCaseRes
	// 获取指定目录下的文件列表
	GetFiles(id identity.OpId, search string, page int64, pageSize int64) mixin.GetFilesCaseRes
	// 创建目录或文件
	Create(req mixin.CreateCaseReq) mixin.CreateCaseRes
	// 重命名目录或文件
	Rename(req mixin.RenameCaseReq) mixin.RenameCaseRes
	// 移动目录或文件
	Move(req mixin.MoveCaseReq) mixin.MoveCaseRes
	// 删除目录或文件
	Delete(req mixin.DeleteCaseReq) mixin.DeleteCaseRes
}

type MixinService struct {
	fileService      fs.IFileService
	directoryService ds.IDirectoryService
}

func NewMixinService(
	fileService fs.IFileService, directoryService ds.IDirectoryService) *MixinService {
	return &MixinService{
		fileService:      fileService,
		directoryService: directoryService,
	}
}

var _ IMixinService = &MixinService{}

func (svc *MixinService) GetInfo(id identity.OpId) (mixin.GetInfoCaseRes, error) {
	if id.IsDirectory {
		directory, err := svc.directoryService.GetInfo(id.Real)
		if err != nil {
			return mixin.GetInfoCaseRes{}, err
		}

		res := mixin.GetInfoCaseRes{IsDirectory: true}
		copier.Copy(&res, &directory)
		return res, nil
	}

	file, err := svc.fileService.GetInfo(id.Real)
	if err != nil {
		return mixin.GetInfoCaseRes{}, err
	}

	res := mixin.GetInfoCaseRes{IsDirectory: false}
	copier.Copy(&res, &file)
	return res, nil
}

func (svc *MixinService) GetDirectories(
	id identity.OpId, search string, page int64, pageSize int64) mixin.GetDirectoriesCaseRes {
	res := mixin.GetDirectoriesCaseRes{}
	return res
}

func (svc *MixinService) GetFiles(
	id identity.OpId, search string, page int64, pageSize int64) mixin.GetFilesCaseRes {
	res := mixin.GetFilesCaseRes{}
	return res
}

func (svc *MixinService) Create(req mixin.CreateCaseReq) mixin.CreateCaseRes {
	res := mixin.CreateCaseRes{}
	return res
}

func (svc *MixinService) Rename(req mixin.RenameCaseReq) mixin.RenameCaseRes {
	res := mixin.RenameCaseRes{}
	return res
}

func (svc *MixinService) Move(req mixin.MoveCaseReq) mixin.MoveCaseRes {
	res := mixin.MoveCaseRes{}
	return res
}

func (svc *MixinService) Delete(req mixin.DeleteCaseReq) mixin.DeleteCaseRes {
	res := mixin.DeleteCaseRes{}
	return res
}
