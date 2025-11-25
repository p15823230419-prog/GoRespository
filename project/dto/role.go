package dto

type CreateRoleRequest struct {
	RoleName string `json:"roleName" binding:"required"`
}
