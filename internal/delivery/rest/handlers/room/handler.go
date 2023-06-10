package room

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	input "logity/internal/delivery/rest/handlers/room/dto/input"
	"logity/internal/delivery/rest/handlers/room/dto/output"
	"logity/internal/domain/usecase/room"
	inputUsecase "logity/internal/domain/usecase/room/dto/input"
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
	createRoom := &input.CreateRoom{}
	if err := json.NewDecoder(r.Body).Decode(createRoom); err != nil {
		w.WriteHeader(400)
		w.Write([]byte(fmt.Sprintf("request body error: %s", err)))
	}
	if err := createRoom.Validate(); err != nil {
		w.WriteHeader(422)
		w.Write([]byte(fmt.Sprintf("validation error: %s", err)))
		return
	}

	newRoom, err := h.usecase.CreateNewRoom(r.Context(), inputUsecase.CreateRoomDto{
		Name: createRoom.Name,
	})
	if err != nil {
		w.WriteHeader(400) //todo separate erros by codeStatus
		w.Write([]byte(fmt.Sprintf("error: %s", err)))
		return
	}

	out := output.NewCreateRoomOutputDto(newRoom)

	resp, err := json.Marshal(out)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("error marshal response: %s", err)))
		return
	}

	w.WriteHeader(201)
	w.Write(resp)
	return
}

func (h *Handler) handleInviteRoom(w http.ResponseWriter, r *http.Request) {
	inviteRoom := &input.InviteRoom{}
	if err := json.NewDecoder(r.Body).Decode(inviteRoom); err != nil {
		w.WriteHeader(400)
		w.Write([]byte(fmt.Sprintf("request body error: %s", err)))
	}
	if err := inviteRoom.Validate(); err != nil {
		w.WriteHeader(422)
		w.Write([]byte(fmt.Sprintf("validation error: %s", err)))
		return
	}

	err := h.usecase.InviteToRoom(r.Context(), inputUsecase.InviteToRoomDto{
		RoomId: inviteRoom.RoomId,
		UserId: inviteRoom.UserId,
	})
	if err != nil {
		w.WriteHeader(400) //todo separate erros by codeStatus
		w.Write([]byte(fmt.Sprintf("error: %s", err)))
		return
	}

	w.WriteHeader(200)
	return
}

func (h *Handler) handleJoinRoom(w http.ResponseWriter, r *http.Request) {
	roomId := chi.URLParam(r, "room_id")

	updatedRoom, err := h.usecase.JoinToRoom(r.Context(), inputUsecase.JoinToRoomDto{
		RoomId: roomId,
	})
	if err != nil {
		w.WriteHeader(400) //todo separate erros by codeStatus
		w.Write([]byte(fmt.Sprintf("error: %s", err)))
		return
	}

	out := output.NewJoinOutputDto(updatedRoom)
	resp, err := json.Marshal(out)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("error marshal response: %s", err)))
		return
	}

	w.WriteHeader(200)
	w.Write(resp)
	return
}

func (h *Handler) handleShowRoom(w http.ResponseWriter, r *http.Request) {
	roomId := chi.URLParam(r, "room_id")

	foundedRoom, err := h.usecase.GetRoom(r.Context(), roomId)
	if err != nil {
		w.WriteHeader(400) //todo separate erros by codeStatus
		w.Write([]byte(fmt.Sprintf("error: %s", err)))
		return
	}

	out := output.NewJoinOutputDto(foundedRoom)
	resp, err := json.Marshal(out)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("error marshal response: %s", err)))
		return
	}

	w.WriteHeader(200)
	w.Write(resp)
	return
}

func (h *Handler) Register(r chi.Router) {
	r.Group(func(r chi.Router) {
		r.Route("/room", func(r chi.Router) {
			r.Post("/", h.handleCreateRoom)
			r.Patch("/", h.handleInviteRoom)
			r.Patch("/{room_id}", h.handleJoinRoom)
			r.Get("/{room_id}", h.handleShowRoom)
		})
	})
}
