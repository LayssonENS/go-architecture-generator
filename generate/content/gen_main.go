package content

import (
	"fmt"
	"os"

	"github.com/LayssonENS/go-architecture-generator/domain"
)

var mainGinTemplate = `package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run() 
}
`

func CreateMainGinFile(config domain.ProjectConfig) error {
	filePath := fmt.Sprintf("%s/app/main.go", config.ProjectName)

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(mainGinTemplate)
	return err
}
