package actor

import (
	"github.com/mailru/easyjson"
	"net/http"
)

func (d *Delivery) CreateActorPage(w http.ResponseWriter, r *http.Request) {
	var dto dtoCreateRequest

	err := easyjson.UnmarshalFromReader(r.Body, &dto)
	if err != nil {
		return
	}

	_, err = d.usecase.CreateActorPage(r.Context(), dto.toDomain())
	if err != nil {
		return
	}
}
