package model

type AlbumID int

type Album struct {
	ID       AlbumID  `json:"id,omitempty"`
	Title    string   `json:"title,omitempty"`
	Singer   Singer   `json:"singer,omitempty"`
}

func (a *Album) Validate() error {
	if a.Title == "" {
		return ErrInvalidParam
	}
	if len(a.Title) > 255 {
		return ErrInvalidParam
	}
	return nil
}
