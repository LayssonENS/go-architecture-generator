package content

import (
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/LayssonENS/go-architecture-generator/domain"
)

var mainGinTemplate = `package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"{{.ProjectPath}}/config"
	"{{.ProjectPath}}/config/database"
	{{.StructName}}HttpDelivery "{{.ProjectPath}}/{{.StructNameLower}}/delivery/http"
	{{.StructName}}Repository "{{.ProjectPath}}/{{.StructNameLower}}/repository"
	{{.StructName}}UCase "{{.ProjectPath}}/{{.StructNameLower}}/usecase"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title {{.ProjectName}} API
// @version 1.0
// @description This is {{.ProjectName}} API in Go.

func main() {
	ctx := context.Background()
	log := logrus.New()

	dbInstance, err := database.New{{.Database}}Connection()
	if err != nil {
		log.WithError(err).Fatal("failed connection database")
		return
	}

	//err = database.DBMigrate(dbInstance, config.GetEnv().DbConfig)
	//if err != nil {
	//	log.WithError(err).Fatal("failed to migrate")
	//	return
	//}

	router := gin.Default()

	{{.StructNameLower}}Repository := {{.StructName}}Repository.New{{.Database}}{{.StructName}}Repository(dbInstance)
	{{.StructNameLower}}Service := {{.StructName}}UCase.New{{.StructName}}UseCase({{.StructNameLower}}Repository)

	{{.StructName}}HttpDelivery.New{{.StructName}}Handler(router, {{.StructNameLower}}Service)
	router.GET("/{{.ProjectName}}/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	gin.SetMode(gin.ReleaseMode)
	if config.GetEnv().Debug {
		gin.SetMode(gin.DebugMode)
	}

	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%v", config.GetEnv().Port),
		Handler: router,
	}

	go func() {
		if err := httpServer.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("Shutting down API...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatal("API Server forced to shutdown:", err)
	}

	log.Println("API Server exiting")
}
`

func CreateMainGinFile(config domain.ProjectConfig) error {
	t, err := template.New("main").Parse(mainGinTemplate)
	if err != nil {
		return err
	}

	filePath := fmt.Sprintf("%s/app/main.go", config.ProjectName)
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	data := map[string]string{
		"ProjectName":     config.ProjectName,
		"ProjectPath":     config.ProjectPath,
		"Framework":       config.Framework,
		"Database":        strings.ToUpper(string(config.Database[0])) + config.Database[1:],
		"StructName":      config.StructName,
		"StructNameLower": strings.ToLower(config.StructName),
	}

	return t.Execute(file, data)
}
