package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	submit := router.Group("")
	{
		submit.GET("/eval",
			validateEvaluationArgs(),
			eval)
		submit.GET("/uploadSubmission",
			validateX("fileName"),
			oproute(1))
		submit.GET("/uploadProblem",
			validateX("fileName"),
			oproute(3))
		submit.GET("/rmSubmission",
			validateX("ID"),
			oproute(2))
		submit.GET("/rmProblem",
			validateX("problemID"),
			oproute(4))
	}
	port := ":7070"
	router.Run(port)
}
