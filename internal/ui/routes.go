package ui

import (
	"github.com/Johnman67112/gpt-client/internal/core"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()

	r.POST("/gpt", core.ChatText)

	r.Run()
}
