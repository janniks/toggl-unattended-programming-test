package main

import (
	"github.com/gin-gonic/gin"
	"janniks.com/toggl/initial/api"
	"janniks.com/toggl/initial/model"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	r := gin.Default()

	// Use file database for initial project (This can easily be changed later without a lot of code changes)
	db, err := gorm.Open("sqlite3", "initial.db")
	if err != nil {
		panic("Failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&model.Deck{})

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	// Deck API routes
	r.POST("/deck", api.CreateDeck())
	r.GET("/deck/:deck_id", api.OpenDeck())
	r.GET("/deck/:deck_id/cards", api.Draw())

	r.Run(":8080")
}
