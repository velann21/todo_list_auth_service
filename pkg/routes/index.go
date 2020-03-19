package routes

import (
	"github.com/gorilla/mux"
	"github.com/todo_list_auth_service/pkg/controller"
)

func RoutesIntialize(indexRoute *mux.Router){
	indexRoute.HandleFunc("/newtoken", controller.NewTokenController).Methods("POST")
	indexRoute.HandleFunc("/authenticate", controller.AuthenticateController).Methods("GET")
}