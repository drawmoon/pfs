package postgres

import (
	"github.com/PowerReport/pfs/src/domain/directory/model"
	"github.com/PowerReport/pfs/src/util/user"
	"gorm.io/gorm"
)

type DirectoryRepository struct {
	db *gorm.DB
}

func NewDirectoryRepository(db *gorm.DB) *DirectoryRepository {
	return &DirectoryRepository{db: db}
}

func (repo *DirectoryRepository) UserViewed() (*gorm.DB, error) {
	current, err := user.Current()
	if err != nil {
		return nil, err
	}

	db := repo.db.Where(&model.Directory{OwnerId: current.Id, DataState: 0})
	return db, nil
}

func (repo *DirectoryRepository) Get(id uint64) (model.Directory, error) {
	db, err := repo.UserViewed()
	if err != nil {
		return model.Directory{}, err
	}

	var directory model.Directory
	err = db.First(&directory, id).Error
	return directory, err
}

func (repo *DirectoryRepository) GetAll() ([]model.Directory, error) {
	db, err := repo.UserViewed()
	if err != nil {
		return []model.Directory{}, err
	}

	var directories []model.Directory
	err = db.Find(&directories).Error
	return directories, err
}

func (repo *DirectoryRepository) Save(directory model.Directory) (model.Directory, error) {
	err := repo.db.Create(&directory).Error
	if err != nil {
		return model.Directory{}, err
	}

	return directory, nil
}

func (repo *DirectoryRepository) Update(directory model.Directory) (model.Directory, error) {
	err := repo.db.Save(&directory).Error
	if err != nil {
		return model.Directory{}, err
	}

	return directory, nil
}

func (repo *DirectoryRepository) Delete(directory model.Directory) (model.Directory, error) {
	directory.DataState = 1
	err := repo.db.Save(&directory).Error
	if err != nil {
		return model.Directory{}, err
	}

	return directory, nil
}
