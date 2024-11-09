package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/pulse227/server-recruit-challenge-sample/model"
	"github.com/pulse227/server-recruit-challenge-sample/service"
)

type albumController struct {
	service service.AlbumService
}

func NewAlbumController(s service.AlbumService) *albumController {
	return &albumController{service: s}
}

// GET /albums のハンドラー
func (c *albumController) GetAlbumListHandler(w http.ResponseWriter, r *http.Request) {
	albums, err := c.service.GetAlbumListService(r.Context())
	if err != nil {
		errorHandler(w, r, 500, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(albums)
}

// GET /albums/{id} のハンドラー
func (c *albumController) GetAlbumDetailHandler(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	albumID, err := strconv.Atoi(idString)
	if err != nil {
		err = fmt.Errorf("invalid path param: %w", err)
		errorHandler(w, r, 400, err.Error())
		return
	}

	album, err := c.service.GetAlbumService(r.Context(), model.AlbumID(albumID))
	if err != nil {
		errorHandler(w, r, 500, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(album)
}
