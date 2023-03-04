package service

import (
	"github.com/PowerReport/pfs/src/domain/file/model"
	"github.com/PowerReport/pfs/src/domain/file/repository"
	"github.com/PowerReport/pfs/src/infra/obs"
)

// 文件清单服务
type IManifestService interface {
	Create(fileId uint64, data []byte) (model.Manifest, error)
}

type ManifestService struct {
	fileRepository     repository.IFileRepository
	manifestRepository repository.IManifestRepository
}

func NewManifestService(
	manifestRepository repository.IManifestRepository, fileRepository repository.IFileRepository) *ManifestService {
	return &ManifestService{fileRepository: fileRepository, manifestRepository: manifestRepository}
}

func (svc *ManifestService) Create(fileId uint64, data []byte) (model.Manifest, error) {
	file, err := svc.fileRepository.Get(fileId)
	if err != nil {
		return model.Manifest{}, err
	}

	// 推送文件
	location, err := obs.Put(file.Name, data, "application/octet-stream")
	if err != nil {
		return model.Manifest{}, err
	}

	manifest := model.Manifest{
		FileId:    fileId,
		Location:  location,
		DataState: 0,
	}

	manifest, err = svc.manifestRepository.Save(manifest)
	if err != nil {
		return model.Manifest{}, err
	}

	return model.Manifest{}, nil
}
