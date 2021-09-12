package main

import (
	"github.com/abba5/sample-user/models"
	"github.com/abba5/sample-user/pkg/routes"
	"github.com/abba5/sample-user/utils"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	routes.InitPublicRoutes(router)

	utils.StartServerWithGracefulShutdown(models.Server{
		Router: router,
	})
}
