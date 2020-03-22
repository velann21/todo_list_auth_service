package dependency_manager

import (
	"github.com/todo_list_auth_service/pkg/service"
)

const(
	AUTHSERVICE = "AuthService"
)

func NewService(objectType string)service.AuthServiceInterface{
	if objectType == AUTHSERVICE{
		return &service.AuthServiceStruct{}
	}
	return nil
}

