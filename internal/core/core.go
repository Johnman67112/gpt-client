package core

import (
	"net/http"
	"os"

	"github.com/Johnman67112/gpt-client/internal/domain"
	"github.com/Johnman67112/gpt-client/internal/infra"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-retryablehttp"
)

func ChatCrudeText(ctx *gin.Context) {
	var chatRequest domain.ChatRequest

	if err := ctx.ShouldBindJSON(&chatRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	client := infra.NewClient(&infra.Config{
		Endpoint: os.Getenv("CHATGPT_URL"),
		ApiKey:   os.Getenv("CHATGPT_APIKEY"),
	}, retryablehttp.NewClient())

	response, err := client.ChatRequest(ctx, chatRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, response)
}

func ChatText(ctx *gin.Context) {
	var chatRequest domain.ChatRequest

	if err := ctx.ShouldBindJSON(&chatRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	client := infra.NewClient(&infra.Config{
		Endpoint: os.Getenv("CHATGPT_URL"),
		ApiKey:   os.Getenv("CHATGPT_APIKEY"),
	}, retryablehttp.NewClient())

	response, err := client.ChatRequest(ctx, chatRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
	}

	parsedResponse := domain.ParsedResponse{
		Message: response.Choices[0].Message.Content,
	}

	ctx.JSON(http.StatusOK, parsedResponse)
}
