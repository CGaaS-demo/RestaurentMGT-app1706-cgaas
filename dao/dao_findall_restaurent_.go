package dao

import (
	"RestaurentMGT/dbConfig"
	"RestaurentMGT/dto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
    "context"
    "errors"
)

func DB_FindallRestaurent () (*[]dto.Restaurent, error) {
	var objects []dto.Restaurent
	results, err := dbConfig.DATABASE.Collection("Restaurents").Find(context.Background(), bson.M{ "deleted":false})
	if err != nil {
        if err == mongo.ErrNoDocuments {
        	return nil, nil
        } else {
        	return nil, err
        }
    }
	for results.Next(context.Background()) {
		var object dto.Restaurent
		if err = results.Decode(&object); err != nil {
			return nil, errors.New("Error when Decoding Restaurent")
		}
		objects = append(objects, object)
	}
	return &objects, nil
}
