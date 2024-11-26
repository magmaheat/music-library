package model

import "time"

type Song struct {
	Name      *string
	GroupName *string
	Detail    SongDetail
}

type SongDetail struct {
	ReleaseDate *time.Time
	Lyrics      []string
	Link        *string
}
