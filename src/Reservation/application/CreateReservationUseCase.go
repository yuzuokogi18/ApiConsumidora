package application

import (
	"apiConsumer/src/reservation/domain"
	"fmt"
	"log"
)

type CreateReservationUseCase struct {
	rabbitmqRepository domain.IReservationRabbitmq
	postgresRepository domain.IReservationPostgres
}

func NewCreateReservationUseCase(rabbitmqRepository domain.IReservationRabbitmq, postgresRepository domain.IReservationPostgres) *CreateReservationUseCase {
	return &CreateReservationUseCase{rabbitmqRepository: rabbitmqRepository, postgresRepository: postgresRepository}
}

func (usecase *CreateReservationUseCase) Run(reservation *domain.Reservation) error {
	if err := usecase.postgresRepository.Save(reservation); err != nil {
		log.Printf("Error al guardar en PostgreSQL: %v", err)
		return fmt.Errorf("error al guardar la reservación en PostgreSQL: %w", err)
	}

	if err := usecase.rabbitmqRepository.Save(reservation); err != nil {
		log.Printf("Error al enviar mensaje a RabbitMQ: %v", err)
		return fmt.Errorf("error al enviar la reservación a RabbitMQ: %w", err)
	}

	log.Println("Reservación guardada exitosamente en PostgreSQL y mensaje enviado a RabbitMQ")
	return nil
}
