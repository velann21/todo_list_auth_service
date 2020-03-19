package controller

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/todo_list_auth_service/pkg/entities/requests"
	"github.com/todo_list_auth_service/pkg/entities/responses"
	dm "github.com/todo_list_auth_service/pkg/service/dependency_manager"
	"net/http"
	"time"
)

func NewTokenController(rw http.ResponseWriter, req *http.Request){
	logrus.WithField("EventType", "NewTokenController").WithField("Action","Request").Info("NewTokenController Start")
	authStruct := requests.NewTokenRequestsStruct{}
	successResponse := responses.Response{}
	serviceObj := dm.NewService(dm.AUTHSERVICE)
	ctx, cancel := context.WithTimeout(req.Context(), time.Second*10)
	defer cancel()
	err := authStruct.PopulateNewTokenRequestsStruct(req.Body)
	if err != nil{
		logrus.WithField("EventType", "NewTokenController").WithError(err).Error("PopulateNewTokenRequestsStruct Failed")
		responses.HandleError(rw, err)
		return
	}
	token, err := serviceObj.NewTokenService(ctx, authStruct)
	if err != nil{
		logrus.WithField("EventType", "NewTokenController").WithError(err).Error("NewTokenService Failed")
		responses.HandleError(rw, err)
		return
	}
	successResponse.NewTokenResposne(token)
	successResponse.SendResponse(rw, http.StatusOK)
	logrus.WithField("EventType", "NewTokenController").WithField("Action","Request").Info("NewTokenController Ends")
	return
}

func AuthenticateController(rw http.ResponseWriter, req *http.Request){
	logrus.WithField("EventType", "AuthenticateController").WithField("Action","Request").Info("AuthenticateController Start")
	authStruct := requests.AuthRequestsStruct{}
	successResponse := responses.Response{}
	requestsHeader := req.Header
	token := requestsHeader.Get("Authorization")
	serviceObj := dm.NewService(dm.AUTHSERVICE)
	ctx, cancel := context.WithTimeout(req.Context(), time.Second*10)
	defer cancel()
	err := authStruct.PopulateAuthRequestsStruct(token)
	if err != nil{
		logrus.WithField("EventType", "AuthenticateController").WithError(err).Error("PopulateAuthRequestsStruct Failed")
		responses.HandleError(rw, err)
		return
	}
	err = authStruct.ValidateAuthRequestsStruct()
	if err != nil{
		logrus.WithField("EventType", "AuthenticateController").WithError(err).Error("ValidateAuthRequestsStruct Failed")
		responses.HandleError(rw, err)
		return
	}
	err = serviceObj.AuthService(ctx, authStruct)
	if err != nil {
		logrus.WithField("EventType", "AuthenticateController").WithError(err).Error("AuthService Failed")
		responses.HandleError(rw, err)
		return
	}
	successResponse.SendResponse(rw, http.StatusOK)
	logrus.WithField("EventType", "AuthenticateController").WithField("Action","Request").Info("AuthenticateController Start")
	return
}
