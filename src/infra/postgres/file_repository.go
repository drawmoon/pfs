package postgres

import (
	"github.com/PowerReport/pfs/src/domain/file/model"
	"gorm.io/gorm"
)

type FileRepository struct {
	db *gorm.DB
}

func NewFileRepository(db *gorm.DB) *FileRepository {
	return &FileRepository{db: db}
}

func (repo *FileRepository) Get(id uint64) (model.File, error) {
	var file model.File
	err := repo.db.First(&file, id).Error
	return file, err
}
