package api

import (
  
"RestaurentMGT/utils"
"github.com/gofiber/fiber/v2"

  "github.com/google/uuid"
    "RestaurentMGT/functions"
    
  "RestaurentMGT/dto"
    "github.com/go-playground/validator/v10"
    
    "RestaurentMGT/dao"
    
  

  
  
  
)

// @Summary      CreateCustomer 
// @Description   This API performs the POST operation on Customer. It allows you to create Customer records.
// @Tags          Customer
// @Accept       json
// @Produce      json
// @Param        data body dto.Customer false "string collection" 
// @Success      200  {array}   dto.Customer "Status OK"
// @Success      202  {array}   dto.Customer "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /CreateCustomer [POST]

    func CreateCustomerApi(c *fiber.Ctx) error {





    
  
    inputObj := dto.Customer{}


    if err := c.BodyParser(&inputObj); err != nil {
    		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }
    
inputObj.CustomerId = uuid.New().String()
        if err := functions.UniqueCheck(inputObj, "Customers", []string{ "CustomerId"}) ; err!=nil{
          return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
        }
    
    validate := validator.New()
    if validationErr := validate.Struct(&inputObj); validationErr != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, validationErr.Error())
    }
err := dao.DB_CreateCustomer(&inputObj)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }



        return utils.SendSuccessResponse(c)
        
    
}

