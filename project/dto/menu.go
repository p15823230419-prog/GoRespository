package dto

type CreateMenuReq struct {
	Name     string `json:"name" binding:"required"`
	ParentId uint64 `json:"parentId"`
}
