package usecase

import (
	"fmt"

	"github.com/LayssonENS/go-architecture-generator/domain"
	"github.com/LayssonENS/go-architecture-generator/generate/content"
)

type generateUseCase struct{}

func NewUserUseCase() domain.GenerateUseCase {
	return &generateUseCase{}
}

func (f *generateUseCase) Generate(config domain.ProjectConfig) error {
	err := content.CreateProjectStructure(config)
	if err != nil {
		return err
	}

	err = content.CreateMainGinFile(config)
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = content.CreateEnvFile(config)
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = content.CreateDatabaseFile(config)
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = content.CreateHandlerFile(config)
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = content.CreateGenericStructFile(config)
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = content.CreateUseCaseFile(config)
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = content.CreateRepositoryFile(config)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println(config)
	return nil
}
