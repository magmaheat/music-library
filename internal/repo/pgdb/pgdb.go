package pgdb

import (
	def "github.com/magmaheat/music-library/internal/repo"
	"github.com/magmaheat/music-library/pkg/postgres"
)

var _ def.MusicRepo = (*repo)(nil)

type repo struct {
	*postgres.Postgres
}

func NewRepository(pg *postgres.Postgres) *repo {
	return &repo{
		Postgres: pg,
	}
}
