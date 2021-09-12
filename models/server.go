package models

import "github.com/gorilla/mux"

type Server struct {
	Router *mux.Router
}
