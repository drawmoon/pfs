package model

// 文件数据库模型
type File struct {
	Id          uint64 `gorm:"column:id;primaryKey"`                   // 主键
	Name        string `gorm:"column:name;size:40;not null"`           // 文件名称
	Description string `gorm:"column:description;size:256"`            // 文件描述
	DirectoryId uint64 `gorm:"column:directory_id"`                    // 父级目录id
	OwnerId     string `gorm:"column:owner_id;size:36;not null"`       // 文件所属者
	DataState   uint64 `gorm:"column:data_state;precision:1;not null"` // 数据状态
}
