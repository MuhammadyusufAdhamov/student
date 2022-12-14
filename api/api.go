package api

import (
	// v1 "github.com/MuhammadyusufAdhamov/student/api/v1"
	v1 "github.com/MuhammadyusufAdhamov/student/api/v1"
	"github.com/MuhammadyusufAdhamov/student/config"
	"github.com/MuhammadyusufAdhamov/student/storage"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	_"github.com/MuhammadyusufAdhamov/student/api/docs"
)

type RouterOptions struct {
	Cfg     *config.Config
	Storage storage.StorageI
}

// @title 			Swagger for blog api
// @version 		1.0
// description 		This is a blog service api.
// @host 			localhost:8000
// @BasePath 		/v1
func New(opt *RouterOptions) *gin.Engine {
	router := gin.Default()

	handlerV1 := v1.New(&v1.HandlerV1Options{
		Cfg:     opt.Cfg,
		Storage: opt.Storage,
	})

	apiV1 := router.Group("/v1")

	apiV1.POST("/users", handlerV1.CreateUser)
	// apiV1.GET("/users", handlerV1.GetAll)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}