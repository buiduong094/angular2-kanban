package router

import (
	"github.com/gin-gonic/gin"
	"github.com/octoberstorm/angular2-kanban/server/config"
	"sync"
)

var once sync.Once
var router *gin.Engine

func Get() *gin.Engine {
	once.Do(func() {
		if config.Get().Environment == "Production" {
			gin.SetMode(gin.ReleaseMode)
		}
		router = gin.New()
		router.Use(gin.Logger())
		router.Use(gin.Recovery())
	})
	return router
}
