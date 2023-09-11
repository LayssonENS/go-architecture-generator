package content

import (
	"fmt"
	"html/template"
	"os"
	"strings"

	"github.com/LayssonENS/go-architecture-generator/domain"
)

var structTemplate = `package {{.PackageName}}

import "time"

type {{.StructName}} struct {
	ID        int64     ` + "`json:\"id\"`" + `
	Text      string    ` + "`json:\"text\"`" + `
	CreatedAt time.Time ` + "`json:\"createdAt\"`" + `
	UpdatedAt time.Time ` + "`json:\"updatedAt\"`" + `
	IsActive  bool      ` + "`json:\"isActive\"`" + `
}
`

func CreateGenericStructFile(config domain.ProjectConfig) error {
	filePath := fmt.Sprintf("%s/domain/%s.go", config.ProjectName, strings.ToLower(config.StructName))

	tmpl, err := template.New("structTemplate").Parse(structTemplate)
	if err != nil {
		return err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	data := struct {
		PackageName string
		StructName  string
	}{
		PackageName: strings.ToLower(config.StructName),
		StructName:  config.StructName,
	}

	err = tmpl.Execute(file, data)
	return err
}
