package repository

import "github.com/PowerReport/pfs/src/domain/file/model"

// 文件清单仓储服务
type IManifestRepository interface {
	// 获取文件清单信息
	Get(id uint64) (model.Manifest, error)
	// 获取指定文件的文件清单信息
	GetByFileId(fileId uint64) (model.Manifest, error)
	// 保存一条文件清单记录
	Save(manifest model.Manifest) (model.Manifest, error)
	// 更新一条文件清单记录
	Update(manifest model.Manifest) (model.Manifest, error)
	// 删除一条文件清单记录
	Delete(manifest model.Manifest) (model.Manifest, error)
}
