// package category provides category related functionality.
package category

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/emregocer/golang_project/config"
	"github.com/emregocer/golang_project/internal/handler"
	"github.com/emregocer/golang_project/internal/middleware"
	"github.com/go-playground/validator/v10"
	"github.com/lib/pq"

	"github.com/julienschmidt/httprouter"
)

const CATEGORY_ROUTER_BASE = "/categories"

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterHandlers(router *httprouter.Router, config config.Config) {
	router.GET(CATEGORY_ROUTER_BASE, h.Get)
	router.GET(CATEGORY_ROUTER_BASE+"/:id", h.GetOne)
	router.POST(CATEGORY_ROUTER_BASE, middleware.Jwt(h.Create, config.JWTKey))
	router.PUT(CATEGORY_ROUTER_BASE+"/:id", middleware.Jwt(h.Update, config.JWTKey))
	router.DELETE(CATEGORY_ROUTER_BASE+"/:id", middleware.Jwt(h.Delete, config.JWTKey))
}

func (h *Handler) GetOne(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		handler.JsonErrorResponse(w, http.StatusBadRequest, "Bad Request")
		return
	}

	category, err := h.service.GetOne(r.Context(), id)

	if err != nil {
		handler.JsonErrorResponse(w, http.StatusInternalServerError, "Server error")
		return
	}

	if category == nil {
		handler.JsonErrorResponse(w, http.StatusNotFound, "Resource not found")
		return
	}

	handler.JsonSuccessResponse(w, http.StatusOK, handler.NewAPIResponse(CategoryResource{
		category.Id,
		category.Name,
		category.Description,
	}))
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	categories, err := h.service.Get(r.Context())

	if err != nil {
		handler.JsonErrorResponse(w, http.StatusInternalServerError, "Server error")
		return
	}

	categoryResources := make([]CategoryResource, 0, len(categories))

	for _, c := range categoryResources {
		categoryResources = append(categoryResources, CategoryResource{
			c.Id,
			c.Name,
			c.Description,
		})
	}

	handler.JsonSuccessResponse(w, http.StatusOK, handler.NewAPIResponse(categoryResources))
}

type CreateCategoryRequest struct {
	Name        string `validate:"required"`
	Description string `validate:"required"`
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	createCategoryRequest := CreateCategoryRequest{}

	err := json.NewDecoder(r.Body).Decode(&createCategoryRequest)
	if err != nil {
		handler.JsonErrorResponse(w, http.StatusBadRequest, "Data is not valid.")
		return
	}

	validationError := validator.New().Struct(createCategoryRequest)

	if validationError != nil {
		handler.JsonErrorResponse(w, http.StatusBadRequest, "Data is not valid")
		return
	}

	category, err := h.service.Create(r.Context(), createCategoryRequest)

	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			if err.Code == "23505" {
				handler.JsonErrorResponse(w, http.StatusConflict, "Category exists")
				return
			}
		}

		handler.JsonErrorResponse(w, http.StatusInternalServerError, "Server error")
		return
	}

	handler.JsonSuccessResponse(w, http.StatusOK, handler.NewAPIResponse(CategoryResource{
		category.Id,
		category.Name,
		category.Description,
	}))
}

type UpdateCategoryRequest struct {
	Name        string `validate:"required"`
	Description string `validate:"required"`
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		handler.JsonErrorResponse(w, http.StatusBadRequest, "Bad Request")
		return
	}

	updateCategoryRequest := UpdateCategoryRequest{}

	err = json.NewDecoder(r.Body).Decode(&updateCategoryRequest)
	if err != nil {
		handler.JsonErrorResponse(w, http.StatusBadRequest, "Data is not valid")
		return
	}

	validationError := validator.New().Struct(updateCategoryRequest)

	if validationError != nil {
		handler.JsonErrorResponse(w, http.StatusBadRequest, "Data is not valid")
		return
	}

	category, err := h.service.Update(r.Context(), id, updateCategoryRequest)

	if category == nil {
		handler.JsonErrorResponse(w, http.StatusNotFound, "Resource not found")
		return
	}

	if err != nil {
		handler.JsonErrorResponse(w, http.StatusInternalServerError, "Server error")
		return
	}

	handler.JsonSuccessResponse(w, http.StatusOK, handler.NewAPIResponse(CategoryResource{
		category.Id,
		category.Name,
		category.Description,
	}))
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		handler.JsonErrorResponse(w, http.StatusBadRequest, "Bad Request")
		return
	}

	category, err := h.service.Delete(r.Context(), id)

	if err != nil {
		handler.JsonErrorResponse(w, http.StatusInternalServerError, "Server error")
		return
	}

	handler.JsonSuccessResponse(w, http.StatusOK, handler.NewAPIResponse(category))
}
