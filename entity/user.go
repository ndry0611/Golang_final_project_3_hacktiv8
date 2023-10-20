package entity

import (
	"final_project_3/pkg/helpers"
	"final_project_3/pkg/errs"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint   `gorm:"primaryKey;not null"`
	FullName  string `gorm:"not null"`
	Email     string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	Role      string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) BeforeCreate(db *gorm.DB) (err error) {
	hashedPw, hashErr := helpers.GenerateHashedPassword([]byte(u.Password))
	if hashErr != nil {
		return errs.NewInternalServerError(hashErr.Error())
	}
	u.Password = hashedPw
	return nil
}