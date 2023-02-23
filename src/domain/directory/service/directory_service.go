package service

import (
	"github.com/PowerReport/pfs/src/domain/directory/model"
	"github.com/PowerReport/pfs/src/domain/directory/repository"
	"github.com/PowerReport/pfs/src/util/user"
	"gorm.io/gorm"
)

type IDirectoryService interface {
	GetInfo(uint64) (model.Directory, error)
	GetRootDirectories() ([]model.Directory, error)
	GetChildren(uint64, string, int64, int64) ([]model.Directory, error)
	Create(string, uint64) (model.Directory, error)
	Rename(uint64, string) (model.Directory, error)
	Move(uint64, uint64) (model.Directory, error)
	Delete(uint64) (model.Directory, error)
}

type DirectoryService struct {
	directoryRepository repository.IDirectoryRepository
}

func NewDirectoryService(directoryRepository repository.IDirectoryRepository) *DirectoryService {
	return &DirectoryService{directoryRepository: directoryRepository}
}

func (svc *DirectoryService) GetInfo(id uint64) (model.Directory, error) {
	return svc.directoryRepository.Get(id)
}

func (svc *DirectoryService) GetRootDirectories() ([]model.Directory, error) {
	db, err := svc.directoryRepository.UserViewed()
	if err != nil {
		return []model.Directory{}, err
	}

	directories := []model.Directory{}
	err = db.Where("DirectoryId = ?", gorm.Expr("NULL")).Find(&directories).Error
	if err != nil {
		return []model.Directory{}, err
	}

	return directories, nil
}

func (svc *DirectoryService) GetChildren(
	id uint64, search string, page int64, pageSize int64) ([]model.Directory, error) {
	db, err := svc.directoryRepository.UserViewed()
	if err != nil {
		return []model.Directory{}, err
	}

	directories := []model.Directory{}
	err = db.Where(&model.Directory{DirectoryId: id}).Find(&directories).Error
	if err != nil {
		return []model.Directory{}, err
	}

	return directories, nil
}

func (svc *DirectoryService) Create(name string, directoryId uint64) (model.Directory, error) {
	user, err := user.Current()
	if err != nil {
		return model.Directory{}, err
	}

	directory := model.Directory{
		Name:        name,
		DirectoryId: directoryId,
		OwnerId:     user.Id,
		DataState:   0,
	}

	directory, err = svc.directoryRepository.Save(directory)
	if err != nil {
		return model.Directory{}, err
	}

	return directory, nil
}

func (svc *DirectoryService) Rename(id uint64, name string) (model.Directory, error) {
	directory, err := svc.directoryRepository.Get(id)
	if err != nil {
		return model.Directory{}, err
	}

	directory.Name = name
	directory, err = svc.directoryRepository.Update(directory)
	if err != nil {
		return model.Directory{}, err
	}

	return directory, nil
}

func (svc *DirectoryService) Move(id uint64, directoryId uint64) (model.Directory, error) {
	directory, err := svc.directoryRepository.Get(id)
	if err != nil {
		return model.Directory{}, err
	}

	directory.DirectoryId = directoryId
	directory, err = svc.directoryRepository.Update(directory)
	if err != nil {
		return model.Directory{}, err
	}

	return directory, nil
}

func (svc *DirectoryService) Delete(id uint64) (model.Directory, error) {
	directory, err := svc.directoryRepository.Get(id)
	if err != nil {
		return model.Directory{}, err
	}

	directory, err = svc.directoryRepository.Delete(directory)
	if err != nil {
		return model.Directory{}, err
	}

	return directory, nil
}
