package actor

import (
	"context"
	"github.com/goriiin/go-films/internal/domain"
	"log/slog"
)

type usecase interface {
	CreateActorPage(ctx context.Context, actor domain.Actor) (string, error)
}

type Delivery struct {
	usecase usecase
	log     *slog.Logger
}

func NewActorDelivery(usecase usecase, log *slog.Logger) *Delivery {
	return &Delivery{
		usecase: usecase,
		log:     log,
	}
}
