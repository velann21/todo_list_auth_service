package main

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/todo_list_auth_service/pkg/routes"
	"log"
	"net/http"
	"time"
)

func main(){
	logrus.SetFormatter(&logrus.JSONFormatter{TimestampFormat:time.RFC3339,})
	r := mux.NewRouter().StrictSlash(false)
	mainRoutes := r.PathPrefix("/api/v1/auth").Subrouter()
	routes.RoutesIntialize(mainRoutes)
	logrus.WithField("EventType", "Bootup").Info("Booting up server at port : "+"8083")
	if err := http.ListenAndServe(":8083", r); err != nil {
		logrus.WithField("EventType", "Server Bootup").WithError(err).Error("Server Bootup Error")
		log.Fatal(err)
	}
}
