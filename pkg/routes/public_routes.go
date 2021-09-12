package routes

import (
	"net/http"

	"github.com/abba5/sample-user/app/controllers"
	"github.com/gorilla/mux"
)

func InitPublicRoutes(router *mux.Router) {
	router.HandleFunc("/api/user", controllers.GetUser).Methods(http.MethodGet)
}
