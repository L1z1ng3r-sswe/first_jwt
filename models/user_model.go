package models

type TUser struct {
	ID       uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Email    string `gorm:"unique;index:idx_email;type:varchar(128);not null" validate:"email"`
	Password string `gorm:"type:varchar(255);not null" validate:"min=8"`
}
