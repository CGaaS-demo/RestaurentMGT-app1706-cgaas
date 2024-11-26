package api

import (
  
"RestaurentMGT/utils"
"github.com/gofiber/fiber/v2"

  "RestaurentMGT/dto"
    "github.com/go-playground/validator/v10"
    
    "RestaurentMGT/dao"
    
  

  
  
  
)

// @Summary      UpdateCustomer 
// @Description   This API performs the PUT operation on Customer. It allows you to update Customer records.
// @Tags          Customer
// @Accept       json
// @Produce      json
// @Param        data body dto.Customer false "string collection" 
// @Success      200  {array}   dto.Customer "Status OK"
// @Success      202  {array}   dto.Customer "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /UpdateCustomer [PUT]

    func UpdateCustomerApi(c *fiber.Ctx) error {





    
  
    inputObj := dto.Customer{}


    if err := c.BodyParser(&inputObj); err != nil {
    		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }
    

    validate := validator.New()
    if validationErr := validate.Struct(&inputObj); validationErr != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, validationErr.Error())
    }
err := dao.DB_UpdateCustomer(&inputObj)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }



        return utils.SendSuccessResponse(c)
        
    
}

