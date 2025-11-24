package dto //数据和前端的通信接口

// json请求
type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Password string `json:"password" binding:"required,min=6,max=20"`
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

type SelectResponse struct {
	UserId   uint64 `json:"userId"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}
