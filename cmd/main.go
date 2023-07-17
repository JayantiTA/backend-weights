package main

import (
	"log"
	"os"

	"github.com/JayantiTA/backend-weights/config"
	route "github.com/JayantiTA/backend-weights/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := SetupRouter()

	log.Fatal(router.Run(":" + os.Getenv("GO_PORT")))
}

func SetupRouter() *gin.Engine {
	db := config.Connection()

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"*"},
		AllowWildcard: true,
	}))

	route.InitWeightRoutes(db, router)

	return router
}
