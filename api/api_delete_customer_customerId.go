package api

import (
  
"RestaurentMGT/utils"
"github.com/gofiber/fiber/v2"

  
    "RestaurentMGT/dao"
    
  

  
  
  
)

// @Summary      DeleteCustomer 
// @Description   This API performs the DELETE operation on Customer. It allows you to delete Customer records.
// @Tags          Customer
// @Accept       json
// @Produce      json
// @Param        objectId query []string false "string collection"  collectionFormat(multi)
// @Success      200  {array}   dto.Customer "Status OK"
// @Success      202  {array}   dto.Customer "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /DeleteCustomer [DELETE]

    func DeleteCustomerApi(c *fiber.Ctx) error {





    customerId := c.Query("customerId")
        
    
  err := dao.DB_DeleteCustomer(customerId)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }



        return utils.SendSuccessResponse(c)
        
    
}

