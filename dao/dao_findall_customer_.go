package dao

import (
	"RestaurentMGT/dbConfig"
	"RestaurentMGT/dto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
    "context"
    "errors"
)

func DB_FindallCustomer () (*[]dto.Customer, error) {
	var objects []dto.Customer
	results, err := dbConfig.DATABASE.Collection("Customers").Find(context.Background(), bson.M{ "deleted":false})
	if err != nil {
        if err == mongo.ErrNoDocuments {
        	return nil, nil
        } else {
        	return nil, err
        }
    }
	for results.Next(context.Background()) {
		var object dto.Customer
		if err = results.Decode(&object); err != nil {
			return nil, errors.New("Error when Decoding Customer")
		}
		objects = append(objects, object)
	}
	return &objects, nil
}
