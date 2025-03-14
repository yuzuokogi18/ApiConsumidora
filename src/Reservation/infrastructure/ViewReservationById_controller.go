package infrastructure

import (
	"apiConsumer/src/reservation/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ViewReservationByIdController struct {
	useCase *application.ViewReservationByIdUseCase
}

func NewViewReservationByIdController(useCase *application.ViewReservationByIdUseCase) *ViewReservationByIdController {
	return &ViewReservationByIdController{useCase: useCase}
}

func (controller *ViewReservationByIdController) Execute(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	reservation, err := controller.useCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Reserva no encontrada"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reservation": reservation})
}
