package entities

import "time"

type Person struct {
	ID          int       `json:"id"`
	Name        string    `json:"name" validate:"required"`
	Surname     string    `json:"surname" validate:"required"`
	Patronymic  *string   `json:"patronymic,omitempty"`
	Gender      string    `json:"gender,omitempty"`
	Age         int       `json:"age,omitempty"`
	Nationality string    `json:"nationality,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
