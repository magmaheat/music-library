package pgdb

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/magmaheat/music-library/internal/model"
	def "github.com/magmaheat/music-library/internal/repo"
	"github.com/magmaheat/music-library/pkg/postgres"
	log "github.com/sirupsen/logrus"
)

var _ def.MusicRepo = (*repo)(nil)

type repo struct {
	*postgres.Postgres
}

func NewRepository(pg *postgres.Postgres) def.MusicRepo {
	return &repo{
		Postgres: pg,
	}
}

func (r *repo) DeleteSong(ctx context.Context, id int) error {

}

func (r *repo) UpdateSong(ctx context.Context, id int, song model.Song) (model.Song, error) {
	const fn = "repo - pgdb - UpdateSong"

	log.Debug("%s, %v", fn, song)

	queryBuilder := r.Builder.Update("songs").Where("id = ?", id)

	if song.Name != nil {
		queryBuilder = queryBuilder.Set("name", song.Name)
	}

	if song.Detail.Link != nil {
		queryBuilder = queryBuilder.Set("link", song.Detail.Link)
	}

	if song.Detail.Lyrics != nil {
		queryBuilder = queryBuilder.Set("lyrics", song.Detail.Lyrics)

	}

	if song.Detail.ReleaseDate != nil {
		queryBuilder = queryBuilder.Set("release_date", song.Detail.ReleaseDate)
	}

	sql, args, _ := queryBuilder.Prefix("RETURNING name, group_name, release_date, lyrics, link").ToSql()

	var updateSong model.Song

	err := r.Pool.QueryRow(ctx, sql, args...).Scan(
		&updateSong.Name,
		&updateSong.GroupName,
		&updateSong.Detail.ReleaseDate,
		&updateSong.Detail.Lyrics,
		&updateSong.Detail.Link,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			log.Errorf("%s - QueryRow: %v", err)
			return model.Song{}, model.ErrorSongNotFound
		}

		log.Errorf("%s - QueryRow: %v", fn, err)
		return model.Song{}, fmt.Errorf("QueryRow")
	}

	return updateSong, nil
}
