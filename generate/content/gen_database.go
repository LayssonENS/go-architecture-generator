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

const postgresTemplate = `package database

import (
	"database/sql"
	"fmt"
	"log"

	"{{.ProjectPath}}/config"
	_ "github.com/lib/pq"
)

func NewPostgresConnection() (*sql.DB, error) {
	dataSourceName := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.GetEnv().DbConfig.User,
		config.GetEnv().DbConfig.Password,
		config.GetEnv().DbConfig.Host,
		config.GetEnv().DbConfig.Port,
		config.GetEnv().DbConfig.Name,
	)

	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	log.Println("Successfully connected to the PostgreSQL database.")
	return db, nil
}`

const sqliteTemplate = `package database

import (
	"database/sql"
	"log"

	"{{.ProjectPath}}/config"
	_ "github.com/mattn/go-sqlite3"
)

func NewSQLiteConnection() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", config.GetEnv().DbConfig.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}

	log.Println("Successfully connected to the SQLite database.")
	return db, nil
}`

func CreateDatabaseFile(config domain.ProjectConfig) error {
	var t *template.Template
	var err error

	switch config.Database {
	case "postgres":
		t, err = template.New("postgres").Parse(postgresTemplate)
	case "sqlite":
		t, err = template.New("sqlite").Parse(sqliteTemplate)
	default:
		t, err = template.New("db").Parse(dbTemplate)
	}

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
