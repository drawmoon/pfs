package service

import "github.com/PowerReport/pfs/src/domain/directory/model"

type IDirectoryService interface {
	GetInfo(uint64) model.Directory
}

type DirectoryService struct {
}

func (svc *DirectoryService) GetInfo(id uint64) model.Directory {
	directory := model.Directory{}
	return directory
}
