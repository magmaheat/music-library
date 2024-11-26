package repo

import (
	"context"
	"github.com/magmaheat/music-library/internal/model"
)

type MusicRepo interface {
	DeleteSong(ctx context.Context, id int) error
	UpdateSong(ctx context.Context, id int, song model.Song) error
	GetIdGroup(ctx context.Context, group string) (int, error)
}
