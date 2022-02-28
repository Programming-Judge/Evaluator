package main

import (
	"github.com/gin-gonic/gin"
	"github.com/Programming-Judge/Evaluator/evaluation/evaluation"
)

// The main function
func main() {
	router := gin.Default()
	submit := router.Group("")
	{
		submit.GET("/eval",
			evaluation.ValidateAll(),
			eval)
		submit.GET("/uploadSubmission")
	}
	port := ":7070"
	router.Run(port)
}
