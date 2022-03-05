package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func validateX(x string) gin.HandlerFunc {
	validateX := func(next gin.HandlerFunc) gin.HandlerFunc {
		return func(ctx *gin.Context) {
			id := ctx.Query(x)
			if id == "" {
				ctx.AbortWithStatusJSON(http.StatusBadRequest, "Parameter " + x +" is required")
				return
			}
			next(ctx)
		}
	}
	return chainMiddleWareWithDummy(validateX)
}

func processOpRequest(ctx *gin.Context, t int) string {
	switch(t) {
	case 1:
		return ctx.Query("fileName")
	case 2:
		return ctx.Query("ID")
	case 3:
		return ctx.Query("fileName")
	case 4:
		return ctx.Query("problemID")
	}
	return ""
}
