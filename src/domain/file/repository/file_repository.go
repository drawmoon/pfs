package repository

import "github.com/PowerReport/pfs/src/domain/file/model"

type IFileRepository interface {
	Get(uint64) (model.File, error)
	GetAll() ([]model.File, error)
	Save(model.File) (model.File, error)
	Update(model.File) (model.File, error)
	Delete(uint64) (model.File, error)
}
