package content

import (
	"fmt"
	"html/template"
	"os"
	"strings"

	"github.com/LayssonENS/go-architecture-generator/domain"
)

const handlerTemplate = `package http

import (
	"net/http"
	"strconv"

	"{{.ProjectPath}}/domain"
	"github.com/gin-gonic/gin"
)

type {{.StructName}}Handler struct {
	{{.StructShortName}}UseCase domain.{{.StructName}}UseCase
}

func New{{.StructName}}Handler(routerGroup *gin.Engine, us domain.{{.StructName}}UseCase) {
	handler := &{{.StructName}}Handler{
		{{.StructShortName}}UseCase: us,
	}

	routerGroup.GET("/v1/{{.StructLowerName}}/:{{.StructLowerName}}Id", handler.GetByID)
	routerGroup.GET("/v1/{{.StructLowerName}}/all", handler.GetAll{{.StructName}})
	routerGroup.POST("/v1/{{.StructLowerName}}", handler.Create{{.StructName}})
}

// GetByID godoc
// @Summary Get {{.StructName}} by ID
// @Description get {{.StructName}} by ID
// @Tags {{.StructName}}
// @Accept  json
// @Produce  json
// @Param {{.StructLowerName}}Id path int true "{{.StructName}} ID"
// @Success 200 {object} domain.{{.StructName}}
// @Failure 400 {object} domain.ErrorResponse
// @Router /v1/{{.StructLowerName}}/{id} [get]
func (h *{{.StructName}}Handler) GetByID(c *gin.Context) {
	idParam, err := strconv.Atoi(c.Param("{{.StructLowerName}}Id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{ErrorMessage: err.Error()})
		return
	}

	{{.StructLowerName}}Id := int64(idParam)

	response, err := h.{{.StructShortName}}UseCase.GetByID({{.StructLowerName}}Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetAll{{.StructName}} godoc
// @Summary Get all {{.StructName}}s
// @Description get all {{.StructName}}s
// @Tags {{.StructName}}
// @Accept  json
// @Produce  json
// @Success 200 {array} domain.{{.StructName}}
// @Failure 400 {object} domain.ErrorResponse
// @Router /v1/{{.StructLowerName}}/all [get]
func (h *{{.StructName}}Handler) GetAll{{.StructName}}(c *gin.Context) {
	response, err := h.{{.StructShortName}}UseCase.GetAll{{.StructName}}()
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{ErrorMessage: err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

// Create{{.StructName}} godoc
// @Summary Create a new {{.StructName}}
// @Description create new {{.StructName}}
// @Tags {{.StructName}}
// @Accept  json
// @Produce  json
// @Param {{.StructLowerName}} body domain.{{.StructName}}Request true "Create {{.StructName}}"
// @Success 201 {object} string
// @Failure 400 {object} domain.ErrorResponse
// @Failure 422 {object} domain.ErrorResponse
// @Router /v1/{{.StructLowerName}} [post]
func (h *{{.StructName}}Handler) Create{{.StructName}}(c *gin.Context) {
	var {{.StructLowerName}} domain.{{.StructName}}Request
	if err := c.ShouldBindJSON(&{{.StructLowerName}}); err != nil {
		c.JSON(http.StatusUnprocessableEntity, domain.ErrorResponse{ErrorMessage: err.Error()})
		return
	}

	err := h.{{.StructShortName}}UseCase.Create{{.StructName}}({{.StructLowerName}})
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{ErrorMessage: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"msg": "Created"})
}
`

func CreateHandlerFile(config domain.ProjectConfig) error {
	t, err := template.New("handler").Parse(handlerTemplate)
	if err != nil {
		return err
	}

	filePath := fmt.Sprintf("%s/%s/delivery/http/%s_handler.go", config.ProjectName, strings.ToLower(config.StructName), strings.ToLower(config.StructName))
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	data := struct {
		ProjectPath     string
		StructName      string
		StructShortName string
		StructLowerName string
	}{
		ProjectPath:     config.ProjectPath,
		StructName:      config.StructName,
		StructShortName: strings.ToLower(string(config.StructName[0])),
		StructLowerName: strings.ToLower(config.StructName),
	}

	return t.Execute(file, data)
}
