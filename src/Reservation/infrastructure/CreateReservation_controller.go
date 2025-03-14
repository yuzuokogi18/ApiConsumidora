package infrastructure
import (
	"apiConsumer/src/reservation/application"
	"apiConsumer/src/reservation/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateReservationController struct {
	useCase *application.CreateReservationUseCase
}

func NewCreateReservationController(useCase *application.CreateReservationUseCase) *CreateReservationController {
	return &CreateReservationController{useCase: useCase}
}

func (controller *CreateReservationController) Execute(c *gin.Context) {
	var reservation domain.Reservation

	if err := c.ShouldBindJSON(&reservation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := controller.useCase.Run(&reservation); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "Reserva creada correctamente",
		"data":   reservation,
	})
}
