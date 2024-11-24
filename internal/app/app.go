package app

import (
	"github.com/magmaheat/music-library/internal/config"
	"github.com/magmaheat/music-library/internal/http"
	"github.com/magmaheat/music-library/internal/repo"
	"github.com/magmaheat/music-library/internal/service"
	"github.com/magmaheat/music-library/pkg/postgres"
	log "github.com/sirupsen/logrus"
)

func Run() {
	cfg := config.MustLoad()

	setupLogger(cfg.Env)

	log.Infof("Initializing storage...")
	pg, err := postgres.New(cfg.URL)
	if err != nil {
		log.Fatalf("Error init storage: %v", err)
	}

	log.Infof("Initializing repositories...")
	repositories := repo.New(pg)

	log.Infof("Initializing services...")
	services := service.NewServices(repositories)

	router := http.NewRouter(services)

	router.Run(":" + cfg.Port)
}
