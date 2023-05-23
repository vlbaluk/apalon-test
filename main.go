package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/patrickmn/go-cache"
	. "github.com/vlbaluk/apalon-test/controllers"
	. "github.com/vlbaluk/apalon-test/http"
	"github.com/vlbaluk/apalon-test/services"
	. "github.com/vlbaluk/apalon-test/services/cached"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	//Load the .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error: failed to load the env file")
	}

	if os.Getenv("ENV") == "PRODUCTION" {
		gin.SetMode(gin.ReleaseMode)
	}

	//Start the  gin server
	r := gin.Default()

	v1 := r.Group("/")
	{
		c := cache.New(time.Minute, time.Millisecond)
		mathService := &services.MathService{}
		mathCachedService := &MathService{Cache: c, Service: mathService}
		mathController := &MathController{}

		v1.GET("/add", func(c *gin.Context) {
			mathController.Calc(c, mathService.Add, mathCachedService.CachedCalculate, "add")
		})
		v1.GET("/divide", func(c *gin.Context) {
			mathController.Calc(c, mathService.Divide, mathCachedService.CachedCalculate, "divide")
		})
		v1.GET("/substract", func(c *gin.Context) {
			mathController.Calc(c, mathService.Subtract, mathCachedService.CachedCalculate, "subtract")
		})
		v1.GET("/multiply", func(c *gin.Context) {
			mathController.Calc(c, mathService.Multiply, mathCachedService.CachedCalculate, "multiply")
		})
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, &ErrorResponse{Message: "Page is not found"})
	})

	port := os.Getenv("PORT")

	r.Run(":" + port)
}
