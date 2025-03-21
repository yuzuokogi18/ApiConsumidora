package infrastructure

import (
	"apiConsumer/src/room/application"
	"apiConsumer/src/room/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateRoomController struct {
	useCase *application.CreateRoomUseCase
}

func NewCreateRoomController(useCase *application.CreateRoomUseCase) *CreateRoomController {
	return &CreateRoomController{useCase: useCase}
}

func (controller *CreateRoomController) Execute(c *gin.Context) {
	var room domain.Room

	if err := c.ShouldBindJSON(&room); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := controller.useCase.Run(&room); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "Habitación creada correctamente",
		"data":   room,
	})
}
