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

func (service *RoleService) Create(c *gin.Context, req *dto.CreateRoleRequest) error {
	role, err := service.roleDao.FindByName(c.Request.Context(), req.RoleName)
	if err != nil {
		return err
	}
	if role != nil {
		return errors.New("角色已存在")
	}
	err = service.roleDao.Create(c.Request.Context(), *CreateRoleRequestToEntity(req))
	if err != nil {
		return err
	}
	return nil
}

func (service *RoleService) List(c *gin.Context) (interface{}, error) {
	data, err := service.roleDao.FindAll(c.Request.Context())
	if err != nil {
		return nil, err
	}

	return data, nil
}
