package controllers

import (
    "github.com/revel/revel"
    "github.com/cncgl/revel-gorm-todo/app/utils"
    "net/http"
//    "log"
)

type TodoController struct {
    *revel.Controller
}

// Struct for Success Json
type Response struct {
	Result interface{} `json:"results"`
}

// Struct for Error Json
type ErrorResponse struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
}

// Bind Request Parameters to specify interface
func (c *TodoController) BindParams(s interface{}) error {
    return utils.JsonDecode(c.Request.Body, s)
}

// return Bad Request error
func (c *TodoController) HandleBadRequestError(s string) revel.Result {
    c.Response.Status = http.StatusBadRequest
    r := ErrorResponse{c.Response.Status, s}
    return c.RenderJson(r)
}

// return Found Error
func (c *TodoController) HandleNotFoundError(s string) revel.Result {
    c.Response.Status = http.StatusNotFound
    r := ErrorResponse{c.Response.Status, s}
    return c.RenderJson(r)
}

// return Internal Server Error
func (c *TodoController) HandleInternalServerError(s string) revel.Result {
    c.Response.Status = http.StatusInternalServerError
    r := ErrorResponse{c.Response.Status, s}
    return c.RenderJson(r)
}
