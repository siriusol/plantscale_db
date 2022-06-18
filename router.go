package main

import (
	"github.com/gin-gonic/gin"
	"github.com/siriusol/plantscale_db/handler/open"
)

func register(r *gin.Engine) {
	openGroup := r.Group("/open")
	{
		openGroup.GET("/ping", open.Ping)
		openGroup.GET("/beat", open.LatestHeartBeat)
		openGroup.POST("/beat/emit", open.EmitHeartBeat)
	}
}
