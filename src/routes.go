package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func eval(ctx *gin.Context) {
	stdout, stderr, err := evaluate(processEvaluationRequest(ctx))
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"stdout": stdout, "stderr": stderr, "error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"stdout": stdout, "stderr": stderr, "error": ""})
	}
}

func oproute(p int) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		err := op(processOpRequest(ctx, p), p)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"error": ""})
		}
	}
}
