package models

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Status      int         `json:"status"`
	Data        interface{} `json:"data"`
	Message     string      `json:"message"`
	contentType string
	respWrite   http.ResponseWriter
}

func CreateDefaultResponse(rw http.ResponseWriter) Response {
	return Response{
		Status:      http.StatusOK,
		respWrite:   rw,
		contentType: "application/json",
	}
}

func (resp *Response) Send() {
	resp.respWrite.Header().Set("Content-Type", resp.contentType)
	resp.respWrite.WriteHeader(resp.Status)

	output, _ := json.Marshal(&resp)
	fmt.Fprintln(resp.respWrite, string(output))
}

func SendData(rw http.ResponseWriter, data interface{}) {
	response := CreateDefaultResponse(rw)
	response.Data = data
	response.Send()
}

func (resp *Response) NoFound() {
	resp.Status = http.StatusNotFound
	resp.Message = "Resource No Found"
}

func SendNoFound(rw http.ResponseWriter) {
	response := CreateDefaultResponse(rw)
	response.NoFound()
	response.Send()
}

func (resp *Response) UnprocessableEntity() {
	resp.Status = http.StatusUnprocessableEntity
	resp.Message = "UnprocessableEntity No Found"
}

func SendUnprocessableEntity(rw http.ResponseWriter) {
	response := CreateDefaultResponse(rw)
	response.UnprocessableEntity()
	response.Send()
}
