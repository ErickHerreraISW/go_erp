package products

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ErickHerreraISW/go_erp/internal/pkg/response"
	"github.com/go-chi/chi/v5"
)

type Handler struct{ Svc Service }

func NewHandler(s Service) *Handler { return &Handler{Svc: s} }

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var dto CreateProductDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		response.Fail(w, 400, err.Error())
		return
	}
	p, err := h.Svc.Create(dto)
	if err != nil {
		response.Fail(w, 400, err.Error())
		return
	}
	response.JSON(w, http.StatusCreated, p)
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	p, err := h.Svc.List()
	if err != nil {
		response.Fail(w, 500, err.Error())
		return
	}
	response.JSON(w, 200, p)
}

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	p, err := h.Svc.Get(uint(id))
	if err != nil {
		response.Fail(w, 404, "not found")
		return
	}
	response.JSON(w, 200, p)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	var dto UpdateProductDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		response.Fail(w, 400, err.Error())
		return
	}
	p, err := h.Svc.Update(uint(id), dto)
	if err != nil {
		response.Fail(w, 400, err.Error())
		return
	}
	response.JSON(w, 200, p)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	if err := h.Svc.Delete(uint(id)); err != nil {
		response.Fail(w, 400, err.Error())
		return
	}
	w.WriteHeader(204)
}
