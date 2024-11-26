package api

import (
  
"RestaurentMGT/utils"
"github.com/gofiber/fiber/v2"

  
    "RestaurentMGT/dao"
    
  

  
  
  
)

// @Summary      DeleteRestaurent 
// @Description   This API performs the DELETE operation on Restaurent. It allows you to delete Restaurent records.
// @Tags          Restaurent
// @Accept       json
// @Produce      json
// @Param        objectId query []string false "string collection"  collectionFormat(multi)
// @Success      200  {array}   dto.Restaurent "Status OK"
// @Success      202  {array}   dto.Restaurent "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /DeleteRestaurent [DELETE]

    func DeleteRestaurentApi(c *fiber.Ctx) error {





    restaurentId := c.Query("restaurentId")
        
    
  err := dao.DB_DeleteRestaurent(restaurentId)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }



        return utils.SendSuccessResponse(c)
        
    
}

