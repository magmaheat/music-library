package song

import (
	"context"
	"github.com/magmaheat/music-library/internal/model"
	"github.com/magmaheat/music-library/internal/repo"
	def "github.com/magmaheat/music-library/internal/service"
)

var _ def.MusicService = (*service)(nil)

type service struct {
	musicRepo repo.MusicRepo
}

func NewService(musicRepo repo.MusicRepo) def.MusicService {
	return &service{
		musicRepo: musicRepo,
	}
}

func (s *service) DeleteSong(ctx context.Context, id int) error {
	err := s.musicRepo.DeleteSong(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) UpdateSong(ctx context.Context, id int, song model.Song) error {
	err := s.musicRepo.UpdateSong(ctx, id, song)
	if err != nil {
		return model.Song{}, err
	}

	return nil
}
