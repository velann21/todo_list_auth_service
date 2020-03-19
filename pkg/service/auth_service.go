package service

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/todo_list_auth_service/pkg/entities/requests"
	"github.com/todo_list_auth_service/pkg/helpers"
	"time"
)

const (JWT = "jwt")

type Claims struct {
	Username string `json:"username"`
	Email string `json:"email"`
	Roles []int  `json:"roles"`
	Permissions []int  `json:"permissions"`
	PermissionNames []string `json:"permissionNames"`
	jwt.StandardClaims
}

type AuthServiceStruct struct {

}

var jwtKey = []byte("Siar@123&*423")
func (authStruct *AuthServiceStruct) NewTokenService(ctx context.Context, newTokenRequests requests.NewTokenRequestsStruct) (*string,error) {
	token, err := TokenFactory(JWT).GetToken(newTokenRequests)
	if err!= nil{
		return nil, helpers.SomethingWrong
	}
	return token, nil
}

func (authStruct *AuthServiceStruct) AuthService(ctx context.Context, authRequests requests.AuthRequestsStruct) error {
     err := TokenFactory(JWT).ValidateJWTToken(authRequests.Token)
     if err != nil{
     	return err
	 }
     return nil
}




type Token interface {
	GetToken(authRequests requests.NewTokenRequestsStruct) (*string,error);
	ValidateJWTToken(token string) error
}

type JWTTokenGenarator struct {

}

func (jwt *JWTTokenGenarator) GetToken(authRequests requests.NewTokenRequestsStruct) (*string,error){
	token, err := tokenGenerator(authRequests)
	if err != nil{
		return nil, err
	}
	return token, nil
}

func (jwt *JWTTokenGenarator) ValidateJWTToken(token string) error{
	err := validateJWTToken(token)
	if err != nil{
		return err
	}
	return nil
}

func TokenFactory(tokenType string, ) Token {
	switch tokenType {
	case JWT:
		return &JWTTokenGenarator{}
	default:
		return nil
	}
}


func tokenGenerator(authRequests requests.NewTokenRequestsStruct) (*string,error) {
	expirationTime := time.Now().Add(5 * time.Hour)
	claims := &Claims{
		Username: authRequests.UserName,
		Email: authRequests.EmailID,
		Roles:authRequests.Roles,
		Permissions:authRequests.Permissions,
		PermissionNames:authRequests.PermissionName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return nil, helpers.SomethingWrong
	}
	return &tokenString, nil
}


func validateJWTToken(tok string) (error){
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tok, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return helpers.UnAuthorized
		}

		return  helpers.UnAuthorized
	}
	if !tkn.Valid {
		return helpers.UnAuthorized
	}

	return nil

}


