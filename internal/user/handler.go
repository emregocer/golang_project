// package user provides user related functionality.
package user

import (
	"net/http"

	"github.com/emregocer/golang_project/config"
	"github.com/emregocer/golang_project/internal/handler"
	"github.com/emregocer/golang_project/internal/service"
	"github.com/go-playground/validator/v10"

	"github.com/julienschmidt/httprouter"
)

const USER_ROUTER_BASE = "/users"

type Handler struct {
	service    *Service
	jwtService *service.JwtService
}

func NewHandler(service *Service, jwtService *service.JwtService) *Handler {
	return &Handler{service: service, jwtService: jwtService}
}

func (u *Handler) RegisterUserHandlers(router *httprouter.Router, config config.Config) {
	router.POST(USER_ROUTER_BASE+"/login", u.Login)
	router.POST(USER_ROUTER_BASE+"/register", u.Register)
}

type UserRegisterRequest struct {
	Username string `validate:"required"`
	Password string `validate:"required"`
	Email    string `validate:"required,email"`
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := r.ParseForm()
	if err != nil {
		handler.JsonErrorResponse(w, http.StatusInternalServerError, "Server error")
		return
	}

	userRegisterRequest := UserRegisterRequest{Username: r.PostFormValue("username"), Password: r.PostFormValue("password"), Email: r.PostFormValue("email")}

	validationError := validator.New().Struct(userRegisterRequest)

	if validationError != nil {
		handler.JsonErrorResponse(w, http.StatusBadRequest, "Data is not valid")
		return
	}

	user, err := h.service.Create(r.Context(), r.PostFormValue("username"), r.PostFormValue("password"), r.PostFormValue("email"))

	if err != nil {
		handler.JsonErrorResponse(w, http.StatusInternalServerError, "Server error")
		return
	}

	if user.Id == 0 && err == nil {
		handler.JsonErrorResponse(w, http.StatusConflict, "User exists")
		return
	}

	handler.JsonSuccessResponse(w, http.StatusOK, handler.Message(http.StatusOK, "success"))
}

type UserLoginRequest struct {
	Username string `validate:"required"`
	Password string `validate:"required"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := r.ParseForm()
	if err != nil {
		handler.JsonErrorResponse(w, http.StatusInternalServerError, "Server error")
		return
	}

	username := r.PostFormValue("username")
	password := r.PostFormValue("password")

	userRegisterRequest := UserLoginRequest{Username: username, Password: password}

	validationError := validator.New().Struct(userRegisterRequest)

	if validationError != nil {
		handler.JsonErrorResponse(w, http.StatusBadRequest, "Data is not valid")
		return
	}

	user, err := h.service.Login(r.Context(), username, password)
	if user == nil {
		handler.JsonErrorResponse(w, http.StatusUnauthorized, "Login failed")
		return
	}

	if err != nil {
		handler.JsonErrorResponse(w, http.StatusInternalServerError, "Server error")
		return
	}

	_, token, err := h.jwtService.GenerateToken(user.Id)
	if err != nil {
		handler.JsonErrorResponse(w, http.StatusInternalServerError, "Server error")
		return
	}

	handler.JsonSuccessResponse(w, http.StatusOK, handler.NewAPIResponse(token))
}
