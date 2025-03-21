package infrastructure

import (
	"apiConsumer/src/hotel/application"
	"net/http"
	"github.com/gin-gonic/gin"
)

type GetAllHotelsController struct {
	useCase *application.GetAllHotelsUseCase
}

func NewGetAllHotelsController(useCase *application.GetAllHotelsUseCase) *GetAllHotelsController {
	return &GetAllHotelsController{useCase: useCase}
}

func (controller *GetAllHotelsController) Execute(c *gin.Context) {
	
	hotels, err := controller.useCase.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Hoteles obtenidos correctamente",
		"data":   hotels,
	})
}
