package http

import (
	"github.com/gin-gonic/gin"
	_ "github.com/magmaheat/music-library/docs"
	"github.com/magmaheat/music-library/internal/http/middleware"
	"github.com/magmaheat/music-library/internal/service"
	log "github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"os"
)

func NewRouter(service service.MusicService) *gin.Engine {
	router := gin.Default()

	router.Use(gin.Recovery())
	router.Use(middleware.SetRequestID())
	router.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		Output: setLogsFile(),
	}))

	router.GET("/health", func(c *gin.Context) { c.Status(http.StatusOK) })
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	h := newMusicHandler(service)

	router.GET("/library", h.getLibrary)
	router.GET("songs/:id", h.getSongLyrics)
	router.DELETE("/songs/:id", h.deleteSong)
	router.PUT("/songs/:id", h.updateSong)
	router.POST("/songs", h.addSong)

	return router
}

func setLogsFile() *os.File {
	file, err := os.OpenFile("logs/requests.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}

	return file
}
