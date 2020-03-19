package responses

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status string                   `json:"status"`
	Data   []map[string]interface{} `json:"data"`
	Meta   map[string]interface{}   `json:"meta,omitempty"`
}

func (entity *Response)  NewTokenResposne(token *string){
	responseData := make([]map[string]interface{}, 0)
	data := make(map[string]interface{})
	data["Token"] = token
	responseData = append(responseData, data)
	entity.Data = responseData
	metaData := make(map[string]interface{})
	metaData["message"] = "User Token"
}

// SendResponse send http response
func (entity *Response) SendResponse(rw http.ResponseWriter, statusCode int) {
	rw.Header().Set("Content-Type", "application/json")

	switch statusCode {
	case http.StatusOK:
		rw.WriteHeader(http.StatusOK)
		entity.Status = http.StatusText(http.StatusOK)
	case http.StatusCreated:
		rw.WriteHeader(http.StatusCreated)
		entity.Status = http.StatusText(http.StatusCreated)
	case http.StatusAccepted:
		rw.WriteHeader(http.StatusAccepted)
		entity.Status = http.StatusText(http.StatusAccepted)
	default:
		rw.WriteHeader(http.StatusOK)
		entity.Status = http.StatusText(http.StatusOK)
	}

	// send response
	_ = json.NewEncoder(rw).Encode(entity)

	return
}
