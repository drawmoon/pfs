package main

import (
	"os"

	_ "github.com/PowerReport/pfs/docs"
	"github.com/PowerReport/pfs/src/app/controller"
	appSvc "github.com/PowerReport/pfs/src/app/service"
	ds "github.com/PowerReport/pfs/src/domain/directory/service"
	fs "github.com/PowerReport/pfs/src/domain/file/service"
	"github.com/PowerReport/pfs/src/infra/postgres"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	swagger "github.com/swaggo/gin-swagger"
)

// @title Power File Storage API
// @version 1.0
// @description DDD-based file storage management project built with Gin.
func main() {
	// 加载环境变量配置文件
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// 从环境变量中获取数据库配置
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	pwd := os.Getenv("DB_PASSWORD")
	db := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	tablePrefix := os.Getenv("DB_TABLE_PREFIX")

	// 初始化数据库上下文
	dbContext, err := postgres.NewDbContext(host, user, pwd, db, tablePrefix, port)
	if err != nil {
		panic(err)
	}

	// 停止后台服务后关闭数据库连接
	defer postgres.Close(dbContext)

	// 数据库迁移
	err = postgres.AutoMigrate(dbContext)
	if err != nil {
		panic(err)
	}

	// 初始化仓储服务
	directoryRepository := postgres.NewDirectoryRepository(dbContext)
	fileRepository := postgres.NewFileRepository(dbContext)

	// 初始化领域服务
	directoryService := ds.NewDirectoryService(directoryRepository)
	fileService := fs.NewFileService(fileRepository)

	// 初始化应用层服务
	mixinService := appSvc.NewMixinService(fileService, directoryService)

	// 初始化控制器
	mixin := controller.NewMixinController(mixinService)

	// 初始化服务器
	server := gin.Default()

	// 配置接口路由
	fs := server.Group("api/fs")
	{
		fs.GET(":id", mixin.GetInfo)
		fs.GET(":id/files", mixin.GetFiles)
		fs.GET(":id/directories", mixin.GetDirectories)
		fs.POST("", mixin.Create)
		fs.PUT("rename", mixin.Rename)
		fs.PUT("move", mixin.Move)
		fs.DELETE("", mixin.Delete)
	}

	// Swagger 的访问路由，/swagger/index.html
	server.GET("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler))

	// 配置后台服务监听的端口
	server.Run(":3000")
}
