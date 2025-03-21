package infrastructure

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"apiConsumer/src/reservation/domain"
)

type ReservationExchange struct {
	ch *amqp.Channel
}

func NewRabbitRepository(ch *amqp.Channel) *ReservationExchange {
	if err := ch.ExchangeDeclare(
		"reservation", // Nombre del exchange
		"fanout",       // Tipo del exchange
		true,           // Durable
		false,          // Auto-deleted
		false,          // Internal
		false,          // No-wait
		nil,            // Argumentos
	); err != nil {
		log.Fatalf("Error al declarar el exchange: %v", err)
	}

	return &ReservationExchange{ch: ch}
}

func (ch *ReservationExchange) Save(reservation *domain.Reservation) error {
	body, err := json.Marshal(reservation)
	if err != nil {
		return fmt.Errorf("error al serializar la reserva: %v", err)
	}

	log.Printf("Enviando mensaje: %s", body)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := ch.ch.PublishWithContext(ctx,
		"notification", // Exchange
		"",             // Routing key
		false,          // Mandatory
		false,          // Immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		}); err != nil {
		return fmt.Errorf("error al enviar el mensaje a RabbitMQ: %v", err)
	}

	log.Printf(" [x] Enviado: %s", body)
	return nil
}
