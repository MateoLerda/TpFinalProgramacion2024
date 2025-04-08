package models

import (
	"Status418/go/enums"
)

type Filter struct {
	Aproximation string         `bson:"filter_aproximation"`
	Moment       enums.Moment   `bson:"filter_moment"`
	Type         enums.FoodType `bson:"filter_type"`
	All          bool           `bson:"filter_all"`	
	Year		 int			`bson:"filter_year"`	
}

