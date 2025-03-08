package core

import (
	"apiConsumer/src/core/middleware"
	"apiConsumer/src/orders/application"
	"apiConsumer/src/orders/infrastructure"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"  // Importar pq para PostgreSQL
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

	// Casos de uso de la API
	createOrderUseCase := application.NewCreateOrderUseCase(rabbitmqRepository, postgresRepository)
	updateOrderUseCase := application.NewUpdateOrderUseCase(postgresRepository)
	deleteOrderUseCase := application.NewDeleteOrderUseCase(postgresRepository)
	getAllOrdersUseCase := application.NewViewAllOrderUseCase(postgresRepository)
	getOrderByIdUseCase := application.NewViewOrderByIdUseCase(postgresRepository)
	getOrderByCellphoneUseCase := application.NewViewByCellphoneOrderUseCase(postgresRepository)

	// Controladores para cada endpoint
	createOrderController := infrastructure.NewCreateOrderController(createOrderUseCase)
	updateOrderController := infrastructure.NewUpdateOrderController(updateOrderUseCase)
	deleteOrderController := infrastructure.NewDeleteOrderController(deleteOrderUseCase)
	getAllOrdersController := infrastructure.NewViewAllOrderController(getAllOrdersUseCase)
	getOrderByIdController := infrastructure.NewViewByIdOrderController(getOrderByIdUseCase)
	getOrderByCellphoneController := infrastructure.NewViewByCellphoneOrderController(getOrderByCellphoneUseCase)

	// Configurar el router de Gin
	router := gin.Default()
	corsMiddleware := middleware.NewCorsMiddleware()
	router.Use(corsMiddleware)

	// Rutas de la API
	router.POST("/order", createOrderController.Execute)
	router.PUT("/order/:id", updateOrderController.Execute)
	router.DELETE("/order/:id", deleteOrderController.Execute)
	router.GET("/order", getAllOrdersController.Execute)
	router.GET("/order/:id", getOrderByIdController.Execute)
	router.GET("/orders/cellphone/:cellphone", getOrderByCellphoneController.Execute)

	// Iniciar el servidor en el puerto 8082
	if err := router.Run(":8082"); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
