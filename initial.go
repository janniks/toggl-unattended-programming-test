package main

import (
	"github.com/gin-gonic/gin"
	"janniks.com/toggl/initial/api"
	m "janniks.com/toggl/initial/middleware"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	// Use file database for initial project (This can easily be changed later without a lot of code changes)
	db, err := gorm.Open("sqlite3", "initial.db")
	if err != nil {
		panic("Failed to connect database")
	}
	defer db.Close()

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.POST("/deck", m.UseDatabase(db), api.CreateDeck())
	r.GET("/deck/:deck_id", m.UseDatabase(db), api.OpenDeck())
	r.GET("/deck/:deck_id/cards", m.UseDatabase(db), api.Draw())

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
