package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/iamMarkchu/alpha/lib/orm"
	"net/http"
)

type IndexController struct {

}

// router / [get]
func (c *IndexController) Index(ctx *gin.Context)  {
	result, _ := orm.NewEngine().QueryString("SELECT * FROM articles")
	ctx.HTML(http.StatusOK, "index/index.htm", gin.H{
		"title": "homepage",
		"result": result,
	})
}
