package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID         uint64         `gorm:"autoIncrement; primaryKey;" json:"id"`
	Password   string         `json:"-"`
	FirstName  string         `json:"first_name"`
	MiddleName string         `json:"middle_name"`
	LastName   string         `json:"last_name"`
	Email      string         `gorm:"unique" json:"email"`
	Age        uint8          `json:"age"`
	CreatedAt  time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at"`
}
