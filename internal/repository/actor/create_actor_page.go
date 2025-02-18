package actor

import (
	"context"
	"github.com/goriiin/go-films/internal/domain"
)

func (r *Repository) CreateActorPage(ctx context.Context, actor domain.Actor) error {
	const query = `INSERT INTO actors 
    					(id, name, gender, birthday, description, created_time, updated_time) 
					VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err := r.db.Exec(ctx, query, actor.ID, actor.Name, actor.Gender, actor.Birthday, actor.Description)
	if err != nil {
		return err
	}

	return nil
}
