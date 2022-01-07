// package movie provides movie related functionality.
package movie

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/emregocer/golang_project/config"
	"github.com/emregocer/golang_project/internal/category"
	"github.com/emregocer/golang_project/internal/handler"
	"github.com/emregocer/golang_project/internal/middleware"
	"github.com/go-playground/validator/v10"

	"github.com/julienschmidt/httprouter"
)

const MOVIE_ROUTER_BASE = "/movies"

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterHandlers(router *httprouter.Router, config config.Config) {
	router.GET(category.CATEGORY_ROUTER_BASE+"/:id/movies", h.GetByCategoryId)
	router.GET(MOVIE_ROUTER_BASE+"/id/:id", h.GetOne)
	router.POST(MOVIE_ROUTER_BASE, middleware.Jwt(h.Create, config.JWTKey))
	router.PUT(MOVIE_ROUTER_BASE+"/id/:id", middleware.Jwt(h.Update, config.JWTKey))
	router.DELETE(MOVIE_ROUTER_BASE+"/id/:id", middleware.Jwt(h.Delete, config.JWTKey))

	router.GET(MOVIE_ROUTER_BASE+"/favourites", middleware.Jwt(h.GetFavourites, config.JWTKey))
	router.POST(MOVIE_ROUTER_BASE+"/id/:id/favourites", middleware.Jwt(h.AddMovieToFavourites, config.JWTKey))
	router.DELETE(MOVIE_ROUTER_BASE+"/id/:id/favourites", middleware.Jwt(h.RemoveMovieFromFavourites, config.JWTKey))
}

func (h *Handler) GetOne(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		handler.JsonErrorResponse(w, http.StatusBadRequest, "Bad Request")
		return
	}

	movie, err := h.service.GetOne(r.Context(), id)

	if err != nil {
		handler.JsonErrorResponse(w, http.StatusInternalServerError, "Server error")
		return
	}

	if movie == nil {
		handler.JsonErrorResponse(w, http.StatusNotFound, "Resource not found")
		return
	}

	categoryResources := make([]category.CategoryResource, 0, len(movie.Categories))

	for _, c := range movie.Categories {
		categoryResources = append(categoryResources, category.CategoryResource{
			Id:          c.Id,
			Name:        c.Name,
			Description: c.Description,
		})
	}

	movieResource := MovieResource{
		movie.Id,
		movie.Name,
		movie.Plot,
		categoryResources,
	}

	handler.JsonSuccessResponse(w, http.StatusOK, handler.NewAPIResponse(movieResource))
}

func (h *Handler) GetByCategoryId(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	categoryId, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		handler.JsonErrorResponse(w, http.StatusBadRequest, "Bad Request")
		return
	}

	movies, err := h.service.GetByCategoryId(r.Context(), categoryId)

	if err != nil {
		handler.JsonErrorResponse(w, http.StatusInternalServerError, "Server error")
		return
	}

	movieResources := make([]MovieResource, 0, len(movies))

	for _, m := range movies {
		movieResources = append(movieResources, MovieResource{
			Id:         m.Id,
			Name:       m.Name,
			Plot:       m.Plot,
			Categories: nil,
		})
	}

	handler.JsonSuccessResponse(w, http.StatusOK, handler.NewAPIResponse(movieResources))
}

type CreateMovieRequest struct {
	Name       string `validate:"required"`
	Plot       string `validate:"required"`
	Categories []int  `validate:"required,unique"`
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	createMovieRequest := CreateMovieRequest{}

	err := json.NewDecoder(r.Body).Decode(&createMovieRequest)
	if err != nil {
		handler.JsonErrorResponse(w, http.StatusBadRequest, "Data is not valid")
		return
	}

	validationError := validator.New().Struct(createMovieRequest)

	if validationError != nil {
		handler.JsonErrorResponse(w, http.StatusBadRequest, "Data is not valid")
		return
	}

	movie, err := h.service.Create(r.Context(), createMovieRequest)

	if err != nil {
		handler.JsonErrorResponse(w, http.StatusInternalServerError, "Server error")
		return
	}

	handler.JsonSuccessResponse(w, http.StatusOK, handler.NewAPIResponse(movie))
}

