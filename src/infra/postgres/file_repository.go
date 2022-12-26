package postgres

import (
	"github.com/PowerReport/pfs/src/domain/file/model"
	"github.com/PowerReport/pfs/src/util/user"
	"gorm.io/gorm"
)

type FileRepository struct {
	db *gorm.DB
}

func NewFileRepository(db *gorm.DB) *FileRepository {
	return &FileRepository{db: db}
}

func (repo *FileRepository) UserViewed() (*gorm.DB, error) {
	current, err := user.Current()
	if err != nil {
		return nil, err
	}

	db := repo.db.Where(&model.File{OwnerId: current.Id, DataState: 0})
	return db, nil
}

func (repo *FileRepository) Get(id uint64) (model.File, error) {
	db, err := repo.UserViewed()
	if err != nil {
		return model.File{}, err
	}

	var file model.File
	err = db.First(&file, id).Error
	return file, err
}

func (repo *FileRepository) GetAll() ([]model.File, error) {
	db, err := repo.UserViewed()
	if err != nil {
		return []model.File{}, err
	}

	var files []model.File
	err = db.Find(&files).Error
	return files, err
}

func (repo *FileRepository) Save(file model.File) (model.File, error) {
	err := repo.db.Create(&file).Error
	if err != nil {
		return model.File{}, err
	}

	return file, nil
}

func (repo *FileRepository) Update(file model.File) (model.File, error) {
	err := repo.db.Save(&file).Error
	if err != nil {
		return model.File{}, err
	}

	return file, nil
}

func (repo *FileRepository) Delete(id uint64) (model.File, error) {
	db, err := repo.UserViewed()
	if err != nil {
		return model.File{}, err
	}

	var file model.File
	err = db.First(&file, id).Error
	if err != nil {
		return model.File{}, err
	}

	err = db.Delete(&file).Error
	if err != nil {
		return model.File{}, err
	}

	return file, nil
}

func (repo *FileRepository) SoftDelete(id uint64) (model.File, error) {
	db, err := repo.UserViewed()
	if err != nil {
		return model.File{}, err
	}

	var file model.File
	err = db.First(&file, id).Error
	if err != nil {
		return model.File{}, err
	}

	file.DataState = 1
	err = db.Save(&file).Error
	if err != nil {
		return model.File{}, err
	}

	return file, nil
}
