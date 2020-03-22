package routes

import (
	"github.com/gorilla/mux"
	"github.com/todo_list_auth_service/pkg/controller"
	"github.com/todo_list_auth_service/pkg/service"
)

func RoutesIntialize(indexRoute *mux.Router){
	indexRoute.HandleFunc("/newtoken", newController().NewTokenController).Methods("POST")
	indexRoute.HandleFunc("/authenticate", newController().AuthenticateController).Methods("GET")
}

func newController() controller.Controller{
	controllerObj := controller.Controller{Service:&service.AuthServiceStruct{}}
	return controllerObj
}