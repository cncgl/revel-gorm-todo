package controllers

import (
    "github.com/revel/revel"
    "net/http"
)

// Struct for Success Json
type Response struct {
	Result interface{} `json:"results"`
}

// Struct for Error Json
type ErrorResponse struct {
    Code int `json:"code"`
    Message string `json:"message"`
}

// return Internal Server Error
func (c TodoController) HandleInternalServerError(s string) revel.Result {
    c.Response.Status = http.StatusInternalServerError
    r := ErrorResponse{c.Response.Status, s}
    return c.RenderJson(r)
}
