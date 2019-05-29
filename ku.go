package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RouteKu(router *gin.Engine) {
	ku := router.Group("/ku")
	{
		ku.GET("/test", KuTest)
	}
}

func KuTest(c *gin.Context) {
	c.JSON(http.StatusOK, ResultSuccess(nil))
}
