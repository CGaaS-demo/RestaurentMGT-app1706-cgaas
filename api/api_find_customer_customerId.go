package api

import (
  
"RestaurentMGT/utils"
"github.com/gofiber/fiber/v2"

  
    "RestaurentMGT/dao"
    
  

  
  
  
)

// @Summary      FindCustomer 
// @Description   This API performs the GET operation on Customer. It allows you to retrieve Customer records.
// @Tags          Customer
// @Accept       json
// @Produce      json
// @Param        objectId query []string false "string collection"  collectionFormat(multi)
// @Success      200  {array}   dto.Customer "Status OK"
// @Success      202  {array}   dto.Customer "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /FindCustomer [GET]

    func FindCustomerApi(c *fiber.Ctx) error {





    customerId := c.Query("customerId")
        
    
  returnValue, err := dao.DB_FindCustomerbyCustomerId(customerId)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }


return c.Status(fiber.StatusAccepted).JSON(&returnValue)
}

