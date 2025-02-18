package actor

import (
	"context"
	"github.com/google/uuid"
	"github.com/goriiin/go-films/internal/domain"
)

func (u *Usecase) CreateActorPage(ctx context.Context, actor domain.Actor) (string, error) {
	actor.ID = uuid.New().String()

	err := u.repo.CreateActorPage(ctx, actor)
	if err != nil {
		u.log.Error("CreateActorPage", "error", err)

		return "", err
	}

	return actor.ID, nil
}
