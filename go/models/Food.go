package models

import (
	"Status418/go/enums"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Food struct {
	Code            primitive.ObjectID `bson:"_id,omitempty"`
	Type            enums.FoodType     `bson:"type"`
	Moments         []enums.Moment     `bson:"moments"`
	Name            string             `bson:"name"`
	UnitPrice       float64            `bson:"unit_price"`
	CurrentQuantity int                `bson:"current_quantity"`
	MinimumQuantity int                `bson:"minimum_quantity"`
	CreationDate    string             `bson:"creation_date"`
	UpdateDate      string             `bson:"update_date"`
	UserCode        string             `bson:"user_code"`
}
