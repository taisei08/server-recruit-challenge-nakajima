package model

type SingerID int

type Singer struct {
	ID   SingerID `json:"id,omitempty"`
	Name string   `json:"name,omitempty"`
}

func (s *Singer) Validate() error {
	if s.Name == "" {
		return ErrInvalidParam
	}
	if len(s.Name) > 255 {
		return ErrInvalidParam
	}
	return nil
}
