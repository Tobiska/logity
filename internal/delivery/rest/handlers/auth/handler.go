package auth

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"logity/internal/domain/usecase/auth"
	"logity/internal/domain/usecase/auth/dto"
	"net/http"
)

type Handler struct {
	usecase *auth.AuthUsecase
}

func NewHandler(usecase *auth.AuthUsecase) *Handler {
	return &Handler{
		usecase: usecase,
	}
}

func (h *Handler) handleSignIn(w http.ResponseWriter, r *http.Request) {
	signIn := &SignIn{}
	if err := json.NewDecoder(r.Body).Decode(signIn); err != nil {
		w.WriteHeader(400)
	}
	if err := signIn.Validate(); err != nil {
		w.Write([]byte(fmt.Sprintf("validation error: %s", err)))
		w.WriteHeader(422)
	}
	signInDto, err := h.usecase.SignIn(r.Context(), dto.SignInInputDto{
		Login:    signIn.Login,
		Password: signIn.Password,
	})
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error: %s", err)))
		w.WriteHeader(400) //todo separate erros by codeStatus
	}

	resp, err := json.Marshal(signInDto)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error marshal response: %s", err)))
		w.WriteHeader(500)
	}

	w.Write(resp)
	w.WriteHeader(200)
}

func (h *Handler) handleSignUp(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) handleUpdateAccessToken(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) handleResetPassword(w http.ResponseWriter, r *http.Request) {
	panic("implement handleResetPassword me!!!")
}

func (h *Handler) Register(r chi.Router) {
	r.Group(func(r chi.Router) {
		r.Post("/sign-in", h.handleSignIn)
		r.Post("/sign-up", h.handleSignUp)
		r.Patch("/refresh", h.handleUpdateAccessToken)
		r.Post("/reset-password", h.handleResetPassword)
		r.Delete("/revoke", h.handleResetPassword)
	})
}
