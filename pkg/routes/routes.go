package routes

import "github.com/gorilla/mux"

func InitRouter() *mux.Router {
	router := mux.NewRouter()
	initPublicRoutes(router)
	return router
}
