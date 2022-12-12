package service

import "github.com/PowerReport/pfs/src/app/usecases/directory"

type IDirectoryService interface {
	GetInfo(int64) directory.GetInfoResponse
	GetDirectories(int64, string, int64, int64) directory.GetDirectoriesResponse
	GetFiles(int64, string, int64, int64) directory.GetFilesResponse
}

type DirectoryService struct {
}

func (svc *DirectoryService) GetInfo(id int64) directory.GetInfoResponse {
	res := directory.GetInfoResponse{}
	return res
}

func (svc *DirectoryService) GetDirectories(
	id int64, search string, page int64, pageSize int64) directory.GetDirectoriesResponse {
	res := directory.GetDirectoriesResponse{}
	return res
}

func (svc *DirectoryService) GetFiles(
	id int64, search string, page int64, pageSize int64) directory.GetFilesResponse {
	res := directory.GetFilesResponse{}
	return res
}
