package api

import (
  
"RestaurentMGT/utils"
"github.com/gofiber/fiber/v2"

  
    "RestaurentMGT/dao"
    
  

  
  
  
)

// @Summary      FindRestaurent 
// @Description   This API performs the GET operation on Restaurent. It allows you to retrieve Restaurent records.
// @Tags          Restaurent
// @Accept       json
// @Produce      json
// @Param        objectId query []string false "string collection"  collectionFormat(multi)
// @Success      200  {array}   dto.Restaurent "Status OK"
// @Success      202  {array}   dto.Restaurent "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /FindRestaurent [GET]

    func FindRestaurentApi(c *fiber.Ctx) error {





    restaurentId := c.Query("restaurentId")
        
    
  returnValue, err := dao.DB_FindRestaurentbyRestaurentId(restaurentId)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }


return c.Status(fiber.StatusAccepted).JSON(&returnValue)
}

