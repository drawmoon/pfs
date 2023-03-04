package model

// 文件清单数据库模型
type Manifest struct {
	Id        uint64 `gorm:"column:id;primaryKey"`                   // 主键
	FileId    uint64 `gorm:"column:file_id;not null"`                // 文件id
	Location  string `gorm:"column:location;size:256;not null"`      // 文件的物理路径
	DataState uint64 `gorm:"column:data_state;precision:1;not null"` // 数据状态
}
