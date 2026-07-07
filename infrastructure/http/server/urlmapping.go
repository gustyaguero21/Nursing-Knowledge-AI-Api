package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UrlMapping(s *gin.Engine) {
	app := s.Group("/api/v1/nursing-knowledge")

	app.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"statuscheck": "ok",
		})
	})
}
