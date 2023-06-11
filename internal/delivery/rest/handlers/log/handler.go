package log

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"logity/internal/delivery/rest/handlers/log/dto/input"
	"logity/internal/domain/usecase/log"
	inputUsecase "logity/internal/domain/usecase/log/dto/input"
	"net/http"
)

type Handler struct {
	usecase *log.Usecase
}

func NewHandler(usecase *log.Usecase) *Handler {
	return &Handler{
		usecase: usecase,
	}
}

func (h *Handler) handlePushTextLog(w http.ResponseWriter, r *http.Request) {
	logInput := &input.LogTextInput{}
	if err := json.NewDecoder(r.Body).Decode(logInput); err != nil {
		w.WriteHeader(400)
		w.Write([]byte(fmt.Sprintf("request body error: %s", err)))
	}
	if err := logInput.Validate(); err != nil {
		w.WriteHeader(422)
		w.Write([]byte(fmt.Sprintf("validation error: %s", err)))
		return
	}

	if err := h.usecase.PushTextLog(r.Context(), inputUsecase.PushLogTextDto{
		Text:    logInput.Text,
		RoomIds: logInput.RoomIds,
	}); err != nil {
		w.WriteHeader(400)
		w.Write([]byte(fmt.Sprintf("error: %s", err)))
		return
	}
	w.WriteHeader(200)
	return
}
func (h *Handler) Register(r chi.Router) {
	r.Group(func(r chi.Router) {
		r.Route("/log", func(r chi.Router) {
			r.Post("/push-text-log", h.handlePushTextLog)
		})
	})
}
