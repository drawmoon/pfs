package service

import (
	"errors"

	"github.com/PowerReport/pfs/src/app/usecases/mixin"
	ds "github.com/PowerReport/pfs/src/domain/directory/service"
	fs "github.com/PowerReport/pfs/src/domain/file/service"
	"github.com/PowerReport/pfs/src/util/id"
	"github.com/jinzhu/copier"
)

// 混合服务（不区分目录或文件的管理服务）
type IMixinService interface {
	// 获取目录或文件的详细信息
	GetInfo(id id.OpId) (mixin.GetInfoCaseRes, error)
	// 获取指定目录下的目录列表
	GetDirectories(id id.OpId, search string, page int64, pageSize int64) (mixin.GetDirectoriesCaseRes, error)
	// 获取指定目录下的文件列表
	GetFiles(id id.OpId, search string, page int64, pageSize int64) (mixin.GetFilesCaseRes, error)
	// 创建目录或文件
	Create(req mixin.CreateCaseReq) (mixin.CreateCaseRes, error)
	// 重命名目录或文件
	Rename(req mixin.RenameCaseReq) (mixin.RenameCaseRes, error)
	// 移动目录或文件
	Move(req mixin.MoveCaseReq) (mixin.MoveCaseRes, error)
	// 删除目录或文件
	Delete(req mixin.DeleteCaseReq) (mixin.DeleteCaseRes, error)
}

type MixinService struct {
	fileService      fs.IFileService
	manifestService  fs.IManifestService
	directoryService ds.IDirectoryService
}

func NewMixinService(
	fileService fs.IFileService, manifestService fs.IManifestService, directoryService ds.IDirectoryService) *MixinService {
	return &MixinService{
		fileService:      fileService,
		manifestService:  manifestService,
		directoryService: directoryService,
	}
}

func (svc *MixinService) GetInfo(id id.OpId) (mixin.GetInfoCaseRes, error) {
	// 如果是获取目录信息
	if id.IsDir {
		directory, err := svc.directoryService.GetInfo(id.Real)
		if err != nil {
			return mixin.GetInfoCaseRes{}, err
		}

		res := mixin.GetInfoCaseRes{IsDir: true}
		copier.Copy(&res, &directory)
		return res, nil
	}

	file, err := svc.fileService.GetInfo(id.Real)
	if err != nil {
		return mixin.GetInfoCaseRes{}, err
	}

	res := mixin.GetInfoCaseRes{IsDir: false}
	copier.Copy(&res, &file)
	return res, nil
}

func (svc *MixinService) GetDirectories(
	id id.OpId, search string, page int64, pageSize int64) (mixin.GetDirectoriesCaseRes, error) {
	if !id.IsDir {
		return mixin.GetDirectoriesCaseRes{}, errors.New("无法获取子级目录列表，因为指定的不是目录")
	}

	directories, err := svc.directoryService.GetChildren(id.Real, search, page, pageSize)
	if err != nil {
		return mixin.GetDirectoriesCaseRes{}, err
	}

	dirRes := []mixin.GetDirectoryCaseRes{}
	copier.Copy(&dirRes, &directories)

	res := mixin.GetDirectoriesCaseRes{
		PageIndex:    page,
		PageSize:     pageSize,
		TotalCount:   0,
		TotalPages:   0,
		Items:        dirRes,
		HasPrevPages: true,
		HasNextPages: true,
	}
	return res, nil
}

func (svc *MixinService) GetFiles(
	id id.OpId, search string, page int64, pageSize int64) (mixin.GetFilesCaseRes, error) {
	if !id.IsDir {
		return mixin.GetFilesCaseRes{}, errors.New("无法获取子级文件列表，因为指定的不是目录")
	}

	files, err := svc.fileService.GetByDirectoryId(id.Real, search, page, pageSize)
	if err != nil {
		return mixin.GetFilesCaseRes{}, err
	}

	fileRes := []mixin.GetFileCaseRes{}
	copier.Copy(&fileRes, &files)

	res := mixin.GetFilesCaseRes{
		PageIndex:    page,
		PageSize:     pageSize,
		TotalCount:   0,
		TotalPages:   0,
		Items:        fileRes,
		HasPrevPages: true,
		HasNextPages: true,
	}
	return res, nil
}

