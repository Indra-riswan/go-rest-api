package dto

type UpdateUserDto struct {
	ID       uint   `json:"id" form:"id" binding:"required"`
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password,omitempty" form:"password.omitempty"`
}

// type CreateUserDto struct {
// 	Name     string `json:"name" form:"name" binding:"required"`
// 	Email    string `json:"email" form:"email" binding:"required" validate:"email"`
// 	Password string `json:"password,omitempty" form:"password.omitempty" validate:"min:6"`
// }
