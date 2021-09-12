package main

import (
	"log"

	"github.com/abba5/sample-user/models"
	"github.com/abba5/sample-user/pkg/routes"
	"github.com/abba5/sample-user/utils"
	"github.com/abba5/sample-user/utils/config"
)

func main() {

	config, err := config.InitConfig()
	if err != nil {
		log.Panic(err.Error())
	}

	router := routes.InitRouter()

	utils.StartServerWithGracefulShutdown(models.Server{
		Router: router,
		Config: config,
	})
}
