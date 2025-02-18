package actor

import (
	"context"
	"github.com/goriiin/go-films/internal/domain"
	"log/slog"
)

type actorsRepository interface {
	CreateActorPage(ctx context.Context, actor domain.Actor) error
}

type Usecase struct {
	repo actorsRepository
	log  *slog.Logger
}
