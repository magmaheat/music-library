package app

import (
	"github.com/magmaheat/music-library/internal/config"
	"github.com/magmaheat/music-library/internal/http"
	"github.com/magmaheat/music-library/internal/repo/pgdb"
	"github.com/magmaheat/music-library/internal/service/song"
	"github.com/magmaheat/music-library/pkg/postgres"
	log "github.com/sirupsen/logrus"
)

// @title Music Info
// @version 0.0.1
// @description Music Info is a service for searching and adding song to your library.

// @contact.name George Epishev
// @contact.email epishcom@gmail.com

// @host localhost:8090
// @BasePath /
func Run() {
	cfg := config.MustLoad()

	setupLogger(cfg.Env)

	log.Info("Initializing storage...")
	pg, err := postgres.New(cfg.URL)
	if err != nil {
		log.Fatalf("Error init storage: %v", err)
	}

	log.Info("Initializing repositories...")
	repositories := pgdb.NewRepository(pg)

	log.Info("Initializing services...")
	services := song.NewService(repositories)

	log.Info("Initializing router and handlers...")
	router := http.NewRouter(services)

	log.Infof("Starting server, port: %s", cfg.Port)
	router.Run(":" + cfg.Port)
}
