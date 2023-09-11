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

	errChan := make(chan error, 9)

	go func() { errChan <- content.CreateMainGinFile(config) }()
	go func() { errChan <- content.CreateEnvFile(config) }()
	go func() { errChan <- content.CreateDatabaseFile(config) }()
	go func() { errChan <- content.CreateHandlerFile(config) }()
	go func() { errChan <- content.CreateGenericStructFile(config) }()
	go func() { errChan <- content.CreateUseCaseFile(config) }()
	go func() { errChan <- content.CreateRepositoryFile(config) }()
	go func() { errChan <- content.CreateDockerFiles(config) }()

	for i := 0; i < 8; i++ {
		if err := <-errChan; err != nil {
			fmt.Println(err)
			return err
		}
	}

	fmt.Println(config)
	return nil
}
