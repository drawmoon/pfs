package model

type Directory struct {
	Id   uint64 `gorm:"column:id;primaryKey"`
	Name string `gorm:"column:name;size:40;not null"`
}
