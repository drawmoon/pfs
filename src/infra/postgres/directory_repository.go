package postgres

import (
	"github.com/PowerReport/pfs/src/domain/directory/model"
	"gorm.io/gorm"
)

type DirectoryRepository struct {
	db *gorm.DB
}

func NewDirectoryRepository(db *gorm.DB) *DirectoryRepository {
	return &DirectoryRepository{db: db}
}

func (repo *DirectoryRepository) Get(id uint64) (model.Directory, error) {
	var directory model.Directory
	err := repo.db.First(&directory, id).Error
	return directory, err
}
