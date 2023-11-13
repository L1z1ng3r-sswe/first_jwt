package models

type TProduct struct {
	ID          uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Title       string `gorm:"varchar(124);not null" validate:"min=6"`
	Description string `gorm:"text;not null" validate:"min=6"`
	Price       int    `gorm:"not null" validate:"required,number"`
	UserId      int    `gorm:"not null"`
	CreatorId   TUser  `gorm:"foreignKey:UserId" validate:"omitempty"`
	Category    string `gorm:"type:text" validate:"required"`
}
