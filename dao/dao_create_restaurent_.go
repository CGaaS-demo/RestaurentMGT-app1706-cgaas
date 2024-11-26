package dao

import (
    "context"
	"RestaurentMGT/dbConfig"
	"RestaurentMGT/dto"

)

func DB_CreateRestaurent (object *dto.Restaurent) error {

	_, err := dbConfig.DATABASE.Collection("Restaurents").InsertOne(context.Background(), object)
	if err != nil {
		return err
	}
	return nil
}