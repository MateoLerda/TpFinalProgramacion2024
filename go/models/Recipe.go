package models

import (
	"Status418/go/enums"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Recipe struct {
	Id           primitive.ObjectID `bson:"_id,omitempty"`
	Name         string             `bson:"recipe_name"`
	Ingredients  []FoodQuantity     `bson:"recipe_ingredients"`
	Moment       enums.Moment       `bson:"recipe_moment"`
	Description  string             `bson:"recipe_description"`
	CreationDate string             `bson:"creation_date"`
	UpdateDate   string             `bson:"update_date"`
	UserCode     string             `bson:"recipe_usercode"`
}
