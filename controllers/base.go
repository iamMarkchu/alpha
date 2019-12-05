package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/iamMarkchu/alpha/lib/config"
	"net/http"
)

type BaseController struct {

}

func (c *BaseController) Render(code int, name string, h gin.H, ctx *gin.Context)  {
	h["commonData"] = c.CommonData()
	ctx.HTML(http.StatusOK, "index/index.htm", h)
}

func (c *BaseController) CommonData() map[string]interface{} {
	return gin.H{
		"navs": c.Navs(),
	}
}

func (c *BaseController) Navs() interface{} {
	return config.GetConfigInstance().Get("navs")
}
