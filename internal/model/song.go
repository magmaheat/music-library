package model

import "time"

type Song struct {
	Name      string
	GroupName string
	Text      string
	Link      string
}

type SongDetail struct {
	ReleaseDate time.Time
	Text        string
	Link        string
}
