package controller

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/todo_list_auth_service/pkg/entities/requests"
	"github.com/todo_list_auth_service/pkg/entities/responses"
	"github.com/todo_list_auth_service/pkg/service"
	"net/http"
	"time"
)

type Controller struct {
	Service service.AuthServiceInterface
}

func (controller Controller) NewTokenController(rw http.ResponseWriter, req *http.Request){
	logrus.WithField("EventType", "NewTokenController").WithField("Action","Request").Info("NewTokenController Start")
	authStruct := requests.NewTokenRequestsStruct{}
	successResponse := responses.Response{}
	ctx, cancel := context.WithTimeout(req.Context(), time.Second*10)
	defer cancel()
	err := authStruct.PopulateNewTokenRequestsStruct(req.Body)
	if err != nil{
		logrus.WithField("EventType", "NewTokenController").WithError(err).Error("PopulateNewTokenRequestsStruct Failed")
		responses.HandleError(rw, err)
		return
	}
	token, err := controller.Service.NewTokenService(ctx, authStruct)
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

func (controller Controller) AuthenticateController(rw http.ResponseWriter, req *http.Request){
	logrus.WithField("EventType", "AuthenticateController").WithField("Action","Request").Info("AuthenticateController Start")
	authStruct := requests.AuthRequestsStruct{}
	successResponse := responses.Response{}
	requestsHeader := req.Header
	token := requestsHeader.Get("Authorization")
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
	err = controller.Service.AuthService(ctx, authStruct)
	if err != nil {
		logrus.WithField("EventType", "AuthenticateController").WithError(err).Error("AuthService Failed")
		responses.HandleError(rw, err)
		return
	}
	successResponse.SendResponse(rw, http.StatusOK)
	logrus.WithField("EventType", "AuthenticateController").WithField("Action","Request").Info("AuthenticateController Start")
	return
}
