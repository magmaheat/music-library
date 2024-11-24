package http

import (
	"github.com/gin-gonic/gin"
	"github.com/magmaheat/music-library/internal/service"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"

	_ "github.com/magmaheat/music-library/docs"
)

func NewRouter(services *service.Services) *gin.Engine {
	router := gin.Default()

	router.GET("health", func(c *gin.Context) { c.Status(http.StatusOK) })
	router.GET("/swagger/*", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
