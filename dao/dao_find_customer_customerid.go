package dao

import (
	"RestaurentMGT/dbConfig"
	"RestaurentMGT/dto"
	
	"context"
    "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func DB_FindCustomerbyCustomerId (customerId string) (*dto.Customer, error) {
	var object dto.Customer

	err := dbConfig.DATABASE.Collection("Customers").FindOne(context.Background(), bson.M{"customerid": customerId, "deleted":false}).Decode(&object)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		} else {
		    return nil, err
		}
    }
	return &object, nil
}
