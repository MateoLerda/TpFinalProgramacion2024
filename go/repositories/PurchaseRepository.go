package repositories

import (
	"Status418/go/models"
	"context"
	"errors"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PurchaseRepositoryInterface interface {
	Create(newPurchase models.Purchase) (*mongo.InsertOneResult, error)
	GetAll(userCode string, filters models.Filter) ([]models.Purchase, error)
}

type PurchaseRepository struct {
	db DB
}

func NewPurchaseRepository(db DB) *PurchaseRepository {
	return &PurchaseRepository{
		db: db,
	}
}

func (purchaseRepository PurchaseRepository) Create(purchase models.Purchase) (*mongo.InsertOneResult, error) {
	DBNAME := os.Getenv("DB_NAME")

	res, err := purchaseRepository.db.GetClient().Database(DBNAME).Collection("Purchases").InsertOne(context.TODO(), purchase)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (purchaseRepository PurchaseRepository) GetAll(userCode string, filters models.Filter) ([]models.Purchase, error){
	DBNAME := os.Getenv("DB_NAME")
	filter := bson.M{}
	// if filters.Year != 0{
	// 	filter["$expr"] = bson.M{
	// 		"$eq": []interface{}{
	// 			bson.M{"$year": "$purchase_date"}, 
	// 			filters.Year,                     
	// 		},
	// 	}
	// }
	data, err := purchaseRepository.db.GetClient().Database(DBNAME).Collection("Purchases").Find(context.TODO(), filter)
	if err != nil {
		err = errors.New("internal")
		return nil, err
	}
	var purchase []models.Purchase
	data.All(context.TODO(), &purchase)

	return purchase, nil
}