func (svc *MixinService) Create(req mixin.CreateCaseReq) (mixin.CreateCaseRes, error) {
	// 如果是新增目录
	if req.Type == "dir" {
		return svc.handleCreateDirectory(req.Name, req.DirectoryId)
	}

	return svc.handleCreateFile(req.Name, req.DirectoryId)
}

func (svc *MixinService) handleCreateDirectory(name string, directoryId uint64) (mixin.CreateCaseRes, error) {
	directory, err := svc.directoryService.Create(name, directoryId)
	if err != nil {
		return mixin.CreateCaseRes{}, err
	}

	res := mixin.CreateCaseRes{IsDir: true}
	copier.Copy(&res, &directory)
	return res, nil
}

func (svc *MixinService) handleCreateFile(name string, directoryId uint64) (mixin.CreateCaseRes, error) {
	// 新建文件
	file, err := svc.fileService.Create(name, directoryId)
	if err != nil {
		return mixin.CreateCaseRes{}, err
	}

	// 新建文件清单
	_, err = svc.manifestService.Create(file.Id, []byte{})
	if err != nil {
		return mixin.CreateCaseRes{}, err
	}

	res := mixin.CreateCaseRes{IsDir: false}
	copier.Copy(&res, &file)
	return res, nil
}

func (svc *MixinService) Rename(req mixin.RenameCaseReq) (mixin.RenameCaseRes, error) {
	// 如果是重命名目录
	if req.Id.IsDir {
		directory, err := svc.directoryService.Rename(req.Id.Real, req.Name)
		if err != nil {
			return mixin.RenameCaseRes{}, err
		}

		res := mixin.RenameCaseRes{IsDir: true}
		copier.Copy(&res, &directory)
		return res, nil
	}

	file, err := svc.fileService.Rename(req.Id.Real, req.Name)
	if err != nil {
		return mixin.RenameCaseRes{}, err
	}

	res := mixin.RenameCaseRes{IsDir: false}
	copier.Copy(&res, &file)
	return res, nil
}

func (svc *MixinService) Move(req mixin.MoveCaseReq) (mixin.MoveCaseRes, error) {
	// 如果是移动目录
	if req.Id.IsDir {
		directory, err := svc.directoryService.Move(req.Id.Real, req.DirectoryId)
		if err != nil {
			return mixin.MoveCaseRes{}, err
		}

		res := mixin.MoveCaseRes{IsDir: true}
		copier.Copy(&res, &directory)
		return res, nil
	}

	file, err := svc.fileService.Move(req.Id.Real, req.DirectoryId)
	if err != nil {
		return mixin.MoveCaseRes{}, err
	}

	res := mixin.MoveCaseRes{IsDir: false}
	copier.Copy(&res, &file)
	return res, nil
}

func (svc *MixinService) Delete(req mixin.DeleteCaseReq) (mixin.DeleteCaseRes, error) {
	// 如果是删除目录
	if req.Id.IsDir {
		directory, err := svc.directoryService.Delete(req.Id.Real)
		if err != nil {
			return mixin.DeleteCaseRes{}, err
		}

		res := mixin.DeleteCaseRes{IsDir: true}
		copier.Copy(&res, &directory)
		return res, nil
	}

	file, err := svc.fileService.Delete(req.Id.Real)
	if err != nil {
		return mixin.DeleteCaseRes{}, err
	}

	res := mixin.DeleteCaseRes{IsDir: false}
	copier.Copy(&res, &file)
	return res, nil
}
