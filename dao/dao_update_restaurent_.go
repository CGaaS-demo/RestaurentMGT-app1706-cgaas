package dao

import (
	"RestaurentMGT/dbConfig"
    "RestaurentMGT/dto"
    "errors"
    "go.mongodb.org/mongo-driver/bson"
	"context"
)

func DB_UpdateRestaurent (object *dto.Restaurent)  error {

	result, err := dbConfig.DATABASE.Collection("Restaurents").UpdateOne(context.Background(), bson.M{"restaurentid": object.RestaurentId, "deleted":false}, bson.M{"$set": object})
	if err != nil {
		return err
	}
	if result.ModifiedCount < 1 && result.MatchedCount != 1 {
		return errors.New("Specified ID not found!")
	}

	return nil
}