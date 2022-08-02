package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/franso/restaurant-management/database"
	"github.com/franso/restaurant-management/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	// "gopkg.in/mgo.v2/bson"
	"go.mongodb.org/mongo-driver/bson"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func GetFoods() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func GetFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		foodId := c.Params.ByName("food_id")
		var food models.Food

		err := foodCollection.FindOne(ctx, bson.M{"food_id": foodId}).Decode(&food)
		defer cancel()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while fetching the food item"})

		}

		c.JSON(http.StatusOK, food)

	}
}

func CreateFood() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func UpdateFood() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func round(num float64) int {

}

func toFixed(num float64, precision int) float64 {}
