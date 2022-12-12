package service

import "github.com/PowerReport/pfs/src/app/usecases/file"

type IFileService interface {
	GetInfo(int64) file.GetInfoCaseRes
}

type FileService struct {
}

func (svc *FileService) GetInfo(id int64) file.GetInfoCaseRes {
	res := file.GetInfoCaseRes{}
	return res
}
