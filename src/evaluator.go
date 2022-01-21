package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	submit := router.Group("/submit")
	{
		submit.GET("/eval", eval)
	}

	// Server serves requests here
	port := ":7070"
	router.Run(port)
}
