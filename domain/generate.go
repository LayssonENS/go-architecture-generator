package domain

import (
	"time"
)

type GenerateRequest struct {
	Name      string `json:"name" binding:"required,min=3"`
	Email     string `json:"email" binding:"required,email"`
	BirthDate string `json:"birth_date"`
}

type Generate struct {
	ID        int64      `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	BirthDate *time.Time `json:"birth_date"`
	CreatedAt time.Time  `json:"created_at"`
}

type GenerateUseCase interface {
	Generate(person GenerateRequest) error
}
