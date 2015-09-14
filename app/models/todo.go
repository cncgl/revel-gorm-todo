package models

import (
    "github.com/revel/revel"
    "time"
)

type Todo struct {
    Id int64                `gorm:"primary_key" json:"id"`
    Status bool             `json:"status"`
    Title string            `sql:"size:255 json:"title" vlidate:"min=1,max=255"`
    Inserted_at time.Time   `json:"inserted_at"`
    Updated_at time.Time    `json:"updated_at"`
}

func (todo Todo) Validate(v revel.Validation) {
    v.Required(todo.Title)
    v.MaxSize(todo.Title, 50)
}