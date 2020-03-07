package main

import (
	"github.com/gin-gonic/gin"
	"janniks.com/toggl/initial/api"
	m "janniks.com/toggl/initial/middleware"
	"janniks.com/toggl/initial/model"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	// Use file database for initial project (This can easily be changed later without a lot of code changes)
	db, err := gorm.Open("sqlite3", "initial.db")
	if err != nil {
		panic("Failed to connect database")
	}
	defer db.Close()

	r := SetupRouter(db)
	r.Run("0.0.0.0:80")
}

func SetupRouter(db *gorm.DB) *gin.Engine {
	db.AutoMigrate(&model.Deck{})

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Deck API routes
	r.POST("/deck", m.UseDatabase(db), api.CreateDeck())
	r.GET("/deck/:deck_id", m.UseDatabase(db), api.OpenDeck())
	r.GET("/deck/:deck_id/draw", m.UseDatabase(db), api.Draw())

	return r
}
