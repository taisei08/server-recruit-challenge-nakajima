package repository

import (
	"context"

	"github.com/pulse227/server-recruit-challenge-sample/model"
)

type AlbumRepository interface {
	GetAll(ctx context.Context) ([]*model.Album, error)
}
