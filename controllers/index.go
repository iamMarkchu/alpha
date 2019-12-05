package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/iamMarkchu/alpha/lib/config"
	"net/http"
)

type IndexController struct {
	BaseController
}

// router / [get]
func (c *IndexController) Index(ctx *gin.Context) {
	configInstance := config.GetConfigInstance()
	result := configInstance.Get("articles")
	c.Render(http.StatusOK, "index/index.htm", gin.H{
		"title": "index",
		"result": result,
	}, ctx)
}
