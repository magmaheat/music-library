package http

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/magmaheat/music-library/internal/http/converter"
	"github.com/magmaheat/music-library/internal/model"
	"github.com/magmaheat/music-library/internal/service"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"time"
)

type musicHandler struct {
	service service.MusicService
}

func newMusicHandler(service service.MusicService) *musicHandler {
	return &musicHandler{
		service: service,
	}
}

func (h *musicHandler) getLibrary(ctx *gin.Context) {

}

func (h *musicHandler) getSongLyrics(ctx *gin.Context) {

}

func (h *musicHandler) deleteSong(ctx *gin.Context) {
	idParam := ctx.Param("id")
	if idParam == "" {
		errorResponse(ctx, http.StatusBadRequest, "no id param")
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, "not valid id")
		return
	}

	err = h.service.DeleteSong(ctx, id)
	if err != nil {
		if errors.Is(err, model.ErrorSongNotFound) {
			errorResponse(ctx, http.StatusNotFound, err.Error())
			return
		}

		errorResponse(ctx, http.StatusInternalServerError, "internal server error")
		return
	}

	ctx.Status(http.StatusOK)
}

type inputUpdateSong struct {
	NameSong    *string    `json:"song"`
	NameGroup   *string    `json:"group"`
	Lyrics      []string   `json:"lyrics"`
	ReleaseDate *time.Time `json:"release_date"`
	Link        *string    `json:"link"`
}

func (h *musicHandler) updateSong(ctx *gin.Context) {
	idParam := ctx.Param("id")
	if idParam == "" {
		errorResponse(ctx, http.StatusBadRequest, "no id param")
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, "not valid id")
		return
	}

	log.Debug("id update song:")

	var input inputUpdateSong
	if err = ctx.Bind(&input); err != nil {
		errorResponse(ctx, http.StatusBadRequest, "invalid body request")
		return
	}

	err = h.service.UpdateSong(ctx, id, converter.ToSongFromHTTPUpdate(
		input.NameSong,
		input.NameGroup,
		input.Link,
		input.Lyrics,
		input.ReleaseDate,
	))
	if err != nil {
		if errors.Is(err, model.ErrorSongNotFound) {
			errorResponse(ctx, http.StatusNotFound, err.Error())
			return
		}

		errorResponse(ctx, http.StatusInternalServerError, "internal server error")
		return
	}

	ctx.Status(http.StatusOK)

}

func (h *musicHandler) addSong(ctx *gin.Context) {

}
