package api

import (
  
"RestaurentMGT/utils"
"github.com/gofiber/fiber/v2"

  "RestaurentMGT/dto"
    "github.com/go-playground/validator/v10"
    
    "RestaurentMGT/dao"
    
  

  
  
  
)

// @Summary      UpdateRestaurent 
// @Description   This API performs the PUT operation on Restaurent. It allows you to update Restaurent records.
// @Tags          Restaurent
// @Accept       json
// @Produce      json
// @Param        data body dto.Restaurent false "string collection" 
// @Success      200  {array}   dto.Restaurent "Status OK"
// @Success      202  {array}   dto.Restaurent "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /UpdateRestaurent [PUT]

    func UpdateRestaurentApi(c *fiber.Ctx) error {





    
  
    inputObj := dto.Restaurent{}


    if err := c.BodyParser(&inputObj); err != nil {
    		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }
    

    validate := validator.New()
    if validationErr := validate.Struct(&inputObj); validationErr != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, validationErr.Error())
    }
err := dao.DB_UpdateRestaurent(&inputObj)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }



        return utils.SendSuccessResponse(c)
        
    
}

