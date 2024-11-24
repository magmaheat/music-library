package app

import "github.com/magmaheat/music-library/internal/config"

func Run() {
	cfg := config.MustLoad()

	setupLogger(cfg.Env)

}
