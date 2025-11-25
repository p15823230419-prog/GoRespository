package service

import (
	"abc/dao"
	"abc/dto"
	"errors"

	"github.com/gin-gonic/gin"
)

type RoleService struct {
	roleDao *dao.RoleDao
}

func NewRoleService() *RoleService {
	return &RoleService{
		roleDao: dao.NewRoleDao(),
	}
}

func (rs RoleService) CreateRole(c *gin.Context, req *dto.CreateRoleRequest) error {
	role, err := rs.roleDao.FindByRoleName(c.Request.Context(), req.RoleName)
	if err != nil {
		return err
	}
	if role != nil {
		return errors.New("角色已存在")
	}
	err = rs.roleDao.CreateRole(c.Request.Context(), *CreateRoleRequestToEntity(req))
	if err != nil {
		return err
	}
	return nil
}
