package content

import (
	"fmt"
	"os"

	"github.com/LayssonENS/go-architecture-generator/domain"
)

const envTemplate = `package config

import (
	"log"

	"github.com/Netflix/go-env"
)

type Environment struct {
	Port     string ` + "`env:\"PORT,default=9000\"`" + `
	Debug    bool   ` + "`env:\"DEBUG,default=false\"`" + `
	DbConfig DbConfig
	Extras   env.EnvSet
}

type DbConfig struct {
	User     string ` + "`env:\"DB_USER,default=postgres\"`" + `
	Port     string ` + "`env:\"DB_PORT,default=5432\"`" + `
	Password string ` + "`env:\"DB_PASSWORD,default=postgres\"`" + `
	Host     string ` + "`env:\"DB_HOST,default=localhost\"`" + `
	Name     string ` + "`env:\"DB_NAME,default=postgres\"`" + `
	Path     string ` + "`env:\"DB_PATH,default=./sqlite.db\"`" + `
	DbType   string ` + "`env:\"DB_TYPE,default=postgres\"`" + `
}

var Env Environment

func init() {
	_, err := env.UnmarshalFromEnviron(&Env)
	if err != nil {
		log.Fatal(err)
	}
}

func GetEnv() Environment {
	return Env
}
`

func CreateEnvFile(config domain.ProjectConfig) error {
	filePath := fmt.Sprintf("%s/config/env.go", config.ProjectName)

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(envTemplate)
	return err
}
