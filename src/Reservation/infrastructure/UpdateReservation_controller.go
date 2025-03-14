package infrastructure

import (
	"apiConsumer/src/reservation/application"
	"apiConsumer/src/reservation/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateReservationController struct {
	useCase *application.UpdateReservationUseCase
}

func NewUpdateReservationController(useCase *application.UpdateReservationUseCase) *UpdateReservationController {
	return &UpdateReservationController{useCase: useCase}
}

func (controller *UpdateReservationController) Execute(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de reserva no encontrada"})
		return
	}

	var reservation domain.Reservation
	if err := c.ShouldBindJSON(&reservation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	controller.useCase.Run(int32(id), reservation)

	c.JSON(http.StatusOK, gin.H{
		"message": "Reserva actualizada exitosamente",
		"data":    reservation,
	})
}
