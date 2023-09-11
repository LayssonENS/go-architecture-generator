package content

import (
	"fmt"
	"html/template"
	"os"
	"strings"

	"github.com/LayssonENS/go-architecture-generator/domain"
)

const repositoryTemplate = `package {{.StructLowerName}}Repository

import (
	"database/sql"
	"time"

	"{{.ProjectPath}}/domain"
)

const dateLayout = "2006-01-02"

type postgres{{.StructName}}Repo struct {
	DB *sql.DB
}

// NewPostgres{{.StructName}}Repository will create an implementation of {{.StructName}}.Repository
func NewPostgres{{.StructName}}Repository(db *sql.DB) domain.{{.StructName}}Repository {
	return &postgres{{.StructName}}Repo{
		DB: db,
	}
}

// GetByID : Retrieves a {{.StructLowerName}} by ID from the Postgres repository
func (p *postgres{{.StructName}}Repo) GetByID(id int64) (domain.{{.StructName}}, error) {
	var {{.StructLowerName}} domain.{{.StructName}}
	err := p.DB.QueryRow(
		"SELECT id, name, email, birth_date, created_at FROM {{.StructLowerName}} WHERE id = $1", id).Scan(
		&{{.StructLowerName}}.ID, &{{.StructLowerName}}.Name, &{{.StructLowerName}}.Email, &{{.StructLowerName}}.BirthDate, &{{.StructLowerName}}.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return {{.StructLowerName}}, domain.ErrRegistrationNotFound
		}
		return {{.StructLowerName}}, err
	}
	return {{.StructLowerName}}, nil
}

// Create{{.StructName}} : Inserts a new {{.StructLowerName}} into the Postgres repository using the provided {{.StructLowerName}} request data
func (p *postgres{{.StructName}}Repo) Create{{.StructName}}({{.StructLowerName}} domain.{{.StructName}}Request) error {
	date, _ := time.Parse(dateLayout, {{.StructLowerName}}.BirthDate)
	birthDate := date

	query := ` + "`INSERT INTO {{.StructLowerName}} (name, email, birth_date) VALUES ($1, $2, $3)`" + `
	_, err := p.DB.Exec(query, {{.StructLowerName}}.Name, {{.StructLowerName}}.Email, birthDate)
	if err != nil {
		return err
	}

	return nil
}

// GetAll{{.StructName}} : Retrieves all {{.StructLowerName}} data from the Postgres repository
func (p *postgres{{.StructName}}Repo) GetAll{{.StructName}}() ([]domain.{{.StructName}}, error) {
	var {{.StructLowerName}}s []domain.{{.StructName}}

	rows, err := p.DB.Query("SELECT id, name, email, birth_date, created_at FROM {{.StructLowerName}}")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var {{.StructLowerName}} domain.{{.StructName}}
		err := rows.Scan(
			&{{.StructLowerName}}.ID,
			&{{.StructLowerName}}.Name,
			&{{.StructLowerName}}.Email,
			&{{.StructLowerName}}.BirthDate,
			&{{.StructLowerName}}.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		{{.StructLowerName}}s = append({{.StructLowerName}}s, {{.StructLowerName}})
	}

	if len({{.StructLowerName}}s) == 0 {
		return nil, domain.ErrRegistrationNotFound
	}

	return {{.StructLowerName}}s, nil
}
`

func CreateRepositoryFile(config domain.ProjectConfig) error {
	t, err := template.New("repository").Parse(repositoryTemplate)
	if err != nil {
		return err
	}

	filePath := fmt.Sprintf("%s/%s/repository/%s_repository.go", config.ProjectName, strings.ToLower(config.StructName), strings.ToLower(config.StructName))
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
