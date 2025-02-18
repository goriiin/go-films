package actor

import (
	"github.com/goriiin/go-films/internal/domain"
	"time"
)

//easyjson:json
type dtoCreateRequest struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Birthday    time.Time `json:"birthday"`
	Gender      string    `json:"gender"`
}

func (dto dtoCreateRequest) toDomain() domain.Actor {
	return domain.Actor{
		Name:        dto.Name,
		Description: dto.Description,
		Birthday:    dto.Birthday,
		Gender:      dto.Gender,
	}
}
