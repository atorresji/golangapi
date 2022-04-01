package models

import "net/http"

type Response struct {
	Status         int         `json:"status"`
	Data           interface{} `json:"data"`
	Message        string      `json:"message"`
	contenType     string
	responseWriter http.ResponseWriter
}

func CreateDefaultResponse(rw http.ResponseWriter) Response {
	return Response{
		Status:         http.StatusOK,
		responseWriter: rw,
		contenType:     "application/json",
	}
}
