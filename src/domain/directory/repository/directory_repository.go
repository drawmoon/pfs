package repository

import "github.com/PowerReport/pfs/src/domain/directory/model"

type IDirectoryRepository interface {
	Get(uint64) (model.Directory, error)
	GetAll() ([]model.Directory, error)
	Save(model.Directory) (model.Directory, error)
	Update(model.Directory) (model.Directory, error)
	Delete(uint64) (model.Directory, error)
}
