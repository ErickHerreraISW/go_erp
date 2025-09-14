package users

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/youruser/myapp/internal/pkg/response"
)

type Handler struct {
	Svc     Service
	JWTAuth *jwtauth.JWTAuth
}

func NewHandler(s Service, t *jwtauth.JWTAuth) *Handler { return &Handler{Svc: s, JWTAuth: t} }

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var dto CreateUserDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		response.Fail(w, 400, err.Error())
		return
	}
	u, err := h.Svc.Create(dto)
	if err != nil {
		response.Fail(w, 400, err.Error())
		return
	}
	response.JSON(w, http.StatusCreated, u)
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	u, err := h.Svc.List()
	if err != nil {
		response.Fail(w, 500, err.Error())
		return
	}
	response.JSON(w, 200, u)
}

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	u, err := h.Svc.Get(uint(id))
	if err != nil {
		response.Fail(w, 404, "not found")
		return
	}
	response.JSON(w, 200, u)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	var dto UpdateUserDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		response.Fail(w, 400, err.Error())
		return
	}
	u, err := h.Svc.Update(uint(id), dto)
	if err != nil {
		response.Fail(w, 400, err.Error())
		return
	}
	response.JSON(w, 200, u)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	if err := h.Svc.Delete(uint(id)); err != nil {
		response.Fail(w, 400, err.Error())
		return
	}
	w.WriteHeader(204)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var dto LoginDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		response.Fail(w, 400, err.Error())
		return
	}
	tok, err := h.Svc.Login(dto, h.JWTAuth)
	if err != nil {
		response.Fail(w, 401, err.Error())
		return
	}
	response.JSON(w, 200, map[string]string{"token": tok})
}
