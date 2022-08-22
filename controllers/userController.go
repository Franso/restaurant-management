package controllers

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/franso/restaurant-management/database"
	"github.com/franso/restaurant-management/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		// lets create pagination to avoid sending all data
		recordPerPage, err := strconv.Atoi(c.Query("recordsPerPage"))
		if err != nil || recordPerPage < 1 {
			recordPerPage = 10
		}
		// which page do you want the records
		page, err1 := strconv.Atoi(c.Query("page"))
		if err1 != nil || page < 1 {
			page = 1
		}

		// JS skipping limit logic
		startIndex := (page - 1) * recordPerPage
		startIndex, err = strconv.Atoi(c.Query("startIndex"))

		// aggrefatiion to match the user or query parameters
		// in the match stage, we will not do anything but match all documents
		matchStage := bson.D{{"$match", bson.D{{}}}}
		projectStage := bson.D{
			{"$project", bson.D{
				{"_id", 0},
				{"total_count", 1},
				{"user_items", bson.D{{"$slice", []interface{}{"$data", startIndex, recordPerPage}}}},
			}},
		}

		result, err := userCollection.Aggregate(ctx, mongo.Pipeline{
			matchStage, projectStage})

		defer cancel()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing user items"})

		}

		var allUsers []bson.M
		if err := result.All(ctx, &allUsers); err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, allUsers)
	}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		userId := c.Param("user_id")
		var user models.User

		err := userCollection.FindOne(ctx, bson.M{"user_id": userId}).Decode(&user)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing user items"})

		}
		c.JSON(http.StatusOK, user)
	}
}

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		// convert the json data from postman to something Go understands

		// validate the data based on user struct

		// check if the email is used by another user

		// hash the password

		// check if phone number has already been used by anothe user

		// create some extra details for the user object -> created_at and updated_at

		// generate and refresh token from helper

		// if all ok, then insert the new user into the user collection

		// send statusOk and send the result back
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		// convert the login data from postman which is in json to Go lang readable format

		// find a user with that email and see if the user exists

		// then if user exists, verify the password

		// if all goes weel, generate tokens and update tokens

		// return StatusOK
	}
}

func HashPassword(password string) string {
	var res string
	return res
}

func VerifyPassword(userPassword string, providedPassword string) bool {
	return true
}
