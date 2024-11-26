package dao

import (
    "context"
	"RestaurentMGT/dbConfig"
	"RestaurentMGT/dto"

)

func DB_CreateCustomer (object *dto.Customer) error {

	_, err := dbConfig.DATABASE.Collection("Customers").InsertOne(context.Background(), object)
	if err != nil {
		return err
	}
	return nil
}