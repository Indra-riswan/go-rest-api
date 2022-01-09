package entity

type Book struct {
	ID          uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Tittle      string `gorm:"type:varchar(255)" json:"tittle"`
	Description string `gorm:"type:text" json:"description"`
	UserID      uint   `gorm:"not null" json:"-"`
	User        User   `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
}
