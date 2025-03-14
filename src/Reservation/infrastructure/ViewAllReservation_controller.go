package infrastructure

import (
	"apiConsumer/src/reservation/application"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ViewAllReservationsController struct { // Corrección en el nombre (agregado 's' en Reservations)
	useCase *application.ViewAllReservationsUseCase
}

func NewViewAllReservationsController(useCase *application.ViewAllReservationsUseCase) *ViewAllReservationsController {
	return &ViewAllReservationsController{useCase: useCase}
}

func (controller *ViewAllReservationsController) Execute(c *gin.Context) {
	reservations, err := controller.useCase.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener las reservas"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reservations": reservations})
}
