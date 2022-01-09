package dto

type UpdateBookDto struct {
	ID          uint   `json:"id" form:"id" binding:"required"`
	Tittle      string `json:"tittle" form:"tittle" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	UserID      int    `json:"user_id,omitempty" form:"user_id,omitempty"`
}
type CreateBookDto struct {
	Tittle      string `json:"tittle" form:"tittle" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	UserID      int    `json:"user_id,omitempty" form:"user_id,omitempty"`
}
