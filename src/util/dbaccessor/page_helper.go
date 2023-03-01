package dbaccessor

import "gorm.io/gorm"

func Paginate(page int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		const defaultPageNumber = 1

		// 检查页码是否大于零，否则使用默认页码
		if page <= 0 {
			page = defaultPageNumber
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
