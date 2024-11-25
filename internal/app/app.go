package app

import (
	"github.com/magmaheat/music-library/internal/config"
	"github.com/magmaheat/music-library/internal/http"
	"github.com/magmaheat/music-library/internal/repo/pgdb"
	"github.com/magmaheat/music-library/internal/service/music"
	"github.com/magmaheat/music-library/pkg/postgres"
	log "github.com/sirupsen/logrus"
)

// @title Music Info
// @version 0.0.1
// @description Music Info is a service for searching and adding music to your library.

// @contact.name George Epishev
// @contact.email epishcom@gmail.com

// @host localhost:8090
// @BasePath /
func Run() {
	cfg := config.MustLoad()

	setupLogger(cfg.Env)

	log.Infof("Initializing storage...")
	pg, err := postgres.New(cfg.URL)
	if err != nil {
		log.Fatalf("Error init storage: %v", err)
	}

	log.Infof("Initializing repositories...")
	repositories := pgdb.NewRepository(pg)

	log.Infof("Initializing services...")
	services := music.NewService(repositories)

	router := http.NewRouter(services)

	router.Run(":" + cfg.Port)
}
