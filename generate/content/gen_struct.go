package content

import (
	"fmt"
	"html/template"
	"os"
	"strings"

	"github.com/LayssonENS/go-architecture-generator/domain"
)

func CreateProjectStructure(config domain.ProjectConfig) error {
	if err := os.Mkdir(config.ProjectName, 0755); err != nil {
		return err
	}

	subDirs := []string{
		"app",
		"config/database",
		"docs",
		"domain",
		fmt.Sprintf("%s", strings.ToLower(config.StructName)),
		fmt.Sprintf("%s/delivery/http", strings.ToLower(config.StructName)),
		fmt.Sprintf("%s/repository", strings.ToLower(config.StructName)),
		fmt.Sprintf("%s/usecase", strings.ToLower(config.StructName)),
	}

	for _, dir := range subDirs {
		if err := os.MkdirAll(fmt.Sprintf("%s/%s", config.ProjectName, dir), 0755); err != nil {
			return err
		}
	}

	err := createGoModFile(config)
	if err != nil {
		return err
	}

	err = createMakefile(config)
	if err != nil {
		return err
	}

	return nil
}

func createGoModFile(config domain.ProjectConfig) error {
	content := fmt.Sprintf("module %s\n\ngo 1.21", config.ProjectPath)

	filePath := fmt.Sprintf("%s/go.mod", config.ProjectName)

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	return err

}

func createMakefile(config domain.ProjectConfig) error {
	makefileTemplate := `swagger:
	swag init -g app/main.go
	swag init -g {{.HandlerPath}}

mock:
	mockgen -destination={{.MockDestination}} -source=domain/{{.StructLower}}.go

test:
	go test ./...

build:
	docker-compose build
	docker-compose up -d

run:
	go run app/main.go
`

	filePath := fmt.Sprintf("%s/Makefile", config.ProjectName)

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	data := struct {
		HandlerPath     string
		MockDestination string
		StructLower     string
	}{
		HandlerPath:     fmt.Sprintf("%s/delivery/http/%s_handler.go", strings.ToLower(config.StructName), strings.ToLower(config.StructName)),
		MockDestination: "domain/mock/domain_mock.go",
		StructLower:     strings.ToLower(config.StructName),
	}

	tmpl, err := template.New("makefileTemplate").Parse(makefileTemplate)
	if err != nil {
		return err
	}

	err = tmpl.Execute(file, data)
	return err
}
