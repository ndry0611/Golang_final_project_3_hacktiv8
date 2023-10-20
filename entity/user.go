package entity

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey;not null"`
	FullName  string `gorm:"not null"`
	Email     string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	Role      string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
