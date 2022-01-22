package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func eval(ctx *gin.Context) {

	// Get id and lang from request
	id, lang := ctx.Query("id"), ctx.Query("lang")

	// Start execution
	message, err := execute(id, lang)
	
	if err != nil {
		message = "Failed to execute"
	}

	// Return JSON with status
	ctx.JSON(http.StatusOK, gin.H{"message": message})
}
