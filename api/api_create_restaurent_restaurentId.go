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

// @Summary      CreateRestaurent 
// @Description   This API performs the POST operation on Restaurent. It allows you to create Restaurent records.
// @Tags          Restaurent
// @Accept       json
// @Produce      json
// @Param        data body dto.Restaurent false "string collection" 
// @Success      200  {array}   dto.Restaurent "Status OK"
// @Success      202  {array}   dto.Restaurent "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /CreateRestaurent [POST]

    func CreateRestaurentApi(c *fiber.Ctx) error {





    
  
    inputObj := dto.Restaurent{}


    if err := c.BodyParser(&inputObj); err != nil {
    		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }
    
inputObj.RestaurentId = uuid.New().String()
        if err := functions.UniqueCheck(inputObj, "Restaurents", []string{ "RestaurentId"}) ; err!=nil{
          return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
        }
    
    validate := validator.New()
    if validationErr := validate.Struct(&inputObj); validationErr != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, validationErr.Error())
    }
err := dao.DB_CreateRestaurent(&inputObj)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }



        return utils.SendSuccessResponse(c)
        
    
}

