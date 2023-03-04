package model

// 目录数据库模型
type Directory struct {
	Id          uint64 `gorm:"column:id;primaryKey"`                   // 主键
	Name        string `gorm:"column:name;size:40;not null"`           // 目录名称
	DirectoryId uint64 `gorm:"column:directory_id"`                    // 父级目录id
	OwnerId     string `gorm:"column:owner_id;size:36;not null"`       // 目录的所属者
	DataState   uint64 `gorm:"column:data_state;precision:1;not null"` // 数据状态
}
