package repository

import (
	"github.com/PowerReport/pfs/src/domain/file/model"
	"gorm.io/gorm"
)

// 文件管理仓储服务
type IFileRepository interface {
	// 获取当前的数据库上下文
	CurrentDbContext() *gorm.DB
	// 获取用户可见的文件列表
	UserViewed() (*gorm.DB, error)
	// 获取用户可见的文件信息
	Get(uint64) (model.File, error)
	// 获取用户可见的文件列表
	GetAll() ([]model.File, error)
	// 保存一条文件记录
	Save(model.File) (model.File, error)
	// 更新一条文件记录
	Update(model.File) (model.File, error)
	// 硬删除一条文件记录
	Delete(model.File) (model.File, error)
	// 软删除一条文件记录
	SoftDelete(model.File) (model.File, error)
}
