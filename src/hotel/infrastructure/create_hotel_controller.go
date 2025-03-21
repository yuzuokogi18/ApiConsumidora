package infrastructure

import (
	"apiConsumer/src/hotel/application"
	"apiConsumer/src/hotel/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateHotelController struct {
	useCase *application.CreateHotelUseCase
}

func NewCreateHotelController(useCase *application.CreateHotelUseCase) *CreateHotelController {
	return &CreateHotelController{useCase: useCase}
}

func (controller *CreateHotelController) Execute(c *gin.Context) {
	var hotel domain.Hotel

	if err := c.ShouldBindJSON(&hotel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := controller.useCase.Run(&hotel); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "Hotel creado correctamente",
		"data":   hotel,
	})
}
