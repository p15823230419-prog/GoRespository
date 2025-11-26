package service

import (
	"abc/dao"
	"abc/dto"
	"abc/model"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MenuService struct {
	MenuDao *dao.MenuDao
}

func NewMenuService() *MenuService {
	return &MenuService{
		MenuDao: dao.NewMenuDao(),
	}
}

func (ms *MenuService) Create(c *gin.Context, req *dto.CreateMenuReq) error {
	if name, err := ms.MenuDao.FindByName(c.Request.Context(), req.Name); err != nil {
		return errors.New("查询数据库错误")
	} else if name != nil {
		return errors.New("菜单名已存在")
	}
	if err := ms.MenuDao.Create(c.Request.Context(), req); err != nil {
		return errors.New("创建数据库错误")
	}
	return nil
}

func (ms *MenuService) Delete(c *gin.Context) error {
	id := c.Param("id")
	if id == "" {
		return errors.New("请输入要删除的id")
	}
	menuId, _ := strconv.Atoi(id)

	if user, err := ms.MenuDao.FindById(c.Request.Context(), uint64(menuId)); err != nil {
		return errors.New("数据库错误")
	} else if user == nil {
		return errors.New("未找到用户")
	}

	if err := ms.MenuDao.Delete(c, uint64(menuId)); err != nil {
		return errors.New("删除失败,数据库错误")
	}
	return nil
}

func (ms *MenuService) List(c *gin.Context) (interface{}, error) {
	list, err := ms.MenuDao.FindAll(c.Request.Context())
	if err != nil {
		return nil, err
	}
	tree := BuildMenuTree(list, 0)
	return tree, nil
}

// 递归遍历数组添加
func BuildMenuTree(list []model.Menu, parentID uint64) []model.Menu {
	var tree []model.Menu
	for _, menu := range list {
		if menu.ParentId == parentID {
			children := BuildMenuTree(list, menu.Id)
			menu.Children = children
			tree = append(tree, menu)
		}
	}
	return tree
}
