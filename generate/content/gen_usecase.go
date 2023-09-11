package content

import (
	"fmt"
	"html/template"
	"os"
	"strings"

	"github.com/LayssonENS/go-architecture-generator/domain"
)

const useCaseTemplate = `package usecase

import (
	"github.com/{{.ProjectPath}}/domain"
)

type {{.StructLowerName}}UseCase struct {
	{{.StructLowerName}}Repository domain.{{.StructName}}Repository
}

func New{{.StructName}}UseCase({{.StructLowerName}}Repository domain.{{.StructName}}Repository) domain.{{.StructName}}UseCase {
	return &{{.StructLowerName}}UseCase{
		{{.StructLowerName}}Repository: {{.StructLowerName}}Repository,
	}
}

func (a *{{.StructLowerName}}UseCase) GetByID(id int64) (domain.{{.StructName}}, error) {
	{{.StructLowerName}}, err := a.{{.StructLowerName}}Repository.GetByID(id)
	if err != nil {
		return {{.StructLowerName}}, err
	}

	return {{.StructLowerName}}, nil
}

func (a *{{.StructLowerName}}UseCase) Create{{.StructName}}({{.StructLowerName}} domain.{{.StructName}}Request) error {
	err := a.{{.StructLowerName}}Repository.Create{{.StructName}}({{.StructLowerName}})
	if err != nil {
		return err
	}

	return nil
}

func (a *{{.StructLowerName}}UseCase) GetAll{{.StructName}}() ([]domain.{{.StructName}}, error) {
	{{.StructLowerName}}, err := a.{{.StructLowerName}}Repository.GetAll{{.StructName}}()
	if err != nil {
		return {{.StructLowerName}}, err
	}

	return {{.StructLowerName}}, nil
}
`

func CreateUseCaseFile(config domain.ProjectConfig) error {
	t, err := template.New("useCase").Parse(useCaseTemplate)
	if err != nil {
		return err
	}

	filePath := fmt.Sprintf("%s/%s/usecase/%s_usecase.go", config.ProjectName, strings.ToLower(config.StructName), strings.ToLower(config.StructName))
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	data := map[string]string{
		"ProjectPath":     config.ProjectPath,
		"StructName":      config.StructName,
		"StructLowerName": strings.ToLower(config.StructName),
	}

	return t.Execute(file, data)
}
