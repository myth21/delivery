package main

import (
	"delivery/entity/model"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"os"
)

// Init is invoked before main()
func init() {
	// Loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	route := gin.Default()
	route.GET("/deliveries/", getDeliveriesHandler)
	route.POST("/deliveries/", createDeliveryHandler)

	Host := os.Getenv("HOST")
	Port := os.Getenv("PORT")
	route.Run(Host + ":" + Port)
}

func getDeliveriesHandler(context *gin.Context) {

	deliveries, err := GetDeliveries()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"count": len(deliveries),
		"data":  deliveries,
	})
}

func createDeliveryHandler(context *gin.Context) {

	// TODO validation
	// see docs for context.Request

	var delivery model.Delivery

	err := context.ShouldBindJSON(&delivery)

	if err != nil {
		// TODO logger
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	delivery.SetPendingStatus()

	deliveryFromRepository, err := CreateDelivery(delivery)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, deliveryFromRepository)
}
