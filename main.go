package main

import (
	"github.com/gin-gonic/gin"
	"translate-svc/core/db"
	"translate-svc/core/logger"
	"translate-svc/routes"
)

func main() {
	// Initialize Logger
	logger.Init()

	// Initialize DB connection
	db.InitDB()
	defer func() {
		err := db.DB.Close()
		if err != nil {
			logger.Log.Println(err)
		}
	}()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	routes.RegisterRoutes(r)

	// listen and serve on 0.0.0.0:8080
	err := r.Run()

	if err != nil {
		logger.Log.Error(err)
	}
}
