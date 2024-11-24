package service

import "github.com/magmaheat/music-library/internal/repo"

type MusicLibrary interface {
}

type Services struct {
	MusicLibrary
}

func NewServices(rp *repo.Repositories) *Services {
	return &Services{
		MusicLibrary: NewServicesMusicLibrary(rp),
	}
}
