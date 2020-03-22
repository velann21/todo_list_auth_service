package helpers
import "errors"

var (
	// ErrParamMissing required field missing error
	NotValidRequestBody = errors.New("Invalid Request")
	ErrUserNotFound = errors.New("User Not found")
	InvalidRequest = errors.New("Invalid Request")
	UnAuthorized = errors.New("UnAuthorized")
	SomethingWrong = errors.New("Something went wrong")
)
