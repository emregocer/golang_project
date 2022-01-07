// package server provides the router initialization and listening functionality.
package server

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/emregocer/golang_project/config"
	"github.com/emregocer/golang_project/internal/category"
	"github.com/emregocer/golang_project/internal/movie"
	"github.com/emregocer/golang_project/internal/service"
	"github.com/emregocer/golang_project/internal/user"
	"github.com/jmoiron/sqlx"
	"github.com/julienschmidt/httprouter"
)

type Server struct {
	Router *httprouter.Router
	DB     *sqlx.DB
	Config config.Config
}

// Initialize registers the router on the router.
func (server *Server) Initialize() {
	router := httprouter.New()

	server.Router = router

	// register user handlers to the router
	userHandler := user.NewHandler(user.NewService(user.NewRepo(server.DB)), service.NewJwtService(server.Config))
	userHandler.RegisterUserHandlers(router, server.Config)

	categoryHandler := category.NewHandler(category.NewService(category.NewRepo(server.DB)))
	categoryHandler.RegisterHandlers(router, server.Config)

	movieHandler := movie.NewHandler(movie.NewService(movie.NewRepo(server.DB)))
	movieHandler.RegisterHandlers(router, server.Config)
}

// Listen starts up a http server and listens on the given port.
//
// Parameters:
//   - `port` : int
func (server *Server) Listen(port int) {
	err := http.ListenAndServe(":"+strconv.Itoa(port), server.Router)
	if err != nil {
		log.Fatal("Could not start the http server on the given port: " + strconv.Itoa(port))
		os.Exit(-1)
	}
}
