package service

import "github.com/PowerReport/pfs/src/domain/directory/model"

type IDirectoryService interface {
	GetInfo(uint64) model.Directory
	GetRootDirectories() []model.Directory
	GetChildren(uint64, string, int64, int64) []model.Directory
	Create(string, uint64) model.Directory
	Rename(uint64, string) model.Directory
	Move(uint64, uint64) model.Directory
	Delete(uint64) model.Directory
}

type DirectoryService struct {
}

func (svc *DirectoryService) GetInfo(id uint64) model.Directory {
	directory := model.Directory{}
	return directory
}

func (svc *DirectoryService) GetRootDirectories() []model.Directory {
	directories := []model.Directory{}
	return directories
}

func (svc *DirectoryService) GetChildren(
	id uint64, search string, page int64, pageSize int64) []model.Directory {
	directories := []model.Directory{}
	return directories
}

func (svc *DirectoryService) Create(name string, directoryId uint64) model.Directory {
	directory := model.Directory{}
	return directory
}

func (svc *DirectoryService) Rename(id uint64, name string) model.Directory {
	directory := model.Directory{}
	return directory
}

func (svc *DirectoryService) Move(id uint64, directoryId uint64) model.Directory {
	directory := model.Directory{}
	return directory
}

func (svc *DirectoryService) Delete(id uint64) model.Directory {
	directory := model.Directory{}
	return directory
}
