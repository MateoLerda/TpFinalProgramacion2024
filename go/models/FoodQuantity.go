package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type FoodQuantity struct {
	FoodCode primitive.ObjectID `bson:"_id"`
	Name     string             `bson:"name"`
	Quantity int                `bson:"quantity"`
}
