package service

import (
	"github.com/PowerReport/pfs/src/domain/file/model"
	"github.com/PowerReport/pfs/src/domain/file/repository"
	"github.com/PowerReport/pfs/src/domain/file/validator"
	"github.com/PowerReport/pfs/src/util/dbaccessor"
	"github.com/PowerReport/pfs/src/util/user"
	"gorm.io/gorm"
)

// 文件管理服务
type IFileService interface {
	// 获取文件的详细信息
	GetInfo(id uint64) (model.File, error)
	// 获取所有根文件列表
	GetRootFiles() ([]model.File, error)
	// 获取指定目录下的文件列表
	GetByDirectoryId(directoryId uint64, search string, page int64, pageSize int64) ([]model.File, error)
	// 创建新的文件
	Create(name string, directoryId uint64) (model.File, error)
	// 重命名文件
	Rename(id uint64, name string) (model.File, error)
	// 移动文件
	Move(id uint64, directoryId uint64) (model.File, error)
	// 删除文件
	Delete(id uint64) (model.File, error)
}

type FileService struct {
	fileRepository repository.IFileRepository
}

func NewFileService(fileRepository repository.IFileRepository) *FileService {
	return &FileService{fileRepository: fileRepository}
}

func (svc *FileService) GetInfo(id uint64) (model.File, error) {
	return svc.fileRepository.Get(id)
}

func (svc *FileService) GetRootFiles() ([]model.File, error) {
	db, err := svc.fileRepository.UserViewed()
	if err != nil {
		return []model.File{}, err
	}

	files := []model.File{}
	err = db.Where("DirectoryId = ?", gorm.Expr("NULL")).Find(&files).Error
	if err != nil {
		return []model.File{}, err
	}

	return files, nil
}

func (svc *FileService) GetByDirectoryId(
	directoryId uint64, search string, page int64, pageSize int64) ([]model.File, error) {
	db, err := svc.fileRepository.UserViewed()
	if err != nil {
		return []model.File{}, err
	}

	// 设置分页
	db.Scopes(dbaccessor.Paginate(int(page), int(pageSize)))

	files := []model.File{}
	err = db.Where("DirectoryId = ? AND Name LIKE %?%", directoryId, search).Find(&files).Error
	if err != nil {
		return []model.File{}, err
	}

	return files, nil
}

func (svc *FileService) Create(name string, directoryId uint64) (model.File, error) {
	user, err := user.Current()
	if err != nil {
		return model.File{}, err
	}

	file := model.File{
		Name:        name,
		DirectoryId: directoryId,
		OwnerId:     user.Id,
		DataState:   0,
	}

	// 校验文件添加是否合法
	validator.ValidateForAdding(&file)

	file, err = svc.fileRepository.Save(file)
	if err != nil {
		return model.File{}, err
	}

	return file, nil
}

func (svc *FileService) Rename(id uint64, name string) (model.File, error) {
	file, err := svc.fileRepository.Get(id)
	if err != nil {
		return model.File{}, err
	}

	file.Name = name
	// 校验文件名称是否合法
	validator.ValidateName(&file)

	file, err = svc.fileRepository.Update(file)
	if err != nil {
		return model.File{}, err
	}

	return file, nil
}

func (svc *FileService) Move(id uint64, directoryId uint64) (model.File, error) {
	file, err := svc.fileRepository.Get(id)
	if err != nil {
		return model.File{}, err
	}

	file.DirectoryId = directoryId
	// 校验文件移动是否合法
	validator.ValidateForMove(&file)

	file, err = svc.fileRepository.Update(file)
	if err != nil {
		return model.File{}, err
	}

	return file, nil
}

func (svc *FileService) Delete(id uint64) (model.File, error) {
	file, err := svc.fileRepository.Get(id)
	if err != nil {
		return model.File{}, err
	}

	// 校验当前用户是否有删除文件的权限
	validator.ValidatePermissionForDeleting(&file)

	file, err = svc.fileRepository.Delete(file)
	if err != nil {
		return model.File{}, err
	}

	return file, nil
}
