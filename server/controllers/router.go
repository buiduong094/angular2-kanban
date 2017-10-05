package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/octoberstorm/angular2-kanban/server/controllers/router"
	"net/http"
)

var Router *gin.Engine

func init() {
	Router = router.Get()
	Router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})
}
