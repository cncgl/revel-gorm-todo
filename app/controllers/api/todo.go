package controllers

import (
    "github.com/revel/revel"
    "github.com/cncgl/revel-gorm-todo/app/models"
    "gopkg.in/validator.v2"
    "log"
)

type ApiTodo struct {
    TodoController
}

func (c ApiTodo) Index() revel.Result {
    todos := []models.Todo{}
    if err := DB.Order("id asc").Find(&todos).Error; err != nil {
        return c.HandleInternalServerError("Record Find Failure")
    }
    r := Response{todos}
    return c.RenderJson(r)
}

func (c ApiTodo) Show(id int) revel.Result {
    todo := &models.Todo{}
    if err := DB.First(&todo, id).Error; err != nil {
        return c.HandleNotFoundError(err.Error())
    }
    r := Response{todo}
    return c.RenderJson(r)
}

func (c ApiTodo) Create() revel.Result {
    todo := &models.Todo{}
    log.Println(todo)
    if err := c.BindParams(todo); err != nil {
        log.Println( err.Error() )
        return c.HandleBadRequestError(err.Error())
    }
    log.Println(todo)
    if err := validator.Validate(todo); err != nil {
        log.Println( err.Error() )
        return c.HandleBadRequestError(err.Error())
    }
    if err := DB.Create(todo).Error; err != nil {
        log.Println( "Record Create Failure" )
        return c.HandleInternalServerError("Record Create Failure")
    }
    r := Response{todo}
    return c.RenderJson(r)
}

func (c ApiTodo) Update(id int) revel.Result {
    todo := &models.Todo{}
    if err := c.BindParams(todo); err != nil {
        return c.HandleBadRequestError(err.Error())
    }
    if err := validator.Validate(todo); err != nil {
        return c.HandleBadRequestError(err.Error())
    }
    log.Println(todo)
    if err := DB.Save(&todo).Error; err != nil {
        return c.HandleInternalServerError("Record Update Failure")
    }
    r := Response{todo}
    return c.RenderJson(r)
}

func (c ApiTodo) Delete(id int) revel.Result {
    todo := models.Todo{}
    if err := DB.First(&todo, id).Error; err != nil {
        return c.HandleNotFoundError(err.Error())
    }
    if err := DB.Delete(&todo).Error; err != nil {
        return c.HandleInternalServerError("Record Delete Failure")
    }
    r := Response{"success"}
    return c.RenderJson(r)
}
