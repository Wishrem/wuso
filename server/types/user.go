package types

type User struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type UserRegisterReq struct {
	Email    string `json:"email" form:"email" binding:"required"`
	Name     string `json:"name" form:"name" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type UserLoginReq struct {
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type UserResp struct {
	User  User   `json:"user"`
	Token string `json:"token"`
}