type UpdateMovieRequest struct {
	Name string `validate:"required"`
	Plot string `validate:"required"`
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		handler.JsonErrorResponse(w, http.StatusBadRequest, "Bad Request")
		return
	}

	updateMovieRequest := UpdateMovieRequest{}

	err = json.NewDecoder(r.Body).Decode(&updateMovieRequest)
	if err != nil {
		handler.JsonErrorResponse(w, http.StatusBadRequest, "Data is not valid")
		return
	}

	validationError := validator.New().Struct(updateMovieRequest)

	if validationError != nil {
		handler.JsonErrorResponse(w, http.StatusBadRequest, "Data is not valid")
		return
	}

	movie, err := h.service.Update(r.Context(), id, updateMovieRequest)

	if err != nil {
		fmt.Println(err)
		handler.JsonErrorResponse(w, http.StatusInternalServerError, "Server error")
		return
	}

	handler.JsonSuccessResponse(w, http.StatusOK, handler.NewAPIResponse(movie))
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		handler.JsonErrorResponse(w, http.StatusBadRequest, "Bad Request")
		return
	}

	movie, err := h.service.Delete(r.Context(), id)

	if err != nil {
		handler.JsonErrorResponse(w, http.StatusInternalServerError, "Server error")
		return
	}

	handler.JsonSuccessResponse(w, http.StatusOK, handler.NewAPIResponse(movie))
}

func (h *Handler) GetFavourites(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userId := r.Context().Value(middleware.ContextUserId)

	movies, err := h.service.GetFavourites(r.Context(), userId.(int))

	if err != nil {
		handler.JsonErrorResponse(w, http.StatusInternalServerError, "Server error")
		return
	}

	movieResources := make([]MovieResource, 0, len(movies))

	for _, m := range movies {
		categoryResources := make([]category.CategoryResource, 0, len(m.Categories))

		for _, c := range m.Categories {
			categoryResources = append(categoryResources, category.CategoryResource{
				Id:   c.Id,
				Name: c.Name,
			})
		}

		movieResources = append(movieResources, MovieResource{
			Id:         m.Id,
			Name:       m.Name,
			Plot:       m.Plot,
			Categories: categoryResources,
		})
	}

	handler.JsonSuccessResponse(w, http.StatusOK, handler.NewAPIResponse(movieResources))
}

func (h *Handler) AddMovieToFavourites(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	movieId, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		handler.JsonErrorResponse(w, http.StatusBadRequest, "Bad Request")
		return
	}

	userId := r.Context().Value(middleware.ContextUserId)

	res, err := h.service.AddMovieToFavourites(r.Context(), userId.(int), movieId)

	if err != nil {
		handler.JsonErrorResponse(w, http.StatusInternalServerError, "Server error")
		return
	}

	if err == nil && !res {
		handler.JsonErrorResponse(w, http.StatusBadRequest, "Movie is already in favourites")
		return
	}

	handler.JsonSuccessResponse(w, http.StatusOK, handler.NewAPIResponse(res))
}

func (h *Handler) RemoveMovieFromFavourites(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	movieId, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		handler.JsonErrorResponse(w, http.StatusBadRequest, "Bad Request")
		return
	}

	userId := r.Context().Value(middleware.ContextUserId)

	res, err := h.service.RemoveMovieFromFavourites(r.Context(), userId.(int), movieId)

	if err != nil {
		handler.JsonErrorResponse(w, http.StatusInternalServerError, "Server error")
		return
	}

	handler.JsonSuccessResponse(w, http.StatusOK, handler.NewAPIResponse(res))
}
