package core

import (
	"apiConsumer/src/core/middleware"
	"apiConsumer/src/reservation/application"
	"apiConsumer/src/reservation/infrastructure"	
	"log"
	_ "github.com/lib/pq" 

	"github.com/gin-gonic/gin"
)

func InitRoutes() {
	// Conectar con PostgreSQL
	postgresConn, err := GetDBPool()
	if err != nil {
		log.Fatalf("Error al obtener la conexión a la base de datos: %v", err)
	}

	// Conectar a RabbitMQ
	rabbitmqCh, err := GetChannel()
	if err != nil {
		log.Fatalf("Error al obtener la conexión a RabbitMQ: %v", err)
	}

	// Repositorios para manejar las entidades
	postgresRepository := infrastructure.NewPostgresRepository(postgresConn.DB)
	rabbitmqRepository := infrastructure.NewRabbitRepository(rabbitmqCh.ch)

	// Casos de uso de la API de Reservaciones
	createReservationUseCase := application.NewCreateReservationUseCase(rabbitmqRepository, postgresRepository)
	updateReservationUseCase := application.NewUpdateReservationUseCase(postgresRepository)
	deleteReservationUseCase := application.NewDeleteReservationUseCase(postgresRepository)
	getAllReservationsUseCase := application.NewViewAllReservationsUseCase(postgresRepository)
	getReservationByIdUseCase := application.NewViewReservationByIdUseCase(postgresRepository)

	// Controladores para cada endpoint
	createReservationController := infrastructure.NewCreateReservationController(createReservationUseCase)
	updateReservationController := infrastructure.NewUpdateReservationController(updateReservationUseCase)
	deleteReservationController := infrastructure.NewDeleteReservationController(deleteReservationUseCase)
	getAllReservationsController := infrastructure.NewViewAllReservationsController(getAllReservationsUseCase)
    getReservationByIdController := infrastructure.NewViewReservationByIdController(getReservationByIdUseCase)


	// Configurar el router de Gin
	router := gin.Default()
	corsMiddleware := middleware.NewCorsMiddleware()
	router.Use(corsMiddleware)

	// Rutas de la API de Reservaciones
	router.POST("/reservation", createReservationController.Execute)
	router.PUT("/reservation/:id", updateReservationController.Execute)
	router.DELETE("/reservation/:id", deleteReservationController.Execute)
	router.GET("/reservation", getAllReservationsController.Execute)
	router.GET("/reservation/:id", getReservationByIdController.Execute)

	// Iniciar el servidor en el puerto 8082
	if err := router.Run(":8082"); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
