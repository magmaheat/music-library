package converter

import (
	"github.com/magmaheat/music-library/internal/model"
	"time"
)

func ToSongFromHTTPUpdate(song, group, link *string, lyrics []string, releaseDate *time.Time) model.Song {
	return model.Song{
		Name:      song,
		GroupName: group,
		Detail: model.SongDetail{
			ReleaseDate: releaseDate,
			Link:        link,
			Lyrics:      lyrics,
		},
	}
}
