package service

import (
	"github.com/PowerReport/pfs/src/app/usecases/mixin"
	ds "github.com/PowerReport/pfs/src/domain/directory/service"
	fs "github.com/PowerReport/pfs/src/domain/file/service"
	"github.com/PowerReport/pfs/src/util/identity"
	"github.com/jinzhu/copier"
)

type IMixinService interface {
	GetInfo(identity.OpId) (mixin.GetInfoCaseRes, error)
	GetDirectories(identity.OpId, string, int64, int64) mixin.GetDirectoriesCaseRes
	GetFiles(identity.OpId, string, int64, int64) mixin.GetFilesCaseRes
	Create(mixin.CreateCaseReq) mixin.CreateCaseRes
	Rename(mixin.RenameCaseReq) mixin.RenameCaseRes
	Move(mixin.MoveCaseReq) mixin.MoveCaseRes
	Delete(mixin.DeleteCaseReq) mixin.DeleteCaseRes
}

type MixinService struct {
	fileService      fs.IFileService
	directoryService ds.IDirectoryService
}

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
