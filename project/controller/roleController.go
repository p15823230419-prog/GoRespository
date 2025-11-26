package controller

import (
	"abc/dto"
	"abc/service"
	"abc/utils"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
	roleService *service.RoleService
}

func NewRoleController() *RoleController {
	return &RoleController{
		roleService: service.NewRoleService(),
	}
}

// 添加角色接口
func (rc *RoleController) Create(c *gin.Context) {
	var req *dto.CreateRoleRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.ReturnBindError(c, err)
		return
	}
	err := rc.roleService.Create(c, req)
	if err != nil {
		utils.ReturnError(c, err)
		return
	}
	utils.ReturnSuccess(c, "添加成功")
	return
}
