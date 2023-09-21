package controller

import (
	"AnaProductService/request"
	"AnaProductService/response"
	"AnaProductService/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.ProductService
}

func NewUserController(userService *service.ProductService) *UserController {
	return &UserController{userService}
}

// @Summary		Get User By Id
// @Description	Get User By Id
// @Produce		json
// @Param product_name body request.ProductRequest false "product_name"
// @Success		200	{object} model.Product
// @Router			/product [post]
func (uc *UserController) GetListProductByName(c *gin.Context) {
	request := request.ProductRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, response.Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	data, err := uc.userService.GetListProductByName(request.Name)
	if err != nil {
		c.JSON(400, response.Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	c.JSON(200, response.Response{
		Status:  "success",
		Data:    data,
		Message: "success",
	})
}
