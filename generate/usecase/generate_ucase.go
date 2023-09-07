package usecase

import (
	"fmt"

	"github.com/LayssonENS/go-architecture-generator/domain"
)

type generateUseCase struct{}

func NewUserUseCase() domain.GenerateUseCase {
	return &generateUseCase{}
}

func (f *generateUseCase) Generate(config domain.ProjectConfig) error {
	fmt.Println(config)
	return nil
}
