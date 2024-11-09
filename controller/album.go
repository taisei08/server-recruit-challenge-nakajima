package controller

import (
	"encoding/json"
	"net/http"

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
