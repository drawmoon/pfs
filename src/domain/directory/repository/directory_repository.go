package repository

import (
	"github.com/PowerReport/pfs/src/domain/directory/model"
	"gorm.io/gorm"
)

// 文件管理仓储服务
type IDirectoryRepository interface {
	// 获取用户可见的目录列表
	UserViewed() (*gorm.DB, error)
	// 获取用户可见的目录信息
	Get(id uint64) (model.Directory, error)
	// 获取用户可见的目录列表
	GetAll() ([]model.Directory, error)
	// 保存一条目录记录
	Save(directory model.Directory) (model.Directory, error)
	// 更新一条目录记录
	Update(directory model.Directory) (model.Directory, error)
	// 删除一条目录记录
	Delete(directory model.Directory) (model.Directory, error)
}
