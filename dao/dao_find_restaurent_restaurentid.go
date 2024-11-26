package dao

import (
	"RestaurentMGT/dbConfig"
	"RestaurentMGT/dto"
	
	"context"
    "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func DB_FindRestaurentbyRestaurentId (restaurentId string) (*dto.Restaurent, error) {
	var object dto.Restaurent

	err := dbConfig.DATABASE.Collection("Restaurents").FindOne(context.Background(), bson.M{"restaurentid": restaurentId, "deleted":false}).Decode(&object)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		} else {
		    return nil, err
		}
    }
	return &object, nil
}
