package room

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"logity/internal/domain/usecase/room"
	"logity/internal/domain/usecase/room/dto/input"
	"net/http"
)

type Handler struct {
	usecase *room.Usecase
}

func NewHandler(usecase *room.Usecase) *Handler {
	return &Handler{
		usecase: usecase,
	}
}

func (h *Handler) handleCreateRoom(w http.ResponseWriter, r *http.Request) {
	createRoom := &CreateRoom{}
	if err := json.NewDecoder(r.Body).Decode(createRoom); err != nil {
		w.WriteHeader(400)
		w.Write([]byte(fmt.Sprintf("request body error: %s", err)))
	}
	if err := createRoom.Validate(); err != nil {
		w.WriteHeader(422)
		w.Write([]byte(fmt.Sprintf("validation error: %s", err)))
		return
	}

	roomOutput, err := h.usecase.CreateNewRoom(r.Context(), input.CreateRoomDto{
		Name: createRoom.Name,
	})
	if err != nil {
		w.WriteHeader(400) //todo separate erros by codeStatus
		w.Write([]byte(fmt.Sprintf("error: %s", err)))
		return
	}

	resp, err := json.Marshal(roomOutput)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("error marshal response: %s", err)))
		return
	}

	w.WriteHeader(201)
	w.Write(resp)
	return
}

func (h *Handler) Register(r chi.Router) {
	r.Group(func(r chi.Router) {
		r.Route("/room", func(r chi.Router) {
			r.Post("/", h.handleCreateRoom)
		})
	})
}
