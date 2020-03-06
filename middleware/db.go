package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func UseDatabase(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}
