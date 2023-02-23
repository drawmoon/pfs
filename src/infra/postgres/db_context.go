package postgres

import (
	"fmt"

	dm "github.com/PowerReport/pfs/src/domain/directory/model"
	fm "github.com/PowerReport/pfs/src/domain/file/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gs "gorm.io/gorm/schema"
)

func NewDbContext(host, user, pwd, db, tablePrefix, port string) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		host, user, pwd, db, port)
	config := gorm.Config{
		NamingStrategy: gs.NamingStrategy{
			TablePrefix:   tablePrefix, // 表的统一前缀
			SingularTable: false,
		},
	}
	return gorm.Open(postgres.Open(dsn), &config)
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&dm.Directory{}, &fm.File{})
}

func Close(db *gorm.DB) error {
	database, err := db.DB()
	if err != nil {
		return err
	}

	return database.Close()
}
