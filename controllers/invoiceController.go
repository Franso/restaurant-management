package controllers

import (
	"time"

	"github.com/franso/restaurant-management/database"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type InvoiceViewFormat struct {
	Invoice_id       string
	Payment_method   string
	Order_id         string
	Payment_status   *string
	Payment_due      interface{}
	Table_number     interface{}
	Payment_due_date time.Time
	Order_details    interface{}
}

var invoiceCollection *mongo.Collection = database.OpenCollection(database.Client, "invoice")

func GetInvoices() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func GetInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func CreateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func UpdateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {}
}
