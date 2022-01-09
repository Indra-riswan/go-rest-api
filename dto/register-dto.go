package dto

type RegisterDto struct {
	Name     string `json:"name" form:"name" binding:"required" validate:"min:2"`
	Email    string `json:"email" form:"email" binding:"required" validate:"email"`
	Password string `json:"password" form:"password" binding:"required" validate:"min:6"`
}