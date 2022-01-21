package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func eval(ctx *gin.Context) {

	// Get id and lang from request
	id, lang := ctx.Query("id"), ctx.Query("lang")

	// Get file paths from request
	codeFile, inputFile, outputFile := getPaths(id, lang)

	message, err := execute(codeFile, inputFile, outputFile, lang)
	if err != nil {
		message = "Failed to execute"
	}

	// Return JSON with status
	ctx.JSON(http.StatusOK, gin.H{"message": message})
}
