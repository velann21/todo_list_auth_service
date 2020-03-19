package requests

import (
	"encoding/json"
	"github.com/todo_list_auth_service/pkg/helpers"
	"io"
	"log"
	"strings"
)

type NewTokenRequestsStruct struct {
	UserName string `json:"userName"`
	EmailID string `json:"emailID"`
	Roles []int `json:"roles"`
	PermissionName []string `json:"permissionsName"`
	Permissions []int `json:"permissions"`
}
func (newTokenRequestsStruct *NewTokenRequestsStruct) PopulateNewTokenRequestsStruct(body io.ReadCloser) error{
	decoder := json.NewDecoder(body)
	err := decoder.Decode(newTokenRequestsStruct)
	if err != nil {
		return helpers.NotValidRequestBody
	}
	return nil
}

func (newTokenRequestsStruct *NewTokenRequestsStruct) ValidateNewTokenRequestsStruct() error {
	if newTokenRequestsStruct.EmailID == ""{
		return helpers.InvalidRequest
	}
	if len(newTokenRequestsStruct.Permissions) <= 0{
		return helpers.InvalidRequest
	}
	if len(newTokenRequestsStruct.Roles) <= 0{
		return helpers.InvalidRequest
	}
	return nil
}

type AuthRequestsStruct struct {
	Token string `json:"token"`
}
func (authRequestsStruct *AuthRequestsStruct) PopulateAuthRequestsStruct(token string) error{
	if token == ""{

		return helpers.UnAuthorized
	}
	tokenizedStr := strings.Split(token, " ")
	if len(tokenizedStr)<2{
		log.Println("Error in token it can't br empty")
		return helpers.UnAuthorized
	}
	authRequestsStruct.Token = tokenizedStr[1]
	return nil
}

func (authRequestsStruct *AuthRequestsStruct) ValidateAuthRequestsStruct() error {
	if authRequestsStruct.Token == ""{
		return helpers.UnAuthorized
	}
	return nil
}


