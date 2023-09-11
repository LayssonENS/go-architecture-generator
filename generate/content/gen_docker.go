package content

import (
	"fmt"
	"os"

	"github.com/LayssonENS/go-architecture-generator/domain"
)

func CreateDockerFiles(config domain.ProjectConfig) error {
	dockerfileContent := `FROM golang:1.21 as builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./app/main.go

FROM scratch

COPY --from=builder /app/main /main

EXPOSE 9000

CMD ["/main"]
`
	dockerComposerContent := `version: '3.8'

services:
  app:
    build:
      context: .
      dockerfile: Dockerfile
    ports:
      - "9000:9000"
    environment:
      - ENVIRONMENT=dev
      - DEBUG=true
      - PORT=9000
      - DB_HOST=db
      - DB_PORT=5432
      - DB_USER=postgres
      - DB_PASSWORD=postgres
      - DB_NAME=postgres
    depends_on:
      - db
  db:
    image: postgres:alpine
    environment:
      - POSTGRES_USER=postgres
      - POSTGRES_PASSWORD=postgres
      - POSTGRES_DB=postgres
    ports:
      - "5432:5432"
    volumes:
      - postgres-db:/var/lib/postgresql/data

volumes:
  postgres-db:
`

	// Create Dockerfile
	dockerfilePath := fmt.Sprintf("%s/Dockerfile", config.ProjectName)
	if err := createFileWithContent(dockerfilePath, dockerfileContent); err != nil {
		return err
	}

	// Create docker-compose.yml
	dockerComposePath := fmt.Sprintf("%s/docker-compose.yml", config.ProjectName)
	if err := createFileWithContent(dockerComposePath, dockerComposerContent); err != nil {
		return err
	}

	return nil
}

func createFileWithContent(filePath, content string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	return err
}
