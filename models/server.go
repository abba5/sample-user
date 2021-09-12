package models

import (
	"github.com/abba5/sample-user/utils/config"
	"github.com/gorilla/mux"
)

type Server struct {
	Router *mux.Router
	Config *config.Config
}
