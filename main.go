package main

import (
	"main/constants"
	"main/controllers"
	"main/db"
	"main/middlewares"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	db.Connect()

	gin.SetMode(gin.DebugMode)
	engine := gin.Default()
	engine.Use(middlewares.Cors())
	routes(engine)

	host := os.Getenv("HOST") + ":" + os.Getenv("PORT")

	err = engine.Run(host)
	if err != nil {
		print(err)
		return
	}
}

func routes(r *gin.Engine) {
	r.Static("/static/", "./static/")
	r.Static("/uploads/", "./static/uploads/")

	r.GET("/examples", controllers.Examples)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"status":  false,
			"message": "Route not found",
			"data":    nil,
			"code":    constants.ErrNotFound,
		})
	})
}
