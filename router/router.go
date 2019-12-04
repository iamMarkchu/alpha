package router

import (
	"github.com/gin-gonic/gin"
	"github.com/iamMarkchu/alpha/controllers"
)

var engine *gin.Engine

func init()  {
	engine = gin.Default()
	engine.LoadHTMLGlob("templates/**/*")
	SetupRouter()
}

func SetupRouter()  {
	engine.Static("/public", "./public")
	index := controllers.IndexController{}
	engine.GET("/", index.Index)
}

func NewEngine() *gin.Engine {
	return engine
}
