package operating

import (
	"fmt"
	"github.com/go-chi/chi"
	"logity/internal/domain/usecase/operating"
	"net/http"
)

type Handler struct {
	usecase *operating.Usecase
}

func NewHandler(usecase *operating.Usecase) *Handler {
	return &Handler{
		usecase: usecase,
	}
}

// @Summary updateSubscribes
// @Tags operating
// @Security ApiKeyAuth
// @Description when a user loses connection with centrifugo, subscriptions to all channels are automatically lost, in order to restore the subscription when the token expires or disconnects, this route is used
// @ID update-subscribes
// @Accept json
// @Produce json
// @Success 200 {string} string  "just status code"
// @Failure 401 {string} string "unauth"
// @Router /op/update-subscribes [patch]
func (h *Handler) handleUpdatedSubscribes(w http.ResponseWriter, r *http.Request) {
	if err := h.usecase.UpdateSubscribes(r.Context()); err != nil {
		w.WriteHeader(400)
		w.Write([]byte(fmt.Sprintf("error: %s", err)))
		return
	}
	w.WriteHeader(200)
	return
}
func (h *Handler) Register(r chi.Router) {
	r.Group(func(r chi.Router) {
		r.Route("/op", func(r chi.Router) {
			r.Patch("/update-subscribes", h.handleUpdatedSubscribes)
		})
	})
}
