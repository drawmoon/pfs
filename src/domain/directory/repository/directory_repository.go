package repository

import (
	"github.com/PowerReport/pfs/src/domain/directory/model"
	"gorm.io/gorm"
)

type IDirectoryRepository interface {
	UserViewed() (*gorm.DB, error)
	Get(uint64) (model.Directory, error)
	GetAll() ([]model.Directory, error)
	Save(model.Directory) (model.Directory, error)
	Update(model.Directory) (model.Directory, error)
	Delete(model.Directory) (model.Directory, error)
	SoftDelete(model.Directory) (model.Directory, error)
}
