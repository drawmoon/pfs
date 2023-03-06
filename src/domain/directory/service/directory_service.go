package service

import (
	"github.com/PowerReport/pfs/src/domain/directory/model"
	"github.com/PowerReport/pfs/src/domain/directory/repository"
	"github.com/PowerReport/pfs/src/domain/directory/validator"
	"github.com/PowerReport/pfs/src/util/pages"
	"github.com/PowerReport/pfs/src/util/user"
	"gorm.io/gorm"
)

// 目录管理服务
type IDirectoryService interface {
	// 获取目录详细信息
	GetInfo(id uint64) (model.Directory, error)
	// 获取所有根目录列表
	GetRootDirectories() ([]model.Directory, error)
	// 获取子级目录列表
	GetChildren(id uint64, search string, page int64, pageSize int64) ([]model.Directory, error)
	// 创建新的目录
	Create(name string, directoryId uint64) (model.Directory, error)
	// 重命名目录
	Rename(id uint64, name string) (model.Directory, error)
	// 移动目录
	Move(id uint64, directoryId uint64) (model.Directory, error)
	// 删除目录
	Delete(id uint64) (model.Directory, error)
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

	// 设置分页
	db.Scopes(pages.Paginate(int(page), int(pageSize)))

	directories := []model.Directory{}
	err = db.Where("DirectoryId = ? AND Name LIKE %?%", id, search).Find(&directories).Error
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

	// 校验目录添加是否合法
	validator.ValidateForAdding(&directory)

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
	// 校验目录名称是否合法
	validator.ValidateName(&directory)

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
	// 校验目录移动是否合法
	validator.ValidateForMove(&directory)

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

	// 校验目录删除是否合法
	validator.ValidatePermissionForDeleting(&directory)

	directory, err = svc.directoryRepository.Delete(directory)
	if err != nil {
		return model.Directory{}, err
	}

	return directory, nil
}
