package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func eval(ctx *gin.Context)  {

	fileName := ctx.Param("fileName")
	timeLimit := ctx.Query("timelimit")

	codeFile , inputFile , outputFile := getPaths(fileName)

	message, err:= execute(codeFile, inputFile, outputFile, timeLimit);

	if err != nil {
		ctx.JSON(200 , gin.H{
			"message" : "Failed to execute",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
        "message": message,
    })



}