package repositories

import (
	"Status418/go/enums"
	"Status418/go/models"
	"context"
	"errors"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type FoodRepositoryInterface interface {
	GetAll(userCode string, filters models.Filter) ([]models.Food, error)
	GetByCode(foodCode primitive.ObjectID, userCode string) (models.Food, error)
	Create(newFood models.Food) (*mongo.InsertOneResult, error)
	Update(updateFood models.Food, cook bool) (*mongo.UpdateResult, error)
	Delete(foodcode primitive.ObjectID) (*mongo.DeleteResult, error)
}

type FoodRepository struct {
	db DB
}

func NewFoodRepository(db DB) *FoodRepository {
	return &FoodRepository{
		db: db,
	}
}

func (foodRepository FoodRepository) GetAll(userCode string, filters models.Filter) ([]models.Food, error) {
	DBNAME := os.Getenv("DB_NAME")
	filter := bson.M{
		"user_code": userCode,
	}
	if filters.Aproximation != "" {
		filter["name"] = bson.M{
			"$regex":   filters.Aproximation,
			"$options": "i",
		}
	}
	if filters.Type != enums.InvalidFoodType {
		filter["type"] = filters.Type
	}
	if !filters.All {
		filter = bson.M{
			"$and": bson.A{
				filter,
				bson.M{
					"$expr": bson.M{
						"$lt": bson.A{"$current_quantity", "$minimum_quantity"},
					},
				},
			},
		}
	}

	cursor, err := foodRepository.db.GetClient().Database(DBNAME).Collection("Foods").Find(context.TODO(), filter)

	if err != nil {
		return nil, err
	}

	var foods []models.Food
	cursor.All(context.TODO(), &foods)

	if len(foods) == 0 {
		err = errors.New("nocontent")
		return nil, err
	}
	return foods, nil
}

func (foodRepository FoodRepository) GetByCode(foodCode primitive.ObjectID, userCode string) (models.Food, error) {
	DBNAME := os.Getenv("DB_NAME")

	filter := bson.M{
		"_id":       foodCode,
		"user_code": userCode,
	}
	data := foodRepository.db.GetClient().Database(DBNAME).Collection("Foods").FindOne(context.TODO(), filter)
	var food models.Food
	err := data.Decode(&food)

	if err == mongo.ErrNoDocuments {
		err = errors.New("could not find the food with the given code ")
	}
	return food, err
}

func (foodRepository FoodRepository) Create(food models.Food) (*mongo.InsertOneResult, error) {
	DBNAME := os.Getenv("DB_NAME")
	res, err := foodRepository.db.GetClient().Database(DBNAME).Collection("Foods").InsertOne(context.TODO(), food)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (foodRepository FoodRepository) Update(food models.Food, cook bool) (*mongo.UpdateResult, error) {
	DBNAME := os.Getenv("DB_NAME")
	food.UpdateDate = time.Now().String()
	filter := bson.M{
		"_id": food.Code,
	}
	update := toBSONUpdate(food, cook)
	res, err := foodRepository.db.GetClient().Database(DBNAME).Collection("Foods").UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func toBSONUpdate(food models.Food, cook bool) bson.M {
	update := bson.M{"$set": bson.M{}}
	if food.Type != 0 {
		update["$set"].(bson.M)["type"] = food.Type
	}
	if food.Name != "" {
		update["$set"].(bson.M)["name"] = food.Name
	}
	if food.UnitPrice != 0 {
		update["$set"].(bson.M)["unit_price"] = food.UnitPrice
	}
	if cook {
		update["$inc"] = bson.M{"current_quantity": food.CurrentQuantity}
	} else {
		update["$set"].(bson.M)["current_quantity"] = food.CurrentQuantity
	}
	if food.MinimumQuantity != 0 {
		update["$set"].(bson.M)["minimum_quantity"] = food.MinimumQuantity
	}
	update["$set"].(bson.M)["update_date"] = food.UpdateDate
	return update
}

func (foodRepository FoodRepository) Delete(foodCode primitive.ObjectID) (*mongo.DeleteResult, error) {
	DBNAME := os.Getenv("DB_NAME")
	filter := bson.M{"_id": foodCode}
	res, err := foodRepository.db.GetClient().Database(DBNAME).Collection("Foods").DeleteOne(context.TODO(), filter)
	if res.DeletedCount == 0 {
		err = errors.New("notfound")
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return res, nil
}
