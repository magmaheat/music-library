package service

import (
	"context"
	"github.com/magmaheat/music-library/internal/model"
)

type MusicService interface {
	DeleteSong(ctx context.Context, id int) error
	UpdateSong(ctx context.Context, id int, song model.Song) error
}
