package db

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Uid       int64
	Password  string
	Username  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Task struct {
	Id        int64
	Uid       int64
	Title     string
	Content   string
	Status    int16
	StartAt   time.Time
	EndAt     time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
