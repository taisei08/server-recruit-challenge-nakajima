package service

import (
	"context"

	"github.com/pulse227/server-recruit-challenge-sample/model"
	"github.com/pulse227/server-recruit-challenge-sample/repository"
)

type AlbumService interface {
	GetAlbumListService(ctx context.Context) ([]*model.Album, error)
	GetAlbumService(ctx context.Context, albumID model.AlbumID) (*model.Album, error)
}

type albumService struct {
	albumRepository repository.AlbumRepository
}

var _ AlbumService = (*albumService)(nil)

func NewAlbumService(albumRepository repository.AlbumRepository) *albumService {
	return &albumService{albumRepository: albumRepository}
}

func (s *albumService) GetAlbumListService(ctx context.Context) ([]*model.Album, error) {
	albums, err := s.albumRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return albums, nil
}

func (s *albumService) GetAlbumService(ctx context.Context, albumID model.AlbumID) (*model.Album, error) {
	album, err := s.albumRepository.Get(ctx, albumID)
	if err != nil {
		return nil, err
	}
	return album, nil
}
