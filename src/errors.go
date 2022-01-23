package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// returns false if validation failed
func validateSubmitRequest(id, lang string, ctx *gin.Context) bool {
	if len(id) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error" : "Query 'id' is missing"})
	}else if len(lang) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error" : "Query 'lang' is missing"})
	}else{
		return true
	}
	return false
}