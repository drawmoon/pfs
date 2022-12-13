package service

import "github.com/PowerReport/pfs/src/domain/file/model"

type IFileService interface {
	GetInfo(uint64) model.File
}

type FileService struct {
}

func (svc *FileService) GetInfo(id uint64) model.File {
	file := model.File{}
	return file
}
