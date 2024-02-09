package core

import (
	"net/http"

	"github.com/Johnman67112/gpt-client/internal/domain"
	"github.com/gin-gonic/gin"
)

func ChatText(ctx *gin.Context) {
	var chatRequest domain.ChatRequest

	if err := ctx.ShouldBindJSON(&chatRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, chatRequest)
}
