package models

import (
	"time"
)

type Person struct {
	ID          uint      `gorm:"primaryKey" json:"id" example:"1"`
	Name        string    `gorm:"not null" json:"name" example:"dmitry"`
	Surname     string    `gorm:"not null" json:"surname" example:"vasiliev"`
	Patronymic  string    `json:"patronymic,omitempty" example:"vasilyevich"`
	Gender      string    `json:"gender,omitempty" example:"male"`
	Age         int       `json:"age,omitempty" example:"40"`
	Nationality string    `json:"nationality,omitempty" example:"RU"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

type PersonCreateRequest struct {
	Name       string `json:"name" binding:"required" example:"dmitry"`
	Surname    string `json:"surname" binding:"required" example:"vasiliev"`
	Patronymic string `json:"patronymic,omitempty" example:"vasilyevich"`
}

type PersonUpdateRequest struct {
	Name        string `json:"name,omitempty" example:"dmitriy"`
	Surname     string `json:"surname,omitempty" example:"vasiliev"`
	Patronymic  string `json:"patronymic,omitempty" example:"vasilyevich"`
	Gender      string `json:"gender,omitempty" example:"male"`
	Age         int    `json:"age,omitempty" example:"41"`
	Nationality string `json:"nationality,omitempty" example:"RU"`
}
