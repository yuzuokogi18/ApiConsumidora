package core

import (
	"apiConsumer/src/core/middleware"
	reservationApp "apiConsumer/src/reservation/application"
	reservationInfra "apiConsumer/src/reservation/infrastructure"
	hotelApp "apiConsumer/src/hotel/application"
	hotelInfra "apiConsumer/src/hotel/infrastructure"
	roomApp "apiConsumer/src/room/application"
	roomInfra "apiConsumer/src/room/infrastructure"
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
	postgresRepository := reservationInfra.NewPostgresRepository(postgresConn.DB)
	rabbitmqRepository := reservationInfra.NewRabbitRepository(rabbitmqCh.ch)

	// Casos de uso de la API de Reservaciones
	createReservationUseCase := reservationApp.NewCreateReservationUseCase(rabbitmqRepository, postgresRepository)
	updateReservationUseCase := reservationApp.NewUpdateReservationUseCase(postgresRepository)
	deleteReservationUseCase := reservationApp.NewDeleteReservationUseCase(postgresRepository)
	getAllReservationsUseCase := reservationApp.NewViewAllReservationsUseCase(postgresRepository)
	getReservationByIdUseCase := reservationApp.NewViewReservationByIdUseCase(postgresRepository)

	// Repositorio de hoteles
	hotelPostgresRepository := hotelInfra.NewHotelPostgresRepository(postgresConn.DB)

	// Casos de uso de la API de Hoteles
	createHotelUseCase := hotelApp.NewCreateHotelUseCase(hotelPostgresRepository)
	getAllHotelsUseCase := hotelApp.NewGetAllHotelsUseCase(hotelPostgresRepository) // Ahora solo pasamos el repositorio de hoteles

	// Repositorio de habitaciones
	roomPostgresRepository := roomInfra.NewRoomPostgresRepository(postgresConn.DB)

	// Caso de uso de creación de habitación
	createRoomUseCase := roomApp.NewCreateRoomUseCase(roomPostgresRepository)

	// Controladores para cada endpoint
	createReservationController := reservationInfra.NewCreateReservationController(createReservationUseCase)
	updateReservationController := reservationInfra.NewUpdateReservationController(updateReservationUseCase)
	deleteReservationController := reservationInfra.NewDeleteReservationController(deleteReservationUseCase)
	getAllReservationsController := reservationInfra.NewViewAllReservationsController(getAllReservationsUseCase)
	getReservationByIdController := reservationInfra.NewViewReservationByIdController(getReservationByIdUseCase)
	createHotelController := hotelInfra.NewCreateHotelController(createHotelUseCase)
	getAllHotelsController := hotelInfra.NewGetAllHotelsController(getAllHotelsUseCase) 
	createRoomController := roomInfra.NewCreateRoomController(createRoomUseCase)

	// Configurar el router de Gin
	router := gin.Default()
	corsMiddleware := middleware.NewCorsMiddleware()
	router.Use(corsMiddleware)


	router.POST("/reservation", createReservationController.Execute)
	router.PUT("/reservation/:id", updateReservationController.Execute)
	router.DELETE("/reservation/:id", deleteReservationController.Execute)
	router.GET("/reservation", getAllReservationsController.Execute)
	router.GET("/reservation/:id", getReservationByIdController.Execute)

	
	router.POST("/hotel", createHotelController.Execute)
	router.GET("/hotel", getAllHotelsController.Execute) 

	
	router.POST("/room", createRoomController.Execute)

	if err := router.Run(":8082"); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
