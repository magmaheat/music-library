package pgdb

import "github.com/magmaheat/music-library/pkg/postgres"

type Repo struct {
	*postgres.Postgres
}

func NewRepoMusicLibrary(pg *postgres.Postgres) *Repo {
	return &Repo{
		Postgres: pg,
	}
}
