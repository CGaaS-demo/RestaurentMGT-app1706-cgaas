package dao

import (
	"RestaurentMGT/dbConfig"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"context"
)

func DB_DeleteRestaurent (restaurentId string)  error {
  
        result, err := dbConfig.DATABASE.Collection("Restaurents").UpdateOne(context.Background(), bson.M{"restaurentid": restaurentId}, bson.M{"$set": bson.M{"deleted": true}})
        if err != nil {
            return err
        }
        if result.ModifiedCount < 1 {
            return errors.New("Specified Id not found!")
        }
        return nil
  
}

