package service

import "github.com/PowerReport/pfs/src/app/usecases/directory"

type IDirectoryService interface {
	GetInfo(int64) directory.GetInfoCaseRes
	GetDirectories(int64, string, int64, int64) directory.GetDirectoriesCaseRes
	GetFiles(int64, string, int64, int64) directory.GetFilesCaseRes
}

type DirectoryService struct {
}

func (svc *DirectoryService) GetInfo(id int64) directory.GetInfoCaseRes {
	res := directory.GetInfoCaseRes{}
	return res
}

func (svc *DirectoryService) GetDirectories(
	id int64, search string, page int64, pageSize int64) directory.GetDirectoriesCaseRes {
	res := directory.GetDirectoriesCaseRes{}
	return res
}

func (svc *DirectoryService) GetFiles(
	id int64, search string, page int64, pageSize int64) directory.GetFilesCaseRes {
	res := directory.GetFilesCaseRes{}
	return res
}
