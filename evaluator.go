package main

import "github.com/gin-gonic/gin"

func main()  {
	router := gin.Default();

	submit := router.Group("/submit")
	{
		submit.GET("/eval/:fileName" , eval)
	}

	router.Run(":7070")
}