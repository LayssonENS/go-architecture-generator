package content

import (
	"fmt"
	"html/template"
	"os"

	"github.com/LayssonENS/go-architecture-generator/domain"
)

const dbTemplate = `package database

// Conteúdo para {{.Database}}
`

func createDatabaseFile(config domain.ProjectConfig) error {
	t, err := template.New("db").Parse(dbTemplate)
	if err != nil {
		return err
	}

	filePath := fmt.Sprintf("%s/config/database/%s.go", config.ProjectName, config.Database)
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	return t.Execute(file, config)
}
