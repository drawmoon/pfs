package service

import (
	"github.com/PowerReport/pfs/src/app/usecases/mixin"
	ds "github.com/PowerReport/pfs/src/domain/directory/service"
	fs "github.com/PowerReport/pfs/src/domain/file/service"
	"github.com/PowerReport/pfs/src/util/identity"
)

type IMixinService interface {
	GetInfo(identity.OpId) mixin.GetInfoCaseRes
	GetDirectories(identity.OpId, string, int64, int64) mixin.GetDirectoriesCaseRes
	GetFiles(identity.OpId, string, int64, int64) mixin.GetFilesCaseRes
}

type MixinService struct {
	fileService      fs.IFileService
	directoryService ds.IDirectoryService
}

func (svc *MixinService) GetInfo(id identity.OpId) mixin.GetInfoCaseRes {
	if id.IsDirectory {
		directory := svc.directoryService.GetInfo(id.Real)
		return mixin.GetInfoCaseRes{
			Id:          directory.Id,
			Name:        directory.Name,
			IsDirectory: true,
		}
	}

	file := svc.fileService.GetInfo(id.Real)
	return mixin.GetInfoCaseRes{
		Id:          file.Id,
		Name:        file.Name,
		IsDirectory: false,
	}
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
