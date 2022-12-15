package service

import "github.com/PowerReport/pfs/src/domain/file/model"

type IFileService interface {
	GetInfo(uint64) model.File
	GetRootFiles() []model.File
	GetByDirectoryId(uint64, string, int64, int64) []model.File
	Create(string, uint64) model.File
	Rename(uint64, string) model.File
	Move(uint64, uint64) model.File
	Delete(uint64) model.File
}

type FileService struct {
}

func (svc *FileService) GetInfo(id uint64) model.File {
	file := model.File{}
	return file
}

func (svc *FileService) GetRootFiles() []model.File {
	files := []model.File{}
	return files
}

func (svc *FileService) GetByDirectoryId(
	directoryId uint64, search string, page uint64, pageSize uint64) []model.File {
	files := []model.File{}
	return files
}

func (svc *FileService) Create(name string, directoryId uint64) model.File {
	file := model.File{}
	return file
}

func (svc *FileService) Rename(id uint64, name string) model.File {
	file := model.File{}
	return file
}

func (svc *FileService) Move(id uint64, directoryId uint64) model.File {
	file := model.File{}
	return file
}

func (svc *FileService) Delete(id uint64) model.File {
	file := model.File{}
	return file
}
