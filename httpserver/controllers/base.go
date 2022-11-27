package controllers

import (
	"github.com/deevarindu/final-project-3/httpserver/views"
	"github.com/gin-gonic/gin"
)

func WriteJsonResponse(ctx *gin.Context, resp *views.Response) {
	ctx.JSON(resp.Status, resp)
}
