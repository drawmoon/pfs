package main

import (
	_ "github.com/PowerReport/pfs/docs"
	"github.com/PowerReport/pfs/src/app/controller"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	swagger "github.com/swaggo/gin-swagger"
)

// @title Power File Storage API
// @version 1.0
// @description DDD-based file storage management project built with Gin.
func main() {
	server := gin.Default()

	file := controller.NewFileController()
	directory := controller.NewDirectoryController()

	fs := server.Group("api/fs")
	{
		files := fs.Group("files")
		{
			files.GET(":id", file.GetInfo)
		}
		directories := fs.Group("directories")
		{
			directories.GET(":id", directory.GetInfo)
			directories.GET(":id/files", directory.GetFiles)
			directories.GET(":id/directories", directory.GetDirectories)
		}
	}
	server.GET("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler))
	server.Run(":3000")
}
