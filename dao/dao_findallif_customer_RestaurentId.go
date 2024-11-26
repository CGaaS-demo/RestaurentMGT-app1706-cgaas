package dao

import (
    "context"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "RestaurentMGT/dbConfig"
    "RestaurentMGT/dto"
    "errors"
)

func DB_FindallifCustomerbyRestaurentId (RestaurentId string) (*[]dto.Customer, error) {
	var objects []dto.Customer
   	
    
    results, err := dbConfig.DATABASE.Collection("Customers").Find(context.Background(), bson.M{"restaurentid" :RestaurentId, "deleted" : false})
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
            return nil, errors.New("Error When Decoding Customer")
        }
        objects = append(objects, object)
    }
    return &objects, nil
}