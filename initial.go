package main

import (
	"github.com/gin-gonic/gin"
	"janniks.com/toggl/initial/api"
	"net/http"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.POST("/deck", api.CreateDeck())
	r.GET("/deck/:deck_id", api.OpenDeck())
	r.GET("/deck/:deck_id/cards", api.Draw())

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
