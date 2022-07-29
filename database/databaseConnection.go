package database

import "go.mongodb.org/mongo-driver/mongo"

func OpenCollection(client *mongo.Client, connectionName string) *mongo.Collection {}
