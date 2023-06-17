package auth

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"logity/internal/domain/entity/user"
	"logity/internal/domain/usecase/auth"
	"logity/internal/domain/usecase/auth/dto"
	"logity/internal/domain/usecase/room"
	"net/http"
	"strings"
)

type Handler struct {
	authUsecase *auth.Usecase
	roomUsecase *room.Usecase
}

func NewHandler(usecase *auth.Usecase, roomUsecase *room.Usecase) *Handler {
	return &Handler{
		authUsecase: usecase,
		roomUsecase: roomUsecase,
	}
}

func (h *Handler) handleSignIn(w http.ResponseWriter, r *http.Request) {
	signIn := &SignIn{}
	if err := json.NewDecoder(r.Body).Decode(signIn); err != nil {
		w.WriteHeader(400)
		w.Write([]byte(fmt.Sprintf("request body error: %s", err)))
		return
	}
	if err := signIn.Validate(); err != nil {
		w.WriteHeader(422)
		w.Write([]byte(fmt.Sprintf("validation error: %s", err)))
		return
	}
	signInDto, err := h.authUsecase.SignIn(r.Context(), dto.SignInInputDto{
		Login:    signIn.Login,
		Password: signIn.Password,
	})
	if err != nil {
		w.WriteHeader(400) //todo separate erros by codeStatus
		w.Write([]byte(fmt.Sprintf("error: %s", err)))
		return
	}

	resp, err := json.Marshal(signInDto)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("error marshal response: %s", err)))
		return
	}

	w.WriteHeader(200)
	w.Write(resp)
	return
}

func (h *Handler) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.WriteHeader(401)
			w.Write([]byte("this method secure authenticated, request doesn't contained Authorization token"))
			return
		}
		splitedAuthHeader := strings.Split(authHeader, " ")
		if len(splitedAuthHeader) != 2 || splitedAuthHeader[0] != "Bearer" {
			w.WriteHeader(401)
			w.Write([]byte("content of Authorization header isn't valid"))
			return
		}
		u, err := h.authUsecase.FindUserByAccessToken(r.Context(), splitedAuthHeader[1])
		if err != nil {
			w.WriteHeader(401)
			w.Write([]byte(fmt.Sprintf("error find user: %s", err)))
			return
		}

		next.ServeHTTP(w, r.WithContext(user.PutToCtx(r.Context(), u)))
	})
}

func (h *Handler) handleMe(w http.ResponseWriter, r *http.Request) {
	out, err := h.authUsecase.Me(r.Context())
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("%s", err)))
		return
	}

	resp, err := json.Marshal(out)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("error marshal json: %s", err)))
		return
	}

	w.WriteHeader(200)
	w.Write(resp)
	return
}

func (h *Handler) handleSignUp(w http.ResponseWriter, r *http.Request) {
	in := dto.SignUpByEmailInputDto{}
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		w.WriteHeader(400)
		w.Write([]byte(fmt.Sprintf("error decode request body: %s", err)))
		return
	}

	if err := in.Validate(); err != nil {
		w.WriteHeader(422)
		w.Write([]byte(fmt.Sprintf("validation error: %s", err)))
		return
	}

	out, err := h.authUsecase.SignUpByEmail(r.Context(), in)
	if err != nil {
		w.WriteHeader(400) //todo separate erros by codeStatus
		w.Write([]byte(fmt.Sprintf("error: %s", err)))
		return
	}
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

func (h *Handler) handleUpdateAccessToken(w http.ResponseWriter, r *http.Request) {
	in := dto.UpdateTokenInputDto{}
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		w.WriteHeader(400)
		w.Write([]byte(fmt.Sprintf("error decode request body: %s", err)))
		return
	}

	if err := in.Validate(); err != nil {
		w.WriteHeader(422)
		w.Write([]byte(fmt.Sprintf("validation error: %s", err)))
		return
	}

	out, err := h.authUsecase.UpdateAccessToken(r.Context(), in)
	if err != nil {
		w.WriteHeader(400) //todo separate erros by codeStatus
		w.Write([]byte(fmt.Sprintf("error: %s", err)))
		return
	}

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

func (h *Handler) handleResetPassword(w http.ResponseWriter, r *http.Request) {
	panic("implement handleResetPassword me!!!")
}

func (h *Handler) Register(r chi.Router) {
	r.Route("/auth", func(r chi.Router) {
		r.Post("/sign-in", h.handleSignIn)
		r.Post("/sign-up", h.handleSignUp)
		r.Patch("/refresh", h.handleUpdateAccessToken)
		r.Post("/reset-password", h.handleResetPassword)
		r.Delete("/revoke", h.handleResetPassword)
	})

	r.Group(func(r chi.Router) {
		r.Use(h.AuthMiddleware)
		r.Get("/me", h.handleMe)
	})
}
