package pgdb

import (
	"context"
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
	const fn = "repo - pgdb - DeleteSong"
	sql, args, _ := r.Builder.Delete("songs").Where("id = ?", id).ToSql()

	_, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		log.Errorf("%s - Pool.Query: %v", fn, err)
		return err
	}

	return nil
}

func (r *repo) UpdateSong(ctx context.Context, id int, song model.Song) error {
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

	_, err := r.Pool.Query(ctx, sql, args...)

	if err != nil {
		log.Errorf("%s - QueryRow: %v", fn, err)
		return err
	}

	return nil
}

func (r *repo) GetIdGroup(ctx context.Context, group string) (int, error) {
	sql, args, _ := r.Builder.Select("id").
		From("groups").
		Where("name = ?", group).
		ToSql()

	var id int

	err := r.Pool.QueryRow(ctx, sql, args...).Scan(&id)
	if err != nil {
		log.Errorf("repo - pgdb - GetIdGroup: %v", err)
		return 0, err
	}

	return id, nil
}
