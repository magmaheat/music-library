package music

import (
	"github.com/magmaheat/music-library/internal/repo"
	def "github.com/magmaheat/music-library/internal/service"
)

var _ def.MusicService = (*service)(nil)

type service struct {
	musicRepo repo.MusicRepo
}

func NewService(musicRepo repo.MusicRepo) *service {
	return &service{
		musicRepo: musicRepo,
	}
}
