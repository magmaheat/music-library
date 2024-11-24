package repo

import (
	"github.com/magmaheat/music-library/internal/repo/pgdb"
	"github.com/magmaheat/music-library/pkg/postgres"
)

type MusicLibrary interface {
}

type Repositories struct {
	MusicLibrary
}

func New(pg *postgres.Postgres) *Repositories {
	return &Repositories{
		MusicLibrary: pgdb.NewRepoMusicLibrary(pg),
	}
}
