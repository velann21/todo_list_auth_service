package service

import (
	"context"
	"github.com/todo_list_auth_service/pkg/entities/requests"
)

type AuthServiceInterface interface {
	NewTokenService(ctx context.Context, newTokenRequests requests.NewTokenRequestsStruct) (*string,error)
	AuthService(ctx context.Context, authRequests requests.AuthRequestsStruct) error
}
