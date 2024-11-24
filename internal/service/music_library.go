package service

import "github.com/magmaheat/music-library/internal/repo"

type MusicLibraryService struct {
	*repo.Repositories
}

func NewServicesMusicLibrary(rp *repo.Repositories) *MusicLibraryService {
	return &MusicLibraryService{
		Repositories: rp,
	}
}
