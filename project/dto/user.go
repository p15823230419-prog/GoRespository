package dto //数据和前端的通信接口

// json请求
type RegisterRequest struct {
	Username string   `json:"username" binding:"required"`
	Nickname string   `json:"nickname"`
	Avatar   string   `json:"avatar"`
	Phone    string   `json:"phone"`
	Email    string   `json:"email"`
	Password string   `json:"password" binding:"required,min=6,max=20"`
	RoleIDs  []uint64 `json:"roleIDs"`
}

type RegisterResponse struct {
	UserId uint64 `json:"userId"`
}

// 登录请求
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 登录回应
type LoginResponse struct {
	Token string `json:"token"`
}

// 更新请求
type UpdateRequest struct {
	Id       uint64   `json:"id" binding:"required"`
	Password string   `json:"password"`
	Username string   `json:"username"`
	Nickname string   `json:"nickname"`
	Avatar   string   `json:"avatar"`
	Phone    string   `json:"phone"`
	Email    string   `json:"email"`
	RoleIDs  []uint64 `json:"roleIds"`
}

// 查询请求
type SelectRequest struct {
	Username string `form:"username"`
	PageNum  int    `form:"pageNum"`
	PageSize int    `form:"pageSize"`
}

// 查询返回
type SelectResponse struct {
	UserId   uint64   `json:"userId"`
	Username string   `json:"username"`
	Nickname string   `json:"nickname"`
	Avatar   string   `json:"avatar"`
	Phone    string   `json:"phone"`
	Email    string   `json:"email"`
	Status   uint64   `json:"status"`
	Roles    []string `json:"roles"`
}
