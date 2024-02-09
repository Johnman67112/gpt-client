package ui

import (
	"github.com/Johnman67112/gpt-client/internal/core"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()

	r.POST("/gpt/parsed", core.ChatText)
	r.POST("/gpt", core.ChatCrudeText)

	r.Run()
}
