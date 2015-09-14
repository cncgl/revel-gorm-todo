package controllers

import (
    _ "github.com/lib/pq"
    "github.com/jinzhu/gorm"
    "github.com/revel/revel"
    "github.com/cncgl/revel-gorm-todo/app/models"
    "log"
)

var DB *gorm.DB

func InitDB() {
    db, err := gorm.Open("postgres", "user=postgres password=postgres dbname=hello_phoenix_dev sslmode=disable")
    if err != nil {
        log.Panicf("Failed to connect to database: %v\n", err)
    }
    db.DB()
    db.AutoMigrate(&models.Todo{})
    DB = &db
}

func dbInfoString() string {
    s, b := revel.Config.String("db.info")
    if !b {
        log.Panicf("database info not found")
    }
    return s
}

