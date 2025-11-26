package controller

import (
	"abc/dto"
	"abc/service"
	"abc/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

type MenuController struct {
	menusService *service.MenuService
}

func NewMenuController() *MenuController {
	return &MenuController{
		menusService: service.NewMenuService(),
	}
}

func (controller *MenuController) Create(c *gin.Context) {
	var CreateMenuRequest dto.CreateMenuReq
	if err := c.ShouldBind(&CreateMenuRequest); err != nil {
		fmt.Println(err)
		utils.ReturnBindError(c, err)
		return
	}
	if err := controller.menusService.Create(c, &CreateMenuRequest); err != nil {
		fmt.Println(err)
		utils.ReturnError(c, err)
		return
	}
	utils.ReturnSuccess(c, "创建成功")
}

func (controller *MenuController) Delete(c *gin.Context) {
	if err := controller.menusService.Delete(c); err != nil {
		utils.ReturnError(c, err)
		return
	}
	utils.ReturnSuccess(c, "删除成功")
}

func (controller *MenuController) List(c *gin.Context) {
	data, err := controller.menusService.List(c)
	if err != nil {
		utils.ReturnError(c, err)
		return
	}
	utils.ReturnSuccess(c, data)
}
