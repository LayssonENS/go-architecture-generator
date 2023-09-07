package http

import (
	"net/http"

	"github.com/LayssonENS/go-architecture-generator/domain"
	"github.com/gin-gonic/gin"
)

type GenerateHandler struct {
	GUseCase domain.GenerateUseCase
}

func NewGenerateHandler(routerGroup *gin.Engine, us domain.GenerateUseCase) {
	handler := &GenerateHandler{
		GUseCase: us,
	}

	routerGroup.POST("/v1/generate", handler.GenerateArchitecture)
}

// GenerateArchitecture godoc
// @Summary Route generate
// @Description Create template
// @Tags Generate
// @Accept  json
// @Produce  json
// @Param Payload body domain.ProjectConfig true "Payload"
// @Success 201 {object} string
// @Failure 400 {object} domain.ErrorResponse
// @Failure 422 {object} domain.ErrorResponse
// @Router /v1/generate [POST]
func (h *GenerateHandler) GenerateArchitecture(c *gin.Context) {
	var config domain.ProjectConfig

	if err := c.ShouldBindJSON(&config); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.GUseCase.Generate(config)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{ErrorMessage: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"msg": "Created"})
}
