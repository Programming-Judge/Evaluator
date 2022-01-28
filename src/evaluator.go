package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	submit := router.Group("/submit")
	{
		submit.GET(
			"/eval",
			validateId,
			validateLang,
			validateTimelimit,
			validateMemoryLimit,
			eval)
	}
	port := ":7070"
	router.Run(port)
}
