package controllers

import (
    "github.com/revel/revel"
    "github.com/cncgl/go-todo/app/models"
//    "gopkg.in/validator.v2"
)

type TodoController struct {
    *revel.Controller

    // TODO:
    //db.Transactional
}

func (c TodoController) Index() revel.Result {
    // TODO:pending
    //rows, _ := c.Txn.Query("select * from todos")
    // convert database rows to your entity

    todos := []models.Todo{}
    if err := DB.Order("id asc").Find(&todos).Error; err != nil {
        return c.HandleInternalServerError("Record Find Failure")
    }
    r := Response{todos}
    return c.RenderJson(r)
}

func (c TodoController) Show() revel.Result {
    r := Response{"show"}
    return c.RenderJson(r)
}

func (c TodoController) Create() revel.Result {
    //o.Validate(c.Validation)
    //if c.Validation.HasErrors() {
    // return errors
    //}
    r := Response{"create"}
    return c.RenderJson(r)
}

func (c TodoController) Update() revel.Result {
    r := Response{"update"}
    return c.RenderJson(r)
}

func (c TodoController) Delete() revel.Result {
    r := Response{"delete"}
    return c.RenderJson(r)
}
