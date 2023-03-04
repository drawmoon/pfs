package postgres

import (
	"github.com/PowerReport/pfs/src/domain/file/model"
	"gorm.io/gorm"
)

type ManifestRepository struct {
	db *gorm.DB
}

func NewManifestRepository(db *gorm.DB) *ManifestRepository {
	return &ManifestRepository{db: db}
}

func (repo *ManifestRepository) Get(id uint64) (model.Manifest, error) {
	var manifest model.Manifest
	err := repo.db.First(&manifest, id).Error
	if err != nil {
		return model.Manifest{}, err
	}

	return manifest, nil
}

func (repo *ManifestRepository) GetByFileId(fileId uint64) (model.Manifest, error) {
	var manifest model.Manifest
	err := repo.db.Where(&model.Manifest{FileId: fileId}).First(&manifest).Error
	if err != nil {
		return model.Manifest{}, err
	}

	return manifest, nil
}

func (repo *ManifestRepository) Save(manifest model.Manifest) (model.Manifest, error) {
	err := repo.db.Create(manifest).Error
	if err != nil {
		return model.Manifest{}, err
	}

	return manifest, nil
}

func (repo *ManifestRepository) Update(manifest model.Manifest) (model.Manifest, error) {
	err := repo.db.Save(manifest).Error
	if err != nil {
		return model.Manifest{}, err
	}

	return manifest, nil
}

func (repo *ManifestRepository) Delete(manifest model.Manifest) (model.Manifest, error) {
	manifest.DataState = 1
	err := repo.db.Save(manifest).Error
	if err != nil {
		return model.Manifest{}, err
	}

	return manifest, nil
}
