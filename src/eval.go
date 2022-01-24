package main

import (
	"net/http"
	"strconv"

	// "strconv"

	"github.com/gin-gonic/gin"
)

func eval(ctx *gin.Context) {

	// Get id and lang from request
	id, lang, timelimit, memorylimit := ctx.Query("id"), ctx.Query("lang"), ctx.Query("timelimit"), ctx.Query("memorylimit")
	if !validateSubmitRequest(id, lang, ctx) {
		return
	}
	// setting default time limit to 1s
	if len(timelimit) == 0 {
		timelimit = strconv.Itoa(DEFAULT_TIME_LIMIT) + SECONDS
	}

	//set default memory limit to 64MB
	if len(memorylimit) == 0 {
		memorylimit = strconv.Itoa(int(DEFAULT_MEMORY_LIMIT))
	} else {
		l := len(memorylimit)
		memorylimit = memorylimit[:l-2] //strip "mb" from parameter
	}

	// Start execution
	message, err := execute(id, lang, timelimit, memorylimit)

	if err != nil {
		message = "Failed to execute"
	}

	// Return JSON with status
	ctx.JSON(http.StatusOK, gin.H{"message": message})
}
