package model

type File struct {
	Id          uint64 `gorm:"column:id;primaryKey"`
	Name        string `gorm:"column:name;size:40;not null"`
	DirectoryId uint64 `gorm:"column:directory_id"`
	OwnerId     string `gorm:"column:owner_id;size:36;not null"`
	DataState   uint64 `gorm:"column:data_state;precision:1;not null"`
}
