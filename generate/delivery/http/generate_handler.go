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
// @Param Payload body domain.GenerateRequest true "Payload"
// @Success 201 {object} string
// @Failure 400 {object} domain.ErrorResponse
// @Failure 422 {object} domain.ErrorResponse
// @Router /v1/generate [POST]
func (h *GenerateHandler) GenerateArchitecture(c *gin.Context) {
	var generate domain.GenerateRequest
	if err := c.ShouldBindJSON(&generate); err != nil {
		c.JSON(http.StatusUnprocessableEntity, domain.ErrorResponse{ErrorMessage: err.Error()})
		return
	}

	err := h.GUseCase.Generate(generate)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{ErrorMessage: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"msg": "Created"})
}
