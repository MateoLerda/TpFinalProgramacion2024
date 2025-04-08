package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Purchase struct {
	Id           primitive.ObjectID `bson:"_id,omitempty"`
	UserCode     string             `bson:"user_code"`
	PurchaseDate string             `bson:"purchase_date"`
	TotalCost    float64            `bson:"total_cost"`
	Foods        []FoodQuantity     `bson:"foods"`
}
